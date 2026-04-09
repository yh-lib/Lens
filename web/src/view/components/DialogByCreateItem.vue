<script setup>
import { onBeforeMount, reactive, ref } from 'vue'
import ElCard from './ElCard.vue';

const props = defineProps(['openDialog'])
const emit = defineEmits(['closeDialog'])
const activeName = ref('Basic')
const switchHostNetwork = ref(false)
const switchAddService = ref(false)
const data = reactive({
  privateRepoSecret:[],
})

onBeforeMount(()=>{
  // 获取私有仓库密钥，并赋值给data.privateRepoSecret
  
})

</script>

<template>
  <!-- 资源创建 dialog -->
  <el-dialog
    v-model="props.openDialog"
    width="1600px"
    style="height: 700px;"
    @close="emit('closeDialog')"
  >
    <el-tabs v-model="activeName" @tab-click="handleClick">
        <!-- 标签：基本配置 -->
        <el-tab-pane label="基本配置" name="Basic">
            <ElCard :op-cluster="true" :op-ns="true" style="border-radius: 0px;width: 1560px;">
              <template #mainData>
                <el-form label-width="150px" label-position="left" style="height: 450px;width: 1460px;">
                  <el-row :gutter="100">
                    <!-- 第一行 -->
                    <el-col :span="8" class="form-item">
                      <el-form-item label="名称" required>
                        <el-input placeholder="请输入资源名称"/>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8" class="form-item">
                      <el-form-item label="副本数" >
                        <el-input placeholder="请输入副本数"/>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8" class="form-item">
                      <el-form-item label="私有仓库密钥">
                        <el-select placeholder="请选择私有仓库密钥">
                          <el-options 
                          
                          />
                        </el-select>
                      </el-form-item>
                    </el-col>
                    <!-- 第二行 -->
                    <el-col :span="8" class="form-item">
                      <el-form-item label="DNS 策略">
                        <el-select placeholder="请选择 DNS 策略">
                          <el-options 
                          
                          />
                        </el-select>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8" class="form-item">
                      <el-form-item label="更新策略">
                        <el-select placeholder="请选择更新策略">
                          <el-options 
                          
                          />
                        </el-select>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8" class="form-item">
                      <div style="display: flex;justify-content: space-between;">
                        <el-form-item label="最大不可用" label-width="100px">
                          <el-input placeholder="" style="width: 100px"/>
                        </el-form-item>
                        <el-form-item label="可超出最大值" label-width="100px">
                          <el-input placeholder="" style="width: 100px"/>
                        </el-form-item>
                      </div>
                      
                    </el-col>
                    <!-- 第三行 -->
                    <el-col :span="8" class="form-item">
                      <el-form-item label="使用宿主机网络">
                        <el-switch
                          v-model="switchHostNetwork"
                          style="--el-switch-on-color: #13ce66;"
                        />
                      </el-form-item>
                    </el-col>
                    <el-col :span="8" class="form-item">
                      <el-form-item label="自动添加 Servcie">
                        <el-switch
                          v-model="switchAddService"
                        />
                      </el-form-item>
                    </el-col>
                  </el-row>
                </el-form>
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

<style scoped>
.form-item{
  margin-bottom: 15px;
}
</style>

