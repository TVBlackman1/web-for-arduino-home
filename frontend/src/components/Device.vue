<template>
    <div class="device">
        <header>{{ device.name }}</header>
        <router-link :to="getLink()">Перейти</router-link>
        <device-picker :device="device"/>
    </div>
</template>

<script>
    import DevicePicker from "./DevicePicker";

    export default {
        name: "Device",
        props: { device: Object },
        mounted() {
            // console.log("Created device:", this.device)
        },
        components: {
            DevicePicker
        },
        methods: {
            getLink() {
                return "/device/" + this.device.id
            }
        }
    }
</script>

<style scoped>
    header {
        font-weight: 700;
        font-size: 1.3em;
        cursor: pointer;
    }

    .device {
        width: 400px;
        height: 180px;
        padding: 0.4em;
        box-sizing: border-box;

        border-style: solid;
        border-color: #40b883;
        border-radius: 15px;
        border-width: 8px 1px 3px 1px;
    }

</style>