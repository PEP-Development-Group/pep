<template>
  <div class="center">
    <el-upload
      :action="`${path}/fileUploadAndDownload/upload`"
      :before-remove="beforeRemove"
      :file-list="fileList"
      :headers="{'x-token':token}"
      :limit="1"
      :on-exceed="handleExceed"
      :on-remove="handleRemove"
      class="upload-demo upload-con"
      accept=".xls,.xlsx"
      drag
    >
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
      <div class="el-upload__tip" slot="tip">只能上传.xls和.xlsx格式文件</div>
    </el-upload>
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
const path = process.env.VUE_APP_BASE_API
export default {
  name: 'Excel',
  data() {
    return {
      path: path,

      fileList: [
        // {
        //   name: 'food.jpeg',
        //   url:
        //     'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'
        // }
      ]
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'token'])
  },
  methods: {
    handleRemove(file) {
      this.$message.warning(
        `移除了${file.name}`
      )
    },
    handleExceed() {
      this.$message.warning(
        `每次只能上传一个文件`
      )
    },
    beforeRemove(file) {
      return this.$confirm(
        `确定移除 ${file.name}？`
      )
    }
  }
}
</script>
<style type="text/css">
  .center{
    text-align: center;
  }
  .upload-con{
    width: 360px;
    display: inline-block;
  }
  .upload-demo .el-upload-list__item-status-label, .upload .el-upload-list__item-status-label{
    left: auto;
  }
</style>