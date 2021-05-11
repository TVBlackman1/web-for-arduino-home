<template>
  <chicken-coop-m v-if="device.id === 'chicken-coop'" :device="device"/>
  <weather-station-m v-else-if="device.id === 'weather-station'" :device="device"/>
  <greenhouse-m v-else-if="device.id === 'greenhouse'" :device="device"/>
  <glasses-m v-else-if="device.id === 'glasses'" :device="device"/>
  <garden-m v-else-if="device.id === 'garden'" :device="device"/>

  <p v-else>No, 404</p>
</template>

<script>
import ChickenCoopM from "../Devices/ChickenCoop/ChickenCoopM";
import WeatherStationM from "../Devices/WeatherStation/WeatherStationM";
import GreenhouseM from "../Devices/Greenhouse/GreenhouseM";
import GlassesM from "../Devices/Glasses/GlassesM";
import GardenM from "../Devices/Garden/GardenM";

export default {
  name: "DevicePicker",
  props: {device: Object},
  components: {
    GardenM,
    GlassesM,
    ChickenCoopM,
    WeatherStationM,
    GreenhouseM,
  }
}
</script>

<style scoped>

</style>