export default {
    data() {
        return {
            ip: "5.164.202.94:80",
            reqIsReady: false
        }
    },
    methods: {
        getServerAddress() {
            return "http://" + this.ip
        },
        async serverRequest(apiAddress, body = {}) {
            const address = this.getServerAddress() + apiAddress
            const requestOptions = {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(body)
            };

            const response = await fetch(address, requestOptions);
            const data =  await response.json();
            this.reqIsReady = true
            return data
        }
    }
}