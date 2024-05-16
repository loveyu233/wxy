import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router'
import Index from "../components/Index.vue";
import Login from "../components/Login.vue";
import {verify} from "../api/user.ts";

const routes:Array<RouteRecordRaw> = [
    {
      path:"/",
      component: Index,
    },
    {
        path:"/login",
        component:Login,
        meta:{
            "skipAuth":true,
        }
    }
]

const router =  createRouter({
    history:createWebHashHistory(),
    routes
})

router.beforeEach((to, _, next) => {
    if (to.meta.skipAuth) {
        verify().then(res=>{
            if (res.status !== 200){
                localStorage.removeItem("token");
                router.replace("/login")
            }
        })
    }
    next()
})

export default router