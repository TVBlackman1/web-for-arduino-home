<template>
  <div class="mode-section">
    <ul class="modes">
      <li class="mode" :class="{'picked' :picked_index === 1}" @click="onclick(1)">Нормальный</li>
      <li class="mode" :class="{'picked' :picked_index === 2}" @click="onclick(2)">Засуха</li>
      <li class="mode" :class="{'picked' :picked_index === 3}" @click="onclick(3)">Холода</li>
      <li class="mode" :class="{'picked' :picked_index === 4}" @click="onclick(4)">Свой режим 1</li>
      <li class="mode" :class="{'picked' :picked_index === 5}" @click="onclick(5)">Свой режим 2</li>
    </ul>
    <footer>
      <div class="add btn"></div>
      <div class="edit btn"></div>
    </footer>
  </div>
</template>

<script>
export default {
  name: "ModeSection",
  data() {
    return {
      picked_index: 1,
    }
  },
  methods: {
    onclick(ind) {
      this.picked_index = ind
    }
  }
}
</script>

<style scoped>
.mode-section {
  text-align: left;
  position: relative;
  height: 300px;
}

.modes {
  font-size: 20px;
  padding: 0.18em 0.3em;
}

.mode {
  padding: 0.2em 0.1em;
  transition: 0.24s;
}

.mode:hover {
  cursor: pointer;
  background-color: #c5f3df;
}

.picked {
  background-color: #7cefbe;
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