import request from "./request.ts";
import {UserLoginReq, UserLoginResp} from "../custom_type/user.ts";

enum user {
    login = "/login"
}

export const loginReq = (userForm:UserLoginReq)=>request.post<string,UserLoginResp>(user.login,userForm)