<template>
  <el-container>
    <!--Setting Header-->
    <el-header class="setting-index-header" :style="settingHeader">
        <el-button-group class="setting-header-button" style="float:right">
          <el-button icon="RefreshRight"  color="#2B2D30">默认恢复</el-button>
          <el-button icon="Check" bg color="#2B2D30" >保存参数</el-button>
          <el-button icon="Close" color="#2B2D30">放弃退出</el-button>
        </el-button-group>

    </el-header>
    <!--Setting Main-->
    <el-main class="setting-index-main" :style="settingMain">
      <el-tabs>
        <el-tab-pane label="测试信息">
          <!--  测试信息-->
          <div class="system-main" :style="displayWindowSize">
            <el-row class="display-param-row" >
            <el-form :model="DisplayParam" label-width="120px" style="width: 35%">
              <el-form-item label="测试人员:">
                <el-input v-model="DisplayParam.TestUser" />
              </el-form-item>
              <el-form-item label="测试地点:">
                <el-input v-model="DisplayParam.TestAddress" />
              </el-form-item>
              <el-form-item label="测试单位:">
                <el-input v-model="DisplayParam.TestCompany" />
              </el-form-item>
            </el-form>
          </el-row>

            <!--  伽马探管设置-->
            <el-row>
              <el-form :model="DisplayParam"  label-width="120px">
                <el-form-item label="探管设置:" >
                  <el-table
                      class="collection-footer-table"
                      :data="DisplayParam.ProbePipeSetlList"
                      border
                      :row-style="{height:'20px'}"
                      :cell-style="{padding:'0px'}"
                  >
                    <el-table-column prop="ChannelName" label="通道名称" width="120" />
                    <el-table-column prop="PowerStatus" label="供电开关状态设置" width="150" />
                    <el-table-column prop="VoltageThreshold" label="电压阈值设置(V)" width="150"/>
                    <el-table-column prop="CurrentThreshold" label="电流阈值设置(mA)" width="150" />
                    <el-table-column prop="PipeNumber" label="探管编号" width="120" />
                  </el-table>
                </el-form-item>
              </el-form>
            </el-row>
            <!--  测试概要-->
            <el-row>
              <el-form :model="DisplayParam" label-width="120px">
                <el-form-item label="测试概要:" >
                  <el-input
                      v-model="DisplayParam.TestDescription"
                      :rows="2"
                      type="textarea"
                      placeholder="Please input"
                      input-style="width: 690px; height: 180px"
                  />
                </el-form-item>
              </el-form>
            </el-row>
          </div>

        </el-tab-pane>
        <el-tab-pane label="显示参数">
          <div class="system-main" :style="displayWindowSize"></div>
        </el-tab-pane>
        <el-tab-pane label="系统参数">
          <div class="system-main" :style="displayWindowSize">
            <el-row class="system_param-row">
              <el-form :model="SystemParam" label-width="120px" style="width: 35%">
                <el-form-item label="通讯端口">
                  <el-input v-model="SystemParam.PortName" />
                </el-form-item>
                <el-form-item label="存储间隔">
                  <el-input v-model="SystemParam.TimeInterval" >
                  <!--TODO: 单位添加-->
                  </el-input>
                </el-form-item>
                <el-form-item label="文件前缀">
                  <el-input v-model="SystemParam.FilePrefix" />
                </el-form-item>
                <el-form-item label="工作路径">
                  <el-input v-model="SystemParam.WorkPath"  >
                    <template #append>
                      <el-button slot="append" icon="Folder" type="success"  />
                    </template>
                  </el-input>
                </el-form-item>
              </el-form>
            </el-row>
          </div>
        </el-tab-pane>
      </el-tabs>
    </el-main>
  </el-container>
</template>

<script>
import { DisplayAPI, SystemAPI, TestInfoAPI } from "../../../wailsjs/go/endinterface/Backend"

export default {
  name: "index",
  data() {
    return {
      settingHeader: {
        width: "",
      },
      settingMain:{
        height: "",
        width: "",
      },
      displayWindowSize:{
        height: "",
        width: "",
      },
      DisplayParam: {
        TestUser: "",
        TestAddress: "",
        TestCompany: "",
        ProbePipeSetlList: [],
        TestDescription: "",
      },
      TestInfo: {

      },
      SystemParam: {
        PortName: "",
        TimeInterval: "",
        FilePrefix: "",
        WorkPath: "",
      }
    };
  },
  mounted() {
    this.DisplayDataFunc();
    this.TestInfoDataFunc();
    this.SystemDataFunc()
  },
  methods: {
    //定义方法，获取高度减去头尾
    getWindowSize() {
      this.settingHeader.width = window.innerWidth - 40 - 2 - 40 - 2 + "px"
      this.settingHeader.width = window.innerWidth - 40 - 2 - 40 - 2 + "px"
      this.settingMain.height = window.innerHeight - 70 -60 - 4- 4 + "px";
      this.displayWindowSize.width = window.innerWidth - 40 - 2 - 40 -2 - 40 + "px"
      this.displayWindowSize.height = window.innerHeight - 70 - 2 - 60 - 2 - 2 - 90 + "px"

    },
    DisplayDataFunc() {
      DisplayAPI().then(result => {
        this.DisplayParam = result
      })
    },
    TestInfoDataFunc() {
      TestInfoAPI().then(result => {
        this.TestInfo= result
      })
    },
    SystemDataFunc() {
      SystemAPI().then(result => {
        this.SystemParam = result
      })
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
.setting-index-header{
  padding: 0;
  background: #2B2D30;
  height: 60px;
}
.setting-header-button{
  width: 350px;
  margin-right: 20px;
  padding: 15px 0 15px 0;
}

.display-param-row{
  padding-top: 20px;
}

.system-main{
  color: #eeeeee;
  background: #2B2D30;
  border-radius: 8px;
}
.system_param-row{
  padding-top: 30px;
}


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