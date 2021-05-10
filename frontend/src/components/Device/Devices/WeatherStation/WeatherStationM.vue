<template>
  <!--    <p>Some information about weather</p>-->
  <p>Температура воздуха: {{ device.additional.temperature }} °C</p>
  <p>Влажность: {{ device.additional.humidity }}%</p>
</template>

<script>
import imageWeatherStation from "../../../../assets/weather-station.jpg"

export default {
  name: "WeatherStationM",
  props: {device: Object},
  data() {
    return {
      imgSrc: imageWeatherStation
    }
  },
  mounted() {
    // console.log(this.imgSrc)
    // this.imgSrc = "edited"
    this.setDeviceImageSrc(this.imgSrc)
    // console.log(this.imgSrc)
  },
  inject: ['setDeviceImageSrc']
}
</script>

<style scoped>
p {
  padding: 0;
  margin: 0.1em;
}
</style>