import axios from "axios";
import {ElMessage} from "element-plus";

const request = axios.create({
    baseURL: import.meta.env.VITE_BASE_URL,
    timeout: 2000
});

request.interceptors.response.use((res) => {
    return res.data;
}, error => {
    if (error && error.response) {
        // 1.公共错误处理
        // 2.根据响应码具体处理
        switch (error.response.status) {
            case 400:
                error.message = "错误请求";
                break;
            case 401:
                error.message = "未授权，请重新登录";
                break;
            case 403:
                error.message = "拒绝访问";
                break;
            case 404:
                error.message = "请求错误,未找到该资源";
                window.location.href = "/NotFound";
                break;
            case 405:
                error.message = "请求方法未允许";
                break;
            case 408:
                error.message = "请求超时";
                break;
            case 500:
                error.message = "服务器端出错";
                break;
            case 501:
                error.message = "网络未实现";
                break;
            case 502:
                error.message = "网络错误";
                break;
            case 503:
                error.message = "服务不可用";
                break;
            case 504:
                error.message = "网络超时";
                break;
            case 505:
                error.message = "http版本不支持该请求";
                break;
            default:
                error.message = `连接错误${error.response.status}`;
        }
    } else {
        // 超时处理
        if (JSON.stringify(error).includes("timeout")) {
            ElMessage.error("服务器响应超时，请刷新当前页");
        }
        console.log(error);
        error.message = "连接服务器失败";
    }
    ElMessage.error(error.message);
    /***** 处理结束 *****/
    //如果不需要错误处理，以上的处理过程都可省略
    return Promise.resolve(error.response);
});


// request.interceptors.request.use((config) => {
//     let token = GET_TOKEN();
//     if (token) {
//         config.headers.token = token;
//     }
//     return config;
// });


export default request;