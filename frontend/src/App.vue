<template>

  <site-header/>
  <router-view/>
  <microphone-section/>
  <device-pop-menu/>
</template>

<script>
import MicrophoneSection from "./components/MicrophoneSection/MicrophoneSection";
import SiteHeader from "./components/SiteHeader";
import DevicePopMenu from "./components/Device/DevicePopMenu";

export default {
  name: 'App',
  components: {
    SiteHeader,
    MicrophoneSection,
    DevicePopMenu
  }
}
</script>

<style>
@import "css/reset.css";

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
