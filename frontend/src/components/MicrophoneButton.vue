<template>
  <div class="ghost-bottom-panel">
    <div class="micro-image bottom-fixed" :style="styleObject" />
  </div>
</template>

<script>
import MicroSVG from "../assets/svg/micro-24.svg"
import MicroSVGEnabled from "../assets/svg/micro-enabled-24.svg"

export default {
  name: "MicrophoneButton",
  data() {
    return {
      MicroSVG,
      MicroSVGEnabled,
    }
  },
  computed: {
    styleObject() {
      return {
        '--width': '84px',
        '--height': '84px',
        '--background-image': `url(${this.MicroSVG})`,
        '--background-image-hover': `url(${this.MicroSVGEnabled})`
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

.ghost-bottom-panel {
  position: fixed;
  bottom: 0;
  width: 100%;
  height: 100px;
  display: flex;
  justify-content: center;
  align-items: center;
  padding-bottom: 10px;

}
</style>