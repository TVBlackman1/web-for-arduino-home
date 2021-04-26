export default {
    data() {
        return {
            ip: "5.164.198.55",
            reqIsReady: false
        }
    },
    methods: {
        getServerAddress() {
            return "http://" + this.ip
        },
        async serverRequest(apiAddress) {
            const address = this.getServerAddress() + apiAddress
            const requestOptions = {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                // body: JSON.stringify({id: this.$route.params.id})
            };

            const response = await fetch(address, requestOptions);
            const data =  await response.json();
            this.reqIsReady = true
            return data
        }
    }
}