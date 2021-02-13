<template>
  <div>
    <el-row :gutter="10">
      <el-col :xs="{span:24,offset:0}" :sm="{span:8}">
        <el-card style="margin-top:20px;">
          <div slot="header">前端</div>
          <div>
            <el-row>
              <el-col :span="8" :offset="8">
                <a href="https://shuangxunian.gitee.io/">
                  <img
                      class="org-img dom-center"
                      src="@/assets/imgwyl.jpg"
                      alt=""
                  />
                </a>
              </el-col>
            </el-row>

            <el-row :gutter="10">
              <el-col :span="8" :offset="4">
                <b>姓&emsp;&emsp;名：</b>
              </el-col>

              <el-col :span="12">
                <b>王跃霖</b>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -20px">
              <el-col :span="8" :offset="4">
                <p>昵&emsp;&emsp;称：</p>
              </el-col>

              <el-col :span="12">
                <p>霜序廿</p>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -20px">
              <el-col :span="8" :offset="4">
                <p>班&emsp;&emsp;级：</p>
              </el-col>

              <el-col :span="12" style="margin-top: 2px">
                <p>CS1705</p>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -20px">
              <el-col :span="8" :offset="4">
                <p>联系方式：</p>
              </el-col>

              <el-col :span="12" style="margin-top: -1px">
                <p>shuangxunian@gmail.com</p>
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
              <el-col :span="8" :offset="8">
                <!-- 你的博客 -->
                <a href="">
                  <img
                      class="org-img dom-center"
                      src="@/assets/imgdjh.jpg"
                      alt=""
                  />
                </a>
              </el-col>
            </el-row>

            <el-row :gutter="10">
              <el-col :span="8" :offset="4">
                <b>姓&emsp;&emsp;名：</b>
              </el-col>

              <el-col :span="12">
                <b>段霁航</b>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -20px">
              <el-col :span="8" :offset="4">
                <p>昵&emsp;&emsp;称：</p>
              </el-col>

              <el-col :span="12">
                <p></p>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -20px">
              <el-col :span="8" :offset="4">
                <p>班&emsp;&emsp;级：</p>
              </el-col>

              <el-col :span="12" style="margin-top: 2px">
                <p></p>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -20px">
              <el-col :span="8" :offset="4">
                <p>联系方式：</p>
              </el-col>

              <el-col :span="12" style="margin-top: -1px">
                <p></p>
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
              <el-col :span="8" :offset="8">
                <!-- 你的博客 -->
                <a href="">
                  <img
                      class="org-img dom-center"
                      src="@/assets/imgljk.jpg"
                      alt=""
                  />
                </a>
              </el-col>
            </el-row>

            <el-row :gutter="10">
              <el-col :span="8" :offset="4">
                <b>姓&emsp;&emsp;名：</b>
              </el-col>

              <el-col :span="12">
                <b>刘继坤</b>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -20px">
              <el-col :span="8" :offset="4">
                <p>昵&emsp;&emsp;称：</p>
              </el-col>

              <el-col :span="12">
                <p></p>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -20px">
              <el-col :span="8" :offset="4">
                <p>班&emsp;&emsp;级：</p>
              </el-col>

              <el-col :span="12" style="margin-top: 2px">
                <p></p>
              </el-col>
            </el-row>

            <el-row :gutter="10" style="margin-top: -20px">
              <el-col :span="8" :offset="4">
                <p>联系方式：</p>
              </el-col>

              <el-col :span="12" style="margin-top: -1px">
                <p></p>
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>

    </el-row>
  </div>
</template>

<script>
import {Commits, Members} from "@/api/github";
import Timeline from "timeline-vuejs";

export default {
  name: "Test",
  components: {
    Timeline,
  },
  data() {
    return {
      messageWhenNoItems: "There arent commits",
      members: [],
      dataTimeline: [],
      page: 0,
    };
  },
  created() {
    this.loadCommits();
    this.loadMembers();
  },
  methods: {
    loadMore() {
      this.page++;
      this.loadCommits();
    },
    loadCommits() {
      Commits(this.page).then(({data}) => {
        data.forEach((element) => {
          if (element.commit.message) {
            this.dataTimeline.push({
              from: new Date(element.commit.author.date),
              title: element.commit.author.name,
              showDayAndMonth: true,
              description: `<a style="color: #26191b" href="${element.html_url}">${element.commit.message}</a>`,
            });
          }
        });
      });
    },
    loadMembers() {
      Members().then(({data}) => {
        this.members = data;
        this.members.sort();
      });
    },
  },
};
</script>

<style scoped>
.load-more {
  margin-left: 120px;
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

.author-name {
  float: left;
  line-height: 65px !important;
  margin-left: 10px;
  color: darkblue;
  line-height: 100px;
  font-family: "Lucida Sans", "Lucida Sans Regular", "Lucida Grande",
  "Lucida Sans Unicode", Geneva, Verdana, sans-serif;
}

.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}
</style>