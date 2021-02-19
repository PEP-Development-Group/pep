<template>
  <div>
    <div class="button-box clearflex">
      <el-radio-group v-model="userType" size="small" @change="typeChanged">
        <el-radio-button :label="1">学生</el-radio-button>
        <el-radio-button :label="2">教师</el-radio-button>
      </el-radio-group>
      <el-button @click="addUser" type="primary">添加用户</el-button>
    </div>

    <el-table ref="userTable" :data="tableData" border stripe>
      <el-table-column :label="getIdType" min-width="150" prop="username"></el-table-column>
      <el-table-column label="姓名" min-width="150" prop="name"></el-table-column>
      <el-table-column label="取消次数" min-width="150" prop="cancel_nums" v-if="userType===1"></el-table-column>
      <el-table-column label="应修学分" min-width="150" prop="total_credits" v-if="userType===1"></el-table-column>
      <el-table-column label="操作" width="300" align="center">
        <template slot-scope="scope">
          <el-button type="warning" icon="el-icon-edit" size="small" slot="reference" @click="modifyUser(scope.row)"
                     class="option-btn">修改
          </el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除此学生吗</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteUser(scope.row)">确定</el-button>
            </div>

            <el-button type="danger" icon="el-icon-delete" size="small" slot="reference" class="option-btn">删除
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
          <el-input v-model="userInfo.name"></el-input>
        </el-form-item>
        <el-form-item :label="getIdType" label-width="80px" prop="userName">
          <el-input v-model="userInfo.username"></el-input>
        </el-form-item>
        <el-form-item label="应修学时" min-width="150" label-width="80px" v-show="userType===1">
          <el-input v-model="userInfo.total_credits" type="number"></el-input>
        </el-form-item>
        <el-form-item label="密码" label-width="80px" prop="password">
          <el-input v-model="userInfo.password" type="password"></el-input>
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
        <el-form-item :label="getIdType" label-width="80px" prop="username">
          <el-input v-model="userInfo.username"></el-input>
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
  setUserInfo
} from "@/api/user";
import {getAuthorityList} from "@/api/authority";
import infoList from "@/mixins/infoList";
import {mapGetters} from "vuex";

export default {
  name: "Api",
  mixins: [infoList],
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
        total_credits: 48
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
      }
    };
  },
  computed: {
    ...mapGetters("user", ["token"]),
    getIdType: function () {
      return this.userType === 2 ? "工号" : "学号";
    }
  },
  methods: {

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
        total_credits: 48
      }
    },
    async deleteUser(row) {
      const res = await deleteUser({id: row.ID});
      if (res.code == 0) {
        this.getTableData();
        row.visible = false;
      }
    },
    async enterAddUserDialog() {
      this.userInfo.authorityId = this.userType.toString();
      this.$refs.userForm.validate(async valid => {
        if (valid) {
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
      this.addUserDialog = false;
    },
    async enterModifyUserDialog() {
      this.$refs.userForm.validate(async valid => {
        if (valid) {
          const res = await setUserInfo(this.userInfo);
          if (res.code == 0) {
            this.$message({type: "success", message: "修改成功"});
          }
          await this.getTableData();
          this.closeModifyUserDialog();
        }
      });
    },
    closeModifyUserDialog() {
      this.$refs.userForm.resetFields();
      this.modifyUserDialog = false;
    },
    addUser() {
      this.addUserDialog = true;
    },
    modifyUser(row) {
      this.userInfo = row
      this.modifyUserDialog = true;
    },
  },
  async created() {
    this.getTableData();
    const res = await getAuthorityList({page: 1, pageSize: 999});
    this.setOptions(res.data.list);
  }
};
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

.option-btn {
  margin-right: 5px;
}

.user-type-radio {
  margin-bottom: 20px;
  margin-left: 80px;
}
</style>