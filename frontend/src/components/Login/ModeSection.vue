<template>
  <div class="mode-section">
    <ul class="modes">
      <li class="mode">Нормальный</li>
      <li class="mode">Засуха</li>
      <li class="mode">Холода</li>
      <li class="mode">Свой режим 1</li>
      <li class="mode">Свой режим 2</li>
    </ul>
    <footer>
      <div class="add btn"></div>
      <div class="edit btn"></div>
    </footer>
  </div>
</template>

<script>
export default {
  name: "ModeSection"
}
</script>

<style scoped>
.mode-section {
  text-align: left;
  position: relative;
  height: 300px;
}

.modes {
  font-size: 26px;
  padding: 0.1em 0.3em;
}

.mode {
  padding: 0.2em 0.1em;
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