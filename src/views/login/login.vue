<template>
    <!-- 最外层的大盒子 -->
    <div class="body">
      <div class="box">
        <!-- 滑动盒子 -->
        <div class="pre-box" ref="preRef">
          <h1>WELCOME</h1>
          <p>JOIN US!</p>
          <div class="img-box">
            <img :src="flag == true ? imgList[1] : imgList[0]" alt="" />
          </div>
        </div>
        <!-- 注册盒子 -->
        <div class="register-form">
          <!-- 标题盒子 -->
          <div class="title-box">
            <h1>注册</h1>
          </div>
          <!-- 输入框盒子 -->
          <el-form
            ref="registerFormRef"
            :model="registerForm"
            :rules="rules"
            label-width="5px"
            class="login-container"
          >
            <el-form-item prop="username">
              <el-input
                type="text"
                placeholder="请输入用户名"
                :suffix-icon="User"
                v-model="registerForm.username"
              />
            </el-form-item>
            <el-form-item prop="password">
              <el-input
                type="password"
                placeholder="请输入密码"
                :suffix-icon="Lock"
                v-model="registerForm.password"
              />
            </el-form-item>
            <el-form-item prop="email">
              <div>
              <el-input
                type="text"
                placeholder="请输入邮箱"
                :suffix-icon="Email"
                v-model="registerForm.email"
              />
              <el-form-item>
                  <el-button @click="sendEmail" class="send-mail-btn" size = "large" :disabled="countdown > 0">{{ countdown > 0 ? `${countdown}秒后重试` : '发送' }}</el-button>
              </el-form-item>
            </div>
            </el-form-item>
            <el-form-item prop="verify_code">
              <el-input
                placeholder="请输入邮箱验证码"
                :suffix-icon="Lock"
                v-model="registerForm.verify_code"
              />
            </el-form-item>
          </el-form>
  
          <!-- 按钮盒子 -->
          <div class="btn-box">
            <button @click="register">注册</button>
            <!-- 绑定点击事件 -->
            <p @click="mySwitch">已有账号?去登录</p>
          </div>
        </div>
        <!-- 登录盒子 -->
        <div class="login-form">
          <!-- 标题盒子 -->
          <div class="title-box">
            <h1>登录</h1>
          </div>
          <!-- 输入框盒子 -->
          <el-form
            ref="loginFormRef"
            :model="loginForm"
            :rules="rules"
            label-width="5px"
            class="login-container"
          >
            <el-form-item prop="username">
              <el-input
                type="text"
                placeholder="用户名"
                :suffix-icon="User"
                v-model="loginForm.username"
              />
            </el-form-item>
            <el-form-item prop="password">
              <el-input
                type="password"
                placeholder="密码"
                :suffix-icon="Lock"
                v-model="loginForm.password"
              />
            </el-form-item>
          </el-form>
          <!-- 按钮盒子 -->
          <div class="btn-box">
            <button @click="login">登录</button>
            <!-- 绑定点击事件 -->
            <p @click="mySwitch">没有账号?去注册</p>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import api from '@/api/main.js'
