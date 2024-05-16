<script setup lang="ts">
import {useRouter} from "vue-router";
import {GetInfoResp, GetShowListReq, GetShowListResp, Topic, User} from "../custom_type/user.ts";
import { reactive, ref} from "vue";
import {CancelTopic, GetInfo, GetTopicList, SelectTopic} from "../api/student.ts";
import {
  Check,
  Delete,
  Edit,
} from '@element-plus/icons-vue'
import {TeacherAddTopic, TeacherDeleteTopic, TeacherEditTopic} from "../api/teacher.ts";
const vueRouter = useRouter();

const logout = () => {
  localStorage.removeItem('token')
  vueRouter.replace("/login")
  ElMessage.success("退出登陆成功")
}

let req = reactive(<GetShowListReq>{})
let topicListResp = ref<GetShowListResp>()

let info = ref<GetInfoResp>()

GetTopicList(req).then(res=>{
  topicListResp.value = res
  if(user.type) {
    topicListResp.value.data.topicList = topicListResp.value.data.topicList.filter((item: Topic) => {
      return item.teacher_id === user.user_id;
    });
  }
})

GetInfo().then(res=>{
  info.value = res
})


let jsonInfo = localStorage.getItem("info");
const user:User = JSON.parse(jsonInfo as string)


const Cancel = (id:number) => {
  CancelTopic(id).then(res=>{
    if (res.status !== 200) {
      ElMessage.error("取消失败请重试！")
    } else {
      ElMessage.success("取消成功！")
      GetTopicList(req).then(res=>{
        topicListResp.value = res
      })

      GetInfo().then(res=>{
        info.value = res
      })
    }
  })
}

const Select = (id:number) => {
  SelectTopic(id).then(res=>{
    if (res.status !== 200) {
      ElMessage.error(res.data)
    } else {
      ElMessage.success("选择成功！")
      GetTopicList(req).then(res=>{
        topicListResp.value = res
      })

      GetInfo().then(res=>{
        info.value = res
      })
    }
  })
}

let topic = reactive(<Topic>{})

const EditTopic = (row) => {
  dialogVisible.value = true
  topic = row
  isEdit = true
}

const DeleteTopic = (id:number) => {
  TeacherDeleteTopic(id).then(res=>{
    if (res.status !== 200) {
      ElMessage.error(res.data)
    } else {
      ElMessage.success("删除成功！")
      topicListResp.value.data.topicList = topicListResp.value.data.topicList.filter((item: Topic) => {
        return item.ID !== id;
      });
    }
  })
}

let isEdit = true

const sub = () => {
  if (isEdit) {
    dialogVisible.value = false
    TeacherEditTopic(topic).then(res=>{
      if (res.status !== 200) {
        ElMessage.error(res.data)
      }else {
        ElMessage.success("修改成功！")
      }
    })
  } else {
    TeacherAddTopic(topic).then(res=>{
      if (res.status !== 200) {
        ElMessage.error(res.data)
      }else {
        ElMessage.success("添加成功！")
        topicListResp.value.data.topicList.push(topic)
        topic =  reactive(<Topic>{})
        dialogVisible.value = false
      }
    })
  }
}


const AddTopic = () => {
  topic =  reactive(<Topic>{})
  dialogVisible.value = true
  isEdit = false
}

const dialogVisible = ref(false)

</script>

