<template>
  <div class="center">
    <p>-- ЛОГИН --</p>
    <input type="text" placeholder="логин" v-model="account.login">
    <input type="password" placeholder="пароль" v-model="account.password">
    <button @click="login">Войти</button>
  </div>
</template>

<script>
import serverDefaultRequests from "../../mixins/serverDefaultRequests";

export default {
  name: "LoginZone",
  data() {
    return {
      account: {
        login: "",
        password: ""
      }
    }
  },
  methods: {
    async login() {
      console.log(this.account)
      const res = await this.register(this.account)

      console.log(res)
    }
  },
  mixins: [serverDefaultRequests]
}
</script>

<style scoped>
.center {
  margin: 12em auto auto;
  -webkit-box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);
    box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);
  width: 600px;
  height: 400px;
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
  margin: 0.2em 0.8em 0.8em;
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
  margin: auto;
}

.center button:hover {
  cursor: pointer;
  background-color: #f3ff0a;

}
</style>