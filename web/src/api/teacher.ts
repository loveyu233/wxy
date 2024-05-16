import {Topic} from "../custom_type/user.ts";
import request from "./request.ts";

enum Teacher {
    AddTopic = "/teacher/add",
    EditTopic = "/teacher/edit",
    DeleteTopic = "/teacher/delete",
    SearchTopic = "/teacher/search",
}

export const TeacherEditTopic = (topic:Topic) => request.post<string,Topic>(Teacher.EditTopic,topic)

export const TeacherDeleteTopic = (id:number) => request.get(Teacher.DeleteTopic,{params:{id:id}})

export const TeacherAddTopic = (topic:Topic) => request.post<string,Topic>(Teacher.AddTopic,topic)
