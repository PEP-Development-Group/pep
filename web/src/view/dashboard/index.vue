<template>
  <div class="big">
    <el-row>
      <el-col :span="24">
        <el-card style="margin-top:-1px">
          <!-- 这个name应该在userInfo里面 -->
          <el-row style="margin:-10px 0">
            <el-col :span="12">
              <p class="welcome">欢迎<span v-auth.not="1">您</span>，</p>
              <p class="name">{{ name }}{{ appellation }}</p>
              <p v-auth.not="1">{{ today }}</p>
            </el-col>
            <el-col :span="12">
              <div class="stu-tips" v-auth="1">
                <p>已修/总学时</p>
                <p>{{ haveCredits }} / {{ totalCredits }}</p>
                <div>{{ today }}</div>
              </div>
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="10">
      <el-col :xs="24" :sm="14">
        <el-card style="height: 149px">
          <b>管理员公告</b>
          <el-button
              v-auth="888"
              @click="dialogVisible = true"
              type="text"
              icon="el-icon-edit"
              plain
          ></el-button>
          <p class="announce-con">{{ adminAnnouncement }}</p>
        </el-card>

        <el-card style="height: 300px; margin-top: 18px; margin-bottom: 18px">
          <el-table
              :data="tableData"
              max-height="260px"
              size="medium"
              :show-header="false"
              class="classTable"
          >
            <el-table-column>
              <template slot-scope="scope">
                <div class="lessonCon">
                <span class="nowarp"><span style="font-weight: bold">{{ scope.row.cname }}</span>
                <el-divider direction="vertical"></el-divider>
                {{ scope.row.hours }}学时</span>
                  <span class="grade fail" v-if="scope.row.grade===101">旷课</span>
                  <span class="grade" v-else-if="isFinished(scope.row.desc)"
                        :class="{fail:scope.row.grade<60,pass:scope.row.grade>=60}">{{
                      scope.row.grade === 102 ? "成绩未出" : scope.row.grade
                    }}</span>
                  <el-divider direction="vertical"></el-divider>
                  <span class="nowarp">{{ scope.row.desc|formatDesc }}
                    <el-divider direction="vertical"></el-divider>
                </span>
                  <span class="nowarp">
                  {{ scope.row.tname }}
                  <el-divider direction="vertical"></el-divider>
                  {{ scope.row.classroom }}
                </span>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="10">
        <el-card>
          <h3 class="courseTitle2">2学时课节</h3>

          <el-table
              :data="tableData2"
              size="medium"
              :show-header="false"
              border
          >
            <el-table-column prop="class" label="课程"></el-table-column>
            <el-table-column prop="date" label="时间"></el-table-column>
          </el-table>

          <h3 class="courseTitle4">4学时课节</h3>
          <el-table
              :data="tableData4"
              size="medium"
              :show-header="false"
              border
          >
            <el-table-column prop="class" label="课程"></el-table-column>
            <el-table-column prop="date" label="时间"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>

    <el-dialog
        title="修改公告"
        :visible.sync="dialogVisible"
        :width="dialogWidth"
        @close="adminClosed"
    >
      <el-form :model="adminFixContent" ref="contentRef" label-width="90px">
        <el-form-item label="请输入内容" prop="content">
          <el-input v-model="adminFixContent.msg"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="adminFix">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import {store} from "@/store";
import {GetPersonalClasses} from "@/api/course";
import {getRecord, updateRecord} from "@/api/globle";
import {realTimeToSchoolTime, schoolTimeToRealTime} from "@/utils/date";

const formatDayOfWeek = ['一', '二', '三', '四', '五', '六', '日']
export default {
  name: "Dashboard",
  data() {
    return {
      dialogWidth: "30",
      dialogVisible: false,
      activeName: "1",
      name: store.state.user.userInfo.name,
      today: "",
      totalCredits: store.state.user.userInfo.total_credits,
      haveCredits: store.state.user.userInfo.have_credits,
      adminAnnouncement: "",
      adminFixContent: {
        msg: "",
      },
      appellation: store.state.user.userInfo.authorityId == 1 ? '同学' : '老师',
      tableData: [],
      tableData2: [
        {
          class: "第一节课",
          date: "08:20-10:00",
        },
        {
          class: "第二节课",
          date: "10:20-12:00",
        },
        {
          class: "第三节课",
          date: "13:30-15:10",
        },
        {
          class: "第四节课",
          date: "15:30-17:10",
        },
        {
          class: "第五节课",
          date: "18:10-19:50",
        },
      ],
      tableData4: [
        {
          class: "第一节课",
          date: "08:20-12:00",
        },
        {
          class: "第二节课",
          date: "13:30-17:10",
        },
        {
          class: "第三节课",
          date: "18:10-21:50",
        },
      ],
    };
  },
  async created() {
    this.tableData = (await GetPersonalClasses()).data.list.crs;
    this.adminAnnouncement = (await getRecord()).data;
    this.getToday();
  },
  mounted() {
    const that = this;
    window.onresize = () => {
      return (() => {
        if (document.body.clientWidth <= 750) that.dialogWidth = "90";
        else that.dialogWidth = "30";
      })()
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
    isFinished(desc) {
      let now = new Date()
      return schoolTimeToRealTime(desc, store.state.user.firstDay) < now
    },
    adminFix() {
      this.$refs.contentRef.validate(async (valid) => {
        if (!valid) return;
        // 可以发起添加用户的网络请求
        const res = await updateRecord(this.adminFixContent)

        if (res.code !== 0) {
          this.$message.error("修改公告失败！");
        } else
          this.$message.success("修改公告成功！");
        this.adminAnnouncement = (await getRecord()).data
        this.dialogVisible = false;
      });
    },
    adminClosed() {
      this.$refs.contentRef.resetFields();
    },
    getToday() {
      let d = new Date();
      const t = realTimeToSchoolTime(d, store.state.user.firstDay);
      if (t)
        this.today = "今天是第" + t.week + "周星期" + formatDayOfWeek[t.day]
      else return "未开学"
    }
  }
};
</script>

<style scoped>
@media screen and (min-width: 320px) and (max-width: 750px) {
  .lessonCon{
    overflow-x: auto;
    width: 80vw;
  }
}
.fail {
  color: #FF6666;
}

.pass {
  color: #00c6ac;
}

.nowarp {
  white-space: nowrap;
}

.grade {
  float: right;
}

b {
  color: rgb(64, 158, 255);
  padding-right: 10px;
}

.big {
  margin: 60px 0 0 0;
  background-color: rgb(243, 243, 243);
  padding-top: 15px;
  overflow-x: hidden;
}

.courseTitle2 {
  color: rgb(64, 158, 255);
  text-align: center;
  margin-bottom: 20px;
}

.courseTitle4 {
  color: rgb(64, 158, 255);
  text-align: center;
  margin-top: 20px;
  margin-bottom: 20px;
}

.welcome {
  font-size: 20px;
}

.name {
  font-size: 30px;
}

.announce-con {
  height: 78px;
  overflow-y: auto;
}

.stu-tips {
  height: 100%;
  text-align: right;
  vertical-align: bottom;
}

.stu-tips p {
  font-size: 22px;
}

.stu-tips p:first-child {
  font-size: 12px;
  color: #aaa;
}
</style>