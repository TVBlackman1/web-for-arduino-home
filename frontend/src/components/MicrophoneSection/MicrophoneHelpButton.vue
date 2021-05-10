<template>
  <div class="micro-image" :style="styleObject" @click="toggleSection">
    <microphone-help-section />
  </div>
</template>

<script>
import MicrophoneHelpSVG from "../../assets/svg/micro-help-24.svg";
import MicrophoneHelpSVGEnabled from "../../assets/svg/micro-help-enabled-24.svg";
import mainEmitter from "../../emitters/mainEmitter";
import MicrophoneHelpSection from "./MicrophoneHelpSection";

export default {
  name: "MicrophoneHelpButton",
  components: {MicrophoneHelpSection},
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
  },
  methods: {
    toggleSection() {
      mainEmitter.$emit('toggle-help-menu')
    }
  }
}
</script>

<style scoped>

.micro-image {
  position: relative;
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