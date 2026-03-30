// 封装 axios
/*
    1. 封装可以简化我的接口调用的代码，可以去掉一些重复的代码
    2. 换掉 Axios 也是比较简单的
*/

import axios from 'axios'
import { API_CONFIG, CONFIG } from '../config/index.js'
import router from '../router/index.js'
import { ElMessage } from 'element-plus'

// axios 全局配置
axios.defaults.baseURL = API_CONFIG.baseUrl

// axios 拦截器
// 添加请求拦截器
axios.interceptors.request.use(
    (config) => {
        // 在发送请求之前做些什么
        // 添加请求时的时间戳，处理浏览器缓存问题
        if (config.method == 'get') {
            let timeStamp = (new Date()).getTime()
            config.params ? config.params.timeStamp = timeStamp : config.params = { timeStamp: timeStamp }
        }
        // 在request header中加入token
        config.headers.Authorization = window.localStorage.getItem(CONFIG.TOKEN_NAME)
        return config;
    },
    (error) => {
        // 对请求错误做些什么
        return Promise.reject(error);
    }
);

// 添加响应拦截器
axios.interceptors.response.use(
    (response) => {
        // 2xx 范围内的状态码都会触发该函数。
        // 对响应数据做点什么
        console.log("响应拦截器:::Response:::", response.data.status)
        if (response.data.status === 200) {
            return response;
        } else if (response.data.status === 401) {
            // 提示信息
            ElMessage({
                message: response.data.message,
                type: 'warning',
            })
            // token已失效,移除本地token
            window.localStorage.removeItem(CONFIG.TOKEN_NAME)
            // 跳转到登录页
            router.currentRoute.value.path != '/login' && router.push('/login')
            return Promise.reject(new Error(response.data.message));
        } else if (response.data.status === 403) {
            // 提示信息
            console.log("响应拦截器403:::Response:::", response.data.status)
            ElMessage.error({ message: response.data.message })
            return Promise.reject(new Error(response.data.message));
        }
    }
    // (error) => {
    //     // 超出 2xx 范围的状态码都会触发该函数。
    //     // 对响应错误做点什么
    //     if (error.response) {
    //         // 请求已发出，但服务器响应的状态码不在 2xx 范围内
    //         ElMessage.error(error.response.data.message || 'An error occurred') // 修改：显示错误信息
    //         return Promise.reject(new Error(error.response.data.message || 'An error occurred')); // 修改：返回 Promise.reject 并传递错误信息
    //     } else if (error.request) {
    //         // 请求已发出，但没有收到响应
    //         ElMessage.error('No response received') // 修改：显示错误信息
    //         return Promise.reject(new Error('No response received')); // 修改：返回 Promise.reject 并传递错误信息
    //     } else {
    //         // 发生了一些设置请求时发生的问题
    //         ElMessage.error(error.message) // 修改：显示错误信息
    //         return Promise.reject(new Error(error.message)); // 修改：返回 Promise.reject 并传递错误信息
    //     }
    // }
);

const request = (url = '', data = {}, method = 'get', timeout = 3000) => {
    return new Promise(
        (resolve, reject) => {
            // GET POST
            const methodLower = method.toLowerCase()
            if (methodLower == 'get') {
                axios({
                    method: methodLower,
                    params: data,
                    timeout: timeout,
                    url: url,
                }).then((response) => {
                    // 能正常拿到数据
                    resolve(response)
                }).catch((error) => {
                    reject(error)
                })
            } else if (methodLower === "post") {
                axios({
                    method: methodLower,
                    data: data,
                    timeout: timeout,
                    url: url,
                }).then((response) => {
                    // 能正常拿到数据
                    resolve(response)
                }).catch((error) => {
                    reject(error)
                })
            }
        }
    )
}

export default request