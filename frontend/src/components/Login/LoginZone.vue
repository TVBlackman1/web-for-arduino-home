<template>
  <div class="center">
    <p>-- ЛОГИН --</p>
    <input type="text" placeholder="логин" v-model="accountInLogin.name">
    <input type="password" placeholder="пароль" v-model="accountInLogin.password">
    <button @click="loginAccount">Войти</button>

    <p>-- РЕГИСТРАЦИЯ --</p>
    <input type="text" placeholder="логин" v-model="accountInReg.name">
    <input type="password" placeholder="пароль" v-model="accountInReg.password">
    <button @click="registerAccount">Зарегистрироваться</button>
    <div v-if="text.length !== 0" class="mini-text">{{text}}</div>


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
      text: "",
      accountInLogin: {
        name: "",
        password: ""
      },
      accountInReg: {
        name: "",
        password: ""
      }
    }
  },
  created() {
    this.isLog1 = this.isLog()
    if(this.isLog1)
      this.redirectToProfile()
  },
  methods: {
    async registerAccount() {
      const res = await this.register(this.accountInReg)

      console.log(res)
      this.text = res.res
    },
    async loginAccount() {
      let login = this.accountInLogin
      const res = await this.login(login)
      if (res.res === "OK") {
        this.loginInFrontend(login.name)
        this.redirectToProfile()
      }
      this.isLog1 = this.isLog()
      console.log(res)
    },
    logout() {
      this.logoutInFrontend()
      this.isLog1 = this.isLogInFrontend()
    },
    redirectToProfile() {
      this.$router.push('profile')
    }
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

.mini-text {
  font-size: 18px;
  background-color: #2c3e50;
  color: #b9f5b9;
  padding: 0.5em;
  border-radius: 20px;
}
</style>