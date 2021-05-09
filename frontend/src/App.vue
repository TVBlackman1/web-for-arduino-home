<template>

  <site-header/>
  <microphone-help-section />
  <router-view/>
  <microphone-section/>
  <device-pop-menu/>
</template>

<script>
import MicrophoneSection from "./components/MicrophoneSection/MicrophoneSection";
import SiteHeader from "./components/SiteHeader";
import DevicePopMenu from "./components/Device/DevicePopMenu";
import MicrophoneHelpSection from "./components/MicrophoneSection/MicrophoneHelpSection";

export default {
  name: 'App',
  components: {
    MicrophoneHelpSection,
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
