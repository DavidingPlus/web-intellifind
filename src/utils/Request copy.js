import axios from 'axios'

import {ElLoading, ElFooter} from 'element-plus'
import router from '@/router'

import Message from './Message'

const contentTypeform = 'application/x-www-form-urlencoded;charset=UTF-8'
const contentTypeJson = 'application/json'

let loading = null;
const instance = axios.create({
    baseURL: '/api',
    timeout: 10*1000,
});

//请求前拦截器
instance.interceptors.request.use(
    (config) =>{
        debugger;
        if(config.showloading)
        {
            loading = ElLoading.service(
                {
                    lock:true,
                    text:'加载中....',
                    background:'rgba(0,0,0,0.7)',
                });
        }
        return config;
    },
    (error) =>
    {
        if(config.showloading && loading)
        {
            loading.close();
        }
        Message.error("请求发送失败");
        return Promise.reject("请求发送失败");
    }
);
// 请求后台拦截器
instance.interceptors.response.use(
    (response) =>
    {
        const{showloading,errorCallback,showError = true} = response.config;
        if(showloading && loading)
        {
            loading.close();    
        }
        const responseData = response.data;
        // 正常请求
        if(responseData.code == 200)
        {
            return responseData;
        }else if(responseData.code == 901)
        {
            //登录超时
            setTimeout(() => {
                router.push("/login")
            }, 2000);
            return Promise.reject({showError:true,msg:"登录超时"});
        }else
        {
            //其他错误
            if(errorCallback)
            {
                errorCallback(responseData.info);
            }
            return Promise.reject({showError:true,msg:responseData.info});
        }
    },
    (error)=>
    {
        if(error.config.showloading && loading)
        {
            loading.close();
        }
        return Promise.reject({showError:true,msg:"网络异常"});
    }
);

const request = (config) =>
{
    const{url,params,dataType,showloading = true} = config;
    let contentType = contentTypeForm;
    let formData = new formData();
    for(let key in params)
    {
        formData.append(key,params[key] == undefined ?"":params[key]);
    }
    if(dataType != null && dataType == 'json')
    {
        contentType = contentTypeJson;
    }
    let headers = 
    {
        'Content-Type':contentType,
        'X-Requested-With':'XMLHttpRequest',
    }
    return instance.post(url,formData,
        {
            headers:headers,
            showloading:showloading,
            errorCallback:config.errorCallback,
            showError:config.showError
        }).catch(error=>
            {
                console.log(error);
                if(error.showError)
                {
                    Message.error(error.msg);
                }
                return null;
            });
};

export default request;