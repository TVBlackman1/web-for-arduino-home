<template>
  <div class="ghost-bottom-panel">
    <div class="mback">
      <microphone-help-button />
      <microphone-button />
      <microphone-settings-button />
    </div>
  </div>
</template>

<script>
import MicrophoneButton from "./MicrophoneButton";
import MicrophoneSettingsButton from "./MicrophoneSettingsButton";
import MicrophoneHelpButton from "./MicrophoneHelpButton";

export default {
  name: "MicrophoneSection",
  components: {
    MicrophoneButton,
    MicrophoneSettingsButton,
    MicrophoneHelpButton
  }
}
</script>

<style scoped>

.ghost-bottom-panel {
  position: fixed;
  bottom: 0;
  width: 100%;
  height: 100px;
  display: flex;
  justify-content: center;
  align-items: center;
  padding-bottom: 26px;
}

.mback {
  border-radius: 50px;
  width: calc(96px + 140px);
  height: 90px;
  background-color: #ffffff;
  -webkit-box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);
  box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);
  display: flex;
  /*flex-grow: 1;*/
  align-content: space-around;
  justify-content: center;
  align-items: center;
}

</style>