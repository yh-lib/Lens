<script setup>
  import { computed, reactive, ref } from 'vue'
  import DialogByYaml from '../components/DialogByYaml.vue'
  import { obj2yaml } from '../../utils/typeConv/type.conv.js'
  import DialogHeaderLabel from '../components/DialogHeaderLabel.vue'
  import DialogOfItem from '../components/workLoads/DialogOfItem.vue'
  import { getdeploymentHandler } from '../../api/deployment.js'
  import { useWorkLoadData } from '../../store/index.js'
  import { storeToRefs } from 'pinia'

  // store from pinia
  const store = useWorkLoadData()
  const { workLoadItem } = storeToRefs(store)

  const emit = defineEmits(['deleteItem'])

  const getAgeText = (row) => {
    const creationTimestamp = row?.metadata?.creationTimestamp
    if (!creationTimestamp) {
      return '-'
    }

    const createdAt = new Date(creationTimestamp).getTime()
    if (Number.isNaN(createdAt)) {
      return '-'
    }

    const diffSeconds = Math.max(0, Math.floor((Date.now() - createdAt) / 1000))
    if (diffSeconds < 60) {
      return `${diffSeconds}秒钟前`
    }

    const diffMinutes = Math.floor(diffSeconds / 60)
    if (diffMinutes < 60) {
      return `${diffMinutes}分钟前`
    }

    const diffHours = Math.floor(diffMinutes / 60)
    if (diffHours < 24) {
      return `${diffHours}小时前`
    }

    const diffDays = Math.floor(diffHours / 24)
    return `${diffDays}天前`
  }

  const props = defineProps(['tableData'])

  // 搜索后的 pod 列表数据
  const filterTableData = computed(() =>
    (props.tableData.items || []).filter(
      (item) =>
        !props.tableData.search ||
        item.metadata.name.toLowerCase().includes(props.tableData.search.toLowerCase())
    )
  )
  // const curItem = ref('')
  // const itemByYaml = ref('')
  // const itemDetailDialog = ref(false)
  // const getItem = (row) => {
  //   curItem.value = row
  //   itemByYaml.value = obj2yaml(row)
  //   itemDetailDialog.value = true
  // }

  const mergeIfExists = (target, source) => {
    Object.keys(source || {}).forEach((key) => {
      const sourceValue = source[key]
      const targetValue = target[key]

      if (sourceValue === undefined) {
        return
      }

      if (Array.isArray(sourceValue)) {
        if (!Array.isArray(targetValue)) {
          target[key] = sourceValue
          return
        }

        target[key] = sourceValue.map((item, index) => {
          const currentTargetItem = targetValue[index]

          if (
            item &&
            typeof item === 'object' &&
            !Array.isArray(item) &&
            currentTargetItem &&
            typeof currentTargetItem === 'object' &&
            !Array.isArray(currentTargetItem)
          ) {
            return mergeIfExists(currentTargetItem, item)
          }

          return item
        })
        return
      }

      if (
        sourceValue &&
        typeof sourceValue === 'object' &&
        !Array.isArray(sourceValue) &&
        targetValue &&
        typeof targetValue === 'object' &&
        !Array.isArray(targetValue)
      ) {
        mergeIfExists(targetValue, sourceValue)
        return
      }

      target[key] = sourceValue
    })

    return target
  }

  const updateItem = (row) => {
    store.resetWorkLoadItem()
    console.log('编辑Item:::', workLoadItem.value)
    // 数据赋值
    getdeploymentHandler(
      props.tableData.clusterId,
      props.tableData.nameSpace,
      row.metadata.name
    ).then((res) => {
      if (res.data.status == 200) {
        console.log('数据核对1:', workLoadItem.value)
        mergeIfExists(workLoadItem.value.item, res.data.data.items)
        console.log('数据核对2:', workLoadItem.value.item.spec.template.spec.imagePullSecrets)
        data.updateItemDialogVisible = true
      }
    })
  }

  const data = reactive({
    updateItemDialogVisible: false,
  })
  const closeDialogOfItem = () => {
    data.updateItemDialogVisible = false
  }
</script>

<template>
  <el-table :data="filterTableData" height="1010px">
    <el-table-column label="名称" prop="metadata.name" width="300px">
      <template #default="scope">
        <el-button type="primary" link @click="updateItem(scope.row)">{{
          scope.row.metadata.name
        }}</el-button>
      </template>
    </el-table-column>
    <el-table-column label="创建时间" prop="metadata.creationTimestamp" />
    <el-table-column label="存活时间" prop="age">
      <template #default="scope">
        {{ getAgeText(scope.row) }}
      </template>
    </el-table-column>
    <el-table-column label="命名空间" prop="metadata.namespace" />
    <el-table-column label="状态" prop="status">
      <template #default="scope">
        <span v-if="scope.row.status.conditions[0].status == 'True'" style="color: green"
          >Available</span
        >
        <span v-else style="color: red">UnAvailable</span>
      </template>
    </el-table-column>
    <el-table-column label="Pods" prop="containerStatus">
      <template #default="scope">
        {{ scope.row.status.availableReplicas || 0 }}/{{ scope.row.status.replicas }}
      </template>
    </el-table-column>
    <el-table-column label="操作" prop="operations">
      <template #default="scope">
        <el-button link type="warning" @click="updateItem(scope.row)">编辑</el-button>
        <el-button link type="danger" @click="emit('deleteItem', scope.row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>

  <!-- <DialogByYaml
    :dialogVisible="itemDetailDialog"
    @closeDialog="itemDetailDialog = false"
    :item-by-yaml="itemByYaml"
  >
    <template #header>
      <DialogHeaderLabel
        :cur-cluster-id="props.tableData.curClusterId"
        :cur-ns-name="curItem.metadata.namespace"
        :cur-resource-name="curItem.metadata.name"
        :cur-node-name="curItem.spec.nodeName"
        cur-resource-kind="Deployment"
      />
    </template>
  </DialogByYaml> -->
  <DialogOfItem :open-dialog="data.updateItemDialogVisible" @close-dialog="closeDialogOfItem" />
</template>
