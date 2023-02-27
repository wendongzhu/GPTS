<template>
  <el-container>
    <!--Collection Header-->
    <el-header class="collection-index-header" :style="collectionHeaderSize">
      <CollectionHeader />
    </el-header>
    <el-container>
      <el-container>
        <!--Collection Main-->
        <el-main class="collection-index-main" :style="collectionMainSize">
          <CollectionMain />
        </el-main>
        <!--Collection Footer-->
        <el-footer class="collection-index-footer" :style="collectionFooterSize">
          <CollectionFooter />
        </el-footer>
      </el-container>
      <!--Collection Right Aside-->
      <el-aside class="collection-index-right-aside" :style="collectionRightAsideSize">
        <CollectionRightAside />
      </el-aside>
    </el-container>

  </el-container>
</template>

<script>
import CollectionHeader from "./header.vue"
import CollectionMain from  "./main.vue"
import CollectionFooter from "./footer.vue"
import CollectionRightAside from "./right_aside.vue"

export default {
  name: "index",
  components: { CollectionHeader, CollectionMain, CollectionFooter, CollectionRightAside },
  data() {
    return {
      collectionHeaderSize: {
        width: "",
      },
      collectionFooterSize: {
        width: "",
      },
      collectionMainSize: {
        width: "",
        height: "",
      },
      collectionRightAsideSize: {
        width: "",
        height: "",
      },
    };
  },
  methods: {
    //定义方法，获取高度减去头尾
    getWindowSize() {
      this.collectionHeaderSize.width = window.innerWidth - 40 - 2 -40 -2+ "px"
      this.collectionFooterSize.width = window.innerWidth - 40 - 2 -40 - 2 - 180 - 2+ "px"
      this.collectionMainSize.width = window.innerWidth - 40 - 2 -40 - 2 - 180 -2 + "px"
      this.collectionMainSize.height = window.innerHeight - 70 - 2 -60 - 250 - 2 + "px"
      this.collectionRightAsideSize.height = window.innerHeight - 70 - 2 -60  - 2  -2 + "px"
      console.log("main-size:", this.collectionMainSize)
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
.collection-index-header{
  padding: 0;
  background: #2B2D30;
  height: 60px;
}

.collection-index-footer{
  padding: 0;
  background: #2B2D30;
  height: 250px;
}
.collection-index-main{
  padding: 1px 1px 1px 1px;
  background: #1E1F22;
}
.collection-index-right-aside{
  width: 180px;
  margin-top: 2px;
  margin-right: 2px;
  background: #2B2D30;

}


</style>