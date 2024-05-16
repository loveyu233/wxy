import request from "./request.ts";
import {UserLoginReq, UserLoginResp} from "../custom_type/user.ts";
import {publicType} from "../custom_type/public.ts";

enum user {
    login = "/login",
    verify = "/verify"
}

export const loginReq = (userForm:UserLoginReq)=>request.post<string,UserLoginResp>(user.login,userForm)

const config = {
    headers: {
        'Authorization': `Bearer ${localStorage.getItem('token')}`,
    }
};
export const verify = () => request.get<string,publicType>(user.verify,config)