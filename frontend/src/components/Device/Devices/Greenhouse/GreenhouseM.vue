<template>
    <p>Температура: {{ device.additional.temperature}} °C</p>
    <p>Влажность почвы: {{ device.additional.soil_moisture}}%</p>
    <p>Влажность воздуха: {{ device.additional.air_humidity}}%</p>

</template>

<script>
    import imageGreenHouse from "../../../../assets/green-house.jpg";

    export default {
        name: "GreenhouseM",
        props: { device: Object },
        data() {
            return {
                imgSrc: imageGreenHouse
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