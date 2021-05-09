import serverHandler from "./serverHandler";

export default {
    methods: {
        async getDeviceById(id) {
            return await this.serverRequest("/api/device/" + id)
        },
        async getDevices() {
            return await this.serverRequest("/devices/all-devices")
        },
        async register(account =  {login: "", password: ""}) {
            return await this.serverRequest("/auth/sign-up", account)
        },
        async login(account =  {login: "", password: ""}) {
            return await this.serverRequest("/auth/sign-in", account)
        },
        async getNews() {
            return await this.serverRequest("/news")
        }
    },
    mixins: [serverHandler]
}