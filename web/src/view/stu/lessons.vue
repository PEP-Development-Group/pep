<template>
  <div>
    <div class="grade-tips">本学期需要修满<strong>{{ totalCredits }}</strong>学时,当前已修<strong>{{
        haveCredits
      }}</strong>学时,平均分为<strong>{{ ave }}</strong>
    </div>
    <el-card v-for="(item,i) in sortedList" v-bind:key="i" class="lesson-card">
      <div class="lesson-title">
        <span class="cname">{{ item.cname }}</span>
        <el-tag size="mini">{{ item.hours }}学时</el-tag>
        <span class="grade fail" v-if="item.grade===101">旷课</span>
        <span class="grade" v-else-if="isFinished(item.desc)"
              :class="{fail:item.grade<60,pass:item.grade>=60}">{{
            item.grade === 102 ? "成绩未出" : item.grade
          }}</span>
        <span class="grade" v-else>
          <el-button type="danger" size="mini" plain @click="deleteCourse(item.id)">退课</el-button>
        </span>
      </div>
      <div>
        {{ item.desc|formatDesc }}
        <el-divider direction="vertical"></el-divider>
        <span class="nowarp">{{ item.tname }}</span>
        <el-divider direction="vertical"></el-divider>
        <span class="nowarp">{{ item.classroom }}</span>
      </div>
    </el-card>
  </div>
</template>

<script>
import {DeleteSelect, GetPersonalClasses} from "@/api/course";
import {store} from "@/store";
import {schoolTimeToRealTime} from "@/utils/date";

const userInfo = store.getters['user/userInfo']
const formatDayOfWeek = ['一', '二', '三', '四', '五', '六', '日']
export default {
  name: "lessons",
  async created() {
    await this.getList()
    this.getAve()
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
    //TODO 需要根据课程情况调整算法
    getAve() {
      if (this.tableData == null) {
        this.ave = 0
        return
      }
      let a = 0.0
      let cnt = 0
      for (let i = 0; i < this.tableData.length; ++i) {
        if (this.tableData[i].grade <= 100) {
          a += this.tableData[i].grade
          cnt++
        }
      }
      this.ave = cnt ? (a / cnt).toFixed(2) : 0
    },
    isFinished(desc) {
      let now = new Date()
      return schoolTimeToRealTime(desc, store.state.user.firstDay) < now
    }, async deleteCourse(cid) {
      this.$confirm('你确定要退课吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        showClose: false,
      }).then(async () => {
        const res = await DeleteSelect({"username": userInfo.username, "cid": cid})
        if (res.code === 0) {
          this.$message({type: "success", message: "退课成功"})
        }
        await this.getList()
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '未退课'
        });
      });

    }, async getList() {
      this.tableData = (await GetPersonalClasses()).data.list.crs
    }
  },
  computed: {
    sortedList: function () {
      if (this.tableData == null) return
      let orderList = this.tableData
      orderList.sort(function (a, b) {
        let x = schoolTimeToRealTime(a.desc, store.state.user.firstDay)
        let y = schoolTimeToRealTime(b.desc, store.state.user.firstDay)
        return ((x > y) ? -1 : (x < y) ? 1 : 0)
      })
      return orderList
    }
  },
  data() {
    return {
      tableData: [],
      totalCredits: userInfo.total_credits,
      haveCredits: userInfo.have_credits,
      ave: null
    }
  }
}
</script>

<style scoped>

.grade-tips {
  margin: 10px;
}

.grade-tips strong {
  font-weight: bold;
  margin: 0 2px;
}

.fail {
  color: #FF6666;
}

.pass {
  color: #00c6ac;
}

.grade {
  float: right;
  font-size: 20px;
}

.lesson-card {
  margin-bottom: 5px;
  width: min(100%, 600px);
}

.lesson-card .el-card__body {

  margin-left: 5%;
  margin-right: 5%;

  box-sizing: border-box;
}

.cname {
  font-size: 18px;
  margin-right: 5px;
}

.lesson-title {
  margin-bottom: 5px;
}

.el-tag {
  vertical-align: top;
}

.nowarp {
  white-space: nowrap;
}
</style>