<template>
  <div>
    <el-form :model="config" label-width="100px">
      <el-form-item label="最大退课次数">
        <el-input-number v-model="config.system.cancel_nums"></el-input-number>
      </el-form-item>
      <el-form-item label="学期开始时间">
        <el-date-picker v-model="config.system.first_day" value-format="yyyy-MM-dd"></el-date-picker>
        <div style="line-height: 15px"> * 请选择本学期的第一周周一,将作为课程系统时间计算依据<br>* 请在添加课程之前修改此项目</div>
      </el-form-item>
      <el-form-item>
        <el-button @click="update" type="primary">提交</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import {getSystemConfig, setSystemConfig} from "@/api/system";

export default {
  name: "courseConfig",
  data() {
    return {
      config: {
        system: {}
      }
    }
  }, async created() {
    await this.initForm();
  },
  methods: {
    async initForm() {
      const res = await getSystemConfig();
      if (res.code == 0) {
        this.config = res.data.config;
      }
    },
    reload() {
    },
    async update() {
      const res = await setSystemConfig({config: this.config});
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "配置文件设置成功"
        });
        await this.initForm();
      }
    }
  }
}
</script>

<style scoped>

</style>