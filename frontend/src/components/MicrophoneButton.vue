<template>
<!--  <div class="ghost-bottom-panel">-->
  <div class="micro-image bottom-fixed" :style="styleObject" @click="startListening">
    <div class="indicator" :style="styleIndicator"></div>
  </div>
<!--  </div>-->
</template>

<script>
import MicroSVG from "../assets/svg/micro-24.svg"
import MicroSVGEnabled from "../assets/svg/micro-enabled-24.svg"
import voiceHandler from "../mixins/voiceHandler";

export default {
  name: "MicrophoneButton",
  data() {
    return {
      MicroSVG,
      MicroSVGEnabled,
      isListening: false
    }
  },
  computed: {
    styleObject() {
      return {
        '--background-image': `url(${this.MicroSVG})`,
        '--background-image-hover': `url(${this.MicroSVGEnabled})`
      }
    },
    styleIndicator() {
      return {
        '--indicatorColor': this.isListening ? "green" : "red"
      }
    }
  },
  created() {
    this.initVoiceHandle(this.onEndListening)
  },
  methods: {
    onEndListening() {
      this.isListening = false
    },
    startListening(e) {
      this.isListening = true
      this.listenStart(e)
    }
  },
  mixins: [voiceHandler]
}
</script>

<style scoped>
.micro-image {
  background-image: var(--background-image);
  position: relative;
  /*var(--background-image-hover);*/
  width: 66px;
  height: 66px;
  background-size: cover;
  margin: 0 0.6em;
}
.micro-image:hover {
  cursor: pointer;
  background-image: var(--background-image-hover);
}

.indicator {
  position: absolute;
  top: 0;
  right: 0;
  width: 8px;
  height: 8px;
  border-radius: 100%;
  background-color: var(--indicatorColor);
}

</style>