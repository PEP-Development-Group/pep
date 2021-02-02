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
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createStu,
    updateStu,
    findStu
} from "@/api/stu";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Stu",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
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
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createStu(this.formData);
          break;
        case "update":
          res = await updateStu(this.formData);
          break;
        default:
          res = await createStu(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findStu({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.restu
       this.type == "update"
     }
    }else{
     this.type == "create"
   }
  
}
};
</script>

<style>
</style>