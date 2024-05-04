import {defineStore} from 'pinia'

export const use_index_store = defineStore("index",{
    state() {
        return {
            user:{
                name:"",
                age:0
            },
        }
    },
    getters:{
        getUser():string{
            return this.user.name +"--"+this.user.age
        }
    },
    actions:{
        updateUserName(name:string){
            this.user.name = name
            this.user.age++
        }
    }
})