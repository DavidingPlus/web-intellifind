import request from '@/utils/requestzxy'


// // 注册接口
// export const userRegisterService = ({ username, password, repassword }) =>
//   request.post('/api/reg', { username, password, repassword })

// // 登录接口
// export const userLoginService = ({ username, password }) =>
//   request.post('/auth/login', { username, password })

// 获取用户基本信息
export const userGetInfoService = () => 
  request.get('/user/info' )

// 更新用户基本信息
export const userUpdateInfoService = ({ email, user_name, city, Tel, Birthday, Gender, create_time }) =>
  request.put('/user/update', { email, user_name, city, Tel, Birthday, Gender, create_time })

// 更新用户头像
export const userUpdateAvatarService = (formData) =>
  request.post('/user/upload-avatar', formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })

// // 更新用户密码
// export const userUpdatePasswordService = ({ old_pwd, new_pwd, re_pwd }) =>
//   request.patch('/my/updatepwd', { old_pwd, new_pwd, re_pwd })
