<template>
  <div class="card-device">
    <img :src="imgSrc" class="card-add-device-picture">
  </div>
</template>

<script>
import imageAdd from "../../assets/add2.png"

export default {
  name: "DeviceEmptyCard",
  data() {
    return {
      imgSrc: imageAdd
    }
  }
}
</script>

<style scoped>
@import "../../css/device-card.css";

.card-device {
  /*add features*/
  overflow: visible;
  display: flex;
  justify-content: center;
  align-items: center;
}

.card-add-device-picture {
  max-width: 100%;
  height: 60%;
  /*max-height: 140px;*/
  background-position: center;
  background-size: cover;
  background-origin: content-box;
  /*background-repeat: repeat;*/
  object-fit: cover;
  object-position: 10% 20%;

  transition: 0.24s;
}

.card-device:hover .card-add-device-picture{
  height: 50%;
  transition: 0.3s;
  /*height: 60%;*/
  /*transition: 0.4s;*/
}

</style>