<template>
  <div>
    <el-page-header @back="goBack" :content="lesson.cname"></el-page-header>

  </div>
</template>

<script>
import {getTeacherAClassStuList, findClass} from "@/api/course";

export default {
  name: "lessonView",
  data() {
    return {
      cid: 0,
      lesson: null,
      stuList: []
    }
  },
  methods: {
    goBack() {
      this.$router.back()
    }
  },
  async created() {
    this.cid = this.$router.currentRoute.query.cid
    this.lesson = (await findClass({ID: this.cid})).data.reclass
    this.stuList = (await getTeacherAClassStuList({"cid": this.cid})).data
  }
}
</script>

<style scoped>

</style>