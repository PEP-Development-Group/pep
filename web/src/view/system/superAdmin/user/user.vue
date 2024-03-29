<template>
  <div>
    <div class="button-box clearflex">
      <el-radio-group v-model="userType" size="small" @change="typeChanged">
        <el-radio-button :label="1">学生</el-radio-button>
        <el-radio-button :label="2">教师</el-radio-button>
      </el-radio-group>
      <el-popover placement="top" v-model="MultiDeleteVisible" width="160">
        <p>确定要删除吗？</p>
        <div style="text-align: right; margin: 0">
          <el-button @click="MultiDeleteVisible = false" size="mini" type="text">取消</el-button>
          <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
        </div>
        <el-button icon="el-icon-delete" slot="reference" type="danger">批量删除</el-button>
      </el-popover>
      <el-button @click="addUser" type="success" style="margin-right: 5px">添加用户</el-button>
    </div>

    <el-table ref="userTable" :data="tableData" border stripe @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="50" align="center"></el-table-column>
      <el-table-column :label="getIdType" min-width="150" prop="username"></el-table-column>
      <el-table-column label="姓名" min-width="150" prop="name"></el-table-column>
      <el-table-column label="剩余取消次数" min-width="150" prop="cancel_nums" v-if="userType===1"></el-table-column>
      <el-table-column label="已修 / 应修学时" min-width="150" prop="total_credits" v-if="userType===1">
        <template slot-scope="scope">
          {{ scope.row.have_credits ? scope.row.have_credits : 0 }} / {{ scope.row.total_credits }}
        </template>
      </el-table-column>
      <el-table-column label="操作" width="350" align="center">
        <template slot-scope="scope">
          <el-button type="primary" icon="el-icon-plus" size="small" slot="reference"
                     @click="modifyUserCancel(scope.row)"
                     v-if="userType===1">增加取消次数
          </el-button>
          <el-button type="warning" icon="el-icon-edit" size="small" slot="reference" @click="modifyUser(scope.row)"
                     class="opt-btn">修改信息
          </el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除此用户吗</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteUser(scope.row)">确定</el-button>
            </div>

            <el-button type="danger" icon="el-icon-delete" size="small" slot="reference" class="opt-btn">删除
            </el-button>
          </el-popover>
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

    <el-dialog :visible.sync="addUserDialog" custom-class="user-dialog" title="添加用户">
      <el-radio-group v-model="userType" class="user-type-radio">
        <el-radio :label="1">学生</el-radio>
        <el-radio :label="2">教师</el-radio>
      </el-radio-group>
      <el-form :rules="rules" ref="userForm" :model="userInfo">
        <el-form-item label="姓名" label-width="80px" prop="name">
          <el-input v-model="userInfo.name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="应修学时" min-width="150" label-width="80px" prop="total_credits" v-show="userType===1">
          <el-input v-model.number="userInfo.total_credits" type="number" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item :label="getIdType" label-width="80px" prop="username">
          <el-input v-model.number="userInfo.username" autocomplete="off" type="tel"></el-input>
        </el-form-item>
        <el-form-item label="密码" label-width="80px" prop="password">
          <el-input v-model="userInfo.password" type="password" autocomplete="new-password"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeAddUserDialog">取 消</el-button>
        <el-button @click="enterAddUserDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
    <el-dialog :visible.sync="modifyUserDialog" custom-class="user-dialog" title="修改信息">
      <el-form :rules="rulesModify" ref="userForm" :model="userInfo">
        <el-form-item label="姓名" label-width="80px" prop="name">
          <el-input v-model="userInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="应修学时" label-width="80px" prop="total_credits" v-show="userType===1">
          <el-input v-model.number="userInfo.total_credits" type="tel"></el-input>
        </el-form-item>
        <el-form-item :label="getIdType" label-width="80px" prop="username">
          <el-input v-model.number="userInfo.username" type="tel"></el-input>
        </el-form-item>
        <el-form-item label="密码" label-width="80px" prop="password">
          <el-input placeholder="未修改" v-model="userInfo.password" type="password"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeModifyUserDialog">取 消</el-button>
        <el-button @click="enterModifyUserDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import {
  getUserList,
  register,
  deleteUser,
  setUserInfo,
  deleteUserByIds
} from "@/api/user";
import infoList from "@/mixins/infoList";
import {mapGetters} from "vuex";
import {addUserCancelNums} from "@/api/course";

