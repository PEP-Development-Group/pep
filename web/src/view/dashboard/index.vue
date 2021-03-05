<template>
  <div class="big">
    <!-- 教师，学生部分 -->
    <div v-if="authId!=888">
      <el-row>
        <el-col :span="24">
          <el-card style="margin-top:-6px">
            <el-row style="margin:-10px 0">
              <el-col :span="16">
                <p class="welcome">欢迎<span v-if="authId!=1">您</span>，</p>
                <p class="name">{{ name }}{{ appellation }}</p>
                <p v-if="authId!=1">{{ today }}</p>
              </el-col>
              <el-col :span="8">
                <div class="stu-tips" v-if="authId==1">
                  <p>已修/总学时</p>
                  <p>{{ haveCredits }} / {{ totalCredits }}</p>
                  <div class="nowarp">{{ today }}</div>
                </div>
              </el-col>
            </el-row>
          </el-card>
        </el-col>
      </el-row>

      <el-row :gutter="10">
        <el-col :xs="24" :sm="14">
          <el-card style="height: 149px">
            <div class="button-box clearflex">
              <b>管理员公告</b>
            </div>
            <!-- <el-button
                v-auth="888"
                @click="dialogVisible = true"
                type="text"
                icon="el-icon-edit"
                plain
            ></el-button> -->
            <p class="announce-con">{{ adminAnnouncement }}</p>
          </el-card>
          <!-- 学生部分 -->
          <el-card v-if="authId==1" style="height: 294px; margin-top: 18px; margin-bottom: 18px">
            <div class="button-box clearflex">
              <b>课程</b>
              <span style="color: #999999" v-if="todoClass">还有{{ todoClass }}节实验已选未上</span>
              <b class="right">
                <el-link type="primary" :underline="false" @click="goToStuLesson">查看详情</el-link>
              </b>
            </div>
            <el-table
                :data="tableData"
                max-height="260px"
                size="medium"
                :show-header="false"
                class="classTable"
                empty-text="暂无课程"
            >
              <el-table-column>
                <template slot-scope="scope">
                  <div class="lessonCon">
                  <span class="nowarp"><span style="font-weight: bold">{{ scope.row.cname }}</span>
                  <el-divider direction="vertical"></el-divider>
                  {{ scope.row.hours }}学时</span>
                    <span class="grade" v-if="isFinished(scope.row.desc)"
                          :class="{fail:scope.row.grade<60||scope.row.grade===101,pass:scope.row.grade>=60}">
            {{ scope.row.grade === 102 ? "成绩未出" : scope.row.grade === 101 ? "旷课" : scope.row.grade }}
                    </span>
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


          <!-- 老师部分 -->
          <el-card v-if="authId==2" style="height: 294px; margin-top: 18px; margin-bottom: 18px">
            <!-- <div class="button-box clearflex">
              <el-radio-group v-model="classType" size="small">
                <el-radio-button :label="1">已上课程</el-radio-button>
                <el-radio-button :label="2">未来课程</el-radio-button>
              </el-radio-group>
            </div> -->

            <div class="button-box clearflex">
              <b>待上课程</b>
              <b class="right">
                <el-link type="primary" :underline="false" @click="goToLessonManagement">查看详情</el-link>
              </b>
            </div>
            <el-table ref="userTable" :data="[courseList,classType]|classFilter" max-height="260px"
                      :show-header="false">
              <el-table-column min-width="150">
                <template slot-scope="scope">
                  {{ scope.row.cname }}
                  <el-tag effect="dark" size="mini" class="space" type="success">{{ scope.row.ccredit }}学时
                  </el-tag>
                  <el-tag effect="light" size="mini" class="space">{{ scope.row.classroom }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column min-width="150">
                <template slot-scope="scope">
                  <el-popover trigger="hover" :content="scope.row.desc|descFormatDateDay" placement="right">
                    <div slot="reference" class="text-wrapper">
                      {{ scope.row.desc | formatDesc }}
                    </div>
                  </el-popover>
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


    </div>


    <!-- 管理员 -->
    <div v-if="authId==888">
      <el-row :gutter="10">
        <el-col :xs="24" :sm="12">
          <el-card style="height: 100px;">
            <p class="welcome">欢迎您，</p>
            <p class="name">管理员</p>
          </el-card>
        </el-col>
        <el-col :xs="24" :sm="12">
          <el-card style="height: 100px;">
            <b>管理员公告</b>
            <el-button
                @click="dialogVisible = true"
                type="text"
                icon="el-icon-edit"
                plain
            ></el-button>
            <p class="announce-con">{{ adminAnnouncement }}</p>
          </el-card>
        </el-col>
      </el-row>
      <el-row :gutter="10">
        <el-col :xs="24" :sm="12">
          <el-card style="height:462px">
            <div class="button-box clearflex">
              <b>服务器状态</b>
              <b class="right">
                <el-link type="primary" :underline="false" @click="goToSystemState">查看详情</el-link>
              </b>
            </div>

            <!--            <el-row :gutter="10" style="margin:60px 0px">-->
            <!--              <el-col :xs="24" :sm="12">-->
            <!--                <el-card v-if="state.disk" class="card_item">-->
            <!--                  <div slot="header">Disk</div>-->
            <!--                  <div>-->
            <!--                    <el-row>-->
            <!--                      &lt;!&ndash; <el-col :span="12">-->
            <!--                        <el-row :gutter="10">-->
            <!--                          <el-col :span="12">total (MB)</el-col>-->
            <!--                          <el-col :span="12" v-text="state.disk.totalMb"></el-col>-->
            <!--                        </el-row>-->
            <!--                        <el-row :gutter="10">-->
            <!--                          <el-col :span="12">used (MB)</el-col>-->
            <!--                          <el-col :span="12" v-text="state.disk.usedMb"></el-col>-->
            <!--                        </el-row>-->
            <!--                        <el-row :gutter="10">-->
            <!--                          <el-col :span="12">total (GB)</el-col>-->
            <!--                          <el-col :span="12" v-text="state.disk.totalGb"></el-col>-->
            <!--                        </el-row>-->
            <!--                        <el-row :gutter="10">-->
            <!--                          <el-col :span="12">used (GB)</el-col>-->
            <!--                          <el-col :span="12" v-text="state.disk.usedGb"></el-col>-->
            <!--                        </el-row>-->
            <!--                      </el-col> &ndash;&gt;-->
            <!--                      <el-col :span="12" :offset="6">-->
            <!--                        <el-progress-->
            <!--                            type="dashboard"-->
            <!--                            :percentage="state.disk.usedPercent"-->
            <!--                            :color="colors"-->
            <!--                        ></el-progress>-->
            <!--                      </el-col>-->
            <!--                    </el-row>-->
            <!--                  </div>-->
            <!--                </el-card>-->
            <!--              </el-col>-->
            <!--              <el-col :xs="24" :sm="12">-->
            <!--                <el-card v-if="state.ram" class="card_item">-->
            <!--                  <div slot="header">Ram</div>-->
            <!--                  <div>-->
            <!--                    <el-row>-->
            <!--                      <el-col :span="12" :offset="6">-->
            <!--                        <el-progress-->
            <!--                            type="dashboard"-->
            <!--                            :percentage="state.ram.usedPercent"-->
            <!--                            :color="colors"-->
            <!--                        ></el-progress>-->
            <!--                      </el-col>-->
            <!--                    </el-row>-->
            <!--                  </div>-->
            <!--                </el-card>-->
            <!--              </el-col>-->
            <!--            </el-row>-->


          </el-card>
        </el-col>
        <el-col :xs="24" :sm="12">
          <el-card style="height:462px">
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
    </div>


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
import {GetPersonalClasses, getTeacherClassList, getUserCreditInfo} from "@/api/course";
import {getRecord, updateRecord} from "@/api/globle";
import {realTimeToSchoolTime, schoolTimeToRealTime, formatTimeToStr} from "@/utils/date";
import {mapGetters} from "vuex";
// import {getSystemState} from "@/api/system.js";

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
      haveCredits: '-',
      selectedCredits: null,
      adminAnnouncement: "",
      courseList: [],
      classType: 2,
      timer: null,
      state: {},
      colors: [
        {color: "#5cb87a", percentage: 20},
        {color: "#e6a23c", percentage: 40},
        {color: "#f56c6c", percentage: 80},
      ],
      adminFixContent: {
        msg: "",
      },
      appellation: store.state.user.userInfo.authorityId == 1 ? '同学' : '老师',
      authId: store.state.user.userInfo.authorityId,
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
    if (this.$store.state.user.userInfo.authorityId === '1') {
      this.tableData = (await GetPersonalClasses()).data.list.crs;
      await this.updateInfo()
    }
    this.adminAnnouncement = (await getRecord()).data;
    this.getToday();
    if (this.$store.state.user.userInfo.authorityId === '2') await this.getList();

  },
  beforeDestroy() {
    clearInterval(this.timer)
    this.timer = null
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
    classFilter(list) {
      let currTime = new Date()
      if (list[0])
        return list[0].filter((item) => {
          let o = schoolTimeToRealTime(item.desc, store.state.user.firstDay) < currTime
          return ((list[1] === 1) && o) || ((list[1] === 2) && !o)
        })
      return []
    },
    formatDesc: function (d) {
      if (d) {
        let descList = d.split('-')
        return "第" + descList[0] + "周 周" + formatDayOfWeek[descList[1] - 1] + " 第" + descList[2] + "节";
      }
    },
    descFormatDateDay: function (desc) {
      let time = schoolTimeToRealTime(desc, store.state.user.firstDay)
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy年MM月dd日");
      } else {
        return "";
      }
    }
  },
  computed: {
    ...mapGetters("user", ["token"]),
    todoClass: function () {
      const t = new Date()
      if (store.state.user.userInfo.authorityId === '1') {
        if (this.tableData) {
          let cnt = 0
          for (let i = 0; i < this.tableData.length; ++i) {
            if (schoolTimeToRealTime(this.tableData[i].desc, store.state.user.firstDay) > t) {
              ++cnt
            }
          }
          return cnt
        }
      }
      return null
    }
  },
  methods: {
    async updateInfo() {
      const res = await getUserCreditInfo()
      this.haveCredits = res.data.have_credits
    },
    // async reload() {
    //   const {data} = await getSystemState();
    //   this.state = data.server;
    // },
    goToSystemState() {
      this.$router.push('/system/state')
    },
    goToLessonManagement() {
      this.$router.push('/lessonManagement')
    },
    goToStuLesson() {
      this.$router.push('/lessons')
    },
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
    },
    async getList() {
      const list = await getTeacherClassList();
      if (list.code === 0) {
        this.courseList = list.data.courses.classes
      }
    },
    viewLesson(cid) {
      this.$router.push({
        path: 'lesson',
        query: {
          cid: cid
        }
      })
    }
  }
};
</script>

<style scoped>
@media screen and (min-width: 320px) and (max-width: 750px) {
  .lessonCon {
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
  margin: 60px 0 -16px 0;
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

.space {
  float: right;
  margin-left: 10px;
}

.button-box {
  position: relative;
  padding: 10px 0px;
}

.button-box .right {
  position: absolute;
  right: 0;
}

.button-box .el-button {
  float: right;
}

.user-dialog .header-img-box {
  width: 200px;
  height: 200px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}

.user-dialog .avatar-uploader .el-upload:hover {
  border-color: #409eff;
}

.user-dialog .avatar-uploader-icon {
  border: 1px dashed #d9d9d9 !important;
  border-radius: 6px;
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}

.user-dialog .avatar {
  width: 178px;
  height: 178px;
  display: block;
}

.option-btn {
  margin-right: 5px;
}

.user-type-radio {
  margin-bottom: 20px;
  margin-left: 80px;
}
</style>