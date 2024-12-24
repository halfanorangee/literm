<template>
  <div class="wrapper">
    <div class="collection-list-side">
      <div class="collection-head-class" id="collection-head-id">
        集合列表
        <el-button class="create-button" circle >
          <el-icon><CirclePlusFilled /></el-icon>
        </el-button>
      </div>
      <div class="collection-list-class">
        <div v-for="connInfo in connInfos" :key="connInfo.ID">
          <el-button class="collection-button" circle>
            <el-icon><Collection /></el-icon>
          </el-button>
          <span class="collection-name">{{ connInfo.ConnName }}</span>
        </div>
      </div>
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

.collection-head-class {
  line-height: 30px;
  padding: 5px;
  height: 30px;
  border: solid lightgray;
}

#collection-head-id {
  margin-bottom: 5px;
}

.create-button {
  border-width: 0;
  padding: 5px;
  float: right;

}
</style>