<script setup>
import { onBeforeMount, reactive, readonly, ref } from 'vue';
import { useWorkLoadData } from '../../../../store';
import { storeToRefs } from 'pinia';

// store from pinia
const store = useWorkLoadData()
const { workLoadItem } = storeToRefs(store)

const data = reactive({
    volumeItem:{
        'name':'',
        'nfs':{
            'server':'',
            'path':'',
            'readOnly':false
        }
    },
    'readOnlyOptions':[
        {'label':'是','value':true},
        {'label':'否','value':false}
    ]
})

const validate = () => ruleFormRef.value?.validate()

// 暴露 data 给父组件
defineExpose({
    data,
    validate
})

// 接收父组件参数
const props = defineProps({
    method: String,
    parentVolumeItem: {
        type: Object,
        default: {},
    }
})

// 数据渲染
onBeforeMount(()=>{
    if (props.method != 'add') {
        data.volumeItem = props.parentVolumeItem
    }
})

// 表单校验：存储卷名称不能重复
const validateVolumeName = (rule, value, callback) => {
    // 如果没有输入，视为必填不通过，返回对应错误信息
    if (!value) {
        callback(new Error('请输入存储卷名称'))
        return
    }

    // 从当前 workload 中取出 volumes 列表，未取到时用空数组兜底
    const volumeList = workLoadItem.value?.item?.spec?.template?.spec?.volumes || []

    // 查找是否存在与当前输入同名且不是正在编辑对象的 volume
    const duplicated = volumeList.some(item => item !== data.volumeItem && item?.name === value)
    
    if (duplicated) {
        callback(new Error('存储卷名称不能重复'))
        return
    }
    callback()
}

// 表单校验
const ruleFormRef = ref()
const rules = reactive({
    name: [
        { validator: validateVolumeName, trigger: 'blur' },
    ],
    'nfs.server':[
        { required: true, message: '请输入NFS服务器地址', trigger: 'blur' },
    ],
    'nfs.path': [
        { required: true, message: '请输入NFS路径地址', trigger: 'blur' },
        { pattern: '^(?=/)', message: 'NFS路径地址必须以 / 开头', trigger: 'blur' },
    ]
})

</script>

<template>
    <el-form 
        label-width="120px" 
        label-position="left" 
        size="large"
        :rules="rules"
        ref="ruleFormRef"
        :model="data.volumeItem"
    >
        <!-- 名称 -->
        <el-form-item label="名称" prop="name" v-show="props.method">
            <el-input v-model="data.volumeItem.name" placeholder="请输入Volume名称"/>                    
        </el-form-item>
        <el-form-item label="NFS服务器地址" prop="nfs.server">
            <el-input v-model="data.volumeItem.nfs.server" placeholder="请输入NFS服务器地址"/>                    
        </el-form-item>
        <el-form-item label="NFS路径地址" prop="nfs.path">
            <el-input v-model="data.volumeItem.nfs.path" placeholder="请输入NFS路径地址"/>                    
        </el-form-item>
        <el-form-item label="是否只读" prop="nfs.readOnly">
            <el-select v-model="data.volumeItem.nfs.readOnly" placeholder="是否只读">
                <el-option
                    v-for="item in data.readOnlyOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                />
            </el-select>                    
        </el-form-item>
    </el-form>
</template>    