<template>
  <div class="micro-image bottom-fixed" :style="styleObject" />
</template>

<script>
import MicrophoneSettingsSVG from "../assets/svg/micro-settings-24.svg";
import MicrophoneSettingsSVGEnabled from "../assets/svg/micro-settings-enabled-24.svg";

export default {
  name: "MicrophoneButton",
  data() {
    return {
      MicrophoneSettingsSVG,
      MicrophoneSettingsSVGEnabled
    }
  },
  computed: {
    styleObject() {
      return {
        '--width': '44px',
        '--height': '44px',
        '--background-image': `url(${this.MicrophoneSettingsSVG})`,
        '--background-image-hover': `url(${this.MicrophoneSettingsSVGEnabled})`
      }
    }
  }
}
</script>

<style scoped>

.micro-image {
  background-image: var(--background-image);
  /*var(--background-image-hover);*/
  width: var(--width);
  height: var(--height);
  background-size: cover;
}
.micro-image:hover {
  cursor: pointer;
  background-image: var(--background-image-hover);
}

</style>