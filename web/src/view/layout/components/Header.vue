<script setup>
import { SwitchButton, Fold, Expand } from '@element-plus/icons-vue'
import { logout } from '../../../api/login';
import { useIsCollapse } from '../../../store/index.js'
import { storeToRefs } from 'pinia';

// 获取当前菜单栏的状态
const {isCollapse} = storeToRefs(useIsCollapse())

// 变更菜单栏折叠状态
const collapseChange = () => {
    const isCollapse = useIsCollapse().changeisCollapse()
}
</script>

<template>
    <el-header class="el-header">
        <el-card class="header-shell">
            <div class="header-actions">
                <el-button class="icon-button" @click="collapseChange">
                    <el-icon size="20" v-show="!isCollapse"><Fold /></el-icon>
                    <el-icon size="20" v-show="isCollapse"><Expand /></el-icon>
                </el-button>
                <el-button class="icon-button" @click="logout">
                    <el-icon size="20"><SwitchButton /></el-icon>
                </el-button>
            </div>
        </el-card>
    </el-header>
</template>   

<style lang="less" scoped>
.el-header{
    display: flex;
    align-items: center; 
    padding: 0 16px;
}

.header-shell{
    width: 100%;
    border: 1px solid #e6eaf0;
    border-radius: 22px;
    box-shadow: 0 12px 24px rgba(15, 23, 42, 0.04);
}

.header-actions{
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 12px;
}

.icon-button{
    border: 1px solid #d7dde6;
    border-radius: 12px;
}

.icon-button:hover{
    border-color: #cbd5e1;
}
</style>
