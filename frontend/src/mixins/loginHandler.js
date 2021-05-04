import mainEmitter from "../emitters/mainEmitter";
import localStorageHandler from "./localStorageHandler";

export default {
    methods: {
        loginInFrontend(login) {
            mainEmitter.$emit('login', login)
            this.setLastLogin(login)
            this.setLog('true')
        },
        logoutInFrontend() {
            mainEmitter.$emit('logout')
            this.setLog('false')
        },
        isLogInFrontend() {
            return this.isLog()
        },
        getLastLoginInFrontend() {
            return this.getLastLogin()
        },
    },
    mixins: [localStorageHandler]
}