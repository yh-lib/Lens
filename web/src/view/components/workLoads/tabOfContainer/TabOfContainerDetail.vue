<script setup>
  import { reactive } from 'vue'
  import TabOfBasic from './TabOfBasic.vue'
  import TabOfPort from './TabOfPort.vue'
  import TabOfHealth from '../tabOfHealth/TabOfHealth.vue'
  import TabOfVolumeMounts from './TabOfVolumeMounts.vue'
  import { useWorkLoadData } from '../../../../store'
  import { storeToRefs } from 'pinia'
  import { ElMessage } from 'element-plus'
  import TableOfEnv from './TableOfEnv.vue'
  import TabOfLifeCycle from './TabOfLifeCycle.vue'

  // store from pinia
  const store = useWorkLoadData()
  const { workLoadItem } = storeToRefs(store)

  const data = reactive({
    activeName: 'Basic',
    tabOfVolumeMountDisabled: false,
  })
  const props = defineProps(['containerItem', 'containerType'])
  const beforeTabLeave = (activeName) => {
    if (
      activeName == 'VolumeMount' &&
      workLoadItem.value.item.spec.template.spec.volumes.length == 0
    ) {
      ElMessage.error('请先添加Volume卷')
      return false
    }

    return true
  }
</script>

<template>
  <el-tabs
    tab-position="left"
    style="height: 560px"
    v-model="data.activeName"
    :before-leave="beforeTabLeave"
  >
    <el-tab-pane label="基础配置" name="Basic">
      <TabOfBasic style="padding: 20px 30px" :container-item="props.containerItem" />
    </el-tab-pane>
    <el-tab-pane label="端口配置" name="Port" v-if="props.containerType == 'container'">
      <TabOfPort style="padding: 10px 30px" :container-item="props.containerItem" />
    </el-tab-pane>
    <el-tab-pane label="健康检查" name="Health" v-if="props.containerType == 'container'">
      <TabOfHealth :container-item="props.containerItem" />
    </el-tab-pane>
    <el-tab-pane label="环境变量" name="Env">
      <TableOfEnv :container-item="props.containerItem" />
    </el-tab-pane>
    <el-tab-pane label="存储挂载" name="VolumeMount" :disabled="data.tabOfVolumeMountDisabled">
      <TabOfVolumeMounts :container-item="props.containerItem" />
    </el-tab-pane>
    <el-tab-pane label="生命周期" name="LifeCycle" v-if="props.containerType == 'container'">
      <TabOfLifeCycle />
    </el-tab-pane>
  </el-tabs>
</template>
