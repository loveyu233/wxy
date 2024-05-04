import {publicType} from "./public.ts";

export type UserLoginReq = {
    user_id:number,
    password:string,
}

export type UserLoginResp = {
    data:string
} & publicType