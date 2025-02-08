<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <el-radio-group v-model="connType" style="margin: 20px 30px">
          <el-radio-button value="1">新建ssh</el-radio-button>
          <el-radio-button value="2">新建ftp</el-radio-button>
        </el-radio-group>
      </el-header>
      <el-main v-if="connType == '1'">
        <CreateSSH/>
      </el-main>
      <el-main v-if="connType == '2'">
        <CreateFTP/>
      </el-main>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import CreateSSH from "./CreateSSH.vue";
import CreateFTP from "./CreateFTP.vue";

//1-ssh 2-ftp
const connType = ref('1')
</script>

<style scoped>
.common-layout {
  display: grid;
  background-color: var(--cust-bg-defualt-color);
  height: 100vh;
}
</style>