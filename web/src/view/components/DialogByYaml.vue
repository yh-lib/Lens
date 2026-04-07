<script setup>
import YamlEdit from './YamlEdit.vue';

const props = defineProps({
    dialogVisible: Boolean,
    itemByYaml: String,
})

const emit = defineEmits(['closeDialog'])

</script>

<template>
    <el-dialog 
            v-model="props.dialogVisible"
            title="配置编辑"
            width="1600px"
            style="height: 700px;"
            @close="emit('closeDialog')"
            destroy-on-close
    >
        <yaml-edit :code="props.itemByYaml"></yaml-edit>
    </el-dialog>
</template>