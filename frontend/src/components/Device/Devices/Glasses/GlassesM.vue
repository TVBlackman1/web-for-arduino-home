<template>
  <p>Очки: {{ getEnableStatus() }}</p>
</template>

<script>
import imageGlasses from "../../../../assets/glasses.jpg"

export default {
  name: "GlassesM",
  props: { device: Object },
  data() {
    return {
      imgSrc: imageGlasses
    }
  },
  methods: {
    getEnableStatus() {
      if(this.device.additional.enable_statis)
        return "подключены"
      return "отключены"
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