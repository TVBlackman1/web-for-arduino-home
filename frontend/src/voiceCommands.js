import serverDefaultRequests from "./mixins/serverDefaultRequests";
import popMenuHandler from "./mixins/popMenuHandler";

export default {
    data() {
        return {

        }
    },
    methods: {
        async openDevicePoMenu(textCommand) {
            let deviceName = ""

            if(textCommand === "метеостанции" || textCommand === "метеостанция") {
                deviceName = "weather-station"
            } else if (textCommand === "курятник") {
                deviceName = "chicken-coop"
            }

            if(deviceName === "")
                return

            let device = await this.getDeviceById("weather-station")
            this.__setContent(device)
            this.__toggle()
        },
        async do(textCommand) {
            await this.openDevicePoMenu(textCommand)
        }
    },
    mixins: [serverDefaultRequests, popMenuHandler]
}