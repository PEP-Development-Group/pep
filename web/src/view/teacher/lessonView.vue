<template>
  <div>
    <el-page-header @back="goBack">
      <template slot="title">
        <div></div>
      </template>
      <template slot="content">
        <div class="title-main" v-if="lesson">{{ lesson.cname }}</div>
        <div class="title-desc" v-if="lesson">{{ lesson.desc |formatDesc }} <span v-auth="888">{{ lesson.tname }}</span>
        </div>
      </template>
    </el-page-header>
    <el-table :data="stuList">
      <el-table-column label="学号" prop="username"></el-table-column>
      <el-table-column label="姓名" prop="name"></el-table-column>
      <el-table-column label="成绩">
        <template slot-scope="scope">
          <grade-form :username="scope.row.username" :grade="scope.row.grade" :cid="cid"></grade-form>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {getTeacherAClassStuList, findClass} from "@/api/course";
import GradeForm from "@/view/teacher/gradeForm";
import {store} from "@/store";

const formatDayOfWeek = ['一', '二', '三', '四', '五', '六', '日']
export default {
  name: "lessonView",
  components: {GradeForm},
  data() {
    return {
      cid: 0,
      lesson: null,
      stuList: [],
      authID: store.state.user.userInfo.authorityId
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


</style>