import router from '@/router'
import { useUserStore } from '@/stores'
import axios from 'axios'
import { ElMessage } from 'element-plus'
// const baseURL = 'http://big-event-vue-api-t.itheima.net'
// const baseURL = 'http://baidu.com'
const baseURL = 'http://8.137.100.0:8080'


const instance = axios.create({
  // TODO 1. 基础地址，超时时间
  baseURL,
  timeout: 10000
})

// 请求拦截器
instance.interceptors.request.use(
  (config) => {
    // TODO 2. 携带token
    // const useStore = useUserStore()
    // if (useStore.token) {
    //   config.headers.Authorization = useStore.token
    // }
    config.headers.Authorization = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTMiLCJ1c2VyX25hbWUiOiJtYXgiLCJleHBpcmVfdGltZSI6MjQyOTU2MzgwMSwiZXhwIjoyNDI5NTYzODAxLCJpYXQiOjE3MDk1NjM4MDEsImlzcyI6InRlc3RfYmFja2VuZCIsIm5iZiI6MTcwOTU2MzgwMX0.V9Zwvwsy1dIc74T_fqj5qYyolFQWYk0_47OTzDa4CjQ'
    return config
  },
  (err) => Promise.reject(err)
)

// 响应拦截器
instance.interceptors.response.use(
  (res) => {
    // TODO 4. 摘取核心响应数据
    if (res.data.code === 0) {
      return res
    }
    // TODO 3. 处理业务失败
    // 处理业务失败, 给错误提示，抛出错误
    ElMessage.error(res.data.message || '服务异常')
    return Promise.reject(res.data)
  },
  (err) => {
    // TODO 5. 处理401错误
    // 错误的特殊情况 => 401 权限不足 或 token 过期 => 拦截到登录
    if (err.response?.status === 401) {
      router.push('/login')
    }

    // 错误的默认情况 => 只要给提示
    // ElMessage.error(err.response.data.message || '服务异常')
    return Promise.reject(err)
  }
)

export default instance
export { baseURL }

