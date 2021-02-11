<template>
  <div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-for="item in tableData" class="class-con">
        <template slot="title">
          <div class="class-title-con">
            <span class="class-title">{{ item.cname }}</span>
            <el-tag effect="dark" size="small" class="hours-tag" type="success">{{ item.ccredit }}学时</el-tag>
          </div>
        </template>
        <el-card v-for="l in item.lessons" class="lesson" shadow="hover">

          <el-progress class="lesson-progress" type="circle" :width="18" :show-text="false"
                       :percentage="l.nowSize/l.maxSize*100"
                       :stroke-width="2"></el-progress>
          <span class="lesson-info">
          <span class="space">{{ l.time }}</span>
          <span class="space">{{ l.tname }}</span>
          <el-tag effect="dark" size="mini" color="#79baca" class="space">{{ l.classroom }}</el-tag>
          <el-tag effect="dark" size="mini" type="warning" class="space">本周</el-tag>
          </span>
          <span class="lesson-op">
            <span class="progress">{{ l.nowSize }}/{{ l.maxSize }}</span>
          <el-button type="primary" v-if="l.selected===0">选课</el-button>
          <el-button type="danger" v-else>推选</el-button>
          </span>
        </el-card>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script>
import {
  getClassList
} from "@/api/boats.js";  //  此处请自行替换地址
import {formatTimeToStr} from "@/utils/date";
import infoList from "@/mixins/infoList";
import {store} from '@/store'

export default {
  name: "Boats",
  mixins: [infoList],
  data() {
    return {
      listApi: getClassList,
      activeNames: [1],
      courseList: [
        {
          id: 1,
          name: "椭圆偏振光法测定介质薄膜的厚度和折射率",
          hours: 2,
          lessons: [
            {
              id: 2,
              teacher: "栾玉国",
              time: "12周 周四 34节",
              classroom: "A205",
              maxSize: 40,
              nowSize: 26,
              startTime: "2021-11-1",
              endTime: "2021-11-15"
            }, {
              id: 2,
              type: "lesson",
              teacher: "栾玉国",
              time: "12周 周四 34节",
              classroom: "A205",
              maxSize: 50,
              nowSize: 34,
              startTime: "2021-11-1",
              endTime: "2021-11-15"
            }, {
              id: 2,
              type: "lesson",
              teacher: "栾玉国",
              time: "12周 周四 34节",
              classroom: "A205",
              maxSize: 20,
              nowSize: 16,
              startTime: "2021-11-1",
              endTime: "2021-11-15"
            }, {
              id: 2,
              type: "lesson",
              teacher: "栾玉国",
              time: "12周 周四 34节",
              classroom: "A205",
              maxSize: 40,
              nowSize: 26,
              startTime: "2021-11-1",
              endTime: "2021-11-15"
            }, {
              id: 2,
              type: "lesson",
              teacher: "王老师",
              time: "12周 周四 34节",
              classroom: "A205",
              maxSize: 40,
              nowSize: 1,
              startTime: "2021-11-1",
              endTime: "2021-11-15"
            }
          ]
        }, {
          id: 1,
          type: "course",
          name: "分光计",
          hours: 4,
          lessons: [
            {
              id: 2,
              type: "lesson",
              teacher: "栾玉国",
              time: "12周 周四 34节",
              classroom: "A205",
              maxSize: 40,
              nowSize: 26,
              startTime: "2021-11-1",
              endTime: "2021-11-15"
            }, {
              id: 2,
              type: "lesson",
              teacher: "王老师",
              time: "12周 周四 34节",
              classroom: "A205",
              maxSize: 40,
              nowSize: 26,
              startTime: "2021-11-1",
              endTime: "2021-11-15"
            }
          ]
        }
      ]
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
    }
  },
  methods: {
    // filterHandler(value, row, column) {
    //   const property = column['property'];
    //   return row[property] === value;
    // },
    // selectClass(id) {
    //   console.log(id)
    //   console.log(store.state.user.userInfo.username)
    // },
    // async deleteBoats(row) {
    //   const res = await deleteBoats({ID: row.ID});
    //   if (res.code == 0) {
    //     this.$message({
    //       type: "success",
    //       message: "删除成功"
    //     });
    //     if (this.tableData.length == 1) {
    //       this.page--;
    //     }
    //     this.getTableData();
    //   }
    // },
  },
  async created() {
    console.log("created")
    await this.getTableData();
  }

}
</script>

<style>
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

</style>
