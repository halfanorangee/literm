<template>
  <div class="wrapper">
    <div class="collection-list-side">
      <div class="head-line">
        <span>
          集合列表
        </span>
        <el-button style="border-width: 0; background: none; float: right" circle >
          <el-icon><CirclePlusFilled /></el-icon>
        </el-button>
      </div>
      <el-menu>
        <template v-for="connInfo in connInfos">
          <el-menu-item :index="connInfo.ID" class="menu-item">
            <span class="span-line">
              <el-icon><Connection /></el-icon>
              {{ connInfo.ConnName }}
            </span>
            <el-button class="optin-button" style="border-width: 0; float: right" >
              <el-popover
                  placement="bottom"
                  title="Title"
                  :width="200"
                  trigger="click"
                  content="this is content, this is content, this is content"
              >
                <template #reference>
                  <el-icon><MoreFilled /></el-icon>
                </template>
              </el-popover>
            </el-button>
          </el-menu-item>
        </template>
      </el-menu>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import {onMounted} from "vue";

import {QueryConnInfos} from "../../../wailsjs/go/service/CollectionService"

const connInfos = ref([]);
const loading = ref(true);

onMounted(async () => {
  try {
    connInfos.value = await QueryConnInfos();
    console.log("查询完成:" + connInfos.value)
  } catch (error) {
    console.error('发生错误', error);
  } finally {
    loading.value = false;
  }
});
</script>

<style scoped>
.wrapper {
  display: grid;
  grid-template-columns: 300px 1fr; /* 左侧菜单固定宽度，右侧自适应 */
  height: 100vh; /* 让容器占据整个视口高度 */
}

.collection-list-side {
  background-color: #ffffff;
  overflow-y: auto;
}

.head-line {
  flex-grow: 1;
  height: 40px;
  line-height: 40px;
  padding-left: 10px;
  border-bottom: 1px solid #e6e6e6;
}

.menu-item {
  height: 40px;
  padding: 0 !important;
  border-width: 0 !important;
  width: 100%;
}

.span-line {
  float: left;
  width: 100%;
}

.optin-button {
  border-width: 0;
  width: 20px;
  height: 20px;
  float: right;
}
</style>