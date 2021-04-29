<template>
    <chicken-coop-page v-if="device.id === 'chicken-coop'" :device="device"/>
    <weather-station-page v-else-if="device.id === 'weather-station'" :device="device"/>
    <greenhouse-page v-else-if="device.id === 'greenhouse'" :device="device"/>
</template>

<script>
    import WeatherStationPage from "../Devices/WeatherStation/WeatherStationPage";
    import ChickenCoopPage from "../Devices/ChickenCoop/ChickenCoopPage";
    import GreenhousePage from "../Devices/Greenhouse/GreenhousePage";

    export default {
        name: "DevicePagePicker",
        props: { device: Object },
        components: {
            WeatherStationPage,
            ChickenCoopPage,
            GreenhousePage
        }
    }
</script>

<style scoped>

</style>