import loginImg from '@/assets/img/login.jpg'
import registerImg from '@/assets/img/register.jpg'
import { Lock, Phone, User } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
  const preRef = ref('')
  const imgList = ref([ registerImg, loginImg ])
  let flag = ref(true)
  /* 登录框与注册框滑动切换效果 */
  const mySwitch = () => {
    if (flag.value) {
      preRef.value.style.background = 'rgb(191, 227, 241)'
      preRef.value.style.transform = 'translateX(100%)'
    } else {
      preRef.value.style.background = 'rgb(247, 243, 209)'
      preRef.value.style.transform = 'translateX(0%)'
    }
    flag.value = !flag.value
  }
  // 登录表单
  const loginForm = reactive({
    // username: 'znz',
    // password: 'abcd',
    username: '',
    password: '',
  })
  // const username = loginForm.username
  // const password = loginForm.password
  // 注册表单
  const registerForm = reactive({
    username: '',
    password: '',
    tel: '',
  })
  const loginFormRef = ref('')
  const registerFormRef = ref('')
  // 检验规则
  const rules = reactive({
    username: [
      { required: true, message: '请输入用户名', trigger: 'blur' },
      {
        min: 3,
        max: 10,
        message: '用户名长度需要在3~10个字符之间',
        trigger: 'blur',
      },
    ],
    password: [
      { required: true, message: '请输入密码', trigger: 'blur' },
      { min: 4, max: 20, message: '请输入正确的密码', trigger: 'blur' },
    ],
    tel: [
      { required: true, message: '请输入电话号码', trigger: 'blur' },
      { min: 8, max: 11, message: '请输入正确的电话号码', trigger: 'blur' },
    ],
  })
  const router = useRouter()
  const login = () => {
    loginFormRef.value.validate((valid) => {
      if (valid) {
        api
          .loginApi(loginForm)
          .then((res) => {
            console.log('login', res)
            const username = loginForm.username
            const password = loginForm.password
            if (res.flag == true) {
              ElMessage.success(res.message)
              window.sessionStorage.setItem('token', res.token)
              router.push({
                path: '/illegaldetail',
                // path: '/user/profile',
                query: { username, password },
              }) //登录成功后跳转
            }
          })
          .catch((error) => {
            console.log(error)
          })
      } else {
        return
      }
    })
  }
  
  const register = () => {
    registerFormRef.value.validate((valid) => {
      if (valid) {
        api
          .registerApi(registerForm)
          .then((res) => {
            if (res.flag == true) {
              ElMessage.success(res.message)
            }
          })
          .catch((error) => {
            console.log(error)
          })
      } else {
        return
      }
    })
  }
  </script>
  
  <style lang="less" scoped>
  /* 去除input的轮廓 */
  input {
    outline: none;
  }
  
  .body {
    /* 溢出隐藏 */
    height: 100vh;
    overflow-x: hidden;
    display: flex;
    /* 渐变方向从左到右 */
    background: linear-gradient(to right, rgb(247, 243, 209), rgb(191, 227, 241));
  }
  
  /* 最外层的大盒子 */
  .box {
    width: 1050px;
    height: 600px;
    display: flex;
    /* 相对定位 */
    position: relative;
    z-index: 2;
    margin: auto;
    /* 设置圆角 */
    border-radius: 8px;
    /* 设置边框 */
    border: 1px solid rgba(255, 255, 255, 0.6);
    /* 设置盒子阴影 */
    box-shadow: 2px 1px 19px rgba(0, 0, 0, 0.1);
  }
  
  /* 滑动的盒子 */
  .pre-box {
    /* 宽度为大盒子的一半 */
    width: 50%;
    /* width: var(--width); */
    height: 100%;
    /* 绝对定位 */
    position: absolute;
    /* 距离大盒子左侧为0 */
    left: 0;
    /* 距离大盒子顶部为0 */
    top: 0;
    z-index: 99;
    border-radius: 4px;
    background-color: rgb(247, 243, 209);
    box-shadow: 2px 1px 19px rgba(0, 0, 0, 0.1);
    /* 动画过渡，先加速再减速 */
    transition: 0.5s ease-in-out;
  }
  
  /* 滑动盒子的标题 */
  .pre-box h1 {
    margin-top: 150px;
    text-align: center;
    /* 文字间距 */
    letter-spacing: 5px;
    color: white;
    /* 禁止选中 */
    user-select: none;
    /* 文字阴影 */
    text-shadow: 4px 4px 3px rgba(0, 0, 0, 0.1);
  }
  
  /* 滑动盒子的文字 */
  .pre-box p {
    height: 30px;
    line-height: 30px;
    text-align: center;
    margin: 20px 0;
    /* 禁止选中 */
    user-select: none;
    font-weight: bold;
    color: white;
    text-shadow: 4px 4px 3px rgba(0, 0, 0, 0.1);
  }
  
  /* 图片盒子 */
  .img-box {
    width: 200px;
    height: 200px;
    margin: 20px auto;
    /* 设置为圆形 */
    border-radius: 50%;
    /* 设置用户禁止选中 */
    user-select: none;
    overflow: hidden;
    box-shadow: 4px 4px 3px rgba(0, 0, 0, 0.1);
  }
  
  /* 图片 */
  .img-box img {
    width: 100%;
    transition: 0.5s;
  }
  
  /* 登录和注册盒子 */
  .login-form,
  .register-form {
    flex: 1;
    height: 100%;
  }
  /* 登录盒子 */
  .login-container {
    width: 500px;
    background-color: rgba(255, 255, 255, 0);
    border-radius: 15px;
    padding: 0px 35px 15px 35px;
    margin: 0px auto;
    :deep(.el-input__wrapper) {
      background-color: rgba(255, 255, 0, 0);
      border: 1px solid #000;
    }
  }
  
  /* 标题盒子 */
  .title-box {
    height: 250px;
    line-height: 400px;
  }
  
  /* 标题 */
  .title-box h1 {
    text-align: center;
    color: white;
    /* 禁止选中 */
    user-select: none;
    letter-spacing: 5px;
    text-shadow: 4px 4px 3px rgba(0, 0, 0, 0.1);
  }
  
  /* 输入框盒子 */
  .el-form {
    display: flex;
    /* 纵向布局 */
    flex-direction: column;
    /* 水平居中 */
    align-items: center;
  }
  
  /* 输入框 */
  .el-form-item {
    width: 65%;
    height: 40px;
  }
  
  input {
    /* width: 60%; */
    /* height: 40px; */
    margin-bottom: 20px;
    text-indent: 10px;
    border: 1px solid #fff;
    background-color: rgba(255, 255, 255, 0.3);
    border-radius: 120px;
    /* 增加磨砂质感 */
    backdrop-filter: blur(10px);
  }
  
  input:focus {
    /* 光标颜色 */
    color: #b0cfe9;
  }
  
  /* 聚焦时隐藏文字 */
  input:focus::placeholder {
    opacity: 0;
  }
  
  /* 按钮盒子 */
  .btn-box {
    display: flex;
    justify-content: center;
  }
  
  /* 按钮 */
  button {
    width: 100px;
    height: 30px;
    margin: 0 7px;
    line-height: 30px;
    border: none;
    border-radius: 4px;
    background-color: #69b3f0;
    color: white;
  }
  
  /* 按钮悬停时 */
  button:hover {
    /* 鼠标小手 */
    cursor: pointer;
    /* 透明度 */
    opacity: 0.8;
  }
  
  /* 按钮文字 */
  .btn-box p {
    height: 30px;
    line-height: 30px;
    /* 禁止选中 */
    user-select: none;
    font-size: 14px;
    color: white;
  }
  
  .btn-box p:hover {
    cursor: pointer;
    border-bottom: 1px solid white;
  }
  </style>
  