<script setup>
import { ref } from 'vue'
import ElCard from './ElCard.vue';

const props = defineProps(['openDialog'])
const emit = defineEmits(['closeDialog'])
const activeName = ref('Basic')

</script>

<template>
  <!-- 资源创建 dialog -->
  <el-dialog
    v-model="props.openDialog"
    width="1600px"
    style="height: 840px;"
    @close="emit('closeDialog')"
  >
    <el-tabs v-model="activeName" @tab-click="handleClick">
        <!-- 标签：基本配置 -->
        <el-tab-pane label="基本配置" name="Basic">
            <ElCard :op-cluster="true" :op-ns="true" style="border-radius: 0px;width: 1560px;">
              <template #mainData>
                <el-row :gutter="20">
                  <el-col :span="8"><div class="grid-content ep-bg-purple" />名称</el-col>
                  <el-col :span="8"><div class="grid-content ep-bg-purple" />副本数</el-col>
                  <el-col :span="8"><div class="grid-content ep-bg-purple" />私有仓库密钥</el-col>
                </el-row>
              </template>                
            </ElCard>
        </el-tab-pane>
        <el-tab-pane label="存储卷配置" name="Volume">存储卷配置</el-tab-pane>
        <el-tab-pane label="调度配置" name="Schedule">调度配置</el-tab-pane>
        <el-tab-pane label="容器配置" name="Container">容器配置</el-tab-pane>
        <el-tab-pane label="初始化容器" name="InitContainer">初始化容器</el-tab-pane>
    </el-tabs>
  </el-dialog>
</template>


