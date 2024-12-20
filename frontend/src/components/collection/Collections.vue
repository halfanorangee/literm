<template>
  <div class="wrapper">
    <div class="collection-list-side">
    </div>
  </div>
</template>

<script setup>

</script>

<style scoped>
.wrapper {
  display: grid;
  grid-template-columns: 300px 1fr; /* 左侧菜单固定宽度，右侧自适应 */
  height: 100vh; /* 让容器占据整个视口高度 */
}
</style>