<script setup>
import { reactive } from 'vue';
import El_card from '../components/El_card.vue';
import Table from './Table.vue';
import { getPodListHandler } from '../../api/pod.js';

// 渲染表格数据
const getList = () => {
    if (!data.curClusterId || !data.curNsName) {
        data.items = []
        return
    }
    getPodListHandler(data.curClusterId, data.curNsName).then((res) => {
        data.items = res.data.data.items || []
    })
}

// 从子组件 EL_card 中获取所需数据
const getSelectValue = (selectValue) =>{
    Object.assign(data, selectValue)
    getList()
}

// 接受子组件传递的参数
const data = reactive({
    items: [],
    curClusterId: '',
    curNsName: '',
    search: '',
})
</script>

<template>
    <!-- 卡片主体 -->
    <El_card 
        title="pod 列表" 
        :op-cluster="true" 
        :op-ns="true" 
        :op-search="true" 
        :op-create="true"
        @change="getSelectValue"
    >
        <!-- 卡片 main 部分 table 数据 -->
        <template #table>
            <Table :table-data="data"></Table>
        </template>
    </El_card> 
</template>