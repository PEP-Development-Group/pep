<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="bid字段:"><el-input v-model.number="formData.bid" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="bname字段:">
                <el-input v-model="formData.bname" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="color字段:">
                <el-input v-model="formData.color" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="sex字段:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.sex" clearable ></el-switch>
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
    createBoats,
    updateBoats,
    findBoats
} from "@/api/boats";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Boats",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            bid:0,
            bname:"",
            color:"",
            sex:false,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createBoats(this.formData);
          break;
        case "update":
          res = await updateBoats(this.formData);
          break;
        default:
          res = await createBoats(this.formData);
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
    const res = await findBoats({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reboats
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