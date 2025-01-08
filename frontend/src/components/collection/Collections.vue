<template>
  <div class="wrapper">
    <div class="collection-list-side">
      <div class="head-line">
        <span class="span-line" style="width: 250px;">
          集合列表
        </span>
        <el-dropdown trigger="click" id="common-line-transition">
          <el-icon  class="option-button" style="padding: 13px 10px;">
            <CirclePlusFilled />
          </el-icon>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="newCollectionVisible = true"><el-icon><FolderAdd /></el-icon>新增集合</el-dropdown-item>
              <el-dropdown-item><el-icon><DocumentAdd /></el-icon>新建连接</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>

      <el-menu>
        <template v-for="collection in collectionRels">
        <el-sub-menu :index="collection.ID.toString()" class="sub-menu">
          <template #title>
            <span><el-icon><Collection /></el-icon>
                {{ collection.CollectionName }}</span>
          </template>
          <template v-for="conn in collection.ConnInfos">
            <el-menu-item :index="collection.ID + '_' + conn.ID.toString()" class="menu-item">
                <span class="span-line">
                  <el-icon><Connection /></el-icon>
                  {{ conn.ConnName }}
                </span>
            </el-menu-item>
          </template>
        </el-sub-menu>
        </template>
      </el-menu>
    </div>

<!--新增合集弹窗-->
    <el-dialog
        v-model="newCollectionVisible"
        title="请输入集合名称"
        width="450"
    >
      <el-input v-model="newCollectionName"></el-input>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="newCollectionVisible = false">取消</el-button>
          <el-button type="primary" @click="createCollection">
            确认
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import {onMounted} from "vue";
import {InsertCollection, QueryAllConnectionRel} from "../../../wailsjs/go/service/CollectionService"

const collectionRels = ref([]);
const newCollectionVisible = ref(false);
const newCollectionName = ref("");

//初始页面加载
onMounted(async () => {
  try {
    collectionRels.value = await QueryAllConnectionRel();
  } catch (error) {
    console.error('发生错误', error);
  }
});
//查询合集列表
const queryCollectionList = async () => {
  try {
    collectionRels.value = await QueryAllConnectionRel();
    console.log("查询完成:" + collectionRels.value)
  } catch (error) {
    console.error('发生错误', error);
  }
}
//创建集合
const createCollection = async () => {
  try {
    const response = await InsertCollection(newCollectionName.value);
    console.log(response)
    if (response.ResponseCode === 200) {
      newCollectionVisible.value = false;
      newCollectionName.value = "";
      await queryCollectionList();
    }
  } catch (error) {
    console.error('发生错误', error);
  }
}
</script>

<style scoped>
.wrapper {
  display: grid;
  grid-template-columns: 300px 1fr; /* 左侧菜单固定宽度，右侧自适应 */
  height: 100vh; /* 让容器占据整个视口高度 */
}

.collection-list-side {
  background-color: #ffffff;
  max-height: 100vh;
}

.head-line {
  flex-grow: 1;
  height: 40px;
  line-height: 40px;
  padding-left: 10px;
  border-bottom: 1px solid var(--cust-border-defualt-color);
}

.sub-menu {
  display: block;
  overflow: visible;
}

.menu-item {
  height: 40px;
  padding: 0 !important;
  border-width: 0 !important;
  color: var(--cust-font-icon-defualt-color);
  width: 100%;
}

.span-line {
  float: left;
  height: 40px;
  line-height: 40px;
  width: 260px;
}

.option-button {
  padding: 11px 10px;
  margin: 0;
  width: 20px;
}
</style>