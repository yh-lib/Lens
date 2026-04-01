<script lang="ts" setup>
import { computed, reactive, ref, onBeforeMount  } from 'vue'
import { getnodeListHandler as getListHandler,getnodeHandler as getHandler } from '../../api/node.js'
import { getClusterListHandler} from '../../api/cluster.js'
import { ElSelect } from 'element-plus'
import { Loading } from '@element-plus/icons-vue'
import { obj2list } from '../../utils/typeConv/type.conv.js'

// 需要的数据变量
const data = reactive({
    loading: false,  // 默认关闭加载图标
    items: [],  // 后端返回的 list
    item: {}, // 后端返回的 obj
    search: "", // 接收 header 搜索框中数据
    curHostName: "",
    // 集群选择器
    clusterOptions:[],  // 集群选择 options
    curClusterId: "",   // 当前选择的集群id
    defaultClusterId: "in-cluster", // 默认选择的集群id
    // 表格
    opDialog: false,    // 配置、主机名 对话框
    nodeLabels: [],          // 后端返回的label
    nodeTaints: [],
})

// // // 子组件加载前自动获取数据
onBeforeMount(async () => {
    await getclusterOptions()   // 获取集群列表
    data.curClusterId=data.defaultClusterId // 获取基础设施集群的Node列表
    getList();
})

// 获取节点角色
const getNodeRole = computed(() => {
    const keyList = Object.keys(data.item.metadata.labels)
    if (keyList.length == 0) {
        return "工作节点"
    }
    if (keyList.includes('node-role.kubernetes.io/control-plane') || keyList.includes('node-role.kubernetes.io/controlplane') || keyList.includes('node-role.kubernetes.io/master')) {
        return "控制节点"
    }
    return "工作节点"
})
    

// 获取节点标签
const getLabel = () => {
    // console.log("获取节点调试数据:",data.item.spec.taints)
    // 获取标签
    data.nodeLabels = obj2list(data.item.metadata.labels)
}

// 节点配置
const updateItem = (row) => {
    data.curHostName = row.metadata.name
    data.nodeTaints = row.spec.taints
    // 获取节点详情s
    getHandler(data.curClusterId,data.curHostName).then((res)=>{
        console.log("获取节点详情:",res)
        data.item=res.data.data.items
        data.opDialog = true
    }) 
}

// 关闭dialog时刷新用户列表
const closeDialog = () =>{
    getList()
}

// 内存地址单位转换
const memoryKi2MB = (memory) => {
    if (memory === null || memory === undefined) return ''
    const match = String(memory).trim().match(/^(\d+(?:\.\d+)?)Ki$/i)
    return match ? Math.floor(Number(match[1]) / 1000) : ''
}

// 搜索后的 node 列表数据
const filterTableData = computed(() =>
  data.items.filter(
    (item) =>
      !data.search ||
      item.metadata.name.toLowerCase().includes(data.search.toLowerCase()) ||
      item.status.addresses[0].address.toLowerCase().includes(data.search.toLowerCase())
  )
)

// 获取当前集群Node列表
const getList = () =>{
    data.loading = true
    getListHandler(data.curClusterId).then((res)=>{
        data.items = res.data.data.items;
        data.loading = false
    })
}

// 获取集群列表
const getclusterOptions = async ()=>{
    await getClusterListHandler().then((res)=>{
        if (res.data.status === 200) {
            data.clusterOptions = res.data.data.items;
            console.log('获取集群列表:::', data.clusterOptions);
        } else {
            console.error('获取集群列表失败:', res.data.message);
        }
    })
}
</script>

