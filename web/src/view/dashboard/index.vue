<template>
  <div class="big">
    <el-row :gutter="10">
      <el-col :xs="24" :sm="12">
        <el-card style="height: 160px">
          <!-- 这个name应该在userInfo里面 -->
          欢迎您，
          <p>{{ name }}</p>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12">
        <el-card style="height: 160px">
          <span>管理员公告</span>
          <el-button type="primary" round>修改公告</el-button>
          <el-divider></el-divider>
          <p>{{ adminAnnouncement }}</p>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="10" style="margin-top: 30px">
      <el-col :xs="24" :sm="12">
        <el-card style="height: 360px">
          <el-table
              :data="tableData"
              border
              size="medium"
              max-height="320px"
              class="classtable"
              :show-header="false"
          >
            <el-table-column prop="date" label="时间" width="100%">
            </el-table-column>
            <el-table-column prop="place" label="地点" width="50%">
            </el-table-column>
            <el-table-column prop="classname" label="课程名"></el-table-column>
            <el-table-column prop="results" label="成绩" width="100%">
              <template slot-scope="scope">
                <el-tag v-if="scope.row.results === 102">未上成绩</el-tag>
                <el-tag type="info" v-else-if="scope.row.results === 101"
                >旷课
                </el-tag
                >
                <el-tag type="success" v-else-if="scope.row.results >= 60">{{
                    scope.row.results
                  }}
                </el-tag>
                <el-tag type="warning" v-else>{{ scope.row.results }}</el-tag>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12">
        <el-card style="height: 360px">
          <el-collapse v-model="activeName" accordion>
            <el-collapse-item title="二课时" name="1">
              <el-table :data="tableData2" size="mini">
                <el-table-column prop="class" label="课程"></el-table-column>
                <el-table-column prop="date" label="时间"></el-table-column>
              </el-table>
            </el-collapse-item>
            <el-collapse-item title="四课时" name="2">
              <el-table :data="tableData4" size="mini">
                <el-table-column prop="class" label="课程"></el-table-column>
                <el-table-column prop="date" label="时间"></el-table-column>
              </el-table>
            </el-collapse-item>
          </el-collapse>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
// import musicPlayer from "./component/musicPlayer";
// import TodoList from "./component/todoList";
// import { mapGetters } from "vuex";
import {store} from "@/store";

export default {
  name: "Dashboard",
  data() {
    return {
      activeName: "1",
      name: store.state.user.userInfo.name,
      adminAnnouncement: "测试公告demo",
      tableData: [
        {
          date: "2021-08-07",
          place: "208",
          classname: "分光计分光计分光计分光计分光计分光计分光计分光计",
          results: 101,
        },
        {
          date: "2021-08-07",
          place: "208",
          classname: "分光计",
          results: 102,
        },
        {
          date: "2021-08-07",
          place: "208",
          classname: "分光计",
          results: 0,
        },
        {
          date: "2021-08-07",
          place: "208",
          classname: "分光计",
          results: 60,
        },
        {
          date: "2021-08-07",
          place: "208",
          classname: "分光计",
          results: 70,
        },
        {
          date: "2021-08-07",
          place: "208",
          classname: "分光计",
          results: 80,
        },
        {
          date: "2021-08-07",
          place: "208",
          classname: "分光计",
          results: 1,
        },
        {
          date: "2021-08-07",
          place: "208",
          classname: "分光计",
          results: 2,
        },
        {
          date: "2021-08-07",
          place: "208",
          classname: "分光计",
          results: 3,
        },
        {
          date: "2021-08-07",
          place: "208",
          classname: "分光计",
          results: 4,
        },
      ],
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
  computed: {
    // ...mapGetters("user", ["userInfo"]),
  },
  components: {
    // musicPlayer, //音乐播放器
    // TodoList, //TodoList
    // RaddarChart, //雷达图
    // stackMap, //堆叠图
    // Sunburst, //旭日图
  },
  methods: {
    // toTarget(name) {
    //   this.$router.push({ name });
    // },
  },
};
</script>

<style lang="scss" scoped>
.big {
  margin: 100px 0 0 0;
  padding-top: 0;
  background-color: rgb(243, 243, 243);
  padding-top: 15px;

  .top {
    width: 100%;
    height: 360px;
    margin-top: 20px;
    overflow: hidden;

    .chart-container {
      position: relative;
      width: 100%;
      height: 100%;
      padding: 20px;
      background-color: #fff;
    }
  }

  .mid {
    width: 100%;
    height: 380px;

    .chart-wrapper {
      height: 340px;
      background: #fff;
      padding: 16px 16px 0;
      margin-bottom: 32px;
    }
  }

  .bottom {
    width: 100%;
    height: 300px;
    // margin: 20px 0;
    .el-row {
      margin-right: 4px !important;
    }

    .chart-player {
      width: 100%;
      height: 270px;
      padding: 10px;
      background-color: #fff;
    }
  }
}
</style>
