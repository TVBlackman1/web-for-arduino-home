<template>
  <header>
    <ul>
      <li>Моя дача</li>
      <li>Новости</li>
      <li>Мой аккаунт</li>
    </ul>
  </header>
</template>

<script>
export default {
  name: "SiteHeader"
}
</script>

<style scoped>
  header {
    position: fixed;
    top: 0;
    width: 100%;
    height: 60px;
    background-color: #5f77de;
    display: flex;
    /*justify-content: center;*/
    align-items: center;
  }

  header ul {
    margin: 0;
    display: flex;
    flex-direction: row;
    color: #dddfe7;
    font-weight: 500;
  }

  header li {
    margin-left: 3.1em;
  }

  header li:nth-child(1) {
    margin-left: 1.2em;
  }

  header li:hover {
    color: #ffffff;
    text-decoration: underline;
  }
</style>