<template>
  <div>
    <el-row :gutter="10">
      <el-col :xs="{span:24,offset:0}" :sm="{span:8}">
        <el-card style="margin-top:20px;">
          <div slot="header">前端</div>
          <div>
            <el-row>
              <el-col>

                <a href="https://shuangxunian.gitee.io/">
                  <el-avatar :src="imgwyl" class="org-img" shape="square"></el-avatar>
                </a>
              </el-col>
            </el-row>

            <el-row :gutter="10">
              <el-col>
                <b>霜序廿</b>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -15px">

              <el-col style="margin-top: 2px">
                <p>shuangxunian@gmail.com</p>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -10px">

              <el-col style="margin-top: -1px">
                <p>wls来写一段废话</p>
              </el-col>
            </el-row>
          </div>
        </el-card>


      </el-col>

      <el-col :xs="{span:24,offset:0}" :sm="{span:8}">
        <el-card style="margin-top:20px;">
          <div slot="header">前端</div>
          <div>
            <el-row>
              <el-col>
                <!-- 你的博客 -->
                <a href="">
                  <el-avatar :src="imgdjh" class="org-img" shape="square"></el-avatar>
                </a>
              </el-col>
            </el-row>

            <el-row :gutter="10">
              <el-col>
                <b>BladeHiker</b>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -15px">

              <el-col style="margin-top: 2px">
                <p>bladehiker@outlook.com</p>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -10px">
              <el-col style="margin-top: -1px">
                <p>这里写一段废话</p>
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>

      <el-col :xs="{span:24,offset:0}" :sm="{span:8}">
        <el-card style="margin-top:20px;">
          <div slot="header">后端</div>
          <div>
            <el-row>
              <el-col>
                <!-- 你的博客 -->
                <a href="">
                  <el-avatar :src="imgljk" class="org-img" shape="square"></el-avatar>
                </a>
              </el-col>
            </el-row>

            <el-row :gutter="10">
              <el-col>
                <b>sh1luo</b>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -15px">

              <el-col style="margin-top: 2px">
                <p>silo1999@163.com</p>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -10px">

              <el-col style="margin-top: -1px">
                <p>这里也写一段废话</p>
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>

    </el-row>
  </div>
</template>

<script>
export default {
  name: "About",
  data() {
    return {
      imgwyl: require('@/assets/imgwyl.jpg'),
      imgdjh: require('@/assets/imgdjh.jpg'),
      imgljk: require('@/assets/imgljk.jpg'),
    };
  },
  methods: {}
};
</script>

<style scoped>
.el-card {
  text-align: center;
}

.el-card b {
  font-size: x-large;
  font-weight: lighter;
}

.el-card p {
  font-size: small;
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

.org-img {
  height: 150px;
  width: 150px;
}

.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}
</style>