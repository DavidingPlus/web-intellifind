import request from '@/utils/requestzxy'


// 获取用户本次解析配置信息
export const getUserConfig = (file_name) => 
    request.get('/core/settings/get', { params: {file_name} })
// 更新用户配置信息
export const updateUserConfig = (data) => 
    request.put('/core/settings/edit_my', data)
// 上传json文件进行解析
export const uploadJsonFile = (formData) => 
    request.post('/core/upload-file', formData, {
        headers: {
            'Content-Type': 'multipart/form-data',
        }})
// 获取单次解析信息 
export const getJsonSolveData = (file_name) => 
    request.get('/core/show-result/once', { params: { file_name } })
// 获取解析日志信息
export const getJsonSolveLogData = (params) => 
    request.get('/core/show-history/total',  {params })

// 删除解析记录信息
export const deleteJsonSolveLogData = (file_name) => 
    request.delete('/core/delete-history/once',  {params: { file_name } })






// // 分类：获取文章分类
// export const artGetChannelsService = () => request.get('/my/cate/list')
// // 分类：添加文章分类
// export const artAddChannelService = (data) => request.post('/my/cate/add', data)
// // 分类：编辑文章分类
// export const artEditChannelService = (data) =>
//   request.put('/my/cate/info', data)
// // 分类：删除文章分类
// export const artDelChannelService = (id) =>
//   request.delete('/my/cate/del', {
//     params: { id }
//   })

// // 文章：获取文章列表
// export const artGetListService = (params) =>
//   request.get('/my/article/list', {
//     params
//   })

// // 文章：添加文章
// // 注意：data需要是一个formData格式的对象
// export const artPublishService = (data) => request.post('/my/article/add', data)

// // 文章：获取文章详情
// export const artGetDetailService = (id) =>
//   request.get('/my/article/info', {
//     params: { id }
//   })

// // 文章：编辑文章接口
// export const artEditService = (data) => request.put('/my/article/info', data)

// // 文章：删除文章接口
// export const artDelService = (id) =>
//   request.delete('/my/article/info', { params: { id } })
