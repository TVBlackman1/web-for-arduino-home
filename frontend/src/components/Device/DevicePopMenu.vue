<template>
  <transition name="fade">
    <div v-if="show" class="dark-background" @click="close">
      <div class="pop-menu" @click="popMenuClick">
<!--        Подробная информация об этом устройстве экосистемы {{content}}-->
        <device-page-picker :device="content"/>
      </div>
    </div>
  </transition>
</template>

<script>
import popMenuOpener from "../../emitters/popMenuOpener";
import DevicePagePicker from "./Picker/DevicePagePicker";

export default {
  name: "DevicePopMenu",
  data() {
    return {
      show: false,
      content: {}
    }
  },
  mounted() {
    popMenuOpener.$on('toggle-pop-menu', this.toggle)
    popMenuOpener.$on('set-content-pop-menu', this.setContent)
  },
  methods: {
    close() {
      this.show = false;
    },
    toggle() {
      this.show = !this.show
    },
    popMenuClick(event) {
      event.stopPropagation()
    },
    setContent(content) {
      this.content = content
      console.log(this.content)
    }
  },
  components: {
    DevicePagePicker
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
  padding: 2em 0.4em;
  overflow-y: auto;
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