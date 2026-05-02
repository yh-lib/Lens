import { API_CONFIG } from '../config/index.js'
import request from './axiosEncap.js'

// 获取 service 列表
export const getServiceListHandler = (clusterId, nameSpace) => {
  return request(API_CONFIG.serviceListApi, { clusterId, nameSpace }, 'get')
}

// 获取 service 详情
export const getServiceHandler = (clusterId, nameSpace, name) => {
  return request(API_CONFIG.serviceGetApi, { clusterId, nameSpace, name }, 'get')
}

// 删除 service
export const deleteServiceHandler = (clusterId, nameSpace, name) => {
  return request(API_CONFIG.serviceDeleteApi, { clusterId, nameSpace, name }, 'post')
}

// 添加 service
export const createServiceHandler = (clusterId, item) => {
  return request(API_CONFIG.serviceCreateApi, { clusterId, item }, 'post')
}

// 更新 service
export const updateServiceHandler = (serviceInfo) => {
  return request(API_CONFIG.serviceUpdateApi, serviceInfo, 'post')
}
