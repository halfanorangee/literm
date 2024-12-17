<script setup>

</script>

<template>
  <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect">
    <el-menu-item index="1">SSH</el-menu-item>
    <el-menu-item index="2">SFTP</el-menu-item>
  </el-menu>

</template>

<style scoped>

</style>