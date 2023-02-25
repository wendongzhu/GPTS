
<template>
  <div id="header">
    <el-row :gutter="0">
      <div class="header-logo-avatar" >
        <el-avatar :size="20" shape="square" fit="fill">
          <img src="../../assets/logo/logo-weld.ico" />
        </el-avatar>
      </div>
      <div class="header-title" :style="defaultTittleWidth">
        <font size="3.5" font-weight="bold">伽马探管测试平台</font>
      </div>
      <div class="header-button">
        <el-button-group>
          <el-button
              style="width: 40px; height: 40px; margin: 0; padding: 0;border-radius: 0"
              color="#2B2D30"
              icon="Minus"
              @click="Minimise" />

          <el-button
              style="width: 40px; height: 40px; margin: 0; padding: 0;border-radius: 0"
              color="#2B2D30"
              icon="FullScreen"
              @click="Fullscreen"/>

          <el-button
              style="width: 40px; height: 40px; margin: 0; padding: 0;border-radius: 0"
              color="#2B2D30"
              icon="Close"
              @click="Quit"/>
        </el-button-group>

      </div>
    </el-row>
  </div>
</template>

<script>
import {WindowFullscreen,WindowMinimise,WindowQuit} from '../../../wailsjs/go/main/App';

export default {
  name: "header",
  data() {
    return {
      index: 1,
      defaultTittleWidth: {
        width: "",
      },
    };
  },
  methods: {
    //定义方法，获取高度减去头尾
    getTittleHeight() {
      this.defaultTittleWidth.width = window.innerWidth - 160+ "px";
    },

    Fullscreen(){
      WindowFullscreen();
      this.getTittleHeight();
    },
    Minimise(){
      WindowMinimise();
    },
    Quit(){
      WindowQuit();
    },
  },

  created() {
    //页面创建时执行一次getHeight进行赋值，顺道绑定resize事件
    window.addEventListener("resize", this.getTittleHeight);
    this.getTittleHeight();
  },
};

</script>


<style lang="scss" scoped>
.header-logo-avatar {
  padding: 10px 10px 10px 10px;
}
.header-button{
  width: 120px;
  margin: 0;
  padding-left: 0;
}


.header-title{
  padding-top: 9px;
  padding-bottom: 9px;
}
</style>
