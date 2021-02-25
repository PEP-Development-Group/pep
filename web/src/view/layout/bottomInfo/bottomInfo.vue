<template>
  <div class="bottom-info">
    <div>
      <el-link href="#" :underline="false">沈阳航空航天大学物理实验中心</el-link>
      <el-divider direction="vertical"></el-divider>
      <el-link :underline="false" @click="toTarget">物理实验平台开发团队</el-link>
    </div>
  </div>
</template>
<script>
export default {
  name: "BottomInfo",
  methods: {
    toTarget() {
      this.$router.push('/about')
    },
  },
}
</script>
<style lang="scss">
.bottom-info {
  color: #888;
  height: 30px;
  line-height: 12px;

  a {
    color: #888;
  }

  div {
    display: flex;
    justify-content: center;

    span {
      margin: 0 3px;
    }
  }
}
</style>