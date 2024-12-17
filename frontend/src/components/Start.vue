<template>
  <div class="window">
    <div class="init">
      <el-button type="info" id = "newConButton" @click="createConnection = true">新建连接</el-button>
      <br>
      <el-button type="info">设置</el-button>
    </div>
  </div>


  <el-dialog
      v-model="createConnection"
      title="新建连接"
      width="60%"
      height="60%"
      :before-close="handleClose"
      top="20vh"
  >
    <New-Connection/>
  </el-dialog>

</template>

<script setup>
import { ref } from 'vue'
import { ElMessageBox } from 'element-plus'
import NewConnection from "./NewConnection.vue";

const createConnection = ref(false)
</script>

<style>
</style>