<template>
    <!-- card -->
    <el-card>
        <!-- card header -->
        <template #header>
            <div class="card-header">
                <div>
                    <span style="font-size: 24px;">节点列表</span>
                </div>
                <div>
                    <el-select v-model="data.curClusterId" placeholder="选择集群" style="width: 240px;margin-right: 10px;" @change="getList(data.curClusterId)">
                        <el-option
                            v-for="item in data.clusterOptions"
                            :key="item.clusterId"
                            :label="item.clusterId"
                            :value="item.clusterId"
                            :disabled="item.disabled"
                        />     
                    </el-select>               
                    <el-input v-model="data.search" style="width: 240px" placeholder="搜索节点"/>
                </div>
            </div>
        </template>
        <!-- card middle -->
        <el-table :data="filterTableData" style="width: 100%;"  height="70vh" v-loading="data.loading">
            <el-table-column label="主机名" prop="hostName">
                <template #default="scope">
                    <el-button link type="primary" @click="updateItem(scope.row)">
                        {{ scope.row.status.addresses[1].address}} 
                    </el-button>                        
                </template>
            </el-table-column>
            <el-table-column label="IP地址" prop="ipAddress" >
                <template #default="scope">
                    {{ scope.row.status.addresses[0].address}} 
                </template>
            </el-table-column>
            <el-table-column label="规格" prop="capacity" >
                <template #default="scope">
                    {{ scope.row.status.capacity.cpu}}C
                    |
                    {{ memoryKi2MB(scope.row.status.capacity.memory) }}MB
                </template>
            </el-table-column>
            <el-table-column label="角色" prop="role" >
                <template #default="scope">
                    <el-button link type="primary">
                        {{ scope.row.status.addresses[1].address}} 
                    </el-button>                        
                </template>                  
            </el-table-column>
            <el-table-column label="状态" prop="status" >
                <template #default="scope">
                    <span v-if="scope.row.status.conditions[scope.row.status.conditions.length-1].status == 'True'" style="color: green;">Ready</span>
                    <span v-else  style="color: red;">NotReady</span>
                </template>
            </el-table-column>
            <el-table-column label="禁止调度" prop="scheduling" >
                <template #default="scope">
                    <el-switch v-model="value1" />
                </template>
            </el-table-column>
            <el-table-column label="驱逐" prop="eviction" >
                <template #default="scope">
                    <el-switch v-model="value1" />
                </template>     
            </el-table-column>
            <el-table-column label="操作" prop="edit">
                <template #default="scope">
                    <el-button link type="primary" @click="updateItem(scope.row)">配置</el-button>
                </template>
            </el-table-column>
        </el-table>
        <!-- dialog -->
        <el-dialog 
            v-model="data.opDialog"
            width="1600px"
            style="height: 600px;"
            @close="closeDialog"
            destroy-on-close
        >
            <!-- dialog header -->
            <template #header>
                
                <div style="display: flex; justify-content: flex-start;">
                    <el-tag type="primary" effect="plain" style="margin-right: 10px;">集群: {{ data.curClusterId || '-' }}</el-tag>
                    <el-tag type="primary" effect="plain" style="margin-right: 10px;">节点: {{ data.item?.metadata?.name || '-' }}</el-tag>
                    <span style="font-size: 18px;margin-left: 500px;">节点配置</span>
                </div>
            </template>
            <!-- dialog middle -->
            <el-tabs v-model="nodeLabel" class="demo-tabs" @tab-click="getLabel">
                <!-- 标签   详情 -->
                <el-tab-pane label="详情" name="nodeDetail">
                    <el-descriptions :column="2" size="large" border>
                        <el-descriptions-item label="主机名">{{ data.curHostName }}</el-descriptions-item>
                        <el-descriptions-item label="IP地址">{{ data.item.status.addresses[0].address }}</el-descriptions-item>
                        <el-descriptions-item label="角色">{{ getNodeRole }}</el-descriptions-item>

                        <el-descriptions-item label="节点状态">
                            <el-tag :type="data.item.status.conditions[data.item.status.conditions.length-1].status=='True'?'success':'danger'">
                                {{ data.item.status.conditions[data.item.status.conditions.length-1].reason }}
                            </el-tag>
                        </el-descriptions-item>
                        <el-descriptions-item label="Pods"><el-button link type="primary">{{ data.item.status.capacity.pods }}</el-button></el-descriptions-item>
                        <el-descriptions-item label="操作系统">{{ data.item.status.nodeInfo.osImage }}</el-descriptions-item>

                        <el-descriptions-item label="Runtime">{{ data.item.status.nodeInfo.containerRuntimeVersion }}</el-descriptions-item>
                        <el-descriptions-item label="CPU架构">{{ data.item.status.nodeInfo.architecture }}</el-descriptions-item>
                        <el-descriptions-item label="组件版本">{{ data.item.status.nodeInfo.kubeletVersion }}</el-descriptions-item>

                        <el-descriptions-item label="系统内核">{{ data.item.status.nodeInfo.kernelVersion }}</el-descriptions-item>

                    </el-descriptions>                    
                </el-tab-pane>
                <!-- 标签   标签 -->
                <el-tab-pane label="标签" name="nodeLabel">
                    <el-table :data="data.nodeLabels" style="width: 100%; height:440px">
                        <el-table-column prop="key" label="Key"/>
                        <el-table-column prop="value" label="Value"/>
                        <el-table-column>
                            <template #header>
                                <el-button type="primary" link style="font-size: 16px;margin-bottom: 10px;" >添加</el-button>  
                            </template>
                            <template #default="scope">
                                <el-button type="danger" link style="font-size: 16px;margin-bottom: 10px;" >删除</el-button>
                            </template>                      
                        </el-table-column>
                    </el-table>                        
                </el-tab-pane>
                <!-- 标签   污点 -->
                <el-tab-pane label="污点" name="nodeTaint">
                    <el-table :data="data.nodeTaints" style="width: 100%; height:440px">
                        <el-table-column prop="key" label="Key"/>
                        <el-table-column prop="value" label="Value"/>
                        <el-table-column prop="Effect" label="Effect"/>
                        <el-table-column>
                            <template #header>
                                <el-button type="primary" link style="font-size: 16px;margin-bottom: 10px;" >添加</el-button>  
                            </template>
                            <template #default="scope">
                                <el-button type="danger" link style="font-size: 16px;margin-bottom: 10px;" >删除</el-button>
                            </template>                      
                        </el-table-column>
                    </el-table> 
                </el-tab-pane>
            </el-tabs>
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





