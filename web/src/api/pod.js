import { API_CONFIG } from "../config/index.js"
import request from "./axiosEncap.js"

// 获取 namespace 列表
export const getPodListHandler = (clusterId, nameSpace) => {
    return request(API_CONFIG.podListApi, { clusterId, nameSpace }, 'get', 10000)
}

// // 获取 namespace 详情
// export const getNamespaceHandler = (clusterId, name) => {
//     return request(API_CONFIG.namespaceGetApi, { clusterId, name }, 'get', 10000)
// }

// // 删除 namespace
// export const deleteNamespaceHandler = (clusterId, name) => {
//     return request(API_CONFIG.namespaceDeleteApi, { clusterId, name }, 'get', 10000)
// }

// // 添加 namespace
// export const createNamespaceHandler = (clusterId, item) => {
//     return request(API_CONFIG.namespaceCreateApi, { clusterId, item }, 'post', 10000)
// }


// // 更新 namespace
// export const updateNamespaceHandler = (NamespaceInfo) => {
//     return request(API_CONFIG.namespaceUpdateApi, NamespaceInfo, 'post', 10000)
// }