import serverHandler from "./serverHandler";

export default {
    methods: {
        async getDeviceById(id) {
            return await this.serverRequest("/api/device/" + id)
        },
        async getDevices() {
            return await this.serverRequest("/api/devices")
        },
        async register(account =  {login: "", password: ""}) {
            return await this.serverRequest("/api/register", account)
        }
    },
    mixins: [serverHandler]
}