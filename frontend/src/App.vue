<script setup>
</script>

<template>
  <router-view></router-view>
  <el-row class="tac">
      <el-menu
          class="el-menu-vertical-demo"
          @open="handleOpen"
          @close="handleClose"
      >
        <el-menu-item index="1" id = "index-tab">
          <el-icon><Management /></el-icon>
        </el-menu-item>
        <el-menu-item index="2" id = "index-tab">
          <el-icon><Tools /></el-icon>
        </el-menu-item>
        <el-menu-item index="3" id = "index-tab">
          <el-icon><setting /></el-icon>
        </el-menu-item>
      </el-menu>
  </el-row>

</template>

<style>
.el-menu--horizontal {
  --el-menu-horizontal-height: 30px;
}

#index-tab {
  padding-left: 10px;
  padding-right: 10px;
  width: 40px;
  height: 40px;
}

#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
