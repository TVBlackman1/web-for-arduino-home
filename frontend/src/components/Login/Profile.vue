<template v-if="login===login">
  <div class="in">Вы зашли в аккаунт под именем {{ login }}</div>
  <button class="btnOut" @click="logout">Выйти</button>
  <hr/>
  <!--  <div class="account">-->
  <!--    <div class="left-part">-->
  <!--      <div class="profile-image"></div>-->
  <!--    </div>-->
  <!--    <div class="center"></div>-->
  <!--    <div class="right-part">-->
  <!--      <friend-list />-->
  <!--    </div>-->
  <!--  </div>-->

  <div class="account">
    <div class="title-section">Профиль</div>
    <div class="title-section">Режимы работы</div>
    <div class="title-section">Состояние системы</div>
    <div class="title-section">Умные очки</div>
    <div class="profile-section sect">
      <div class="nick">{{ login }}</div>
    </div>
    <div class="mode-section sect">
      <mode-section/>
    </div>
    <div class="access-section sect">
      <!--      <friend-list />-->
      <ul class="system-state-ul">
        <li class="system-state">Подключено устройств: <p class="state-value">4</p></li>
        <li class="system-state">Работает устройств: <p class="state-value">4</p></li>
        <li class="system-state">Сбой в подключении: <p class="state-value">0</p></li>
      </ul>

      <div class="verdict pos">Всё работает исправно</div>
<!--      <div class="verdict neg">Ошибка в подключении</div>-->
    </div>
    <div class="sect glasses">
      <div class="center pos">
        Подключены
      </div>
    </div>
  </div>


</template>

<script>
import loginHandler from "../../mixins/loginHandler";
import ModeSection from "./ModeSection";

export default {
  name: "Profile",
  components: {ModeSection},
  data() {
    return {
      login: ""
    }
  },
  created() {
    if (this.isLogInFrontend()) {
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

.account {
  display: grid;
  grid-template-columns: 300px 300px 300px 300px;
  grid-column-gap: 100px;
  margin: 1em;
}

.title-section {
  padding: 1.2em;
  text-align: left;
  font-size: 24px;
}

.sect {
  border-left: 1px solid #2c3e50;
}

/*.account {
/*  margin: -2em auto auto;*/
/*  width: 1700px;*/
/*  height: 800px;*/
/*  -webkit-box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);*/
/*  box-shadow: 2px 2px 10px 1px rgba(0,0,0,0.45);*/
/*  padding: 0.2em 0.6em;*/
/*  display: grid;*/
/*  grid-template-columns: 460px 800px 1fr;*/
/*  grid-column-gap: 20px;*/
/*  box-sizing: border-box;*/
/*  background-color: #dee5f6;*/
/*}*/

/*.account > div{*/
/*  background-color: #40b883;*/
/*}*/

.in {
  margin: 1em;
  font-size: 20px;
}

div.in {
  margin-top: 5.2em;
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

/*.profile-image {*/
/*  margin: 1em 1em;*/
/*  height: 280px;*/
/*  width: 200px;*/
/*  background-color: #dddfe7;*/
/*}
/**/

.nick {
  font-size: 20px;
}

.system-state-ul {
  margin: 0.4em 1.2em;
  text-align: left;
  font-size: 20px;
}

.system-state {
  /*text-align: left;*/
  font-size: 20px;
  padding: 0.18em 0;
}

.state-value {
  display: inline-block;
  font-weight: 600;
}

.verdict {
  margin: 2em;
  font-size: 22px;
  cursor: pointer;
}

.pos {
  color: #40b883;
}

.neg {
  color: #b84040;
}

.center {
  font-size: 24px;
}

.glasses {
  display: flex;
  justify-content: center;
  align-items: center;
}

</style>