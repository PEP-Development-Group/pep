<template>
  <div id="userLayout" class="user-layout-wrapper">
    <el-card class="container">
      <div class="top">
        <div class="desc">
          <img class="logo_login" src="@/assets/saulogo.png" alt=""/>
        </div>
        <div class="header">
          <span class="title">物理实验选课系统</span>
        </div>
      </div>
      <div class="main">
        <el-form
            :model="loginForm"
            :rules="rules"
            ref="loginForm"
            @keyup.enter.native="submitForm"
        >
          <el-form-item prop="username">
            <el-input placeholder="请输入用户名" v-model.number="loginForm.username" type="tel">
              <i class="el-input__icon el-icon-user" slot="suffix"></i
              ></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
                :type="lock === 'lock' ? 'password' : 'text'"
                placeholder="请输入密码"
                v-model="loginForm.password"
            >
              <i
                  :class="'el-input__icon el-icon-' + lock"
                  @click="changeLock"
                  slot="suffix"
              ></i>
            </el-input>
          </el-form-item>
          <el-form-item style="position: relative" prop="captcha">
            <el-input
                v-model="loginForm.captcha"
                name="logVerify"
                placeholder="请输入验证码"
                style="width: 60%"
                type="tel"
            />
            <div class="vPic">
              <img
                  v-if="picPath"
                  :src="picPath"
                  width="100%"
                  height="100%"
                  alt="请输入验证码"
                  @click="loginVefify()"
              />
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm" style="width: 100%" :loading="loading"
            >登 录
            </el-button
            >
          </el-form-item>
        </el-form>
      </div>

      <div class="footer">
        <div class="copyright">Copyright &copy; {{ curYear }} <span>沈阳航空航天大学物理实验中心</span></div>
      </div>
    </el-card>
  </div>
</template>

<script>
import {mapActions} from "vuex";
import {captcha} from "@/api/user";

export default {
  name: "Login",
  data() {
    // const checkUsername = (rule, value, callback) => {
    //   if (value.length < 3 || value.length > 15) {
    //     return callback(new Error("请输入正确的用户名"));
    //   } else {
    //     callback();
    //   }
    // };
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6 || value.length > 32) {
        return callback(new Error("请输入正确的密码"));
      } else {
        callback();
      }
    };
    const checkCaptcha = (rule, value, callback) => {
      if (value.length < 5) {
        return callback(new Error("请输入验证码"));
      } else {
        callback();
      }
    };
    return {
      curYear: 0,
      lock: "lock",
      loginForm: {
        username: null,
        password: "",
        captcha: "",
        captchaId: "",
      },
      rules: {
        username: [{required: true, message: "请输入用户名", trigger: "blur"}],
        password: [{validator: checkPassword, trigger: "blur"}],
        captcha: [{validator: checkCaptcha, trigger: "blur"}]
      },
      logVerify: "",
      picPath: "",
      loading: false
    };
  },
  created() {
    this.loginVefify();
    this.curYear = new Date().getFullYear();
  },
  methods: {
    ...mapActions("user", ["LoginIn"]),
    async login() {
      this.loginForm.password = this.$md5(this.loginForm.password)
      return await this.LoginIn(this.loginForm);
    },
    async submitForm() {
      this.loading = true
      this.$refs.loginForm.validate(async (v) => {
        if (v) {
          const pwd = this.loginForm.password
          const flag = await this.login();
          if (!flag) {
            this.loginForm.captcha = ""
            this.loginForm.password = pwd
            this.loginVefify();
          }
        } else {
          this.$message({
            type: "error",
            message: "请正确填写登录信息",
            showClose: true,
          });
          this.loading = false
          return false;
        }
        this.loading = false
      });

    },
    changeLock() {
      this.lock === "lock" ? (this.lock = "unlock") : (this.lock = "lock");
    },
    loginVefify() {
      captcha({}).then((ele) => {
        this.picPath = ele.data.picPath;
        this.loginForm.captchaId = ele.data.captchaId;
      });
    },
  },
};
</script>

<style scoped lang="scss">
@import "@/style/login.scss";
</style>
