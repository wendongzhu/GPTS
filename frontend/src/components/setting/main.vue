<template>
  <div class="tabs" :style="settingWindowSize">
    <el-tabs>
      <el-tab-pane label="测试信息"><DisplayParameter /></el-tab-pane>
      <el-tab-pane label="显示参数"><TestInformation /></el-tab-pane>
      <el-tab-pane label="系统参数"><SystemParameter /></el-tab-pane>
    </el-tabs>
  </div>

</template>

<script>
import DisplayParameter from "./func/display_parameter.vue";
import TestInformation from "./func/test_information.vue";
import SystemParameter from "./func/system_parameter.vue";

export default {
  components: { DisplayParameter, TestInformation, SystemParameter },
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
      this.settingWindowSize.width = window.innerWidth - 40 - 2 - 40 -2 - 40+ "px"
      this.settingWindowSize.height = window.innerHeight - 70 - 2 - 60 - 2 - 2 - 60 + "px"
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
.tabs {
  text-align: center;
  top: 0px;
  left: 0px;;
  position: relative;
  margin: 0;
}
.el-tabs{
  background-color: #2B2D30;
  border-radius: 4px;
  height: 36px;
  position: relative;
}
/* 去除灰色横条 */
::v-deep .el-tabs__nav-wrap::after {
  position: static !important;
}
/* 设置滑块颜色 */
::v-deep .el-tabs__active-bar{
  background-color: #538FFF !important;
  color: #eeeeee;
}
/* 设置滑块停止位置 */
::v-deep .el-tabs__active-bar.is-top{
  height: 36px;
  width: 150px ! important;
  border-radius: 4px;
  top: 0px !important;
  left: -20px !important;
  position: absolute !important;
  z-index: 1;
}
::v-deep .tab-testPaper{
  background-color: #fff;      /*设置修改默认蚊子颜色，背景颜色，等*/
  font-weight: 550;
  font-size: 20px;
  opacity: 1;
  box-shadow: 0 0 0 0;
}

/* 设置当前选中样式 */
::v-deep .el-tabs__item.is-active{
  color:#ffffff !important;
  z-index: 2;
}

/* 设置未被选中样式 */
::v-deep .el-tabs__item{
  color:#ffffff !important;
  font-size: 16px;
  padding: 0 20px !important;
  width: 150px;
  box-sizing: border-box;
  display: inline-block;
  position: relative !important;
  z-index: 2;
}


</style>