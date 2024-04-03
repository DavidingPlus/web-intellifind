import router from '@/router'
import axios from 'axios'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/stores/modules/user'
import { ref } from 'vue'

const baseURL = 'http://8.137.100.0:8080'

const userStore = ref()
const instance = axios.create({
  // TODO 1. 基础地址，超时时间
  baseURL,
  timeout: 10000
})

// 请求拦截器
instance.interceptors.request.use(
  async (config) => {
    // TODO 2. 携带token
    if (!userStore.value) {
      userStore.value = useUserStore()}
    if (userStore.value.token) {
      config.headers.Authorization = userStore.value.token
    }
    // config.headers.Authorization = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTMiLCJ1c2VyX25hbWUiOiJtYXgiLCJleHBpcmVfdGltZSI6MjQzMTUzMjA2MiwiZXhwIjoyNDMxNTMyMDYyLCJpYXQiOjE3MTE1MzIwNjIsImlzcyI6InRlc3RfYmFja2VuZCIsIm5iZiI6MTcxMTUzMjA2Mn0.oQoYmYZk48bHOIS9Wg3VV-iluD36UaKKNRal_a-0M1U'
    return config
  },
  (err) => Promise.reject(err)
)

// 响应拦截器
instance.interceptors.response.use(
  (res) => {
    // TODO 4. 摘取核心响应数据
    // if (res.data.code === 0) {
    if (res.data.code === 1) {
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

