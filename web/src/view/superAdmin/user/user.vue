<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addUser" type="primary">添加用户</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="学号" min-width="150" prop="stu_id"></el-table-column>
      <el-table-column label="姓名" min-width="150" prop="name"></el-table-column>
      <el-table-column label="学院" min-width="150" prop="college"></el-table-column>
      <el-table-column label="专业" min-width="150" prop="major"></el-table-column>
      <el-table-column label="取消次数" min-width="150" prop="cancel_nums"></el-table-column>
      <el-table-column label="操作" min-width="150">
        <template slot-scope="scope">
          <el-button type="warning" icon="el-icon-edit" size="small" slot="reference" @click="modifyUser" class="option-btn">修改</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除此学生吗</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false" >取消</el-button>
              <el-button type="primary" size="mini" @click="deleteUser(scope.row)">确定</el-button>
            </div>

            <el-button type="danger" icon="el-icon-delete" size="small" slot="reference" class="option-btn">删除</el-button>
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
        <el-radio :label="0">学生</el-radio>
        <el-radio :label="1">教师</el-radio>
      </el-radio-group>
      <el-form :rules="rules" ref="userForm" :model="userInfo">
        <el-form-item label="姓名" label-width="80px" prop="name">
          <el-input v-model="userInfo.name"></el-input>
        </el-form-item>
        <el-form-item :label="getIdType" label-width="80px" prop="userName">
        <el-input v-model="userInfo.userName"></el-input>
      </el-form-item>
        <el-form-item label="学院" label-width="80px" prop="college" v-show="!userType">
          <el-input v-model="userInfo.college"></el-input>
        </el-form-item>
        <el-form-item label="专业" label-width="80px" prop="major" v-show="!userType">
          <el-input v-model="userInfo.major"></el-input>
        </el-form-item>
        <el-form-item label="身份证号" label-width="80px" prop="pid" v-show="!userType">
          <el-input v-model="userInfo.pid"></el-input>
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
<!--      TODO-->
      <el-form :rules="rules" ref="userForm" :model="userInfo">
        <el-form-item label="姓名" label-width="80px" prop="name">
          <el-input v-model="userInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="学号" label-width="80px" prop="userName">
          <el-input v-model="userInfo.userName"></el-input>
        </el-form-item>
        <el-form-item label="学院" label-width="80px" prop="college">
          <el-input v-model="userInfo.college"></el-input>
        </el-form-item>
        <el-form-item label="专业" label-width="80px" prop="major">
          <el-input v-model="userInfo.major"></el-input>
        </el-form-item>
        <el-form-item label="身份证号" label-width="80px" prop="pid">
          <el-input v-model="userInfo.pid"></el-input>
        </el-form-item>
        <el-form-item label="密码" label-width="80px" prop="password">
          <el-input v-model="userInfo.password"></el-input>
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
  setUserAuthority,
  register,
  deleteUser
} from "@/api/user";
import { getAuthorityList } from "@/api/authority";
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
// import CustomPic from "@/components/customPic";
// import ChooseImg from "@/components/chooseImg";
export default {
  name: "Api",
  mixins: [infoList],
  // components: { CustomPic,ChooseImg },
  data() {
    return {
      listApi: getUserList,
      path: path,
      authOptions: [],
      addUserDialog: false,
      modifyUserDialog: false,
      userType:0,
      userInfo: {
        userName: "",
        password: "",
        name: "",
        college: "",
        major: "",
        pid: "",
        authorityId: ""
      },
      rules: {
        userName: [
          { required: true, message: "请输入学号/工号", trigger: "blur" }
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          { min: 6, message: "最低6位字符", trigger: "blur" }
        ],
        pid: [
          { required: true, message: "请输入身份证号", trigger: "blur" },
          { min: 18,max: 18, message: "请输入正确的身份证号", trigger: "blur" }
        ],
        name: [
          { required: true, message: "请输入姓名", trigger: "blur" }
        ],
        college: [
          { required: true, message: "请输入学院", trigger: "blur" }
        ],
        major: [
          { required: true, message: "请输入专业", trigger: "blur" }
        ],
        authorityId: [
          { required: true, message: "请选择用户角色", trigger: "blur" }
        ]
      }
    };
  },
  computed: {
    ...mapGetters("user", ["token"]),
    getIdType:function (){
      return this.userType?"工号":"学号";
    }
  },
  methods: {
    openHeaderChange(){
      this.$refs.chooseImg.open()
    },
    setOptions(authData) {
      this.authOptions = [];
      this.setAuthorityOptions(authData, this.authOptions);
    },
    setAuthorityOptions(AuthorityData, optionsData) {
      AuthorityData &&
        AuthorityData.map(item => {
          if (item.children && item.children.length) {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              children: []
            };
            this.setAuthorityOptions(item.children, option.children);
            optionsData.push(option);
          } else {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName
            };
            optionsData.push(option);
          }
        });
    },
    async deleteUser(row) {
      const res = await deleteUser({ id: row.ID });
      if (res.code == 0) {
        this.getTableData();
        row.visible = false;
      }
    },
    async enterAddUserDialog() {
      this.$refs.userForm.validate(async valid => {
        if (valid) {
          const res = await register(this.userInfo);
          if (res.code == 0) {
            this.$message({ type: "success", message: "创建成功" });
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
          const res = await register(this.userInfo);
          if (res.code == 0) {
            this.$message({ type: "success", message: "修改成功" });
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
    modifyUser() {
      this.modifyUserDialog = true;
    },
    async changeAuthority(row) {
      const res = await setUserAuthority({
        uuid: row.uuid,
        authorityId: row.authority.authorityId
      });
      if (res.code == 0) {
        this.$message({ type: "success", message: "角色设置成功" });
      }
    }
  },
  async created() {
    this.getTableData();
    const res = await getAuthorityList({ page: 1, pageSize: 999 });
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
.option-btn{
  margin-right: 5px;
}

.user-type-radio{
  margin-bottom: 20px;
  margin-left: 80px;
}
</style>