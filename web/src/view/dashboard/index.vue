<template>
  <div class="big">
    <el-row :gutter="10">
      <el-col :xs="24" :sm="12" style="height:160px">
        <el-card>
          <!-- 这个name应该在userInfo里面 -->
          欢迎您，<p>{{name}}</p>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" style="height:160px">
        <el-card>
          <p>管理员公告</p>
          <el-divider></el-divider>
          <p>{{adminAnnouncement}}</p>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="10">
      <el-col :xs="24" :sm="12" style="height:160px">
        <el-card>
          <!-- 这个name应该在userInfo里面 -->
          欢迎您，<p>{{name}}</p>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" style="height:160px">
        <el-card>
          <p>管理员公告</p>
          <el-divider></el-divider>
          <p>{{adminAnnouncement}}</p>
        </el-card>
      </el-col>
    </el-row>

  </div>
</template>

<script>
// import musicPlayer from "./component/musicPlayer";
// import TodoList from "./component/todoList";
import { mapGetters } from "vuex";
export default {
  name: "Dashboard",
  data() {
    return {
      name:"初始名字",
      adminAnnouncement:"测试公告demo"
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo"]),
  },
  components: {
    // musicPlayer, //音乐播放器
    // TodoList, //TodoList
    // RaddarChart, //雷达图
    // stackMap, //堆叠图
    // Sunburst, //旭日图
  },
  methods: {
    toTarget(name) {
      this.$router.push({ name });
    },
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
