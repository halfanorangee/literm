<template>
  <div class="menu-side-div">
    <el-menu class="left-side-menu"
    :default-active="onRoutes"
    router
    >
      <template v-for="menu in menuInfo">
        <el-menu-item class="menu-item" :index="menu.index">
          <el-tooltip :content="getMenuTitle(menu.id)"
                      :hide-after="0"
                       placement="right-end"
          >
            <el-icon class="left-side-icon">
              <component :is="menu.icon"></component>
            </el-icon>
          </el-tooltip>
        </el-menu-item>
      </template>
    </el-menu>
  </div>

</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import {menuInfo} from "../../js/types/menus";

const route = useRoute();
const onRoutes = computed(() => {
  return route.path;
});

const getMenuTitle = (id: string):string => {
  const menu = menuInfo.find(item => item.id === id);
  return menu ? menu.title : '';
}

</script>

<style scoped>
  .menu-side-div {
    width: 40px;
  }

  .left-side-icon {
    margin: 0;
  }

  .menu-item {
    height: 40px;
    padding-left: 7px !important;
    padding-right: 33px;
  }

  .left-side-menu {
    width: 40px;
    min-height: 100vh;
    border-right-color: #e6e6e6;
    border-right-width: 1px;
  }
</style>