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
import Device from '../components/Device.vue'
import serverHandler from "../mixins/serverHandler";
import DeviceEmptyCard from "../components/DeviceEmptyCard";

export default {
  name: "DeviceContainer",
  data() {
    return {
      devices: []
    }
  },
  async created() {
    // POST request using fetch with async/await
    this.devices = await this.serverRequest("/api/devices")
  },
  components: {
    Device,
    DeviceEmptyCard
  },
  mixins: [serverHandler],
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