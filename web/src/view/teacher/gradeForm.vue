<template>
  <div>
    <div v-auth="888">
      <span v-if="grade===101">
        <el-tag type="danger" effect="dark" size="mini">旷课</el-tag>
      </span>
      <span v-else-if="grade<=100">{{ grade }}</span>
      <el-form :inline="true" :model="gradeForm" @submit.native.prevent style="display: inline;margin-left: 10px">
          <span class="icon-middle">
            <i :class="{'el-icon-key':!isLoading,'el-icon-loading':isLoading}"></i>
          </span>
        <el-form-item prop="grade">
          <el-input v-model.number="gradeForm.grade" :autofocus="true" type="tel"
                    @keyup.enter.native="submitGrade"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button :disabled="gradeValid" @click="submitGrade"><i class="el-icon-s-custom"></i>提交</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="warning" @click="submitFailGrade"><i class="el-icon-s-custom"></i>旷课</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div v-auth="2">
    <span v-if="grade<=100">
            <i class="el-icon-lock"></i>
      <span :class="{fail:grade<60}">{{ grade }}</span>
          </span>
      <span v-else-if="grade===101">
      <el-tag type="danger" effect="dark" size="mini">旷课</el-tag>
    </span>
      <span v-else-if="grade===102">
    <el-form :inline="true" :model="gradeForm" @submit.native.prevent>
      <span class="icon-middle">
        <i :class="{'el-icon-unlock':!isLoading,'el-icon-loading':isLoading}"></i>
      </span>
      <el-form-item prop="grade">
        <el-input v-model.number="gradeForm.grade" :autofocus="true" type="tel"
                  @keyup.enter.native="submitGrade"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button :disabled="gradeValid" @click="submitGrade">提交</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="warning" @click="submitFailGrade">旷课</el-button>
      </el-form-item>
    </el-form>
    </span>
    </div>
  </div>
</template>

<script>
import {setStuGrade} from "@/api/course";

export default {
  name: "gradeForm",
  props: {
    username: String,
    cid: String,
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
        cid: null,
        username: null,
        grade: ''
      },
    }
  },
  methods: {
    submitGrade() {
      this.isLoading = true
      setStuGrade(this.gradeForm).then((res) => {
        if (res.code === 0)
          this.grade = this.gradeForm.grade
        this.isLoading = false
      })
    },
    submitFailGrade() {
      this.isLoading = true
      this.gradeForm.grade = 101
      setStuGrade(this.gradeForm).then((res) => {
        if (res.code === 0)
          this.grade = this.gradeForm.grade
        this.isLoading = false
      })
    }
  }, created() {
    this.gradeForm.cid = parseInt(this.cid)
    this.gradeForm.username = this.username
  }
}
</script>

<style scoped>
.fail {
  color: #FF6666;
}

.icon-middle {
  line-height: 40px;
  margin-right: 5px;
}

.el-form-item {
  margin-bottom: 0;
}
</style>