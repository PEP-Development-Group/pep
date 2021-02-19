<template>
  <div>
    <el-form :model="formData" label-width="100px">
      <el-form-item prop="cancelTimes" label="最大退课次数">
        <el-input-number></el-input-number>
      </el-form-item>
      <el-form-item prop="startTime" label="学期开始时间">
        <el-date-picker></el-date-picker>
        <span> * 请选择本学期的第一周周一,将作为课程系统时间计算依据</span>
      </el-form-item>
      <el-form-item>
        <el-button type="primary">提交</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: "courseConfig",
  data() {
    return {
      formData: {
        cancelTimes: null,
        startTime: null
      }
    }
  }
}
</script>

<style scoped>

</style>