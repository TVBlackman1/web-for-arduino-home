<template>
  <div v-if="!isLog1" class="center">
    <p>-- ЛОГИН --</p>
    <input type="text" placeholder="логин" v-model="accountInLogin.login">
    <input type="password" placeholder="пароль" v-model="accountInLogin.password">
    <button @click="loginAccount">Войти</button>

    <p>-- РЕГИСТРАЦИЯ --</p>
    <input type="text" placeholder="логин" v-model="accountInReg.login">
    <input type="password" placeholder="пароль" v-model="accountInReg.password">
    <button @click="registerAccount">Зарегистрироваться</button>
  </div>
  <div v-else>
    Вы вошли в аккаунт под логином <b>{{getLastLogin()}}</b>

    <button @click="logout">Выйти</button>
  </div>
</template>

<script>
import serverDefaultRequests from "../../mixins/serverDefaultRequests";
import loginHandler from "../../mixins/loginHandler";

export default {
  name: "LoginZone",
  data() {
    return {
      isLog1: false,
      accountInLogin: {
        login: "",
        password: ""
      },
      accountInReg: {
        login: "",
        password: ""
      }
    }
  },
  created() {
    this.isLog1 = this.isLog()
  },
  methods: {
    async registerAccount() {
      const res = await this.register(this.accountInReg)

      console.log(res)
    },
    async loginAccount() {
      let login = this.accountInLogin
      const res = await this.login(login)
      if (res.res === "OK") {
        this.loginInFrontend(login.login)
      }
      this.isLog1 = this.isLog()

      console.log(res)
    },
    logout() {
      this.logoutInFrontend()
      this.isLog1 = this.isLogInFrontend()
    },
  },
  mixins: [serverDefaultRequests, loginHandler]
}
</script>

<style scoped>
.center {
  margin: 12em auto auto;
  -webkit-box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);
    box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);
  width: 600px;
  height: 520px;
  background-color: #8998d7;
  padding: 2em;
  box-sizing: border-box;
}

.center input {
  font-size: 30px;
  border-radius: 20px;
  border: 0;
  margin: 0.2em;
  padding: 0.2em 0.4em;
  box-sizing: border-box;
}

.center input::placeholder {
  color: #bdbdbd;
  font-size: 0.8em;
}

.center p {
  font-size: 24px;
  color: #f5f5f5;
  margin: 0.4em 0.8em 0.8em;
}

.center button {
  font-size: 30px;
  border-radius: 20px;
  border: 0;
  padding: 0.2em 1.2em;
  box-sizing: border-box;
  background-color: #faff8b;
  transition: .24s;
  color: #597189;
  display: block;
  margin: auto auto 1.4em auto;
}

.center button:hover {
  cursor: pointer;
  background-color: #f3ff0a;

}
</style>