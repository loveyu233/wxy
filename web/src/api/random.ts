import request from "./request.ts";
import {RandomRespType} from "../custom_type/random.ts";

export const mryy = "https://api.oioweb.cn/api/common/yiyan"

enum Random {
    yiyan = "https://api.oioweb.cn/api/common/yiyan",
    suijitupian = "https://bing.img.run/rand.php",
}

export const getRandom = ()=>request.get<string,RandomRespType>(Random.yiyan)


export const suijitupian = ()=>request.get(Random.suijitupian)