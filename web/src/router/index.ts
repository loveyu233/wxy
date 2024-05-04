import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router'

const routes:Array<RouteRecordRaw> = [
    {
      path:"/",
      component:()=>import("../components/Index.vue")
    },
    {
        path:"/login",
        component:()=>import("../components/Login.vue")
    }
]

export default createRouter({
    history:createWebHashHistory(),
    routes
})