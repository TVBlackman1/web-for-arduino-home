import serverDefaultRequests from "./mixins/serverDefaultRequests";
import popMenuHandler from "./mixins/popMenuHandler";

export default {
    data() {
        return {

        }
    },
    methods: {
        async openDevicePoMenu(textCommand) {
            let deviceId = ""

            if(textCommand === "метеостанции" || textCommand === "метеостанция") {
                deviceId = "weather-station"
            } else if (textCommand === "курятник") {
                deviceId = "chicken-coop"
            }

            if(deviceId === "")
                return

            let device = await this.getDeviceById(deviceId)
            this.__setContent(device)
            this.__toggle()
        },
        async do(textCommand) {
            await this.openDevicePoMenu(textCommand)
        }
    },
    mixins: [serverDefaultRequests, popMenuHandler]
}