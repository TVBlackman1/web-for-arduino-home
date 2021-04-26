<template>
<!--    <p>Some information about weather</p>-->
    <p>Температура воздуха: {{ device.additional.temperature }} °C</p>
    <p>Влажность: {{ device.additional.air_humidity }}%</p>
    <p>Давление: {{ device.additional.air_pressure }}мм рт ст</p>
</template>

<script>
    export default {
        name: "WeatherStationM",
        props: { device: Object },
    }
</script>

<style scoped>
    p {
        padding: 0;
        margin: 0.1em;
    }
</style>