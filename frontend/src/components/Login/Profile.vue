<template v-if="login===login">
    <div>{{login}}</div>
    <button @click="logout">Выйти</button>

</template>

<script>
    import loginHandler from "../../mixins/loginHandler";
    export default {
        name: "Profile",
        data() {
            return {
                login: ""
            }
        },
        created() {
            if(this.isLogInFrontend()) {
                this.login = this.getLastLoginInFrontend()
            }
        },
        methods: {
            logout() {
                this.logoutInFrontend()
                this.redirectToReg()
            },
            redirectToReg() {
                this.$router.push('register')
            }
        },
        mixins: [loginHandler]
    }
</script>

<style scoped>

</style>