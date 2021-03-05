<template>
  <div>
    <div class="warning-info">
      <i class="el-icon-warning"></i> 请注意退课有次数限制,你还可以退课 {{ cancelTimes }} 次
    </div>
    <div class="class-info">
      <i class="el-icon-info"></i> 本学期需要修满 {{ totalCredits }} 学时,当前已选 {{ choosedCredits }} 学时
    </div>
    <el-collapse v-model="activeNames" class="class-area">
      <el-collapse-item v-for="(item,i) in courseList" :key="item.id" class="class-con">
        <template slot="title">
          <div class="class-title-con">
            <!--            <el-tag effect="dark" class="hours-tag" type="info"-->
            <!--                    :color="(colors.duration[item.hours-1])">{{ item.hours }}学时-->
            <!--            </el-tag>-->
            <span class="class-hours">{{ item.hours }}学时</span>
            <el-divider direction="vertical"></el-divider>
            <span class="class-title" :class="{tiny:item.learned}">{{ i }}</span>

            <span class="class-title-right">共{{ item.List.length }}节</span>
            <span class="class-title-right">
            <el-tag effect="light" size="small" type="success" v-if="item.learned">
              <i class="el-icon-check"></i>
            已选
            </el-tag>
            </span>
          </div>
        </template>
        <el-card v-for="l in item.List" class="lesson" :key="l.id" shadow="hover">
          <i class="el-icon-check check" v-if="l.selected"></i>
          <el-progress v-else class="lesson-progress" type="circle" :width="18" :show-text="false"
                       :percentage="selectPercent(l.now,l.max)"
                       :stroke-width="2"></el-progress>
          <span class="lesson-info">
          <span class="whitespace nowarp">{{ l.desc|formatDesc }}</span>
          <span class="whitespace nowarp">{{ l.teacher_name }}</span>
          <el-tag effect="light" size="mini" class="whitespace">{{ l.class_room }}</el-tag>
          <el-tag effect="dark" size="mini" type="info" class="whitespace" v-if="countDown(l.desc)!==-1"
                  :color="(colors.time[countDown(l.desc)])">{{ timeTag[countDown(l.desc)] }}</el-tag>
          </span>
          <span class="lesson-op">
            <el-tooltip effect="dark" placement="left" style="outline: none">
            <div slot="content">剩余 {{ l.max - l.now }} 个名额</div>
            <span class="progress" :class="{red:l.max-l.now<=3}">
                {{ l.now }}/{{ l.max }}
            </span>
          </el-tooltip>
          <el-button type="primary" v-if="l.selected===false" :disabled="l.now===l.max||item.learned"
                     @click="selectCourse(i,item,l)">选课</el-button>
          <el-button type="danger" v-else @click="deleteCourse(l.id,l)">退选</el-button>
          </span>
        </el-card>
      </el-collapse-item>
    </el-collapse>
    <el-dialog :show-close="false" :visible.sync="confirmVisible" center class="check-dialog">
      <template slot="title">
        <i class="el-icon-success"></i>
        选课成功
      </template>
      <div class="confirm-con">
        <el-row>
          <el-col :span="12">课程名称</el-col>
          <el-col :span="12">{{ confirmData.name }}</el-col>
        </el-row>
        <el-row>
          <el-col :span="12">教师</el-col>
          <el-col :span="12">{{ confirmData.teacher }}</el-col>
        </el-row>
        <el-row>
          <el-col :span="12">教室</el-col>
          <el-col :span="12">{{ confirmData.classRoom }}</el-col>
        </el-row>
        <el-row>
          <el-col :span="12">学时</el-col>
          <el-col :span="12">{{ confirmData.hours }}</el-col>
        </el-row>
        <el-row>
          <el-col :span="12">时间</el-col>
          <el-col :span="12">{{ confirmData.desc|formatDesc }}</el-col>
        </el-row>
      </div>
      <template slot="footer">
        <el-button class="confirm-btn" type="primary" @click="confirmVisible = false">确认</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  GetClassListWithPerson,
  SelectClass,
  DeleteSelect,
  getUserCreditInfo
} from "@/api/course.js";  //  此处请自行替换地址
import {formatTimeToStr, schoolTimeToRealTime} from "@/utils/date";
import infoList from "@/mixins/infoList";
import {store} from '@/store'
import context from "@/main";

const userInfo = store.getters['user/userInfo']
const formatDayOfWeek = ['一', '二', '三', '四', '五', '六', '日']

