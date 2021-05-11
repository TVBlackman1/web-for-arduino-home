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
      <div class="nick">
        <div class="profile-image"></div>
        <div class="name">{{login}}<div class="edit btn"></div></div>
        <footer>
          <div class="change-pass">Сменить пароль</div>
        </footer>
      </div>
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

.profile-section {
  position: relative;
}

.profile-image {
  margin: 1em;
  border-radius: 6px;
  border: 1px solid #767676;
  width: 120px;
  height: 120px;
  /*background-color: #8998d7;*/

  background-position: center;
  background-size: cover;
  background-origin: content-box;
  background-repeat: repeat;

  background-image: url("../../assets/profile.png");
}

.change-pass {
  margin-top: 0.4em;
  margin-left: 1em;
  transition: 0.24s;
  -webkit-box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);
  box-shadow: 1px 1px 5px 1px rgba(0,0,0,0.45);
  text-align: center;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0.2em 0.4em;
  border-radius: 6px;
}

footer {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 0.2em;
  background-color: white;
  height: 40px;
  display: flex;
  flex-direction: row;
  justify-content: space-around;
}


.change-pass:hover {
  cursor: pointer;
  color: #2c3e50;
  font-size: 1.02em;
  background-color: #f3f3f3;
}

.name {
  text-align: left;
  margin-left: 2.4em;
}

.btn {
  display: inline-block;
  margin-left: 0em;
  background-color: white;
  background-repeat: no-repeat;
  background-size: contain;
  background-position: center;
  height: 20px;
  width: 30px;
  transition: 0.24s;
  /*-webkit-box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);*/
  /*box-shadow: 1px 1px 5px 1px rgba(0,0,0,0.45);*/
  border-radius: 6px;
  opacity: 0.4;
}

.btn:hover {
  background-color: #dddfe7;
  cursor: pointer;
  opacity: 1;
  -webkit-box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);
  box-shadow: 1px 1px 5px 1px rgba(0,0,0,0.45);
}

.edit {
  background-image: url("../../assets/svg/edit_black_24dp.svg");
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