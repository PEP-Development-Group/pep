<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="学分:"><el-input v-model.number="formData.ccredit" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="课程名:">
                <el-input v-model="formData.cname" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="选课结束:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.etime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="选课开始:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.stime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="上课时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.time" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="教师名:">
                <el-input v-model="formData.tname" clearable placeholder="请输入" ></el-input>
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
    createClass,
    updateClass,
    findClass
} from "@/api/cls_class";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Class",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            ccredit:0,
            cname:"",
            etime:new Date(),
            stime:new Date(),
            time:new Date(),
            tname:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createClass(this.formData);
          break;
        case "update":
          res = await updateClass(this.formData);
          break;
        default:
          res = await createClass(this.formData);
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
    const res = await findClass({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reclass
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