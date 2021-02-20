<template>
  <el-table
      :data="tableData"
      border
      size="medium"
      class="classtable"
  >
    <el-table-column prop="desc" label="时间" width="130px">
      <template slot-scope="scope">
        {{ scope.row.desc|formatDesc }}
      </template>
    </el-table-column>
    <el-table-column prop="classroom" label="地点" width="90px">
    </el-table-column>
    <el-table-column prop="tname" label="教师" width="80px"></el-table-column>
    <el-table-column prop="hours" label="学时" width="80px"></el-table-column>
    <el-table-column prop="cname" label="课程名" width="200px"></el-table-column>
    <el-table-column prop="grade" label="成绩" align="center" width="100%">
      <template slot-scope="scope">
        <el-tag v-if="scope.row.grade === 102">未上成绩</el-tag>
        <el-tag type="info" v-else-if="scope.row.grade === 101"
        >旷课
        </el-tag>
        <el-tag type="success" v-else-if="scope.row.grade >= 60"
        >{{ scope.row.grade }}
        </el-tag>
        <el-tag type="warning" v-else>{{ scope.row.grade }}</el-tag>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import {GetPersonalClasses} from "@/api/course";

const formatDayOfWeek = ['一', '二', '三', '四', '五', '六', '日']
export default {
  name: "lessons",
  filters: {
    formatDesc: function (d) {
      if (d) {
        let descList = d.split('-')
        return "第" + descList[0] + "周 周" + formatDayOfWeek[descList[1] - 1] + " 第" + descList[2] + "节";
      }
    }
  },
  data() {
    return {
      tableData: []
    }
  },
  async created() {
    this.tableData = (await GetPersonalClasses()).data.list.crs
  }
}
</script>

<style scoped>

</style>