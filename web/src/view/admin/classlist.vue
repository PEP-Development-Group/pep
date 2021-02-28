<template>
  <div>
    <div class="op-term">
      <el-button @click="openDialog" type="primary">新增实验课</el-button>
      <el-popover placement="top" v-model="deleteVisible" width="160">
        <p>确定要删除吗？</p>
        <div style="text-align: right; margin: 0">
          <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
          <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
        </div>
        <el-button icon="el-icon-delete" slot="reference" type="danger">批量删除</el-button>
      </el-popover>
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
      <el-table-column type="selection" width="55"></el-table-column>

      <el-table-column label="课程名" prop="cname" min-width="120"></el-table-column>

      <el-table-column label="学时" prop="ccredit" min-width="120"></el-table-column>

      <el-table-column label="教室" prop="classroom" min-width="120"></el-table-column>

      <el-table-column label="操作" width="150">
        <template slot-scope="scope">
          <el-button @click="updateClassList(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteClassList(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="课程操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="课程名:">
          <el-input v-model="formData.cname" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="学时:">
          <el-input v-model.number="formData.ccredit" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="教室:">
          <el-input v-model="formData.classroom" clearable placeholder="请输入"></el-input>
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
  createClassList,
  deleteClassList,
  deleteClassListByIds,
  updateClassList,
  findClassList,
  getClassListList
} from "@/api/course";  //  此处请自行替换地址
import {formatTimeToStr} from "@/utils/date";

export default {
  name: "ClassList",
  data() {
    return {
      tableData: [],
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [], formData: {
        cname: "",
        ccredit: "",
        classroom: "",

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
  methods: {
    async getClassTableData() {
      this.tableData = (await getClassListList()).data
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getClassTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    async onDelete() {
      const ids = []
      this.multipleSelection &&
      this.multipleSelection.map(item => {
        ids.push(item.ID)
      })
      const res = await deleteClassListByIds({ids})
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        this.deleteVisible = false
        await this.getClassTableData()
      }
    },
    async updateClassList(row) {
      const res = await findClassList({ID: row.id});
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reclasslist;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        cname: "",
        ccredit: null,
        classroom: "",

      };
    },
    async deleteClassList(row) {
      this.visible = false;
      const res = await deleteClassList({ID: row.id});
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        await this.getClassTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createClassList(this.formData);
          break;
        case "update":
          res = await updateClassList(this.formData);
          break;
        default:
          res = await createClassList(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "创建/更改成功"
        })
        this.closeDialog();
        await this.getClassTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getClassTableData();

  }
};
</script>

<style scoped>
.op-term {
  float: right;
  margin-bottom: 10px;
}

.op-term button {
  margin-left: 5px;
}
</style>