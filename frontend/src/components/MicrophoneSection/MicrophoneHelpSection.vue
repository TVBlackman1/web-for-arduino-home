<template>
  <div class="help-section">
    <header>Помощь</header>
    <ul>
      <li class="point">Курятник</li>
      <li class="point">Метеостанция</li>
      <li class="point">Теплица</li>
      <li class="point">Новости</li>
      <li class="point">Помощь</li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "MicrophoneHelpSection"
}
</script>

<style scoped>
.help-section {
  width: 460px;
  height: 220px;
  overflow-y: auto;
  background-color: white;
  border-radius: 8px;
  border-top-width: 10px;
  border-bottom-width: 4px;
  border-color: #40b883;
  border-style: solid;
  -webkit-box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);
  box-shadow: 2px 2px 10px 1px rgba(0,0,0,0.45);
  padding: 1em 0.4em;
}

header {
  margin-bottom: 0.4em;
}

.point {
  font-size: 20px;
  text-align: left;
  transition: 0.15s;
  /*border-bottom: 1px solid black;*/
}

.point:before {
  content: "✿ «";
  /*padding-right: 7px;*/
  color: #767676;
}

.point:after {
  content: "»";
}

.point:hover {
  cursor: pointer;
  color: #2c3e50;
  font-size: 22px;
}

.point:hover:before {
  color: #2c3e50;
}

</style>