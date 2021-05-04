<template>
  <header>
    <ul>
      <li><router-link :to="'/'" @click="e=>e.stopPropagation()"><p>Моя дача</p></router-link></li>
      <li><router-link :to="'/news'" @click="e=>e.stopPropagation()"><p>Новости</p></router-link></li>
      <li><router-link :to="'/register'" @click="e=>e.stopPropagation()"><p>{{accountLinkText}}</p></router-link></li>
    </ul>
  </header>
</template>

<script>
import mainEmitter from "../emitters/mainEmitter";
import loginHandler from "../mixins/loginHandler";

export default {
  name: "SiteHeader",
  data() {
    return {
      accountLinkText: "Войти"
    }
  },
  created() {
    if(this.isLogInFrontend()) {
      this.setAccountLinkLogin(this.getLastLoginInFrontend())
    }
    mainEmitter.$on('login', this.setAccountLinkLogin)
    mainEmitter.$on('logout', this.setAccountLinkLogout)
  },
  methods: {
    setAccountLinkLogin(login) {
      this.accountLinkText = login
    },
    setAccountLinkLogout() {
      this.accountLinkText = "Войти"
    }
  },
  mixins: [loginHandler]
}
</script>

<style scoped>
  header {
    position: fixed;
    top: 0;
    width: 100%;
    height: 60px;
    background-color: #5f77de;
    display: flex;
    /*justify-content: center;*/
    align-items: center;
  }

  header ul {
    margin: 0;
    display: flex;
    flex-direction: row;
    color: #dddfe7;
    font-weight: 500;
  }

  header li {
    margin-left: 3.1em;
  }

  header li:nth-child(1) {
    margin-left: 1.2em;
  }

  header li:hover {
    color: #ffffff;
    text-decoration: underline;
  }
</style>