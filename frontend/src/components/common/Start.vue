<template>
  <div class="wrapper">
    <menu-side/>
    <div class="content-box" >
      <router-view v-slot="{ Component }">
        <transition name="move" mode="out-in">
          <component :is="Component"></component>
        </transition>
      </router-view>
    </div>
  </div>
</template>

<script setup>
import MenuSide from '../common/MenuSide.vue'
</script>

<style scoped>
.wrapper {
  display: grid;
  grid-template-columns: 40px 1fr; /* 左侧菜单固定宽度，右侧自适应 */
  height: 100vh; /* 让容器占据整个视口高度 */
}
</style>