<template>
    <div class="device-page">
        <h1>
            {{ title }}
        </h1>
<!--        <p>{{ description }}</p>-->
        <p v-if="!reqIsReady">Ожидание сервера..</p>
        <device-page-picker :device="device"/>
    </div>
</template>

<script>
    import serverDefaultRequests from "../../mixins/serverDefaultRequests";
    import DevicePagePicker from "./Picker/DevicePagePicker";

    export default {
        name: "DevicePage",
        data() {
            return {
                title: "None",
                description: "None",
                device: {}
            }
        },
        async created() {
            const data = await this.getDeviceById(this.$route.params.id)
            this.title = data.name
            this.description = data.description
            this.device = data
        },
        components: {
            DevicePagePicker
        },
        mixins: [serverDefaultRequests],
    }


</script>

<style scoped>
    .device-page {
        width: 800px;
        margin: auto;
        padding: 1.4em;
        box-sizing: border-box;
        border-style: solid;
        border-width: 0 1.5em;
        border-color: #40b883;
        border-radius: 10px;
    }
</style>