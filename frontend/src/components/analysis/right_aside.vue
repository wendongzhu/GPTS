<!-- vue + element 树形选择框（多选框）实现 -->
<!-- 思路：隐藏el-select的下拉框，改用el-popover弹出框和el-tree控件组合-->
<template>
  <div class="collection-right-aside-tree">
      <!-- popover+tree用于选择，树形控件放在弹出框中 -->
      <el-tree
          :data="treeOptions"
          :props="defaultProps"
          :default-expanded-keys="[1,2,3,4,5]"
          :default-checked-keys="[0,1,2,3,4,5,6,7,8]"
          show-checkbox
          @check="checkChange"
          ref="tree"
          node-key="id"
          style="background: #2B2D30; color: #c8c9cc"
      />

  </div>
</template>

<script>
export default {
  name: 'Home',
  components: {},
  data () {
    return {
      // 控制popover弹出框是否展示
      popoverVisible: false,
      // 需把数据整理成以下结构
      // tree数据(children的id第一位为父级id，用于在select中清除某一点，可找到其父级去掉全选)
      treeOptions: [
        {
          id: '0',
          name: '环境参数',
          children: [
            {
              id: '0-1',
              name: '测试盒温度'
            }
          ]
        }, {
          id: '1',
          name: '第1路',
          children: [
            {
              id: '1-1',
              name: '供电电压'
            }, {
              id: '1-2',
              name: '供电电流'
            }, {
              id: '1-3',
              name: '计数'
            }, {
              id: '1-4',
              name: '计数率'
            }, {
              id: '1-5',
              name: '供电开关'
            }
          ]
        },{
          id: '2',
          name: '第2路',
          children: [
            {
              id: '2-1',
              name: '供电电压'
            }, {
              id: '2-2',
              name: '供电电流'
            }, {
              id: '2-3',
              name: '计数'
            }, {
              id: '2-4',
              name: '计数率'
            }, {
              id: '2-5',
              name: '供电开关'
            }
          ]
        }, {
          id: '3',
          name: '第3路',
          children: [
            {
              id: '3-1',
              name: '供电电压'
            }, {
              id: '3-2',
              name: '供电电流'
            }, {
              id: '3-3',
              name: '计数'
            }, {
              id: '3-4',
              name: '计数率'
            }, {
              id: '3-5',
              name: '供电开关'
            }
          ]
        }, {
          id: '4',
          name: '第4路',
          children: [
            {
              id: '4-1',
              name: '供电电压'
            }, {
              id: '4-2',
              name: '供电电流'
            }, {
              id: '4-3',
              name: '计数'
            }, {
              id: '4-4',
              name: '计数率'
            }, {
              id: '4-5',
              name: '供电开关'
            }
          ]
        }, {
          id: '5',
          name: '第5路',
          children: [
            {
              id: '5-1',
              name: '供电电压'
            }, {
              id: '5-2',
              name: '供电电流'
            }, {
              id: '5-3',
              name: '计数'
            }, {
              id: '5-4',
              name: '计数率'
            }, {
              id: '5-5',
              name: '供电开关'
            }
          ]
        }, {
          id: '6',
          name: '第6路',
          children: [
            {
              id: '6-1',
              name: '供电电压'
            }, {
              id: '6-2',
              name: '供电电流'
            }, {
              id: '6-3',
              name: '计数'
            }, {
              id: '6-4',
              name: '计数率'
            }, {
              id: '6-5',
              name: '供电开关'
            }
          ]
        }, {
          id: '7',
          name: '第7路',
          children: [
            {
              id: '7-1',
              name: '供电电压'
            }, {
              id: '7-2',
              name: '供电电流'
            }, {
              id: '7-3',
              name: '计数'
            }, {
              id: '7-4',
              name: '计数率'
            }, {
              id: '7-5',
              name: '供电开关'
            }
          ]
        }, {
          id: '8',
          name: '第8路',
          children: [
            {
              id: '8-1',
              name: '供电电压'
            }, {
              id: '8-2',
              name: '供电电流'
            }, {
              id: '8-3',
              name: '计数'
            }, {
              id: '8-4',
              name: '计数率'
            }, {
              id: '8-5',
              name: '供电开关'
            }
          ]
        }
      ],
      // select选择框选项
      typeOption: [
        {
          id: '1',
          name: '父级1'
        },
        {
          id: '1-1',
          name: '类型1'
        }, {
          id: '1-2',
          name: '类型2'
        }, {
          id: '1-3',
          name: '类型3'
        }, {
          id: '1-4',
          name: '类型4'
        }, {
          id: '1-5',
          name: '类型5'
        }, {
          id: '2',
          name: '父级2'
        }, {
          id: '2-1',
          name: '类型6'
        }, {
          id: '2-2',
          name: '类型7'
        }
      ],
      defaultProps: {
        children: 'children',
        label: 'name'
      },
      // select的值
      typeValue: []
    }
  },
  created () {},
  computed: {},
  watch: {},
  mounted () {},
  methods: {
    // tree选择值修改时
    checkChange () {
      this.typeValue = []
      // 将tree选择的id赋值给select
      this.$refs['tree'].getCheckedNodes().forEach(value => {
        // 父级在select中不展示
        if (value.id.indexOf('-') > 0) {
          this.typeValue.push(value.id)
        }
      })
    },
    // 清空select
    clearTag () {
      // 清空tree选择
      this.$refs['tree'].setCheckedKeys([])
    },
    // 从select中单个移除时，保持tree选择值同步移除
    removeTag (data) {
      // 获取tree目前选择的值
      var chooseData = this.$refs['tree'].getCheckedKeys()
      var deleteIndex = ''
      // 找到chooseData中与清除的data相同的值
      chooseData.forEach((value, index) => {
        if (value === data) {
          deleteIndex = index
        }
      })
      // 从tree目前选择值中去掉
      chooseData.splice(deleteIndex, 1)
      // 若有全选情况，tree的选择值中有父级id，而select中无父级id，需用children的id找到父级id并去掉
      // 查找其父级id是否在chooseData中（即原来此父级是否全选），若在则去掉
      var findFatherData = chooseData.find(element => element === (data.split('-')[0]))
      if (findFatherData) {
        chooseData.splice(chooseData.indexOf(findFatherData), 1)
      }
      // 将修改后的值再赋给tree
      this.$refs['tree'].setCheckedKeys(chooseData)
    }
  }
}
</script>
<style>
/* 把select的下拉框隐藏 */
.hiddenSel{
  /*display: none;*/
}
.collection-right-aside-tree{
  background: #2B2D30;
  color: #2B2D30;
}
</style>

