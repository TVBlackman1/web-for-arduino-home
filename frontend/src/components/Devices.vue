<template>
    <h1>Ваши устройства:</h1>
    <device-container />
</template>

<script>
    import DeviceContainer from '../components/DeviceContainer.vue'

    export default {
        name: "Devices",
        components: {
            DeviceContainer
        }
    }
</script>

<style scoped>
  h1 {
    margin-top: 3em;
  }
</style>