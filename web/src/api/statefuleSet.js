import { API_CONFIG } from '../config/index.js'
import request from './axiosEncap.js'

// 获取 statefulSet 列表
export const getStatefulSetListHandler = (clusterId, nameSpace) => {
  return request(API_CONFIG.statefulSetListApi, { clusterId, nameSpace }, 'get')
}

// 获取 statefulSet 详情
export const getStatefulSetHandler = (clusterId, nameSpace, name) => {
  return request(API_CONFIG.statefulSetGetApi, { clusterId, nameSpace, name }, 'get')
}

// 删除 statefulSet
export const deleteStatefulSetHandler = (clusterId, nameSpace, name) => {
  return request(API_CONFIG.statefulSetDeleteApi, { clusterId, nameSpace, name }, 'post')
}

// 创建 statefulSet
export const createStatefulSetHandler = (itemForm) => {
  return request(API_CONFIG.statefulSetCreateApi, itemForm, 'post')
}

// 更新 statefulSet
export const updateStatefulSetHandler = (itemFrom) => {
  return request(API_CONFIG.statefulSetUpdateApi, itemFrom, 'post')
}
