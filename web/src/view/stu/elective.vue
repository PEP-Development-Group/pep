<template>
  <div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-for="(item,i) in courseList" :key="item.id" class="class-con">
        <template slot="title">
          <div class="class-title-con">
            <span class="class-title">{{ i }}</span>
            <el-tag effect="dark" size="small" class="hours-tag" type="info"
                    :color="(colors.duration[item.hours-1])">{{ item.hours }}学时
            </el-tag>
          </div>
        </template>
        <el-card v-for="l in item.List" class="lesson" :key="l.id" shadow="hover">
          <i class="el-icon-check check" v-if="l.selected"></i>
          <el-progress v-else class="lesson-progress" type="circle" :width="18" :show-text="false"
                       :percentage="selectPercent(l.now,l.max)"
                       :stroke-width="2"></el-progress>
          <span class="lesson-info">
          <span class="space">{{ l.desc|formatDesc }}</span>
          <span class="space nowarp">{{ l.teacher_name }}</span>
          <el-tag effect="light" size="mini" class="space">{{ l.class_room }}</el-tag>
          <el-tag effect="dark" size="mini" type="info" class="space" v-if="countDown(l.time)!==-1"
                  :color="(colors.time[countDown(l.time)])">{{ timeTag[countDown(l.time)] }}</el-tag>
          </span>
          <span class="lesson-op">
            <span class="progress">{{ l.now }}/{{ l.max }}</span>
          <el-button type="primary" v-if="l.selected===false" @click="selectCourse(i,item,l)">选课</el-button>
          <el-button type="danger" v-else @click="deleteCourse(l.id)">退选</el-button>
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
  DeleteSelect
} from "@/api/course.js";  //  此处请自行替换地址
import {formatTimeToStr} from "@/utils/date";
import infoList from "@/mixins/infoList";
import {store} from '@/store'

const userInfo = store.getters['user/userInfo']
const formatDayOfWeek = ['一', '二', '三', '四', '五', '六', '日']

export default {
  name: "Elective",
  mixins: [infoList],
  data() {
    return {
      activeNames: [1],
      courseList: [],
      colors: {
        duration: ["#99CC66", "#67C23A", "#409EFF", "#009999"],
        time: ["#6699CC", "#e6a23c", "#FF6666", "#CC0033"]
      },
      timeTag: ['下周', '本周', '明日', '今日'],
      confirmVisible: true,
      confirmData: {
        name: "null",
        classRoom: "null",
        desc: "null",
        teacher: "null",
        hours: "null"
      }
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
  computed: {},
  methods: {
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
      const res = await SelectClass(d)
      if (res.code === 0) {
        this.confirmData = {
          name: i,
          classRoom: l.class_room,
          desc: l.desc,
          teacher: l.teacher_name,
          hours: item.hours
        }
        this.confirmVisible = true
      }
      await this.getList()
    },
    async deleteCourse(cid) {
      const res = await DeleteSelect({"username": userInfo.username, "cid": cid})
      if (res.code === 0) {
        this.$message({type: "success", message: "退课成功"})
      }
      await this.getList()
    },
    async getList() {
      const list = await GetClassListWithPerson();
      if (list.code === 0) {
        this.courseList = list.data.courses
      }
    },
    countDown(time) {
      let d = new Date().setHours(0, 0, 0, 0)
      let t = new Date(time).setHours(0, 0, 0, 0)
      if (d === t)
        return 3
      else if (parseInt(t / 60 / 60 / 24 / 1000) - parseInt(d / 60 / 60 / 24 / 1000) === 1)
        return 2
      else if (parseInt((parseInt(d / 60 / 60 / 24 / 1000) + 4) / 7) === parseInt((parseInt(t / 60 / 60 / 24 / 1000) + 4) / 7))
        return 1
      else if (parseInt((parseInt(t / 60 / 60 / 24 / 1000) + 4) / 7) - parseInt((parseInt(d / 60 / 60 / 24 / 1000) + 4) / 7) === 1)
        return 0
      return -1
    }
  },
  async created() {
    await this.getList()
  }

}
</script>

<style>

.el-tag {
  border-color: transparent !important;
}

.hours-tag {
  margin-left: 10px;
}

.class-title {
  display: inline-block;
  font-weight: bold;
  margin-left: 20px;
  line-height: normal;
  vertical-align: middle;
  max-width: calc(100% - 80px);
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
  max-width: calc(100% - 120px);
}

.lesson {
  box-sizing: border-box;
  margin: 8px;
  /*font-size: medium;*/
}

.lesson-progress {
  vertical-align: middle;
  margin-right: 2px;
  margin-left: -8px;
}

.lesson-op {
  background-color: #f1f3f4;
  border-radius: 4px;
  float: right;
}

.space {
  vertical-align: middle;
  margin: 6px;
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
  margin-left: -7px;
}

.nowarp {
  white-space: nowrap;
}

.check-dialog .el-dialog {
  width: min(500px, 95%);
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
</style>
