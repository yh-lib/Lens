<script setup>
  import { computed, reactive } from 'vue'
  import { useWorkLoadData } from '../../../../store'
  import { storeToRefs } from 'pinia'

  // store from pinia
  const store = useWorkLoadData()
  const { workLoadItem } = storeToRefs(store)

  const data = reactive({
    containersItem: workLoadItem.value.item.spec.template.spec.containers,
  })
  const props = defineProps(['containerItem'])
  // update data to template
  const syncMinCpu2template = () => {
    console.log('bug1:', props.containerItem.resources)
    const value2String = props.containerItem.resources.requests.cpu.toString() + 'm'
    console.log(value2String)
    // console.log(value2String)
    // props.containerItem.requests.cpu = props.containerItem.requests.cpu.toString()
    props.containerItem.resources.requests.cpu = value2String
    console.log(props.containerItem.resources.requests.cpu)
  }
</script>

<template>
  <el-form label-width="100px" label-position="left" :model="data.containersItem">
    <el-row :gutter="100">
      <el-col :span="8">
        <el-form-item label="容器名称">
          <el-input v-model="props.containerItem.name" />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="工作目录">
          <el-input v-model="props.containerItem.workingDir" />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="分配终端">
          <el-select v-model="props.containerItem.tty" placeholder="是否分配终端">
            <el-option
              v-for="item in [
                { lable: '是', value: 'true' },
                { lable: '否', value: 'false' },
              ]"
              :key="item.value"
              :label="item.lable"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="镜像地址">
          <el-input v-model="props.containerItem.image" />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="镜像拉取策略">
          <el-select v-model="props.containerItem.imagePullPolicy" placeholder="请选择镜像拉取策略">
            <el-option
              v-for="item in ['IfNotPresent', 'Always', 'Never']"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="最小CPU">
          <el-input
            v-model="props.containerItem.resources.requests.cpu"
            placeholder="请输入 CPU 申请值"
          />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="最小内存">
          <el-input
            v-model="props.containerItem.resources.requests.memory"
            placeholder="请输入内存申请值"
          />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="最大CPU">
          <el-input
            v-model="props.containerItem.resources.limits.cpu"
            placeholder="请输入 CPU 限制阈值"
          />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="最大内存">
          <el-input
            v-model="props.containerItem.resources.limits.memory"
            placeholder="请输入内存限制阈值"
          />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="启动命令">
          <el-input
            type="textarea"
            :rows="3"
            v-model="props.containerItem.command"
            placeholder="请输入内存限制阈值"
          />
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="启动参数">
          <el-input
            type="textarea"
            :rows="3"
            v-model="props.containerItem.args"
            placeholder="请输入内存限制阈值"
          />
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>
</template>
