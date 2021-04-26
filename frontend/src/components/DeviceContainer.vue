<template>
    <div class="device-container">
        <device v-for="device in devices"
                :key="device.id"
                :device="device"/>
    </div>
</template>

<script>
    import Device from '../components/Device.vue'
    import serverHandler from "../mixins/serverHandler";

    export default {
        name: "DeviceContainer",
        data() {
            return {
                devices: []
            }
        },
        async created() {
            // POST request using fetch with async/await
            this.devices = await this.serverRequest("/api/devices")
        },
        components: {
            Device
        },
        mixins: [serverHandler],
    }
</script>

<style scoped>
    .device-container {
        width: fit-content;
        /*border: 1px solid black;*/
        margin: auto;

        display: grid;
        grid-template-columns: repeat(2, 1fr);
        grid-column-gap: 2em;
        grid-row-gap: 1.2em;
    }
</style>