<template>
  <div class="wrapper">
    <div class="collection-list-side">
      <div class="collection-head-class" id="collection-head-id">
        这里是集合列表的头部
        <el-button class="create-button" circle >
          <el-icon><CirclePlusFilled /></el-icon>
        </el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { getConnInfos } from "../../js/types/collections";
import {ref} from 'vue'
import {onMounted} from "vue";

const connInfos = ref([]);
const loading = ref(true);

onMounted(async () => {
  try {
    connInfos.value = await getConnInfos();
  } catch (error) {
    console.log('发生错误');
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
  padding: 2px;
  border: solid gray;
  border-right-width: 1.5px;
  border-top-width: 1.5px;
  border-bottom-width: 1.5px;
  border-left-width: 1px;
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