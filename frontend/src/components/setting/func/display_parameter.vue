<template>
<div class="display-parameter" :style="settingWindowSize">
<!--  测试信息-->
  <el-row>
    <el-form :model="testInformation" label-width="120px">
      <el-form-item label="测试人员">
        <el-input v-model="testInformation.testUser" />
      </el-form-item>
      <el-form-item label="测试地点">
        <el-input v-model="testInformation.testAddress" />
      </el-form-item>
      <el-form-item label="测试单位">
        <el-input v-model="testInformation.testCompany" />
      </el-form-item>
    </el-form>
  </el-row>
<!--  伽马探管设置-->
  <el-row>
    <el-form :model="testInformation" label-width="120px">
      <el-form-item label="探管设置" >
        <el-input v-model="testInformation.testDescription"  input-style="width: 800px; height: 250px" />
      </el-form-item>
    </el-form>
  </el-row>
<!--  测试概要-->
  <el-row>
    <el-form :model="testInformation" label-width="120px">
    <el-form-item label="测试概要" >
      <el-input v-model="testInformation.testDescription"  input-style="width: 800px; height: 250px" />
    </el-form-item>
    </el-form>
  </el-row>
</div>
</template>

<script>
export default {
  name: "display_parameter",
  data() {
    return {
      settingWindowSize: {
        width: "",
        height: "",
      },
      testInformation: {
        testUser: "GP",
        testAddress: "",
        testCompany: "",
        testDescription: "",
      },
      gpInformation: {

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
.display-parameter{
  color: #eeeeee;
  background: #2B2D30;
  border-radius: 8px;
}
.el-form{
  color: #eeeeee;
}
</style>