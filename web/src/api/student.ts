import request from "./request.ts";
import {Cancel, GetInfoResp, GetShowListReq, GetShowListResp} from "../custom_type/user.ts";

enum Student {
    SelectTopic = "/student/select",
    ShowTopic = "/student/show",
    InfoTopic = "/student/info",
    CancelTopic = "/student/cancel",
}

export const GetTopicList = (param:GetShowListReq) =>request.post<string,GetShowListResp>(Student.ShowTopic,param)

export const GetInfo = ()=>request.get<string,GetInfoResp>(Student.InfoTopic)

export const CancelTopic = (id:number)=>request.get<string,Cancel>(Student.CancelTopic,{params:{id:id}})

export const SelectTopic = (id:number)=>request.get<string,Cancel>(Student.SelectTopic,{params:{id:id}})
