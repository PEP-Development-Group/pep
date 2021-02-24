<template>
  <div>
    <div class=" operator-box clearflex">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="课程名">
          <el-input placeholder="搜索条件" v-model="searchInfo.cname"></el-input>
        </el-form-item>
        <el-form-item label="教师">
          <el-input placeholder="搜索条件" v-model="searchInfo.tname"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" icon="el-icon-search" type="text">搜索</el-button>
        </el-form-item>
        <div class="right-op">
          <el-form-item>
            <el-button @click="openDialog" type="success" size="mini">新增课程</el-button>
          </el-form-item>
          <el-form-item>
            <el-popover placement="top" v-model="deleteVisible" width="160">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
              <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
            </el-popover>
          </el-form-item>
        </div>

      </el-form>
    </div>
    <el-table
        :data="tableData"
        @selection-change="handleSelectionChange"
        border
        ref="multipleTable"
        stripe
        style="width: 100%"
        tooltip-effect="dark"
    >
      <el-table-column type="selection" width="50" align="center"></el-table-column>
      <el-table-column label="课程名" prop="cname" min-width="180" sortable></el-table-column>
      <el-table-column label="任课教师" prop="tname" min-width="120" sortable></el-table-column>
      <el-table-column label="学时" prop="ccredit" min-width="80" sortable filter-placement="top"
                       :filters="[{text: '2学时',value: 2},{text: '4学时',value: 4}]"
                       :filter-method="filterHandler"></el-table-column>
      <el-table-column label="选课开始" prop="stime" min-width="120" sortable>
        <template slot-scope="scope">
          <div class="text-wrapper">
            {{ scope.row.stime|formatDate }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="选课结束" prop="etime" min-width="120" sortable>
        <template slot-scope="scope">
          <div class="text-wrapper">
            {{ scope.row.etime|formatDate }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="上课时间" prop="time" min-width="120" sortable>
        <template slot-scope="scope">
          <el-popover trigger="hover" :content="scope.row.time|formatDateDay" placement="right">
            <div slot="reference" class="text-wrapper">
              {{ scope.row.desc | formatDesc }}
            </div>
          </el-popover>
        </template>
      </el-table-column>
      <el-table-column label="教室" prop="classroom" min-width="120" sortable></el-table-column>
      <el-table-column label="人数" prop="total" min-width="80" sortable></el-table-column>
      <el-table-column label="操作" width="200" align="center">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateClass(scope.row)" size="small" type="primary"
                     icon="el-icon-edit">变更
          </el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :style="{float:'right',padding:'20px'}"
        :total="total"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
        layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="课程操作">
      <el-form :model="formData" label-position="right" label-width="80px" :inline="false">
        <el-form-item label="课程名:">
          <el-input v-model="formData.cname" clearable placeholder="请输入"></el-input>
        </el-form-item>
        <el-form-item label="教师:">
          <el-input v-model="formData.tname" clearable placeholder="请输入"></el-input>
        </el-form-item>
        <el-form-item label="教室:">
          <el-input v-model="formData.classroom" clearable placeholder="请输入"></el-input>
        </el-form-item>
        <el-form-item label="总人数:">
          <el-input-number v-model="formData.total" clearable placeholder="请输入"></el-input-number>
        </el-form-item>
        <el-form-item label="选课时间:">
          <el-date-picker type="datetimerange" range-separator="至"
                          start-placeholder="开始时间"
                          end-placeholder="结束时间" placeholder="选择时间" v-model="timespan" clearable>

          </el-date-picker>
        </el-form-item>

        <el-form-item label="上课时间:">
          <el-input placeholder="请输入周数" v-model="week" type="number">
            <template slot="append">周</template>
          </el-input>
        </el-form-item>
        <el-form-item label="星期:">
          <el-radio-group size="small" v-model="dayOfWeek">
            <el-radio-button v-for="(item,i) in dayOfWeekOP" :key="i" :label="item.value">
              {{ item.label }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="课时:" class="inline-item">
          <el-input-number v-model="formData.ccredit"></el-input-number>
        </el-form-item>
        <el-form-item label="节数:" class="inline-item">
          <el-input-number v-model="classTime"></el-input-number>
        </el-form-item>
        <el-divider v-if="description"></el-divider>
        <el-card
            v-if="description&&formData.cname&&formData.ccredit&&formData.tname&&formData.classroom&&formData.total"
            class="summary">
          <template slot="header">
            <span>课程 : {{ formData.cname }}</span>&nbsp;
            <el-tag size="mini" effect="dark">{{ formData.ccredit }}学时</el-tag>
            <span style="float: right;color: #666666">摘要</span>
          </template>
          <span>
          {{ description|formatDesc }}
          </span>
          <span>{{ formData.tname }}</span>
          <el-tag effect="light" size="mini">{{ formData.classroom }}</el-tag>
          <br>
          <span>{{ realTime |formatDateDay }}</span>
          <br>
          <span>最大选课人数 : {{ formData.total }}</span>
          <br>
          <span>选课开放时间 : {{ timespan[0]|formatDate }} ~ {{ timespan[1]|formatDate }}</span>
        </el-card>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createClass,
  deleteClass,
  deleteClassByIds,
  updateClass,
  findClass,
  getClass
} from "@/api/course";  //  此处请自行替换地址
import {formatTimeToStr, schoolTimeToRealTime} from "@/utils/date";
import infoList from "@/mixins/infoList";
import {store} from "@/store";

const formatDayOfWeek = ['一', '二', '三', '四', '五', '六', '日']
export default {
  name: "Class",
  mixins: [infoList],
  data() {
    return {
      listApi: getClass,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        ccredit: 2,
        cname: "",
        etime: null,
        stime: null,
        time: null,
        tname: "",
        desc: "",
        classroom: "",
        total: 1
      },
      classTime: 1,
      dayOfWeek: null,
      dayOfWeekOP: [
        {
          "label": "一",
          "value": 1
        }, {
          "label": "二",
          "value": 2
        }, {
          "label": "三",
          "value": 3
        }, {
          "label": "四",
          "value": 4
        }, {
          "label": "五",
          "value": 5
        }, {
          "label": "六",
          "value": 6
        }, {
          "label": "日",
          "value": 7
        }
      ],
      timespan: [],
      week: 1,
    };
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy年MM月dd日 \n hh:mm:ss");
      } else {
        return "";
      }
    },
    formatDateDay: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy年MM月dd日");
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
  computed: {
    description() {
      if (this.week && this.dayOfWeek && this.classTime)
        return this.week + "-" + this.dayOfWeek + "-" + this.classTime
      return null
    },
    realTime() {
      return schoolTimeToRealTime(this.description, store.state.user.firstDay)
    }
  },
  watch: {},
  methods: {
    filterHandler(value, row, column) {
      const property = column['property'];
      return row[property] === value;
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
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
        this.deleteClass(row);
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
      const res = await deleteClassByIds({ids})
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
    async updateClass(row) {
      const res = await findClass({ID: row.ID});
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reclass;
        this.timespan = [this.formData.stime, this.formData.etime]
        let descList = this.formData.desc.split('-')
        this.week = descList[0]
        this.dayOfWeek = descList[1]
        this.classTime = descList[2]
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        ccredit: 0,
        cname: "",
        etime: null,
        stime: null,
        time: null,
        tname: "",

      };
    },
    async deleteClass(row) {
      const res = await deleteClass({ID: row.ID});
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
    pretreatment() {
      this.formData.stime = new Date(this.timespan[0])
      this.formData.etime = new Date(this.timespan[1])
      this.formData.desc = this.description
      this.formData.time = this.realTime
    },
    async enterDialog() {
      let res;
      this.pretreatment()
      switch (this.type) {
        case "create":
          res = await createClass(this.formData);
          break;
        case "update":
          res = await updateClass(this.formData);
          break;
        default:
          res = await createClass(this.formData);
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
    }
  },
  async created() {
    await this.getTableData();
  }
};
</script>

<style lang="scss" scoped>
.el-input {
  width: auto;
}

.operator-box {
  margin-left: 10px;

  .right-op {
    float: right;
  }
}

.text-wrapper, .text-wrapper .cell {
  white-space: pre-wrap !important;
}

.el-date-editor .el-range__icon, .el-date-editor .el-range-separator {
  line-height: 22px;
}

.inline-item {
  display: inline-block;
}

.el-form-item__content .el-input-group {
  vertical-align: middle;
}

.summary .el-card__header span:nth-child(1) {
  font-size: 16px;
  vertical-align: middle;
}

.summary .el-card__body {
  line-height: 30px;
  font-size: 18px;
  font-weight: lighter;
}

.summary span {
  margin-right: 5px;
}

.summary .el-tag {
  vertical-align: middle;
}
</style>
