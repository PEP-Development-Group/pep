<template>
  <div class="big">
    <el-row>
      <el-col :span="24">
        <el-card>
          <!-- 这个name应该在userInfo里面 -->
          <p class="welcome">欢迎您，</p>
          <p class="name">{{ name }}{{ appellation }}</p>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="10" style="margin-top: 10px">
      <el-col :xs="24" :sm="14">
        <el-card style="height: 160px">
          <b>管理员公告</b>
          <el-button
              v-if="adminAnnouncementModification"
              @click="dialogVisible = true"
              type="text"
              icon="el-icon-edit"
              plain
          ></el-button>
          <p>{{ adminAnnouncement }}</p>
        </el-card>

        <el-card style="height: 360px; margin-top: 18px">
          <el-table
              :data="tableData"
              border
              size="medium"
              max-height="320px"
              class="classtable"
          >
            <el-table-column prop="desc" label="时间" width="130px">
              <template slot-scope="scope">
                {{ scope.row.desc|formatDesc }}
              </template>
            </el-table-column>
            <el-table-column prop="classroom" label="地点" width="90px">
            </el-table-column>
            <el-table-column prop="tname" label="教师" width="80px"></el-table-column>
            <el-table-column prop="cname" label="课程名">
              <template slot-scope="scope">
                {{ scope.row.cname }}
                <el-tag effect="light" type="info" size="mini">
                  {{ scope.row.hours }}学时
                </el-tag>
              </template>
            </el-table-column>

            <el-table-column prop="grade" label="成绩" align="center" width="100%">
              <template slot-scope="scope">
                <el-tag v-if="scope.row.grade === 102">未上成绩</el-tag>
                <el-tag type="info" v-else-if="scope.row.grade === 101"
                >旷课
                </el-tag>
                <el-tag type="success" v-else-if="scope.row.grade >= 60"
                >{{ scope.row.grade }}
                </el-tag>
                <el-tag type="warning" v-else>{{ scope.row.grade }}</el-tag>
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

    <el-row :gutter="10" style="margin-top: 30px">
      <el-col :xs="24" :sm="12"></el-col>
    </el-row>

    <el-dialog
        title="修改公告"
        :visible.sync="dialogVisible"
        width="30%"
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

const formatDayOfWeek = ['一', '二', '三', '四', '五', '六', '日']
export default {
  name: "Dashboard",
  data() {
    return {
      dialogVisible: false,
      adminAnnouncementModification: true,
      activeName: "1",
      name: store.state.user.userInfo.name,
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
  filters: {
    formatDesc: function (d) {
      if (d) {
        let descList = d.split('-')
        return "第" + descList[0] + "周 周" + formatDayOfWeek[descList[1] - 1] + " 第" + descList[2] + "节";
      }
    }
  },
  methods: {
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
  },
  async created() {
    this.tableData = (await GetPersonalClasses()).data.list.crs
    this.adminAnnouncement = (await getRecord()).data
  }
};
</script>

<style scoped>
b {
  color: rgb(64, 158, 255);
  padding-right: 10px;
}

.big {
  margin: 60px 0 0 0;
  padding-top: 0;
  background-color: rgb(243, 243, 243);
  padding-top: 15px;
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
</style>