<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="学号:"><el-input v-model.number="formData.stuId" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="姓名:">
                <el-input v-model="formData.stuName" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="学院:">
                <el-input v-model="formData.college" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="专业:">
                <el-input v-model="formData.major" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="身份证:">
                <el-input v-model="formData.PID" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="hash密码:">
                <el-input v-model="formData.hash" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="盐:">
                <el-input v-model="formData.salt" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="取消次数:"><el-input v-model.number="formData.cancelNums" clearable placeholder="请输入"></el-input>
          </el-form-item>
           <el-form-item>
           <el-button v-if="this.wf.clazz == 'start'" @click="start" type="primary">启动</el-button>
           <!-- complete传入流转参数 决定下一步会流转到什么位置 此处可以设置多个按钮来做不同的流转 -->
           <el-button v-if="canShow" @click="complete('yes')" type="primary">提交</el-button>
           <el-button v-if="showSelfNode" @click="complete('')" type="primary">确认</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    startWorkflow,
    completeWorkflowMove
} from "@/api/workflowProcess";
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
export default {
  name: "Stu",
  mixins: [infoList],
  props:{
      business:{
         type:Object,
        default:function(){return null}
      },
      wf:{
        type:Object,
        default:function(){return{}}
      },
      move:{
         type:Object,
         default:function(){return{}}
      },
      workflowMoveID:{
        type:[Number,String],
        default:0
      }
   },
  data() {
    return {formData: {
            stuId:0,
            stuName:"",
            college:"",
            major:"",
            PID:"",
            hash:"",
            salt:"",
            cancelNums:0,
            
      }
    };
  },
  computed:{
      showSelfNode(){
         if(this.wf.assignType == "self" && this.move.promoterID == this.userInfo.ID){
             return true
         }else{
             return false
         }
      },
      canShow(){
         if(this.wf.assignType == "user"){
            if(this.wf.assignValue.indexOf(","+this.userInfo.ID+",")>-1 && this.wf.clazz == 'userTask'){
               return true
            }else{
               return false
            }
         }else if(this.wf.assign_type == "authority"){
            if(this.wf.assignValue.indexOf(","+this.userInfo.authorityId+",")>-1 && this.wf.clazz == 'userTask'){
               return true
            }else{
               return false
            }
         }
      },
      ...mapGetters("user", ["userInfo"])
  },
  methods: {
    async start() {
      const res = await startWorkflow({
            business:this.formData,
            wf:{
              workflowMoveID:this.workflowMoveID,
              businessId:0,
              businessType:"stu",
              workflowProcessID:this.wf.workflowProcessID,
              workflowNodeID:this.wf.id,
              promoterID:this.userInfo.ID,
              operatorID:this.userInfo.ID,
              action:"create",
              param:""
              }
          });
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"启动成功"
        })
       this.back()
      }
    },
    async complete(param){
     const res = await completeWorkflowMove({
            business:this.formData,
            wf:{
              workflowMoveID:this.workflowMoveID,
              businessID:this.formData.ID,
              businessType:"stu",
              workflowProcessID:this.wf.workflowProcessID,
              workflowNodeID:this.wf.id,
              promoterID:this.userInfo.ID,
              operatorID:this.userInfo.ID,
              action:"complete",
              param:param
              }
     })
     if(res.code == 0){
       this.$message({
          type:"success",
          message:"提交成功"
       })
       this.back()
     }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
    if(this.business){
     this.formData = this.business
    }
}
};
</script>

<style>
</style>