<template>
  <div class="friends">
<!--    <div class="title">Настроить доступ</div>-->
    <ul class="scrolling">
      <li class="friend">SilaS</li>
      <li class="friend">zhigulyaka</li>
      <li class="friend">Челебит</li>
      <li class="friend">Челебит</li>
      <li class="friend">Челебит</li>
      <li class="friend">Челебит</li>
      <li class="friend">Челебит</li>
      <li class="friend">Челебит</li>
      <li class="friend">Челебит</li>
      <li class="friend">Челебит</li>
      <li class="friend">Челебит</li>
      <li class="friend">Челебит</li>
      <li class="friend">Челебит</li>
      <li class="friend">Челебит</li>
    </ul>
    <footer>
      <div class="add btn"></div>
      <div class="edit btn"></div>
    </footer>
  </div>
</template>

<script>
export default {
  name: "friend-list"
}
</script>

<style scoped>

.friends {
  width: 100%;
  height: 300px;
  /*background-color: #7692ff;*/
  background-color: white;
  padding: 0.8em 0.6em;
  box-sizing: border-box;
  position: relative;
  border-radius: 4px;
}

.scrolling {
  padding: 0.2em;
  overflow-y: auto;
  max-height: 250px;
}

.friend {
  font-size: 22px;
  border-radius: 4px;
  padding: 0.3em 0.6em;
  background-color: #eaeaea;
  margin: 0.4em 0;
  transition: 0.24s;
}

.friend:hover {
  background-color: #d0d7f8;
  cursor: pointer;
}

.title {
  font-size: 24px;
}

footer {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 0.2em;
  background-color: white;
  height: 40px;
  display: flex;
  flex-direction: row;
  justify-content: space-around;
}

.btn {
  display: block;
  background-color: white;
  background-repeat: no-repeat;
  background-size: contain;
  background-position: center;
  height: 40px;
  width: 60px;
  transition: 0.24s;
  -webkit-box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);
  box-shadow: 1px 1px 5px 1px rgba(0,0,0,0.45);
  border-radius: 6px;
}

.btn:hover {
  background-color: #dddfe7;
  cursor: pointer;
}

.add {
  background-image: url("../../assets/svg/add_black_24dp.svg");
}

.edit {
  background-image: url("../../assets/svg/edit_black_24dp.svg");
}

</style>