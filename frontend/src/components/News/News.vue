<template>
  <div class="container">
    <div class="news" v-for="(_news, index) in news" :key="index">
      <div class="title">{{_news.title}}</div>
      <div class="content">{{_news.content}}</div>
    </div>
  </div>
</template>

<script>
import serverDefaultRequests from "../../mixins/serverDefaultRequests";

export default {
  name: "News",
  created() {
    this._getNews()
  },
  data() {
    return {
      news: []
    }
  },
  methods: {
    async _getNews() {
      const res = await this.getNews()
      console.log(res)
      this.news = res.news
    }
  },
  mixins: [serverDefaultRequests]
}
</script>

<style scoped>
.container {
  margin-top: 6em;
  width: 100%;
  display: flex;
  align-items: center;
  flex-direction: column;
}
.news {
  width: 800px;
  height: 400px;
  -webkit-box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);
  box-shadow: 2px 2px 6px 1px rgba(0,0,0,0.45);
  margin: 1em;
  /*display: inline-flex;*/
  /*justify-content: left;*/
  box-sizing: border-box;
  border-radius: 10px;
}

.news .title {
  border-radius: 10px 10px 0 0;
  font-size: 21px;
  padding: 0.3em;
  width: 100%;
  background-color: #8998d7;
  color: #f6f6f6;
  box-sizing: border-box;
}

.news .content {
  /*font-family: "Arial Narrow";*/
  text-align: justify;
  font-size: 16px;
  line-height: 1.6;
  /*padding: 1.6em;*/
  padding: 1.6em 2.2em;
  color: #2c3e50;
  width: 100%;
  box-sizing: border-box;
}

</style>