<template>
  <div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-for="item in courseList" class="class-con">
        <template slot="title">
          <div class="class-title-con">
            <span class="class-title">{{ item.name }}</span>
            <el-tag effect="dark" size="small" class="hours-tag" type="success">{{ item.hours }}学时</el-tag>
          </div>
        </template>
        <el-card v-for="l in item.lessons" class="lesson" shadow="hover">

          <el-progress class="lesson-progress" type="circle" :width="18" :show-text="false"
                       :percentage="l.nowSize/l.maxSize*100"
                       :stroke-width="2"></el-progress>
          <span class="lesson-info">
          <span class="space">{{ l.time }}</span>
          <span class="space">{{ l.teacher }}</span>
          <el-tag effect="dark" size="mini" color="#79baca" class="space">{{ l.classroom }}</el-tag>
          <el-tag effect="dark" size="mini" type="warning" class="space">本周</el-tag>
          </span>
          <span class="lesson-op">
            <span class="progress">{{ l.nowSize }}/{{ l.maxSize }}</span>
          <el-button type="primary">选课</el-button>
          </span>
        </el-card>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script>
import {
  createClass,
  deleteClass,
  deleteClassByIds,
  updateClass,
  findClass,
  getClassList
} from "@/api/cls_class";  //  此处请自行替换地址
import {formatTimeToStr} from "@/utils/date";
import infoList from "@/mixins/infoList";
import {store} from '@/store/index'

export default {
  name: "Boats",
  mixins: [infoList],
  data() {
    return {
      listApi: getClassList,
      dialogFormVisible: false,
      type: "",
      activeNames: [1],
      deleteVisible: false,
      multipleSelection: [], formData: {
        bid: 0,
        bname: "",
        color: "",
        sex: false,

      },
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
    filterHandler(value, row, column) {
      const property = column['property'];
      return row[property] === value;
    },
    selectClass(id) {
      console.log(id)
      console.log(store.state.user.userInfo.username)
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.sex == "") {
        this.searchInfo.sex = null
      }
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteBoats(row);
      });
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length == 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
      this.multipleSelection.map(item => {
        ids.push(item.ID)
      })
      const res = await deleteBoatsByIds({ids})
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length == ids.length) {
          this.page--;
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateBoats(row) {
      const res = await findBoats({ID: row.ID});
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reboats;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        bid: 0,
        bname: "",
        color: "",
        sex: false,

      };
    },
    async deleteBoats(row) {
      const res = await deleteBoats({ID: row.ID});
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1) {
          this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createBoats(this.formData);
          break;
        case "update":
          res = await updateBoats(this.formData);
          break;
        default:
          res = await createBoats(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    },
    async created() {
      await this.getTableData();

    }
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
  margin-bottom: 3px;
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
  float: right;
}

.space {
  vertical-align: middle;
  margin: 6px;
}

.progress {
  font-weight: bold;
  font-size: small;
  margin: 5px;
}

</style>
