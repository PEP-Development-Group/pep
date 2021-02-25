<template>
  <div class="container">
    <el-row>
      <el-col :xs="{span:24}" :sm="{span:20,offset:2}">
        <h1 class="header">Designers and Developers.</h1>
        <h2 class="hh2">Backend</h2>
      </el-col>
    </el-row>

    <el-row>
      <el-col :xs="{span:24}" :sm="{span:20,offset:2}">
        <el-card>
          <el-row>
            <el-col :xs="{span:24}" :sm="{span:8}">
              <el-row>
                <el-col :offset="8" :span="8">
                  <el-avatar :src="imgljk" class="org-img" shape="square"></el-avatar>
                </el-col>
              </el-row>
            </el-col>
            <el-col :xs="{span:24}" :sm="{span:16}">
              <div class="right">
                <span class="desc">Name:  wasabi</span>
                <br />

                <span class="desc">
                      Blog:  <el-link type="primary" :underline="false" href="https://kcode.icu/" style="font-size: 18px;">https://kcode.icu/</el-link>
                  </span>
                <br />

                <span class="desc">
                      Email:  silo1999@163.com
                  </span>
                <br />

              </div>
            </el-col>
          </el-row>
          <el-row>
            <span class="desc">
                Discriptor:    He is SAU
                <i>Physics Lab Enrollment System</i> 's Maker, from the
                software engineering major 2018th,responsible for the
                backend development of the entire system.
            </span>
          </el-row>
        </el-card>
      </el-col>
    </el-row>

    <el-row>
      <el-col :xs="{span:24}" :sm="{span:20,offset:2}">
        <h2 class="hh2">Frontend</h2>
      </el-col>
    </el-row>



    <el-row>
      <el-col :xs="{span:24}" :sm="{span:20,offset:2}">
        <el-card>
          <el-row>
            <el-col :xs="{span:24}" :sm="{span:8}">
              <el-row>
                <el-col :offset="8" :span="8">
                  <el-avatar :src="imgdjh" class="org-img" shape="square"></el-avatar>
                </el-col>
              </el-row>
            </el-col>
            <el-col :xs="{span:24}" :sm="{span:16}">
              <div class="right">
                <span class="desc">Name:  BladeHiker</span>
                <br />

                <span class="desc">
                      Blog:  <el-link type="primary" :underline="false" href="http://www.bladehiker.com" style="font-size: 18px;">Website</el-link>
                  </span>
                <br />

                <span class="desc">
                      Email:bladehiker@outlook.com
                  </span>
                <br />

              </div>
            </el-col>
          </el-row>
          <el-row>
            <span class="desc">
                Discriptor:Designer and front-end developer of SAU Physics Lab Enrollment System.
            </span>
          </el-row>
        </el-card>
      </el-col>
    </el-row>




    <el-row>
      <el-col :xs="{span:24}" :sm="{span:20,offset:2}">
        <el-card>
          <el-row>
            <el-col :xs="{span:24}" :sm="{span:8}">
              <el-row>
                <el-col :offset="8" :span="8">
                  <el-avatar :src="imgwyl" class="org-img" shape="square"></el-avatar>
                </el-col>
              </el-row>
            </el-col>
            <el-col :xs="{span:24}" :sm="{span:16}">
              <div class="right">
                <span class="desc">Name:  霜序廿</span>
                <br />

                <span class="desc">
                      Blog:  <el-link type="primary" :underline="false" href="https://shuangxunian.gitee.io/" style="font-size: 18px;">Website</el-link>
                  </span>
                <br />

                <span class="desc">
                      Email:  shuangxunian@gmail.com
                  </span>
                <br />

              </div>
            </el-col>
          </el-row>
          <el-row>
            <span class="desc">
                Discriptor:   公者千古，私者一时
            </span>
          </el-row>
        </el-card>
      </el-col>
    </el-row>
    <!-- <el-row>
      <el-col>
        <el-card><el-avatar
              :src="imgdjh"
              class="org-img"
              shape="square"
          ></el-avatar>
          <div class="right">
              <span class="desc">Name: &nbsp;&nbsp;&nbsp; BladeHiker</span>
              <br />

              <span class="desc">
                  Blog: &nbsp;&nbsp;&nbsp; https://kcode.icu/
              </span>
              <br />

              <span class="desc">
                  Email: &nbsp;&nbsp;&nbsp; silo1999@163.com
              </span>
              <br />

              <span class="desc">
                  Discriptor: &nbsp; &nbsp; &nbsp; He is a gopher!
              </span>
          </div>
      </el-card>
      </el-col>
    </el-row> -->
    <!--

        <el-card class="p"
            ><el-avatar
                :src="imgwyl"
                class="org-img"
                shape="square"
            ></el-avatar>
            <div class="right">
                <span class="desc">Name: &nbsp;&nbsp;&nbsp; 霜序廿</span>
                <br />

                <span class="desc">
                    Blog: &nbsp;&nbsp;&nbsp; https://kcode.icu/
                </span>
                <br />

                <span class="desc">
                    Email: &nbsp;&nbsp;&nbsp; silo1999@163.com
                </span>
                <br />

                <span class="desc">
                    Discriptor: &nbsp; &nbsp; &nbsp; He is a gopher!
                </span>
            </div>
        </el-card> -->
  </div>
</template>

<script>
export default {
  name: "About",
  data() {
    return {
      imgwyl: require("@/assets/imgwyl.jpg"),
      imgdjh: require("@/assets/imgdjh.jpg"),
      imgljk: require("@/assets/imgljk.jpg"),
    };
  },
  methods: {},
};
</script>

<style scoped>
.hh2 {
  font-size: 20px;
  margin-bottom: 10px;
}

.p {
  padding-bottom: 15px;
  margin-bottom: 15px;
  margin-left: 10px;
}

/* .container {
  padding: 60px;
} */

.right{
  display:table;
  margin:20px auto;
}

.desc {
  /* margin-left: 30px; */
  font-size: 20px;
  /* margin-bottom: 50px; */
}

.org-img {
  height: 100px;
  width: 100px;
  float: left;
}

.header {
  font-size: 30px;
  font-weight: 800;
  margin-bottom: 30px;
}

.el-card b {
  font-size: x-large;
  font-weight: lighter;
}

.el-card p {
  font-size: small;
  line-height: normal;
}

.avatar-img {
  float: left;
  height: 40px;
  width: 40px;
  border-radius: 50%;
  -webkit-border-radius: 50%;
  -moz-border-radius: 50%;
  margin-top: 15px;
}

.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}
</style>