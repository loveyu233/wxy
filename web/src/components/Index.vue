<script setup lang="ts">
import {useRouter} from "vue-router";

const vueRouter = useRouter();

const logout = () => {
  localStorage.removeItem('token')
  vueRouter.replace("/login")
  ElMessage.success("退出登陆成功")
}
</script>

<template>
  <el-button type="warning" @click="logout">Warning</el-button>

</template>

<style scoped>

</style>