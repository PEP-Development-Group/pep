<template>
  <div>
    <div class=" operator-box clearflex">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="课程名">
          <el-input placeholder="搜索条件" v-model="searchInfo.cname"></el-input>
        </el-form-item>
        <el-form-item label="教师名">
          <el-input placeholder="搜索条件" v-model="searchInfo.tname"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" icon="el-icon-search" type="text">搜索</el-button>
        </el-form-item>
        <div class="right-op">
          <el-form-item>
            <el-button @click="openDialog" type="primary">新增课程</el-button>
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
      <el-table-column label="任课教师" prop="tname" min-width="120" sortable></el-table-column>
      <el-table-column label="课程名" prop="cname" min-width="180" sortable></el-table-column>
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
          <div class="text-wrapper">
            {{ scope.row.time|formatDate }}
          </div>
        </template>
      </el-table-column>
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
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="学分:">
          <el-input v-model.number="formData.ccredit" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="课程名:">
          <el-input v-model="formData.cname" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="选课结束:">
          <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.etime" clearable></el-date-picker>
        </el-form-item>

        <el-form-item label="选课开始:">
          <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.stime" clearable></el-date-picker>
        </el-form-item>

        <el-form-item label="上课时间:">
          <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.time" clearable></el-date-picker>
        </el-form-item>

        <el-form-item label="教师名:">
          <el-input v-model="formData.tname" clearable placeholder="请输入"></el-input>
        </el-form-item>
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
  getClassList
} from "@/api/course";  //  此处请自行替换地址
import {formatTimeToStr} from "@/utils/date";
import infoList from "@/mixins/infoList";

export default {
  name: "Class",
  mixins: [infoList],
  data() {
    return {
      listApi: getClassList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [], formData: {
        ccredit: 2,
        cname: "",
        etime: new Date(),
        stime: new Date(),
        time: new Date(),
        tname: "",

      }
    };
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd \n hh:mm:ss");
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
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        ccredit: 0,
        cname: "",
        etime: new Date(),
        stime: new Date(),
        time: new Date(),
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
    async enterDialog() {
      let res;
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

<style lang="scss">
.operator-box {
  margin-left: 10px;

  .right-op {
    float: right;
  }
}

.text-wrapper, .text-wrapper .cell {
  white-space: pre-wrap !important;
}
</style>
