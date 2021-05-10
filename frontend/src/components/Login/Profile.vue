<template v-if="login===login">
  <div class="in">Вы зашли в аккаунт под именем {{login}}</div>
  <div class="in">Пока на этой странице нет контента, но он обязательно появится по мере развития!</div>
  <button class="btnOut" @click="logout">Выйти</button>
  <hr />
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
    <div class="title-section">Настроить доступ</div>
    <div class="profile-section sect">
      <div class="nick">tvblackman1</div>
    </div>
    <div class="mode-section sect">
      <mode-section />
    </div>
    <div class="access-section sect">
      <friend-list />
    </div>
  </div>


</template>

<script>
    import loginHandler from "../../mixins/loginHandler";
    import ModeSection from "./ModeSection";
    import FriendList from "./friend-list";
    export default {
        name: "Profile",
      components: {FriendList, ModeSection},
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

.account {
  display: grid;
  grid-template-columns: 300px 300px 300px;
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

</style>