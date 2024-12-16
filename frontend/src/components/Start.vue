<template>
  <div class="window">
    <div class="init">
      <lay-button id = "start_button" @click="openNewConnection"><lay-icon type="layui-icon-addition"></lay-icon> 新连接</lay-button>
      <br>
      <lay-button id = "start_button" @click="openSettings" ><lay-icon type="layui-icon-set-sm"></lay-icon> 设置</lay-button>
    </div>
  </div>

  <div id = "newConnection" style="display: none">
  </div>
</template>

<script setup>
import {layer} from "@layui/layui-vue";

const openNewConnection = function() {
  layer.open({
    type: 1,
    title: "新建连接",
    area: ['70%','70%'],
    content: "http://www.layui-vue.com"
  })
}

  const openSettings = function() {
    console.log("弹窗")
    layer.open({
      type: "iframe",
      title: "远程页面",
      area: ['90%','90%'],
      content: "http://www.layui-vue.com"
    })
}

// export default {
//   methods: {
//     openIframe() {
//       layer.open({
//         type: "iframe",
//         title: "远程页面",
//         area: ['90%','90%'],
//         content: "http://www.layui-vue.com"
//       })
//     }
//   }
// }
</script>

<style>
#start_button{

}
</style>