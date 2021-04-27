<template>
    <p>Автопоилка: {{ getAutoWateringValue() }}</p>
</template>

<script>
    import imageChickenCoop from "../assets/chicken-coop.jpg"

    export default {
        name: "ChickenCoopM",
        props: { device: Object },
        data() {
            return {
                imgSrc: imageChickenCoop
            }
        },
        methods: {
            getAutoWateringValue() {
                if(this.device.additional.auto_watering)
                    return "включена"
                return "выключена"
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