<template>
    <button class="Btn" @click="handleclick">{{setText()}}</button>
</template>

<script>
    import imageChickenCoop from "../../../../assets/chicken-coop.jpg"

    export default {
        name: "ChickenCoopM",
        props: { device: Object },
        data() {
            return {
                imgSrc: imageChickenCoop,
                value22: false,
            }
        },
        methods: {
            setText() {
                return this.value22 ? "Выключить рекуператор" : "Включить рекуператор"
            },
            getAutoWateringValue() {
                if(this.device.additional.auto_watering)
                    return "включена"
                return "выключена"
            },
            handleclick($event) {
                $event.stopPropagation()
                this.value22 = !this.value22
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

    .Btn {
        margin: 0.2em;
        font-size: 1.2em;
        padding: 0.3em 0.7em;
        background-color: #40b883;
        color: white;
        border-radius: 20px;
        border: 0;
        transition: 0.24s;
    }

    .Btn:hover {
        cursor: pointer;
        background-color: #b8ac40;
    }
</style>