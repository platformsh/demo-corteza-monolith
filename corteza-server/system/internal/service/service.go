package service

import (
	"context"

	"go.uber.org/zap"

	intAuth "github.com/cortezaproject/corteza-server/internal/auth"
	"github.com/cortezaproject/corteza-server/internal/permissions"
	internalSettings "github.com/cortezaproject/corteza-server/internal/settings"
	"github.com/cortezaproject/corteza-server/pkg/automation"
	"github.com/cortezaproject/corteza-server/pkg/automation/corredor"
	"github.com/cortezaproject/corteza-server/pkg/cli/options"
	"github.com/cortezaproject/corteza-server/system/internal/repository"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	db interface {
		Transaction(callback func() error) error
	}
	permissionServicer interface {
		accessControlPermissionServicer
		Watch(ctx context.Context)
	}

	Config struct {
		Storage          options.StorageOpt
		Corredor         options.CorredorOpt
		GRPCClientSystem options.GRPCServerOpt
	}
)

var (
	DefaultLogger *zap.Logger

	// DefaultPermissions Retrieves & stores permissions
	DefaultPermissions permissionServicer

	DefaultIntSettings internalSettings.Service
	DefaultSettings    SettingsService

	// DefaultAccessControl Access control checking
	DefaultAccessControl *accessControl

	// DefaultAutomationScriptManager manages scripts
	DefaultAutomationScriptManager automationScript

	// DefaultAutomationTriggerManager manages triggerManager
	DefaultAutomationTriggerManager automationTrigger

	// DefaultAutomationRunner runs automation scripts by listening to triggerManager and invoking Corredor service
	DefaultAutomationRunner automationRunner

	DefaultAuthNotification AuthNotificationService
	DefaultAuthSettings     AuthSettings

	DefaultSink *sink

	DefaultAuth         AuthService
	DefaultUser         UserService
	DefaultRole         RoleService
	DefaultOrganisation OrganisationService
	DefaultApplication  ApplicationService
)

func Init(ctx context.Context, log *zap.Logger, c Config) (err error) {
	DefaultLogger = log.Named("service")

	DefaultIntSettings = internalSettings.NewService(internalSettings.NewRepository(repository.DB(ctx), "sys_settings"))

	DefaultPermissions = permissions.Service(
		ctx,
		DefaultLogger,
		permissions.Repository(repository.DB(ctx), "sys_permission_rules"))
	DefaultAccessControl = AccessControl(DefaultPermissions)

	DefaultSettings = Settings(ctx, DefaultIntSettings)

	DefaultUser = User(ctx)
	DefaultRole = Role(ctx)
	DefaultOrganisation = Organisation(ctx)
	DefaultApplication = Application(ctx)

	// Authentication helpers & services
	DefaultAuthSettings, err = DefaultSettings.LoadAuthSettings()
	if err != nil {
		return
	}
	DefaultAuthNotification = AuthNotification(ctx)
	DefaultAuth = Auth(ctx)

	// ias: Internal Automatinon Service
	// handles script & trigger management & keeping runnables cripts in internal cache
	ias := automation.Service(automation.AutomationServiceConfig{
		Logger:        DefaultLogger,
		DbTablePrefix: "sys",
		DB:            repository.DB(ctx),
		TokenMaker: func(ctx context.Context, userID uint64) (jwt string, err error) {
			var u *types.User

			ctx = intAuth.SetSuperUserContext(ctx)
			if u, err = DefaultUser.FindByID(userID); err != nil {
				return
			} else if err = DefaultAuth.LoadRoleMemberships(u); err != nil {
				return
			}

			return intAuth.DefaultJwtHandler.Encode(u), nil
		},
	})

	// Pass automation manager to
	DefaultAutomationTriggerManager = AutomationTrigger(ias)
	DefaultAutomationScriptManager = AutomationScript(ias)

	{
		var scriptRunnerClient corredor.ScriptRunnerClient

		if c.Corredor.Enabled {
			conn, err := corredor.NewConnection(ctx, c.Corredor, DefaultLogger)

			log.Info("initializing corredor connection", zap.String("addr", c.Corredor.Addr), zap.Error(err))
			if err != nil {
				return err
			}

			scriptRunnerClient = corredor.NewScriptRunnerClient(conn)
		}

		DefaultAutomationRunner = AutomationRunner(
			AutomationRunnerOpt{
				ApiBaseURLSystem:    c.Corredor.ApiBaseURLSystem,
				ApiBaseURLMessaging: c.Corredor.ApiBaseURLMessaging,
				ApiBaseURLCompose:   c.Corredor.ApiBaseURLCompose,
			},
			ias,
			scriptRunnerClient,
		)
	}

	DefaultSink = Sink()

	return
}

func Watchers(ctx context.Context) {
	// Reloading automation scripts on change
	DefaultAutomationRunner.Watch(ctx)

	// Reloading permissions on change
	DefaultPermissions.Watch(ctx)
}
