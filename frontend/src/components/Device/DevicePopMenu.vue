<template>
  <transition name="fade">
    <div v-if="show" class="dark-background" @click="close">
      <div class="pop-menu" @click="popMenuClick"></div>
    </div>
  </transition>
</template>

<script>
export default {
  name: "DevicePopMenu",
  data() {
    return {
      show: true
    }
  },
  methods: {
    close() {
      this.show = false;
    },
    popMenuClick(event) {
      event.stopPropagation()
    }
  }
}
</script>

<style scoped>

.dark-background {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background-color: rgba(36, 51, 66, 0.20);
  /*filter: blur(4px);*/
  /*backdrop-filter: blur(2px);*/
  z-index: 10;
  display: flex;
  align-items: center;
  justify-content: center;
}

.pop-menu {
  width: 46%;
  height: 660px;
  background-color: white;
  border-radius: 8px;
  -webkit-box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);
  box-shadow: 2px 2px 10px 1px rgba(0,0,0,0.45);
  /*position: absolute;*/
  /*z-index: 11;*/
}

.fade-enter-active, .fade-leave-active {
  transition: .24s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active до версии 2.1.8 */ {
  opacity: 0;
  /*transform: translateY(-100px);*/
}

/*.fade-enter.pop-menu, .fade-leave-to.pop-menu {*/
/*  !*opacity: 0;*!*/
/*  transform: translateY(-100px);*/
/*}*/

</style>