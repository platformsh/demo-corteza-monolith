<template>
  <div id="roomselection">
    <span>{{ $t('app.jitsi.toStart') }}</span>
    <div id="roomdropdown">
      <select v-model="channelID">
        <option v-for="(c) in channels"
                :key="c.channelID"
                :value="c.channelID">{{ c.name }}</option>
      </select>
    </div>
    <button :disabled="jitsi || (!channelID)"
            @click="onJoin">{{ $t('app.jitsi.join') }}</button>
    <h4>{{ $t('app.jitsi.or') }}</h4>
    <span>{{ $t('app.jitsi.createNewRoom') }}</span>
    <input type="text" id="roomInputField" v-model="roomName" :placeholder="$t('app.jitsi.roomName')" />

    <button :disabled="jitsi || (cleanup(roomName).length === 0)"
            @click="onCreate">{{ $t('app.jitsi.create') }}</button>

    <div ref="jitsiInterface"
         v-show="jitsi"
         class="jitsiInterface"></div>
  </div>
</template>
<script>
import Vue from 'vue'
import LoadScript from 'vue-plugin-load-script'

Vue.use(LoadScript)

Vue.loadScript('https://meet.jit.si/external_api.js')

const domain = 'meet.jit.si'

export default {
  name: 'JitsiBridge',

  params: {
    user: {
      type: Object,
      required: true,
    },
  },

  data () {
    return {
      channelID: null,
      roomName: '',
      jitsi: null,
      channels: null,
    }
  },

  beforeCreate () {
    const userID = this.$auth.user.userID
    const filter = (c) => {
      return (c.type === 'public' || c.type === 'private') &&
        (c.membershipFlag !== 'hidden') &&
        (c.members || []).findIndex((uID) => uID === userID) > -1
    }

    this.$MessagingAPI.channelList().then((cc) => {
      this.channels = cc.filter(filter)
    })
  },

  destroyed () {
    this.dispose()
  },

  methods: {
    dispose () {
      if (this.jitsi) {
        this.jitsi.dispose()
        this.jitsi = null
      }
    },

    cleanup (str) {
      return str.replace(/[^a-z0-9+]+/gi, '')
    },

    onJoin () {
      this.open({
        roomName: this.channelID,
        userDisplayName: this.$auth.user.name || this.$auth.user.email,
      })
    },

    onCreate () {
      this.open({
        roomName: this.roomName,
        userDisplayName: this.$auth.user.name || this.$auth.user.email,
      })
    },

    onClose () {
      this.dispose()
    },

    removeJitsiAfterHangup () {
      this.dispose()
    },

    open ({ roomName, userDisplayName } = {}) {
      this.dispose()

      const $t = (k) => this.$t(k)

      /* eslint-disable no-undef */
      this.jitsi = new JitsiMeetExternalAPI(domain, {
        roomName: `crust_${this.cleanup(roomName || 'unnamed')}`,
        width: '100%',
        height: '100%',
        parentNode: this.$refs.jitsiInterface,
        interfaceConfigOverwrite: {
          DEFAULT_BACKGROUND: '#232323',
          SHOW_JITSI_WATERMARK: true,
          SHOW_WATERMARK_FOR_GUESTS: false,
          SHOW_BRAND_WATERMARK: false,
          BRAND_WATERMARK_LINK: '',
          SHOW_POWERED_BY: false,
          DEFAULT_REMOTE_DISPLAY_NAME: $t('app.jitsi.defaultRemoteDisplayName'),
          DEFAULT_LOCAL_DISPLAY_NAME: userDisplayName || $t('app.jitsi.defaultLocalDisplayName'),
          TOOLBAR_BUTTONS: [
            'microphone',
            'camera',
            'closedcaptions',
            'desktop',
            'fullscreen',
            'fodeviceselection',
            'hangup',
            'profile',
            'info',
            'recording',
            'settings',
            'tileview',
            'videoquality',
            'filmstrip',
            'invite',
            'shortcuts',
          ],
          SETTINGS_SECTIONS: [
            'devices',
            'language',
            'moderator',
            'profile',
            'calendar',
          ],
        },
      })

      // add an event listner to remove jitsi after the local party has hung up the call.
      // this is to remove the page that mentions slack after the rating page.
      this.jitsi.addEventListeners({
        readyToClose: this.removeJitsiAfterHangup,
      })

      window.jitsi = this.jitsi
    },
  },
}
</script>
<style lang="scss" scoped>
html, body {
  margin: 0;
  overflow: hidden;
}

.jitsiInterface {
  position: absolute;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  background: #232323;

  & > iframe {
    flex: 1 1 auto;
  }
}

#roomselection {
  max-width: 600px;
  margin: 50px auto;
  padding: 40px;
  background-color: #fff;
  font-family: Arial;
}

input {
  height: 30px;
  width: 100%;
  border: 1px solid $secondary;
  padding-left: 10px;
  font-size: 14px;
  display: block;
  margin-top: 10px;
  -webkit-box-sizing: border-box;
  -moz-box-sizing: border-box;
  box-sizing: border-box;
}

select {
  height: 30px;
  width: 100%;
  margin-top: 10px;
  background: transparent;
  padding-left: 10px;
  font-size: 14px;
  -webkit-border-radius:0px;
  -moz-border-radius:0px;
  border-radius:0px;
  -webkit-appearance:none;
  -moz-appearance:none;
  appearance:none;
  border: 1px solid $secondary;
}

#roomdropdown::after {
  border: 4px dashed transparent;
  border-top: 4px solid $secondary;
  content: "";
  display: inline-block;
  float: right;
  margin-right: 10px;
  margin-top: -15px;
}

select:focus,
input:focus {
  outline: none;
}

button {
  cursor: pointer;
  background: transparent;
  color: $primary;
  font-size: 14px;
  line-height: 38px;
  text-decoration: none;
  display: block;
  width: 150px;
  text-align: center;
  height: 40px;
  margin: 20px auto 0;
  -webkit-transition: color .2s,background-color .2s;
  transition: color .2s,background-color .2s;
  border: 1px solid $primary;
  &:hover {
    border: 1px solid $primary;
    background: $primary;
    color: #ffffff;
  }
  &:disabled {
    cursor: not-allowed;
    color: $secondary;
    border-color: $secondary;
    &:hover {
      background: transparent;
    }
  }
}

h4 {
  display: flex;
  width: 100%;
  justify-content: center;
  align-items: center;
  text-align: center;
  margin: 30px 0;
  color: $secondary;
  &:before,
  &:after {
    content: '';
    border-top: 1px solid $secondary;
    margin: 0 20px 0 0;
    flex: 1 0 20px;
  }
  &:after {
    margin: 0 0 0 20px;
  }
}

</style>
