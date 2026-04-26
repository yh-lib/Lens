<script setup>
  import { ElMessage } from 'element-plus'
  import { ref, reactive } from 'vue'
  import TabOfContainerDetail from './TabOfContainerDetail.vue'
  import { useWorkLoadData } from '../../../../store'
  import { storeToRefs } from 'pinia'

  // store from pinia
  const store = useWorkLoadData()
  const { workLoadItem } = storeToRefs(store)

  const props = defineProps(['containerType'])

  // 定义初始状态
  const activeTabsValue = ref(0)

  const data = reactive({
    containerItem: workLoadItem.value.item.spec.template.spec.containers,
  })

  // 添加新容器的item模板
  const getContainerTemplate = (newTabIndex) => {
    return {
      name: `Container-${newTabIndex}`,
      ports: [],
      env: [],
      envFrom: [],
      volumeMounts: [],
      resources: {
        requests: {},
        limits: {},
      },
      startupProbe: {
        // 通用参数
        initialDelaySeconds: 0,
        periodSeconds: 10,
        timeoutSeconds: 1,
        successThreshold: 1,
        failureThreshold: 30,
        terminationGracePeriodSeconds: 30,
        // 探测方法
        exec: { command: [] },
        tcpSocket: {},
        httpGet: {
          httpHeaders: [],
        },
        grpc: {},
      },
      readinessProbe: {
        // 通用参数
        initialDelaySeconds: 0,
        periodSeconds: 10,
        timeoutSeconds: 1,
        successThreshold: 1,
        failureThreshold: 30,
        terminationGracePeriodSeconds: 30,
        // 探测方法
        exec: { command: [] },
        tcpSocket: {},
        httpGet: {
          httpHeaders: [],
        },
        grpc: {},
      },
      livenessProbe: {
        // 通用参数
        initialDelaySeconds: 0,
        periodSeconds: 10,
        timeoutSeconds: 1,
        successThreshold: 1,
        failureThreshold: 30,
        terminationGracePeriodSeconds: 30,
        // 探测方法
        exec: { command: [] },
        tcpSocket: {},
        httpGet: {
          httpHeaders: [],
        },
        grpc: {},
      },
      lifecycle: {
        postStart: {},
        preStop: {},
      },
    }
  }

  const handleTabsEdit = (targetName, action) => {
    if (action === 'add') {
      const newTabIndex = data.containerItem.length
      const newContainer = getContainerTemplate(newTabIndex)
      data.containerItem.push(newContainer)
      activeTabsValue.value = newTabIndex
    } else if (action === 'remove') {
      // 如果当前只有一个容器，提示并中止后续删除逻辑
      if (data.containerItem.length === 1 && props.containerType == 'container') {
        ElMessage.error('至少需要保留一个容器')
        return
      }
      const tabs = data.containerItem
      let activeName = activeTabsValue.value

      if (activeName === targetName) {
        tabs.forEach((tab, index) => {
          if (index === targetName) {
            const nextTab = tabs[index + 1] || tabs[index - 1]
            if (nextTab) {
              activeName = tabs.indexOf(nextTab)
            }
          }
        })
      }

      activeTabsValue.value = activeName
      if (Number.isInteger(targetName)) {
        tabs.splice(targetName, 1)
      }
    }
  }
</script>

<template>
  <el-tabs
    v-model="activeTabsValue"
    type="card"
    editable
    @edit="handleTabsEdit"
    style="margin-top: 10px; height: 557px"
  >
    <el-tab-pane
      v-for="(item, index) in data.containerItem"
      :key="item.name"
      :label="item.name"
      :name="index"
    >
      <tab-of-container-detail
        :container-item="data.containerItem[index]"
        :container-type="props.containerType"
      />
    </el-tab-pane>
  </el-tabs>
</template>
