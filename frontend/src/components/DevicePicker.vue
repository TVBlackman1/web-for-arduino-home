<template>
  <chicken-coop-m v-if="device.id === 'chicken-coop'" :device="device"/>
  <weather-station-m v-else-if="device.id === 'weather-station'" :device="device"/>
  <greenhouse-m v-else-if="device.id === 'greenhouse'" :device="device"/>
  <glasses-m v-else-if="device.id === 'glasses'" :device="device"/>

  <p v-else>No, 404</p>
</template>

<script>
import ChickenCoopM from "./ChickenCoopM";
import WeatherStationM from "./WeatherStationM";
import GreenhouseM from "./GreenhouseM";
import GlassesM from "./GlassesM";

export default {
  name: "DevicePicker",
  props: {device: Object},
  components: {
    GlassesM,
    ChickenCoopM,
    WeatherStationM,
    GreenhouseM,
  }
}
</script>

<style scoped>

</style>