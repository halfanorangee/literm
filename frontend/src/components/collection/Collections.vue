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
      <div style="overflow-y: auto">
      <el-menu style="border: 0!important; ">
        <template v-for="collection in collections">
          <el-sub-menu :index="stringIndex(collection.ID)" class="menu-item">
            <template #title>
            <span class="span-line" @click="queryConnInfoList(collection.ID)">
              <el-icon><Collection /></el-icon>
              {{ collection.CollectionName }}
            </span>
            </template>
            <el-dropdown trigger="click" id="common-line-transition">
              <el-icon  class="option-button">
                <MoreFilled />
              </el-icon>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item><el-icon><Edit /></el-icon>编辑</el-dropdown-item>
                  <el-dropdown-item><el-icon><Delete /></el-icon>删除</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>

            <template v-for="conn in connInfos">
            <el-menu-item-group class="menu-item">
                  <template v-if="conn.Collection_ID.Int16 === collection.ID">
                    <span class="span-line" id="common-line-transition">
                      <el-icon><Connection /></el-icon>
                      {{ conn.ConnName }}
                    </span>
                  </template>
            </el-menu-item-group>
            </template>

          </el-sub-menu>
        </template>
      </el-menu>
      </div>
    </div>

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
import {QueryCollections, QueryConnInfos, InsertCollection} from "../../../wailsjs/go/service/CollectionService"
import {menuInfo} from "../../js/types/menus";

const connInfos = ref([]);
const collections = ref([]);
const newCollectionVisible = ref(false);
const newCollectionName = ref("");

//获取string类型index
const stringIndex = (index: number) => {
  return index.toString();
}

//查询合集列表
const queryCollectionList = async () => {
  try {
    collections.value = await QueryCollections();
    console.log("查询完成:" + collections.value)
  } catch (error) {
    console.error('发生错误', error);
  }
}
//初始页面加载
onMounted(async () => {
  try {
    collections.value = await QueryCollections();
    console.log("查询完成:" + collections.value)
  } catch (error) {
    console.error('发生错误', error);
  }
});

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

//查询连接信息列表
const queryConnInfoList = async (collectionId: number) => {
  try {
    connInfos.value = await QueryConnInfos(collectionId);
    console.log("查询连接信息:")
    console.log(connInfos.value)
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