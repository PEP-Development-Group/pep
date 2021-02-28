<template>
  <div>
    <div class="button-box clearflex">
      <el-radio-group v-model="classType" size="small">
        <el-radio-button :label="1">已上课程</el-radio-button>
        <el-radio-button :label="2">未来课程</el-radio-button>
      </el-radio-group>
    </div>

    <el-table ref="userTable" :data="[courseList,classType]|classFilter" border stripe>
      <el-table-column label="课程" min-width="150">
        <template slot-scope="scope">
          {{ scope.row.cname }}
          <el-tag effect="dark" size="mini" class="left-space" type="success">{{ scope.row.ccredit }}学时
          </el-tag>
          <el-tag effect="light" size="mini" class="left-space">{{ scope.row.classroom }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="时间" min-width="150">
        <template slot-scope="scope">
          <el-popover trigger="hover" :content="scope.row.time|formatDateDay" placement="right">
            <div slot="reference" class="text-wrapper">
              {{ scope.row.desc | formatDesc }}
            </div>
          </el-popover>
        </template>
      </el-table-column>
      <el-table-column label="选课人数" min-width="150" prop="selected"></el-table-column>
      <el-table-column label="操作" width="300" align="center">
        <template slot-scope="scope">
          <el-button type="primary" icon="el-icon-s-order" plain size="small" slot="reference"
                     @click="viewLesson(scope.row.cid)"
                     class="option-btn">查看
          </el-button>
        </template>

      </el-table-column>
    </el-table>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
import {getTeacherClassList} from "@/api/course";

import {mapGetters} from "vuex";
import {formatTimeToStr} from "@/utils/date";

const formatDayOfWeek = ['一', '二', '三', '四', '五', '六', '日']
export default {
  name: "courseManage",
  data() {
    return {
      courseList: [],
      classType: 1
    };
  },
  filters: {
    classFilter(list) {
      let currTime = new Date()
      if (list[0])
        return list[0].filter((item) => {
          let o = new Date(item.time) < currTime
          return ((list[1] === 1) && o) || ((list[1] === 2) && !o)
        })
      return []
    }, formatDesc: function (d) {
      if (d) {
        let descList = d.split('-')
        return "第" + descList[0] + "周 周" + formatDayOfWeek[descList[1] - 1] + " 第" + descList[2] + "节";
      }
    }, formatDateDay: function (time) {
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
  },
  methods: {
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
  },
  async created() {
    await this.getList()
  }
};
</script>
<style lang="scss" scoped>
.left-space {
  float: right;
  margin-left: 10px;
}

.button-box {
  padding: 10px 20px;

  .el-button {
    float: right;
  }
}

.user-dialog {
  .header-img-box {
    width: 200px;
    height: 200px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 200px;
    cursor: pointer;
  }

  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }

  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }

  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}

.option-btn {
  margin-right: 5px;
}

.user-type-radio {
  margin-bottom: 20px;
  margin-left: 80px;
}
</style>