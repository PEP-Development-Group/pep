<template>
  <div>
    <el-form :model="config" label-width="100px">
      <el-form-item label="最大退课次数">
        <el-input-number v-model="config.system.cancel_nums"></el-input-number>
      </el-form-item>
      <el-form-item label="学期开始时间">
        <el-date-picker v-model="config.system.first_day" value-format="yyyy-MM-dd"></el-date-picker>
        <div style="line-height: 15px"> * 请选择本学期的第一周周一,将作为课程系统时间计算依据<br>* 请在添加课程之前修改此项目</div>
      </el-form-item>
      <el-form-item>
        <el-button @click="update" type="primary">提交</el-button>
      </el-form-item>
    </el-form>
    <el-card class="danger-zone">
      危险区域
      <el-divider></el-divider>
      <el-popover placement="top" v-model="deleteVisible" width="160">
        <p>确定要删除吗？</p>
        <div style="text-align: right; margin: 0">
          <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
          <el-button @click="initDB" size="mini" type="primary">确定</el-button>
        </div>
        <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">清空数据库</el-button>
      </el-popover>
      <br>
      <br>
      *清空数据表，所有学生，课程记录，选课记录
    </el-card>
  </div>
</template>

<script>
import {getSystemConfig, setSystemConfig} from "@/api/system";
import {delAll} from "@/api/course";

export default {
  name: "courseConfig",
  data() {
    return {
      config: {
        system: {}
      },
      deleteVisible: false
    }
  }, async created() {
    await this.initForm();
  },
  methods: {
    async initDB() {
      this.$confirm('你确定要清空数据库吗? 这个操作不可逆', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        showClose: false,
      }).then(async () => {
        const res = await delAll()
        if (res.code === 0) {
          this.$message({type: "success", message: "清空成功"})
        }
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '操作取消'
        });
      });


    },
    async initForm() {
      const res = await getSystemConfig();
      if (res.code == 0) {
        this.config = res.data.config;
      }
    },
    reload() {
    },
    async update() {
      const res = await setSystemConfig({config: this.config});
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "配置文件设置成功"
        });
        await this.initForm();
      }
    }
  }
}
</script>

<style scoped>
.danger-zone {
  max-width: 800px;
  border-color: #FF6666;
  border-style: dashed;
}
</style>