<template>
  <div class="device-container">
    <device v-for="device in devices"
            :key="device.id"
            :device="device"/>
    <!--      <device-->
    <device-empty-card />
  </div>
</template>

<script>
import Device from './Device.vue'
import serverDefaultRequests from "../../mixins/serverDefaultRequests";
import DeviceEmptyCard from "./DeviceEmptyCard";

export default {
  name: "DeviceContainer",
  data() {
    return {
      devices: []
    }
  },
  async created() {
    this.devices = await this.getDevices()
  },
  components: {
    Device,
    DeviceEmptyCard
  },
  mixins: [serverDefaultRequests],
}
</script>

<style scoped>
.device-container {
  width: fit-content;
  /*border: 1px solid black;*/
  margin: auto;

  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-column-gap: 2em;
  grid-row-gap: 1.2em;
}
</style>