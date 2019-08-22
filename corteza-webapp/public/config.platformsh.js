const config = require("platformsh-config").config();

const system_api = config.credentials("systemapi")
const messaging_api = config.credentials("systemapi")
const compose_api = config.credentials("composeapi")

window.SystemAPI = system_api.host
window.MessagingAPI = messaging_api.host
window.ComposeAPI = compose_api.host
