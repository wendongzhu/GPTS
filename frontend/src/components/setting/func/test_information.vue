<template>
  <div class="test_information" :style="settingWindowSize">
    test_information
  </div>
</template>

<script>
export default {
  name: "test_information",
  data() {
    return {
      settingWindowSize: {
        width: "",
        height: "",
      }
    }
  },
  methods: {
    //定义方法，获取高度减去头尾
    getWindowSize() {
      this.settingWindowSize.width = window.innerWidth - 40 - 2 - 40 -2 - 40 + "px"
      this.settingWindowSize.height = window.innerHeight - 70 - 2 - 60 - 2 - 2 - 90 + "px"
    },
  },

  created() {
    //页面创建时执行一次getHeight进行赋值，顺道绑定resize事件
    window.addEventListener("resize", this.getWindowSize);
    this.getWindowSize();
  },
}
</script>

<style scoped>
.test_information{
  color: #eeeeee;
  background: #2B2D30;
  border-radius: 8px;
}
</style>