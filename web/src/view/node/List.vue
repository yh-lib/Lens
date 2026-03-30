<script lang="ts" setup>
import { computed, reactive, ref, onBeforeMount  } from 'vue'
import { getnodeListHandler as getListHandler,getnodeHandler as getHandler } from '../../api/node.js'
import { getClusterListHandler} from '../../api/cluster.js'
import Add from './Add.vue'
import { ElSelect } from 'element-plus'

// 需要的数据变量
const data = reactive({
  items: [],
  clusterList: [],
  clusterId: "",
  editItem: {},
  editNodeName: "",
  detailItem: {},
  detailNodeName: "",
  createName: "",
  // 集群选择器
  clusterOptions:[],
  value: "",
})

// 加载前
onBeforeMount(async () => {
    console.log("挂载前：：：：：：：：")
    // 获取集群列表
    await getClusterList()
    // 获取集群选择器菜单
    data.clusterList.forEach(item => {
        console.log("看这里:::::::::", item)
    })
})



const options = [
  {
    id: 'Option1',
    label: 'Option1',
  },
  {
    id: 'Option2',
    label: 'Option2',
    disabled: true,
  },
  {
    id: 'Option3',
    label: 'Option3',
  },
  {
    id: 'Option4',
    label: 'Option4',
    disabled: true,
  },
  {
    id: 'Option5',
    label: 'Option5',
  },
]


// 获取集群列表
const getClusterList = async ()=>{
    await getClusterListHandler().then((res)=>{
        if (res.data.status === 200) {
            data.clusterList = res.data.data.items;
            console.log('成功获取集群列表:::', data.clusterList);
        } else {
            console.error('获取集群列表失败:', res.data.message);
        }
    })
}

const getNodeList = ()=>{
    loading.value = true

}

// add_8 获取当前列表
const getList = () =>{
    loading.value = true
    getListHandler().then((response)=>{
        if (response.data.status === 200) {
            data.items = response.data.data.items; // 更新 tableData
            console.log('获取列表成功:', response.data.data.items);
            loading.value = false
        } else {
            console.error('获取列表失败:', response.data.message);
        }
    })
}


// 样式
const props = {
  value: 'id',
  label: 'label',
  options: 'options',
  disabled: 'disabled',
}




// ++++++增

// add_2:打开添加节点的表单
const opDialog = ref(false)

// add_4 定义表单初始数据

// add_7 监听refresh事件，刷新用户列表
const refreshList = () =>{
    opDialog.value = false
    getList()
}

// add_9 关闭dialog时刷新用户列表
const closeDialog = () =>{
    method.value == 'create' && getList()
}

// ++++++改

// update_2 更新
const updateItem = (row) => {
    // 获取节点详情
    getHandler(row.nodeId).then((response)=>{
        data.itemForm=response.data.data.item  
        // 注意下面这两步：如果放在axios外面，则赋值可能失败，因为是异步运行
        // 传递给操作参数给子组件
        method.value='update'
        // 打开表单弹窗
        opDialog.value = true
    }) 
}
// update_3 将method.value、data.itemForm赋值给子组件
// update_4 在子组件中将值渲染到form表单中

// ++++++查

// select_1 子组件加载时自动获取列表
onBeforeMount(() => {
    getList();
})
// select_2 搜索功能
const search = ref('')
const filterTableData = computed(() =>
  data.items.filter(
    (item) =>
      !search.value ||
      item.metadata.name.toLowerCase().includes(search.value.toLowerCase()) ||
      item.status.addresses[0].address.toLowerCase().includes(search.value.toLowerCase())
  )
)
// ++++++独立配置
// Main 标题
const titleName = "节点列表"
// Table 表头
const tableTtile = {
    f1: { prop: "hostName", label: "主机名" },
    f2: { prop: "ipAddress", label: "IP地址" },
    f3: { prop: "role", label: "角色" },
    f4: { prop: "status", label: "状态" },
    f5: { prop: "scheduling", label: "禁止调度" },
    f6: { prop: "eviction", label: "驱逐" },
    f7: { prop: "edit", label: "操作" }
}
// 加载图标默认关闭
const loading = ref(false)
// 集群选择器

</script>

<template>
    <el-card>
        <!-- add_1:添加按钮 -->
        <template #header>
            <div class="card-header">
                <div>
                    <span style="font-size: 24px;">{{ titleName }}</span>
                </div>
                <div>
                    <el-select v-model="data.value" placeholder="选择集群" style="width: 240px;margin-right: 10px;">
                        <el-option
                            v-for="item in data.clusterList"
                            :key="item.clusterId"
                            :label="item.clusterId"
                            :value="item.clusterId"
                            :disabled="item.disabled"
                        />     
                    </el-select>               
                    <el-input v-model="search" style="width: 240px" placeholder="搜索节点"/>
                </div>
            </div>
        </template>
        <!-- 表格 -->
        <el-table :data="filterTableData" style="width: 100%"  height="70vh" v-loading="loading">
            <!-- f1 主机名-->
            <el-table-column :label="tableTtile.f1.label">
                <template #default="scope">
                    <el-button link type="primary" @click="edit(scope.row)">
                        {{ scope.row.status.addresses[1].address}} 
                    </el-button>                        
                </template>
            </el-table-column>
            <!-- f2 IP地址-->
            <el-table-column :label="tableTtile.f2.label" :prop="tableTtile.f2.prop" >
                <template #default="scope">
                    {{ scope.row.status.addresses[0].address}} 
                </template>
            </el-table-column>
            <!-- f3 角色-->
            <el-table-column :label="tableTtile.f3.label" :prop="tableTtile.f3.prop" >
                <template #default="scope">
                    <el-button link type="primary" @click="edit(scope.row)">
                        {{ scope.row.status.addresses[1].address}} 
                    </el-button>                        
                </template>                  
            </el-table-column>
            <!-- f4 状态-->
            <el-table-column :label="tableTtile.f4.label" :prop="tableTtile.f4.prop" >
                <template #default="scope">
                    <span v-if="scope.row.status.conditions[scope.row.status.conditions.length-1].status == 'True'" style="color: green;">Ready</span>
                    <span v-else  style="color: red;">NotReady</span>
                </template>
            </el-table-column>
            <!-- f5 禁止调度-->
            <el-table-column :label="tableTtile.f5.label" prop="" >
                <template #default="scope">
                    <el-switch v-model="value1" />
                </template>
            </el-table-column>
            <!-- f6 驱逐-->
            <el-table-column :label="tableTtile.f6.label" :prop="tableTtile.f6.prop" >
                <template #default="scope">
                    <el-switch v-model="value1" />
                </template>     
            </el-table-column>           
            <!-- f7 操作-->
            <el-table-column :label="tableTtile.f7.label" :prop="tableTtile.f7.prop">
                <template #default="scope">
                    <el-button link type="primary" @click="edit(scope.row)">配置</el-button>
                </template>
            </el-table-column>
        </el-table>  
        <!-- add_3:动态渲染表单抬头/add_5:打开表单-->
        <el-dialog 
            v-model="opDialog"
            :title="method == 'create' ? '添加节点' : '更新节点'"
            width="50vw"
            @close="closeDialog"
            destroy-on-close
        >
            <!-- add_6:添加节点表单 -->
            <!-- update3 将method.value、data.itemForm赋值给子组件 -->
            <Add :subMethod='method' :subRow="data.itemForm" @refresh="refreshList"/>
        </el-dialog>            
    </el-card>
</template>

<style scoped>
.card-header{
    display: flex;
    justify-content: space-between;
    align-items: center;

}
</style>