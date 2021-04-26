<template>
    <p>Температура: {{ device.additional.temperature}} °C</p>
    <p>Влажность почвы: {{ device.additional.soil_moisture}}%</p>
    <p>Влажность воздуха: {{ device.additional.air_humidity}}%</p>

</template>

<script>
    export default {
        name: "GreenhouseM",
        props: { device: Object },
    }
</script>

<style scoped>
    p {
        padding: 0;
        margin: 0.1em;
    }
</style>