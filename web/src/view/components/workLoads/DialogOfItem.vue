<script setup>
import { ref } from 'vue'
import { createdeploymentHandler } from '../../../api/deployment';
import { ElMessage } from 'element-plus';
import { list2obj } from '../../../utils/typeConv/type.conv';
import { useWorkLoadData } from '../../../store';
import { storeToRefs } from 'pinia';
import TabOfBasicConfig from './TabOfBasicConfig/TabOfBasicConfig.vue';
import TabOfScheduleConfig from './tabOfScheduleConfig/TabOfScheduleConfig.vue';
import DialogByYaml from '../DialogByYaml.vue';
import { obj2yaml } from '../../../utils/typeConv/type.conv';
import YamlEdit from '../YamlEdit.vue';

const props = defineProps(['openDialog'])
const emit = defineEmits(['closeDialog'])
const activeName = ref('Basic')

// 创建ref来获取子组件实例
const basicRef = ref(null)

// store from pinia
const store = useWorkLoadData()
const { workLoadItem } = storeToRefs(store)



const createItem = () => {
  if (basicRef.value.data.labelsAndAnnotationsSwtich == 'auto'){
    // 自动生成
    workLoadItem.value.item.metadata.labels.app = workLoadItem.value.item.metadata.name
    workLoadItem.value.item.spec.selector.matchLabels.app = workLoadItem.value.item.metadata.name
    workLoadItem.value.item.spec.template.metadata.labels.app = workLoadItem.value.item.metadata.name
  }else {
    workLoadItem.value.item.metadata.labels = list2obj(basicRef.value.data.controllerLabelsList)
    workLoadItem.value.item.spec.selector.matchLabels = list2obj(basicRef.value.data.controllerLabelsList)
    workLoadItem.value.item.spec.template.metadata.labels = list2obj(basicRef.value.data.controllerLabelsList)
    workLoadItem.value.item.metadata.annotations = list2obj(basicRef.value.data.controllerAnnotationsList)
  }
  // 如果 imagePullSecrets 为空，则删除该字段，否则会报错
  if (
    workLoadItem.value.item.spec.template.spec.imagePullSecrets &&
    workLoadItem.value.item.spec.template.spec.imagePullSecrets.length &&
    workLoadItem.value.item.spec.template.spec.imagePullSecrets[0].name == ''
  ){
    delete workLoadItem.value.item.spec.template.spec.imagePullSecrets
  }
  // 调用后端接口 创建 Deployment
  createdeploymentHandler(workLoadItem.value).then((res)=>{
      if (res.data.status === 200) {
        ElMessage({
                    message: res.data.message,
                    type: 'success',
        })
      }
  })
}
// Yaml
const itemOfYaml = ref('')
const getItemOfYaml = (tab) => {
  if (tab.paneName !== 'Yaml') return
  itemOfYaml.value = obj2yaml(workLoadItem.value)
}
</script>

<template>
  <!-- 资源创建 dialog -->
  <el-dialog
    v-model="props.openDialog"
    width="1600px"
    style="height: 740px;"
    @close="emit('closeDialog')"
  >
    <el-tabs v-model="activeName" @tab-click="getItemOfYaml">
        <el-tab-pane label="基本配置" name="Basic">
          <TabOfBasicConfig ref="basicRef" />
        </el-tab-pane>
        <el-tab-pane label="调度配置" name="Schedule">
          <TabOfScheduleConfig />
        </el-tab-pane>
        <el-tab-pane label="存储卷配置" name="Volume">
          存储卷配置
        </el-tab-pane>        
        <el-tab-pane label="容器配置" name="Container">
          容器配置
        </el-tab-pane>
        <el-tab-pane label="初始化容器" name="InitContainer">
          初始化容器
        </el-tab-pane>
        <el-tab-pane label="Yaml" name="Yaml">
          <YamlEdit 
          :code="itemOfYaml"
          style="height: 560px;"
          />
        </el-tab-pane>
    </el-tabs>
    <el-button type="primary" size="large" style="margin-top: 20px;width: 90px;" @click="createItem">创建</el-button>
  </el-dialog>
  
</template>

<style scoped>
.form-item{
  margin-bottom: 15px;
}
:deep(.no-border-input .el-input__wrapper){
    box-shadow: none;
    border: none;
}
</style>

