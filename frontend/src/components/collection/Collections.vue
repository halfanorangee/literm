<template>
  <div class="wrapper">
    <div class="collection-list-side">
      <div class="head-line">
        <span class="span-line" style="width: 249px">
          集合列表
        </span>
        <el-dropdown trigger="click" id="common-line-transition">
          <el-icon  class="option-button" style="padding: 13px 10px">
            <CirclePlusFilled />
          </el-icon>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="newCollectionVisible = true"><el-icon><FolderAdd /></el-icon>新增集合</el-dropdown-item>
              <el-dropdown-item @click="newConnType = true"><el-icon><DocumentAdd /></el-icon>新建连接
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>

      <el-menu style="border-width: 0">
        <template v-for="collection in collectionRels">
        <el-sub-menu :index="collection.ID.toString()" class="sub-menu" >
          <template #title>
            <span class="span-line" style="width: 300px" @contextmenu="collRightClick($event, collection)">
              <el-icon><Collection /></el-icon>{{ collection.CollectionName }}
            </span>
          </template>
          <template v-for="conn in collection.ConnInfos">

            <el-menu-item :index="collection.ID + '_' + conn.ID.toString()" class="menu-item" @contextmenu="connRightClick($event, conn)">
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

    <div class="right-side">
      <template v-if="newConnType === true">
        <Create/>
      </template>
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
    <!--集合右键菜单-->
    <context-menu
        v-model:show="collRightClickShow"
        :options="optionsComponent"
    >
      <context-menu-item label="删除合集" @click="collDelCheckVisible=true">删除合集</context-menu-item>
      <context-menu-item label="编辑" @click="collEditOption">编辑</context-menu-item>
    </context-menu>
<!--删除合集确认弹窗-->
    <el-dialog
        v-model="collDelCheckVisible"
        title="确认删除？"
        width="400"
    >
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="collDelCheckVisible = false, collDelID = 0">取消</el-button>
          <el-button type="primary" @click="collDelete">
            确认
          </el-button>
        </div>
      </template>
    </el-dialog>

    <!--连接右键菜单-->
    <context-menu
        v-model:show="connRightClickShow"
        :options="optionsComponent"
    >
      <context-menu-item label="删除连接" @click="connDelCheckVisible=true">删除连接</context-menu-item>
      <context-menu-item label="编辑" @click="connEditOption"></context-menu-item>
    </context-menu>
    <el-dialog
        v-model="connDelCheckVisible"
        title="确认删除？"
        width="400"
    >
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="connDelCheckVisible = false, connDelID = 0">取消</el-button>
          <el-button type="primary" @click="connDelete">
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
import {InsertCollection,
        QueryAllConnectionRel,
        DeleteCollection,
        DeleteConnection} from "../../../wailsjs/go/service/CollectionService"
import Create from "./Create.vue";

// 连接和集合关联数据
const collectionRels = ref([]);
// 新增集合弹窗是否可见
const newCollectionVisible = ref(false);
//新增集合名称
const newCollectionName = ref("");
//集合右键菜单是否可见
const collRightClickShow = ref(false);
//连接右键菜单是否可见
const connRightClickShow = ref(false);
//右键菜单配置
const optionsComponent = ref({
  zIndex: 3,
  x: 0,
  y: 0
});
//需要删除的合集信息
const collDelID = ref(0);
//需要删除的连接信息
const connDelID = ref(0);
//删除合集二次确认弹框
const collDelCheckVisible = ref(false);
//删除连接二次确认弹框
const connDelCheckVisible = ref(false);
//新增页面展示
const newConnType = ref(false);

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
    if (response.ResponseCode === '000') {
      newCollectionVisible.value = false;
      newCollectionName.value = "";
      await queryCollectionList();
    }
  } catch (error) {
    console.error('发生错误', error);
  }
}
//新建连接
const createConn = async () => {
  try {
    console.log("新建连接")
  } catch (error) {
    console.error('发生错误', error);
  }
}
//连接合集的右键菜单
const collRightClick = function (e : MouseEvent, collection : any) {
  //prevent the browser's default menu
  e.preventDefault();
  collRightClickShow.value = true;
  optionsComponent.value.x = e.clientX;
  optionsComponent.value.y = e.clientY;
  collDelID.value = collection.ID;
}
//连接信息的右键菜单
const connRightClick = function (e : MouseEvent, conn : any) {
  e.preventDefault();
  connRightClickShow.value = true;
  optionsComponent.value.x = e.clientX;
  optionsComponent.value.y = e.clientY;
  connDelID.value = conn.ID;
}
//删除合集
const collDelete = async () => {
  try {
    console.log("删除合集")
    var response = await DeleteCollection(collDelID.value);
    if (response.ResponseCode === '000') {
      collDelID.value = 0;
      collDelCheckVisible.value = false;
      await queryCollectionList();
    }
  } catch (error) {
    console.error('发生错误', error);
  }
}
//编辑合集
const collEditOption = function () {
  console.log("编辑集合")
}
//删除连接
const connDelete = async () => {
  try {
    console.log("删除连接")
    var response = await DeleteConnection(connDelID.value);
    if (response.ResponseCode === '000') {
      connDelID.value = 0;
      connDelCheckVisible.value = false;
      await queryCollectionList();
    }
  } catch (error) {
    console.error('发生错误', error);
  }
}
//编辑连接
const connEditOption = function () {
  console.log("编辑连接")
}

</script>

<style scoped>
.wrapper {
  display: grid;
  grid-template-columns: 300px 1fr; /* 左侧菜单固定宽度，右侧自适应 */
  height: 100vh; /* 让容器占据整个视口高度 */
}

.collection-list-side {
  border-right: 1px solid var(--cust-border-defualt-color);
  background-color: var(--cust-bg-defualt-color);
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
}

.option-button {
  padding: 11px 10px;
  margin: 0;
  width: 20px;
}
</style>