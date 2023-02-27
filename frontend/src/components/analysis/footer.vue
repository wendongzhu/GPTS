<template>
  <div>
    <el-table class="collection-footer-table" :data="collectionFooterTableData" border >
      <el-table-column prop="名称" label="名称" auto />
      <el-table-column prop="第1路" label="第1路" auto />
      <el-table-column prop="第2路" label="第2路" auto/>
      <el-table-column prop="第3路" label="第3路" auto />
      <el-table-column prop="第4路" label="第4路" auto />
      <el-table-column prop="第5路" label="第5路" auto />
      <el-table-column prop="第6路" label="第6路" auto />
      <el-table-column prop="第7路" label="第7路" auto />
      <el-table-column prop="第8路" label="第8路" auto/>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "header",
  data() {
    return {
      collectionFooterTableData: [
        {
          名称: '供电电压(V)',
          第1路: '22.04',
          第2路: '24.03',
          第3路: '23.01',
          第4路: '23.02',
          第5路: '22.08',
          第6路: '24.01',
          第7路: '23.05',
          第8路: '23.6',
        },
        {
          名称: '供电电流(mA)',
          第1路: '22.04',
          第2路: '24.03',
          第3路: '23.01',
          第4路: '23.02',
          第5路: '22.08',
          第6路: '24.01',
          第7路: '23.05',
          第8路: '23.6',
        },{
          名称: '计数',
          第1路: '22.04',
          第2路: '24.03',
          第3路: '23.01',
          第4路: '23.02',
          第5路: '22.08',
          第6路: '24.01',
          第7路: '23.05',
          第8路: '23.6',
        },{
          名称: '计数率',
          第1路: '22.04',
          第2路: '24.03',
          第3路: '23.01',
          第4路: '23.02',
          第5路: '22.08',
          第6路: '24.01',
          第7路: '23.05',
          第8路: '23.6',
        },{
          名称: '供电开关',
          第1路: 'NO',
          第2路: 'NO',
          第3路: 'NO',
          第4路: 'NO',
          第5路: 'NO',
          第6路: 'NO',
          第7路: 'NO',
          第8路: 'NO',
        },{
          名称: '电压状态',
          第1路: '',
          第2路: '',
          第3路: '',
          第4路: '',
          第5路: '',
          第6路: '',
          第7路: '',
          第8路: '',
        },{
          名称: '电流状态',
          第1路: '',
          第2路: '',
          第3路: '',
          第4路: '',
          第5路: '',
          第6路: '',
          第7路: '',
          第8路: '',
        },

      ]
    };
  },
}
</script>

<style scoped>
/*背景透明*/
::v-deep .el-table__expanded-cell {
  background-color: transparent !important;
}

::v-deep .el-table th,
::v-deep .el-table tr,
::v-deep .el-table td {
  background-color: transparent;
}
.collection-footer-table{
  height: 240px;
  background: #1E1F22;
  color: #c8c9cc;
}

</style>