<template>
  <div class="big">
    <div class="top">
      <div class="left">
        <p>欢迎：<span style="color: aqua">{{user.user_name}}</span> {{user.type?"老师":"同学"}}
          <a @click.prevent="logout" style="cursor: pointer">退出</a> &nbsp;
          <a @click.prevent="AddTopic" style="cursor: pointer" v-if="user.type">添加论文题目</a></p>
      </div>
    </div>
    <div class="info" v-if="!user.type">
      <div>
        <p>学号: {{user.user_id}} 姓名: {{user.user_name}} 学院: {{user.xue_yuan}} 专业: {{user.zhuan_ye}} 年级: {{user.class}}</p>
      </div>
    </div>
    <div class="select" v-if="!user.type">
      <div class="show">
        <p style="text-align: center">已选题目：</p>
        <el-table :data="info?.data" style="width: 100%">
          <el-table-column prop="thesis_name" label="题目名称"/>
          <el-table-column prop="teacher_name" label="指导教师" />
          <el-table-column prop="thesis_type" label="题目类别" />
          <el-table-column prop="nature" label="题目性质"/>
          <el-table-column prop="source" label="题目来源" />
          <el-table-column prop="CreatedAt" label="选题时间" />
          <el-table-column prop="" label="取消选择">
            <template #default="scope">
              <el-button type="danger" :icon="Delete" circle @click="Cancel(scope.row.ID)" />
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
    <div class="show">
      <p style="text-align: center"  v-if="!user.type">待选题目：</p>
      <el-table :data="topicListResp?.data.topicList" style="width: 100%">
        <el-table-column prop="title" label="标题" />
        <el-table-column prop="unit" label="单位" />
        <el-table-column prop="content" label="内容"/>
        <el-table-column prop="thesis_name" label="论文名称" />
        <el-table-column prop="thesis_type" label="论文类型" />
        <el-table-column prop="english_name" label="英文名称" />
        <el-table-column prop="nature" label="题目性质" />
        <el-table-column prop="reference_materials" label="参考文献" />
        <el-table-column prop="source" label="题目来源" />
        <el-table-column prop="special_request" label="特殊要求" />
        <el-table-column prop="student_name" label="选中论文学生名称" v-if="user.type" />
        <el-table-column prop="task_book" label="任务书" />
        <el-table-column prop="teacher_name" label="教师名称" />
        <el-table-column label="选择">
          <template #default="scope">
            <el-button type="success" :icon="Check" circle @click="Select(scope.row.ID)"  v-if="!user.type"/>
            <el-button type="primary" :icon="Edit" circle @click="EditTopic(scope.row)"  v-if="user.type"/>
            <el-button type="danger" :icon="Delete" @click="DeleteTopic(scope.row.ID)" circle  v-if="user.type" />

          </template>
        </el-table-column>
      </el-table>
      <el-dialog v-model="dialogVisible" title="内容修改" width="500" draggable>
        <template #footer>
          <el-form
              label-position="left"
              label-width="auto"
              :model="topic"
              style="max-width: 600px"
          >
            <el-form-item label="内容">
              <el-input v-model="topic.content" />
            </el-form-item>
            <el-form-item label="名称">
              <el-input v-model="topic.thesis_name" />
            </el-form-item>
            <el-form-item label="英文名称">
              <el-input v-model="topic.english_name" />
            </el-form-item>
            <el-form-item label="单位">
              <el-input v-model="topic.unit" />
            </el-form-item>
            <el-form-item label="题目类别">
              <el-input v-model="topic.thesis_type" />
            </el-form-item>
            <el-form-item label="题目性质">
              <el-input v-model="topic.nature" />
            </el-form-item>
            <el-form-item label="参考文献">
              <el-input v-model="topic.reference_materials" />
            </el-form-item>
            <el-form-item label="题目来源">
              <el-input v-model="topic.source" />
            </el-form-item>
            <el-form-item label="特殊要求">
              <el-input v-model="topic.special_request" />
            </el-form-item>
            <el-form-item label="任务书">
              <el-input v-model="topic.task_book" />
            </el-form-item>
          </el-form>
          <div class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="sub()">提交</el-button>
          </div>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<style scoped>
.big{
  width: 100%;
  height: 100vh;
}
.top{
  width: 100%;
  height: 4vh;
  background-color: beige;
}
.info{
  width: 100%;
  height: 2vh;
  margin-top: 10px;
}
.select{
  width: 100%;
}
.show{
  width: 100%;
  height: 100%;
  margin-top: 1vh;
}
.left{
  float: right;
  margin-right: 20px;
  margin-top: 5px;
}
</style>