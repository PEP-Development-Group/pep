<template>
  <div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-for="(item,i) in courseList" class="class-con">
        <template slot="title">
          <div class="class-title-con">
            <span class="class-title">{{ i }}</span>
            <el-tag effect="dark" size="small" class="hours-tag" type="info"
                    :color="(colors.duration[item.hours-1])">{{ item.hours }}学时
            </el-tag>
          </div>
        </template>
        <el-card v-for="l in item.List" class="lesson" shadow="hover">

          <el-progress class="lesson-progress" type="circle" :width="18" :show-text="false"
                       :percentage="selectPercent(l.now,l.max)"
                       :stroke-width="2"></el-progress>
          <span class="lesson-info">
          <span class="space">{{ l.time | formatDate }}</span>
          <span class="space">{{ l.teacher_name }}</span>
          <el-tag effect="light" size="mini" class="space">{{ l.class_room }}</el-tag>
          <el-tag effect="dark" size="mini" type="info" class="space" :color="(colors.time[0])">本周</el-tag>
          </span>
          <span class="lesson-op">
            <span class="progress">{{ l.now }}/{{ l.max }}</span>
          <el-button type="primary" v-if="l.selected===false" @click="selectCourse(l.id)">选课</el-button>
          <el-button type="danger" v-else @click="deleteCourse(l.id)">退选</el-button>
          </span>
        </el-card>
      </el-collapse-item>
    </el-collapse>
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
    }
  },
  computed: {},
  methods: {
    selectPercent: function (now, max) {
      if (now < 0) return 0
      if (max <= 0) return 0
      return now / max * 100;
    },
    async selectCourse(cid) {
      const d = {
        cid: cid,
        username: userInfo.username
      }
      const res = await SelectClass(d)
      if (res.code === 0) {
        this.$message({type: "success", message: "选课成功"})
      }
      await this.getList()
    },
    async deleteCourse(cid) {
      DeleteSelect({"username": userInfo.username, "cid": cid})
      await this.getList()
    },
    async getList() {
      const list = await GetClassListWithPerson();
      if (list.code === 0) {
        this.courseList = list.data.courses
      }
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

</style>
