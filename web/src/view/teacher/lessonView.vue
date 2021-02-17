<template>
  <div>
    <el-page-header @back="goBack" :content="lesson.cname">
      <template slot="title">
        <div></div>
      </template>
      <template slot="content">
        <div class="title-main">{{ lesson.cname }}</div>
        <div class="title-desc">{{ lesson.desc |formatDesc }}</div>
      </template>
    </el-page-header>
    <el-table :data="stuList">
      <el-table-column label="学号" prop="username"></el-table-column>
      <el-table-column label="姓名" prop="name"></el-table-column>
      <el-table-column label="成绩">
        <template slot-scope="scope">
          <span v-if="scope.row.grade<=100">
            <i class="el-icon-lock"></i>
            {{ scope.row.grade }}
          </span>
          <span v-else-if="scope.row.grade===102">
            <el-form inline rules="rules">
              <span class="el-icon-unlock icon-middle"></span>
              <el-form-item>
                <el-input v-model="gradeForm.grade"></el-input>
              </el-form-item>
              <el-form-item>
                <el-button @click="submitGrade(scope.row.cid)">提交</el-button>
              </el-form-item>
            </el-form>
          </span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {getTeacherAClassStuList, findClass, setStuGrade} from "@/api/course";

const formatDayOfWeek = ['一', '二', '三', '四', '五', '六', '日']
export default {
  name: "lessonView",
  data() {
    return {
      cid: 0,
      lesson: null,
      stuList: [],
      gradeForm: {
        cid: null,
        grade: null
      },
      rules: {
        grade: [
          {
            required: true, message: '请输入成绩', trigger: 'blur'
          }, {
            min: 0, max: 100, message: '成绩不合法', trigger: 'blur'
          }
        ]
      }
    }
  },
  filters: {
    formatDesc: function (d) {
      if (d) {
        let descList = d.split('-')
        return "第" + descList[0] + "周 周" + formatDayOfWeek[descList[1] - 1] + " 第" + descList[2] + "节";
      }
    }
  },
  methods: {
    goBack() {
      this.$router.back()
    },
    submitGrade(cid) {
      this.gradeForm.cid = cid
      setStuGrade(this.gradeForm)
    }
  },
  async created() {
    this.cid = this.$router.currentRoute.query.cid
    this.lesson = (await findClass({ID: this.cid})).data.reclass
    this.stuList = (await getTeacherAClassStuList({"cid": this.cid})).data.students
  }
}
</script>

<style scoped>
.title-main {
  font-size: 18px;
}

.title-desc {
  font-size: 10px;
  color: #888;
}

.right {
  float: right;
}

.el-form-item {
  margin-bottom: 0;
}

.icon-middle {
  height: 0px;
  vertical-align: middle;
  margin-right: 5px;
}
</style>