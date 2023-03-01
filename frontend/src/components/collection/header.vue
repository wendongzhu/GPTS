<template>
  <div>
      <el-button-group   class="collection-header-button" style="float:left">
        <el-button icon="VideoPlay"  color="#2B2D30">开始采集</el-button>
        <el-button icon="VideoPause" color="#2B2D30">停止采集</el-button>
        <el-button icon="Document" color="#2B2D30">记录数据</el-button>
      </el-button-group>

<!--    TODO:  测试信息太长情况是需要截取部分-->
      <el-button-group   class="collection-header-text" style="float:right">
        <font size="3" style="padding-left: 30px">测试人员：{{ testInformation.testUser }}</font>
        <font size="3" style="padding-left: 30px">测试地点：{{ testInformation.testAddress }}</font>
        <font size="3" style="padding-left: 30px">测试单位：{{ testInformation.testCompany }}</font>
      </el-button-group>
  </div>
</template>

<script>
export default {
  name: "header",
  data() {
    return {
      defaultHeaderWidth: {
        width: "",
      },
      testInformation: {
        testUser: "朱文东",
        testAddress: "北京市通州区张家湾镇方和正圆工业园",
        testCompany: "北京实耐固连接技术有限公司",
      }
    };
  },


  methods: {
    getWindowSize() {
      this.defaultHeaderWidth.width = window.innerWidth - 40 - 2 - 40 - 2 + "px";
    },
  },
  created() {
    window.addEventListener("resize", this.getWindowSize);
    this.getWindowSize();
  },
}
</script>

<style scoped>
.collection-header-button{
  width: 350px;
  margin-left: 30px;
  padding: 15px 0 15px 0;
}
.collection-header-text{
  padding: 20px 20px 20px 20px;
}

</style>