<template>
<div>
  <el-row :gutter="0">
      <!--Footer-->
      <div class="grid-content bg-purple-dark" :style="defaultFooterWidth">
<!--        预留-->
      </div>
      <!--时间-->
      <div class="footer-time" >
        <font size="2" font-weight="bold">{{ getTime }}</font>
      </div>
</el-row>
</div>
</template>

<script>
export default {
  name: "footer",
  data() {
    return {
      getTime:'',//格式化之后的当前时间
      defaultFooterWidth: {
        width: "",
      },
    };
  },
  created() {
    this.getTimes()
    window.addEventListener("resize", this.getFooterWidth);
    this.getFooterWidth();
  },
  // 在Vue实例销毁前，清除我们的定时器
  beforeDestroy() {
    if (this.getTime) {
      clearInterval(this.getTimesInterval);
    }
  },
  methods: {
    getTimes(){
      setInterval(this.getTimesInterval, 1000);
    },
    getTimesInterval:function() {
      let _this = this;
      let year = new Date().getFullYear(); //获取当前时间的年份
      let month = new Date().getMonth() + 1; //获取当前时间的月份
      let day = new Date().getDate(); //获取当前时间的天数
      let hours = new Date().getHours(); //获取当前时间的小时
      let minutes = new Date().getMinutes(); //获取当前时间的分数
      let seconds = new Date().getSeconds(); //获取当前时间的秒数
      //当小于 10 的是时候，在前面加 0
      if (month < 10) {
        month = "0" + month;
      }
      if (day < 10) {
        day = "0" + day;
      }
      if (hours < 10) {
        hours = "0" + hours;
      }
      if (minutes < 10) {
        minutes = "0" + minutes;
      }
      if (seconds < 10) {
        seconds = "0" + seconds;
      }
      //拼接格式化当前时间
      _this.getTime = year + "-" + month + "-" + day + " " + hours + ":" + minutes + ":" + seconds;
    },
    getFooterWidth() {
      this.defaultFooterWidth.width = window.innerWidth - 180 + "px";
    },
  }
}

</script>

<style scoped>
.footer-time{
  padding: 4px 0 4px 0;
}
</style>