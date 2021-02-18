<template>
  <div>
    <span v-if="grade<=100">
            <i class="el-icon-lock"></i>
            {{ grade }}
          </span>
    <span v-else-if="grade===101">
      <el-tag type="danger" effect="dark" size="mini">旷课</el-tag>
    </span>
    <span v-else-if="grade===102">
    <el-form inline :model="gradeForm">
      <span class="icon-middle">
        <i :class="{'el-icon-unlock':!isLoading,'el-icon-loading':isLoading}"></i>
      </span>
      <!--TODO 旷课选项-->
      <el-form-item prop="grade">
        <el-input v-model.number="gradeForm.grade" type="tel" @keyup.enter.native="submitGrade"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button :disabled="gradeValid" @click="submitGrade">提交</el-button>
      </el-form-item>
    </el-form>
    </span>
  </div>
</template>

<script>
import {setStuGrade} from "@/api/course";

export default {
  name: "gradeForm",
  props: {
    username: String,
    cid: Number,
    grade: Number
  },
  computed: {
    gradeValid: function () {
      return this.gradeForm.grade === '' || !isFinite(this.gradeForm.grade) || this.gradeForm.grade < 0 || this.gradeForm.grade > 100
    }
  },
  data() {
    return {
      isLoading: false,
      gradeForm: {
        cid: 0,
        username: null,
        grade: ''
      },
    }
  },
  methods: {
    submitGrade() {
      this.isLoading = true
      setStuGrade(this.gradeForm).then(() => {
        this.grade = this.gradeForm.grade
        this.isLoading = false
      })

    }
  }, created() {
    this.gradeForm.cid = Number(this.cid)
    this.gradeForm.username = this.username
  }
}
</script>

<style scoped>

.icon-middle {
  line-height: 40px;
  margin-right: 5px;
}

.el-form-item {
  margin-bottom: 0;
}
</style>