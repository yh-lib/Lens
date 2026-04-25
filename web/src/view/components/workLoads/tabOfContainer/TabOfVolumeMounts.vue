<script setup>
  import { onMounted } from 'vue'
  import { useWorkLoadData } from '../../../../store'
  import { storeToRefs } from 'pinia'
  import { ElMessage } from 'element-plus'
  // store from pinia
  const store = useWorkLoadData()
  const { workLoadItem } = storeToRefs(store)

  const props = defineProps(['containerItem'])
</script>

<template>
  <el-table :data="props.containerItem.volumeMounts" border>
    <el-table-column prop="name" label="选择Volume">
      <template #default="scope">
        <el-select v-model="scope.row.name" placeholder="请选择要挂载的Volume" style="width: 240px">
          <el-option
            v-for="item in workLoadItem.item.spec.template.spec.volumes"
            :key="item.name"
            :label="item.name"
            :value="item.name"
          />
        </el-select>
      </template>
    </el-table-column>
    <el-table-column prop="mountPath" label="挂载路径">
      <template #default="scope">
        <el-input v-model="scope.row.mountPath" placeholder="请输入挂载路径" />
      </template>
    </el-table-column>
    <el-table-column prop="subPath" label="请输入挂载子路径">
      <template #default="scope">
        <el-input v-model="scope.row.subPath" placeholder="请输入挂载子路径" />
      </template>
    </el-table-column>
    <el-table-column prop="readOnly" label="是否只读">
      <template #default="scope">
        <el-select v-model="scope.row.readOnly" placeholder="挂载是否只读">
          <el-option
            v-for="item in [
              { label: '是', value: true },
              { label: '否', value: false },
            ]"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </template>
    </el-table-column>
    <el-table-column width="200px" align="center">
      <template #header>
        <el-button
          type="primary"
          link
          style="font-size: 16px; margin-bottom: 10px"
          @click="
            props.containerItem.volumeMounts.push({
              name: '',
              mountPath: '',
              readOnly: '',
              subPath: '',
            })
          "
          >添加</el-button
        >
      </template>
      <template #default="scope">
        <el-button
          type="danger"
          link
          style="font-size: 16px; margin-bottom: 10px"
          @click="props.containerItem.volumeMounts.splice(scope.$index, 1)"
          >删除</el-button
        >
      </template>
    </el-table-column>
  </el-table>
</template>
