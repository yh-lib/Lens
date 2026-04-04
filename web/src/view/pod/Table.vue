<script setup>
import { computed, watch } from 'vue';

const props = defineProps({
    tableData: {
        type: Object,
        default: () => ({
            items: [],
            curClusterId: '',
            curNsName: '',
            search: '',
        }),
    },
})

watch(
    () => props.tableData,
    (value) => {
        console.log('Table.vue 接收到的 data:', value)
    },
    { immediate: true, deep: true }
)

// 搜索后的 pod 列表数据
const filterTableData = computed(() =>
    (props.tableData.items || []).filter(
        (item) =>
            !props.tableData.search ||
            item.metadata.name.toLowerCase().includes(props.tableData.search.toLowerCase())
    )
)
</script>

<template>
    <el-table :data="filterTableData" height="70vh" >
        <el-table-column label="Pod 名称" prop="metadata.name" />
        <el-table-column label="Pod IP" prop="status.podIP" />
        <el-table-column label="创建时间" prop="metadata.creationTimestamp" />
        <el-table-column label="命名空间" prop="metadata.namespace" />
        <el-table-column label="状态" prop="status.phase" />

        <el-table-column label="重启次数" prop="" />
        <el-table-column label="容器状态" prop="" />

        <el-table-column label="宿主机名" prop="spec.nodeName" />
        <el-table-column label="宿主机IP" prop="status.hostIP" />
        <el-table-column label="操作">
            <template #default="scope">
                <el-button link type="danger" @click="deleteItem(scope.row.metadata.name)">删除</el-button>
            </template>   
        </el-table-column>
    </el-table>
</template> 