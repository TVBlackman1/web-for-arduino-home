<template>
  <div class="card-device">
    <div class="card-device-image-holder">
      <div v-if="imgSrc==null">None</div>
      <img v-else :src="imgSrc" class="card-device-picture">
    </div>
    <div class="card-content">
      <header>{{ device.name }}</header>
<!--      <router-link :to="getLink()">Перейти</router-link>-->
      <device-picker :device="device"/>
    </div>
    <div class="card-device-pop-button">
      <router-link :to="getLink()"><p>Перейти</p></router-link>
    </div>
  </div>
</template>

<script>
import DevicePicker from "./DevicePicker";

export default {
  name: "Device",
  props: {device: Object},
  data() {
    return {
      imgSrc: null,
    }
  },
  provide() {
    return {
      setDeviceImageSrc: this.setDeviceImageSrc
    }
  },
  mounted() {
    // console.log("Created device:", this.device)
  },
  components: {
    DevicePicker
  },
  methods: {
    getLink() {
      return "/device/" + this.device.id
    },

    setDeviceImageSrc(imageSrc) {
      this.imgSrc = imageSrc
    }
  }
}
</script>

<style scoped>
  @import "../css/device-card.css";
/*header {*/
/*  font-weight: 700;*/
/*  font-size: 1.3em;*/
/*  cursor: pointer;*/
/*}*/

/*.device {*/
/*  background-color: #fff;*/
/*  width: 480px;*/
/*  height: 280px;*/
/*  box-sizing: border-box;*/

/*  border-style: solid;*/
/*  border-color: #40b883;*/
/*  border-radius: 2px;*/
/*  border: 0;*/
/*  !*border-width: 8px 1px 3px 1px;*!*/

/*  -webkit-box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);*/
/*  box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);*/

/*  transition: 0.24s;*/
/*}*/

/*.image-holder {*/
/*  overflow: hidden;*/
/*  height: 140px;*/
/*  width: 100%;*/
/*}*/

/*.device-picture {*/
/*  width: 100%;*/
/*  max-height: 140px;*/
/*  background-position: center;*/
/*  background-size: cover;*/
/*  background-origin: content-box;*/
/*  background-repeat: repeat;*/
/*  object-fit: cover;*/
/*  object-position: 50% 50%;*/

/*  transition: 0.24s;*/
/*}*/

/*.device:hover {*/
/*  background-color: #f3f6ff;*/
/*  cursor: pointer;*/
/*}*/

/*.device:hover .device-picture {*/
/*  transform: scale(1.1);*/
/*}*/

/*.content {*/
/*  padding: 0.4em;*/
/*}*/

</style>