export default {
  name: "Enroll",
  mixins: [infoList],
  async created() {
    await this.updateInfo()
    await this.getList()
  },
  data() {
    return {
      activeNames: [1],
      courseList: [],
      colors: {
        duration: ["#99CC66", "#67C23A", "#409EFF", "#009999"],
        time: ["#6699CC", "#e6a23c", "#FF6666", "#CC0033"]
      },
      timeTag: ['下周', '本周', '明天', '今天'],
      confirmVisible: false,
      confirmData: {
        name: null,
        classRoom: null,
        desc: null,
        teacher: null,
        hours: null
      },
      totalCredits: userInfo.total_credits,
      choosedCredits: "-",
      cancelTimes: "-"
    };
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function (bool) {
      if (bool != null) {
        return bool ? "是" : "否";
      } else {
        return "";
      }
    },
    formatDesc: function (d) {
      if (d) {
        let descList = d.split('-')
        return "第" + descList[0] + "周 周" + formatDayOfWeek[descList[1] - 1] + " 第" + descList[2] + "节";
      }
    }
  },
  computed: {
    sortClass() {
      let list = this.courseList;
      return list.sort(this.cmp)
    }
  },
  methods: {
    async updateInfo() {
      const res = await getUserCreditInfo()
      this.cancelTimes = res.data.cancel_nums
      this.choosedCredits = res.data.selected_credits
    },
    cmp() {
      return function (a, b) {
        if ((a.selected && !b.selected) || (!a.selected && b.selected))
          return (a.selected && !b.selected) ? 1 : -1
        return a.hours < b.hours ? -1 : 1
      }
    },
    selectPercent: function (now, max) {
      if (now < 0) return 0
      if (max <= 0) return 0
      return now / max * 100;
    },
    async selectCourse(i, item, l) {
      const d = {
        cid: l.id,
        username: userInfo.username
      }
      context.$bus.emit("showLoading")
      const res = await SelectClass(d)
      if (res.code === 0) {
        this.confirmData = {
          name: i,
          classRoom: l.class_room,
          desc: l.desc,
          teacher: l.teacher_name,
          hours: item.hours
        }
        await this.updateInfo()
        this.confirmVisible = true
      }
      context.$bus.emit("closeLoading")
      await this.getList()
    },
    async deleteCourse(cid) {
      this.$confirm('你确定要退课吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        showClose: false,
      }).then(async () => {
        context.$bus.emit("showLoading")
        const res = await DeleteSelect({"username": userInfo.username, "cid": cid})
        if (res.code === 0) {
          this.$message({type: "success", message: "退课成功"})
        }
        await this.updateInfo()
        await this.getList()
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '未退课'
        });
      });
      context.$bus.emit("closeLoading")
    },
    async getList() {
      const list = await GetClassListWithPerson();
      if (list.code === 0) {
        this.courseList = list.data.courses
      }
    },
    countDown(desc) {
      let time = schoolTimeToRealTime(desc, store.state.user.firstDay)
      let d = new Date().setHours(0, 0, 0, 0)
      let t = new Date(time).setHours(0, 0, 0, 0)
      const c = 60 * 60 * 24 * 1000
      if (d === t)
        return 3
      else if (parseInt(t / c) - parseInt(d / c) === 1)
        return 2
      else if (parseInt((parseInt(d / c) + 4) / 7) === parseInt((parseInt(t / c) + 4) / 7))
        return 1
      else if (parseInt((parseInt(t / c) + 4) / 7) - parseInt((parseInt(d / c) + 4) / 7) === 1)
        return 0
      return -1
    }
  }

}
</script>

<style>
.el-message-box {
  width: min(80%, 420px);
}

.el-tag {
  border-color: transparent !important;
}

.hours-tag {
  margin-left: 10px;
}

.class-hours {
  color: #999999;
  margin-left: 10px;
}

.class-title {
  display: inline-block;
  font-weight: bold;
  font-size: 15px;
  /*margin-left: 10px;*/
  line-height: normal;
  vertical-align: middle;
  max-width: calc(100% - 105px);
}

.tiny {
  max-width: calc(100% - 168px);
}

.class-title-right {
  float: right;
  margin-left: 5px;
}

.class-title-con {
  width: 100%;
}

.class-con {
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  margin-bottom: 5px;
}

.class-con .el-collapse-item__header {
  height: auto;
  padding: 5px 0;
}

.class-con .el-collapse-item__wrap {
  background-color: #fafafa;
}

.lesson-info {
  font-size: 18px;
  display: inline-block;
  line-height: normal;
  vertical-align: middle;
  max-width: calc(100% - 90px);
}

.lesson {
  box-sizing: border-box;
  margin: 8px;
}

.lesson-progress {
  vertical-align: middle;
  margin-right: 2px;
  margin-left: -16px;
}

.lesson-op {
  background-color: #f1f3f4;
  border-radius: 4px;
  float: right;
  margin-right: -14px;
}

.whitespace {
  vertical-align: middle;
  margin: 4px;
}

.progress {
  color: #666666;
  font-weight: bold;
  font-size: small;
  margin: 5px;
}

.check {
  font-size: 18px;
  vertical-align: middle;
  margin-left: -17px;
}

.nowarp {
  white-space: nowrap;
}

.check-dialog .el-dialog {
  border-radius: 4px;
  width: min(450px, 95%);
}

.check-dialog .el-icon-success {
  display: block;
  font-size: 80px;
  color: #67C23A;
}

.confirm-con {
  text-align: center;
}

.confirm-btn {
  width: 50%;
}

.confirm-con .el-row:nth-child(odd) {
  background-color: #f0f0f0;
}

.red {
  color: #FF6666;
}

.class-area {
  max-width: 800px;
}

.warning-info {
  margin-bottom: 5px;
  padding: 8px 16px;
  background-color: #fef0f0;
  color: #f56c6c;
  border-radius: 4px;
  max-width: 800px;
  box-sizing: border-box;
}

.class-info {
  margin-bottom: 5px;
  padding: 8px 16px;
  background-color: #f0f8ff;
  color: #1989fa;
  border-radius: 4px;
  max-width: 800px;
  box-sizing: border-box;
}
</style>
