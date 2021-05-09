<template v-if="login===login">
  <div class="in">Вы зашли в аккаунт под именем {{login}}</div>
  <div class="in">Пока на этой странице нет контента, но он обязательно появится по мере развития!</div>
  <button class="btnOut" @click="logout">Выйти</button>

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
.in {
  font-size: 20px;
}

.btnOut {
  font-size: 18px;
  padding: 0.2em 0.6em;
  background-color: #8998d7;
  color: white;
  border: 0;
  border-radius: 10px;
  transition: 0.24s;
}

.btnOut:hover {
  cursor: pointer;
  background-color: #2c3e50;
}
</style>