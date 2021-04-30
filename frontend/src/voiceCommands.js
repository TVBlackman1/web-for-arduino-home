import popMenuOpener from "./transfers/popMenuOpener";
import serverDefaultRequests from "./mixins/serverDefaultRequests";

export default {
    data() {
        return {

        }
    },
    methods: {
        async do(textCommand) {
            if(textCommand === "устройство метеостанции" || textCommand === "устройство метеостанция") {
                let device = await this.serverRequest("/api/device/weather-station")
                popMenuOpener.$emit('set-content-pop-menu', device)
                popMenuOpener.$emit('toggle-pop-menu')
            }
            if(textCommand === "устройство курятник") {
                let device = await this.serverRequest("/api/device/chicken-coop")
                popMenuOpener.$emit('set-content-pop-menu', device)
                popMenuOpener.$emit('toggle-pop-menu')
            }
        }
    },
    mixins: [serverDefaultRequests]
}