export default {
  name: "Api",
  mixins: [infoList],
  async created() {
    await this.getTableData();
  },
  data() {
    return {
      listApi: getUserList,
      path: path,
      authOptions: [],
      addUserDialog: false,
      modifyUserDialog: false,
      userType: 1,
      userInfo: {
        username: "",
        password: "",
        name: "",
        authorityId: "",
        total_credits: null
      },
      rules: {
        username: [
          {required: true, message: "请输入学号/工号", trigger: "blur"}
        ],
        password: [
          {required: true, message: "请输入密码", trigger: "blur"},
          {min: 6, message: "最低6位字符", trigger: "blur"}
        ],
        name: [
          {required: true, message: "请输入姓名", trigger: "blur"}
        ],
      },
      rulesModify: {
        username: [
          {required: true, message: "请输入学号/工号", trigger: "blur"}
        ],
        password: [
          {min: 6, message: "最低6位字符", trigger: "blur"}
        ],
        name: [
          {required: true, message: "请输入姓名", trigger: "blur"}
        ],
      },
      MultiDeleteVisible: false,
      multipleSelection: [],
    };
  },
  computed: {
    ...mapGetters("user", ["token"]),
    getIdType: function () {
      return this.userType === 2 ? "工号" : "学号";
    }
  },
  methods: {
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length == 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的用户'
        })
        return
      }
      this.multipleSelection &&
      this.multipleSelection.map(item => {
        ids.push(item.username)
      })
      const res = await deleteUserByIds({ids})
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length == ids.length) {
          this.page--;
        }
        this.deleteVisible = false
        await this.getTableData()
      }
    },
    async getTableData(page = this.page, pageSize = this.pageSize, param = this.userType) {
      const table = await this.listApi({page, pageSize, param})
      if (table.code == 0) {
        this.tableData = table.data.list
        this.total = table.data.total
        this.page = table.data.page
        this.pageSize = table.data.pageSize
      }
    },
    typeChanged() {
      this.tableData = null
      this.getTableData()
    },
    resetForm() {
      this.userInfo = {
        username: "",
        password: "",
        name: "",
        pid: "",
        authorityId: "",
        total_credits: ""
      }
    },
    async deleteUser(row) {
      const res = await deleteUser({id: row.username});
      if (res.code == 0) {
        this.getTableData();
        row.visible = false;
      }
    },
    async enterAddUserDialog() {
      this.userInfo.authorityId = this.userType.toString();
      this.$refs.userForm.validate(async valid => {
        if (valid) {
          this.userInfo.password = this.$md5(this.userInfo.password)
          const res = await register(this.userInfo);
          if (res.code == 0) {
            this.$message({type: "success", message: "创建成功"});
          }
          await this.getTableData();
          this.closeAddUserDialog();
        }
      });
    },
    closeAddUserDialog() {
      this.$refs.userForm.resetFields();
      this.resetForm()
      this.addUserDialog = false;
    },
    async enterModifyUserDialog() {
      if (this.userInfo.password)
        this.userInfo.password = this.$md5(this.userInfo.password)
      const res = await setUserInfo(this.userInfo);
      if (res.code == 0) {
        this.$message({type: "success", message: "修改成功"});
      }
      await this.getTableData();
      this.closeModifyUserDialog();

    },
    closeModifyUserDialog() {
      this.$refs.userForm.resetFields();
      this.resetForm()
      this.modifyUserDialog = false;
    },
    addUser() {
      this.resetForm()
      this.resetForm()
      this.addUserDialog = true;
    },
    modifyUser(row) {
      this.resetForm()
      this.userInfo = row
      this.modifyUserDialog = true;
    },
    modifyUserCancel(row) {
      this.$prompt("输入要增加的次数", '修改 ' + row.name + " 的取消选课次数", {
        confirmButtonText: '修改',
        cancelButtonText: '取消',
      }).then(async ({value}) => {
        let res = await addUserCancelNums({username: row.username, cnt: parseInt(value)})
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '修改成功'
          });
          this.getTableData()
        } else {
          this.$message.error('修改失败');
        }
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '取消修改'
        });
      });
    }
  }
}
;
</script>
<style lang="scss">

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

.opt-btn {
  margin-left: 5px !important;
}

.user-type-radio {
  margin-bottom: 20px;
  margin-left: 80px;
}
</style>