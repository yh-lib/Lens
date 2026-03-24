import { API_CONFIG } from "../config/index.js"
import request from "./axiosEncap.js"

// 获取集群列表
export const getClusterListHandler = () => {
    return request(API_CONFIG.clusterListApi, {}, 'get', 10000)
}

// 获取集群详情
export const getClusterHandler = (clusterId) => {
    return request(API_CONFIG.clusterGetApi, { clusterId }, 'get', 10000)
}

// 删除集群
export const deleteClusterHandler = (clusterId) => {
    return request(API_CONFIG.clusterDeleteApi, { clusterId }, 'get', 10000)
}

// 添加集群
export const addClusterHandler = (itemFrom) => {
    return request(API_CONFIG.clusterAddApi, itemFrom, 'post', 10000)
}


// 更新集群
export const updateClusterHandler = (clusterInfo) => {
    return request(API_CONFIG.clusterUpdateApi, clusterInfo, 'post', 10000)
}