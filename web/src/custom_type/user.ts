import {publicType} from "./public.ts";

export type UserLoginReq = {
    user_id:number,
    password:string,
}

export type User = {
    ID:number
    class:string,
    xue_yuan:string,
    zhuan_ye:string
    type:boolean,
    user_id:number,
    user_name:string
}

export type UserLoginResp = {
    data:{
        msg:string
        info:User,
        token:string
    }
} & publicType

export type GetShowListReq = {
    pageNum:number,
    pageSize:number
}

export type Topic = {
    ID:number
    CreatedAt:string
    thesis_name: string;
    english_name: string;
    teacher_name: string;
    teacher_id: number;
    title: string;
    unit: string;
    thesis_type: string;
    nature: string;
    source: string;
    content: string;
    task_book: string;
    special_request: string;
    reference_materials: string;
    is_select: boolean;
    student_id: number;
}

export type GetShowListResp = {
    data:{
        total:number,
        topicList:Array<Topic>
    }
}&publicType


export type GetInfoResp = {
    data:Array<Topic>
}&publicType


export type Cancel = {
    data:string
}&publicType