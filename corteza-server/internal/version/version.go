package version

import (
	"net/http"

	"github.com/platformsh-upstream-forks/factory/resputil"
)

var (
	BuildTime = "undefined"
	Version   = "undefined"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	resputil.JSON(w, struct {
		BuildTime string `json:"buildTime"`
		Version   string `json:"version"`
	}{BuildTime, Version})
}
