import serverHandler from "./serverHandler";

export default {
    methods: {
        async getDeviceById(id) {
            return await this.serverRequest("/api/device/" + id)
        },
        async getDevices() {
            return await this.serverRequest("/api/devices")
        }
    },
    mixins: [serverHandler]
}