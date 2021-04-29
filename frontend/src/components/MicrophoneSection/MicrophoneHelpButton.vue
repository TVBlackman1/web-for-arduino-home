<template>
  <div class="micro-image bottom-fixed" :style="styleObject" />
</template>

<script>
import MicrophoneHelpSVG from "../../assets/svg/micro-help-24.svg";
import MicrophoneHelpSVGEnabled from "../../assets/svg/micro-help-enabled-24.svg";

export default {
  name: "MicrophoneHelpButton",
  data() {
    return {
      MicrophoneHelpSVG,
      MicrophoneHelpSVGEnabled
    }
  },
  computed: {
    styleObject() {
      return {
        '--width': '44px',
        '--height': '44px',
        '--background-image': `url(${this.MicrophoneHelpSVG})`,
        '--background-image-hover': `url(${this.MicrophoneHelpSVGEnabled})`
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