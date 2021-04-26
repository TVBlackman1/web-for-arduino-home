<template>
    <p>Автопоилка: {{ getAutoWateringValue() }}</p>
</template>

<script>
    export default {
        name: "ChickenCoopM",
        props: { device: Object },
        methods: {
            getAutoWateringValue() {
                if(this.device.additional.auto_watering)
                    return "включена"
                return "выключена"
            }
        }
    }
</script>

<style scoped>
    p {
        padding: 0;
        margin: 0.1em;
    }
</style>