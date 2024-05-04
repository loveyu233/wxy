<script setup lang="ts">
import {onMounted, reactive, ref} from "vue";
import {getRandom} from "../api/random.ts";
import {FormInstance, FormProps} from "element-plus";
import {UserLoginReq} from "../custom_type/user.ts";
import {loginReq} from "../api/user.ts";
import {useRouter} from "vue-router";

const vueRouter = useRouter();

let image = reactive({
  "pic":"",
  "content":"",
})

onMounted( ()=>{
  let token = localStorage.key("token")
  if (token) {
    ElMessage.info("已登陆，需要重新登录请先退出")
    vueRouter.replace("/")
  }
   getRandom().then(res=>{
     image.pic = res.result.pic_url
     image.content = res.result.content + "---" + res.result.author
  })
})

const labelPosition = ref<FormProps['labelPosition']>('top')

let user = reactive(<UserLoginReq>{})
const formRef = ref<FormInstance>()

const clickLogin = (formEl: FormInstance | undefined) => {
  formEl?.validate((v)=>{
    if (v){
      loginReq(user).then(res=>{
        console.log(user)
        console.log(res)
        if(res.status !== 200) {
          ElMessage.error('账号或密码错误')
        } else {
          ElMessage.success("登陆成功")
          localStorage.setItem("token",res.data)
          vueRouter.replace("/")
        }
      })
    }
  })
}

</script>

<template>
  <div class="f" :style="'background-image: url(' + image.pic+ '); background-size:100% 100%;width:100%;height:100vh'">
    <div class="login">
      <el-form
          ref="formRef"
          :label-position="labelPosition"
          label-width="auto"
          :model="user"
          style="max-width: 600px"
      >
        <el-form-item
            label="user_id"
            prop="user_id"
            :rules="[
        { required: true, message: 'age is required' },
        { type: 'number', message: 'age must be a number' },
      ]"
        >
          <el-input
              v-model.number="user.user_id"
              type="text"
              autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="password" prop="password" :rules="[
        { required: true, message: '密码必须填写' },
      ]">
          <el-input v-model="user.password" type="password" autocomplete="off" />
        </el-form-item>
        <el-form-item>
          <el-button class="button" type="primary" @click="clickLogin(formRef)">Submit</el-button>
        </el-form-item>
      </el-form>
    </div>
    <span class="span">{{image.content}}</span>
  </div>
</template>

<style scoped>
.f{
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
}
.span{
  position: absolute; /* 设置绝对定位 */
  top: 10px; /* 距离顶部的距离 */
  right: 10px; /* 距离右侧的距离 */
  color: white; /* 设置文字颜色 */
  padding: 5px 10px; /* 设置内边距 */
  border-radius: 5px; /* 设置圆角 */
}
.login{
  width: 300px; /* 设置div的宽度 */
  height: 200px; /* 设置div的高度 */
  background-color: rgba(255, 255, 255, 0.3); /* 设置背景颜色和透明度 */
  backdrop-filter: blur(10px); /* 添加毛玻璃效果，数值可以调整模糊程度 */
  border-radius: 10px; /* 添加圆角 */
  display: flex;
  justify-content: center;
  align-items: center;
  color: white; /* 设置文字颜色 */
  font-size: 24px; /* 设置文字大小 */
  padding: 50px;
}
.button{
  margin-top: 30px;
  margin: auto;
}
</style>