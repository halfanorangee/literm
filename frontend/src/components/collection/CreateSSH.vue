<template>
  <div class="wrapper">
    <el-form
        ref="ruleFormRef"
        style="max-width: 800px"
        :model="ruleForm"
        :rules="rules"
        class="new-ssh-form"
        label-width="auto"
        :size="'large'"
        :inline="true"
        status-icon
    >
      <el-form-item label="所属集合：" class="long-input">
<!--        <el-select v-model="ruleForm.collection" placeholder="请选择所属集合">
          <template v-for="collection in collSelections">
            <el-option label="'coll_' + collection.ID" value="collection.CollectionName">

            </el-option>
          </template>
        </el-select>-->
      </el-form-item>
      <el-form-item label="连接名称：" class="long-input" prop="name">
        <el-input
            maxlength="20"
            show-word-limit
            v-model="ruleForm.Name"
            clearable
        >
        </el-input>
      </el-form-item>
      <el-form-item label="备注：" class="long-input" prop="comment">
        <el-input
            maxlength="50"
            show-word-limit
            v-model="ruleForm.Comment"
            clearable
            id="common-form-member">
        </el-input>
      </el-form-item>
      <el-form-item label="地址：" prop="IP" class="short-input">
        <el-input v-model="ruleForm.IP" clearable></el-input>
      </el-form-item>
      <el-form-item label="端口：" prop="port" class="short-input">
        <el-input v-model="ruleForm.Port" clearable></el-input>
      </el-form-item>

      <el-form-item label="验证类型：" >
        <el-radio-group v-model="authType">
          <el-radio-button value=1>用户名密码</el-radio-button>
          <el-radio-button value=2>公钥</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <!--用户名密码登录-->
      <template v-if="authType === AuthType.Password">
        <el-form-item label="用户名：" prop="UserName" class="long-input">
          <el-input v-model="ruleForm.UserName"
                    placeholder="请输入用户名" clearable></el-input>
        </el-form-item>
        <el-form-item label="密码：" prop="Password" class="long-input">
          <el-input v-model="ruleForm.Password"
                    type="password"
                    show-password
                    placeholder="请输入密码"
                    clearable
          >
          </el-input>
        </el-form-item>

      </template>
      <!--公钥登录-->
      <el-form-item v-if="authType === AuthType.PublicKey" label="公钥：" class="long-input" prop="Authkey">
        <el-input v-model="ruleForm.Authkey" type="textarea" />
      </el-form-item>

      <el-form-item style="width: 100%">
        <!--测试连接成功/失败弹窗-->
        <el-popover
            placement="bottom"
            :title=testResult.result
            :width="200"
            trigger="click"
            :content=testResult.description
        >
          <template #reference>
            <el-button @click="testConnection">测试连接</el-button>
          </template>
        </el-popover>
        <el-button @click="saveAndConnect" style="float: right">保存并连接</el-button>
        <el-button @click="saveConnection" style="float: right">保存</el-button>
      </el-form-item>
    </el-form>


  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import {ElMessage, ElNotification, FormInstance, FormRules} from 'element-plus'
import { TestConnection } from '../../../wailsjs/go/service/ShellService'
import { InsertConnection } from '../../../wailsjs/go/service/CollectionService'
import {service, types} from "../../../wailsjs/go/models";



enum AuthType {
  Password = "1",
  PublicKey = "2"
}
const authType = ref(AuthType.Password)
const ruleFormRef = ref<FormInstance>()
const collSelections = ref([])
const testResult = ref({
  result: '',
  description: '',
})
const ruleForm = ref<types.ConnInsertInfo>({
  CollectionId: 0,
  Name: '',
  Comment: '',
  IP: '',
  Port: '',
  UserName: '',
  Password: '',
  Authkey: '',
  AuthType: '',
})

const rules = reactive<FormRules<types.ConnInsertInfo>>({
  CollectionId: [
    { required: true, message: '请选择所属集合', trigger: 'change' },
  ],
  Name: [
    { required: true, message: '请输入连接名称', trigger: 'blur' },
    { max: 20, message: '连接名称应控制在20字及以下', trigger: 'blur' },
  ],
  Comment: [
    { max: 50, message: '备注应控制在50字及以下', trigger: 'blur' },
  ],
  IP: [
    { required: true, message: '请输入IP地址', trigger: 'blur' },
  ],
  Port: [
    { required: true, message: '请输入端口号', trigger: 'blur' },
    {
      pattern: /^([1-9]\d{0,3}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$/,
      message: '端口范围为1-65535'
    }
  ],
})

const testConnection = async () => {
  try {
    ruleForm.value.AuthType = authType.value
    var response = await TestConnection(ruleForm.value);
    console.log("测试连接")
    if (response.ResponseCode == '000') {
      testResult.value.result = response.ResponseMsg
      testResult.value.description = ''
    }else {
      testResult.value.result = response.ResponseMsg
      testResult.value.description = response.ResponseInfo
    }
  }catch (e) {
    console.log('出现错误，'+e)
  }
}

const saveAndConnect = function () {
  console.log("保存并连接")
}

const saveConnection = async () => {
  try {
    const connInfo = service.ConnInfo.createFrom();
    const form = ruleForm.value;
    connInfo.Collection_ID = form.CollectionId
    connInfo.ConnName = form.Name
    connInfo.IP = form.IP
    connInfo.Port = form.Port
    connInfo.AuthType = form.AuthType
    connInfo.UserName = form.UserName
    connInfo.Password = form.Password
    connInfo.AuthKey = form.Authkey
    // var response = await InsertConnection(connInfo);
    const response = {
      ResponseCode: "000",
      ResponseMsg: "保存成功",
      ResponseInfo: "保存成功"
    };
    if (response.ResponseCode == '000') {
      ElMessage({
        type: response.ResponseCode === '000' ? 'success' : 'error',
        message: response.ResponseMsg
      });
      ruleFormRef.value?.resetFields();
    }else {
      testResult.value.result = response.ResponseMsg
      testResult.value.description = response.ResponseInfo
    }
  }catch (e) {
    ElNotification({
      title: '请求异常',
      message: e.message,
      type: 'error'
    });
    console.log('出现错误，'+e)
  }
  console.log("保存连接")
}

</script>

<style scoped>
.wrapper {
  display: flex;
  justify-content: center; /* 水平居中 */
  padding-top: 5rem;
  max-height: 100vh;
}

.new-ssh-form .long-input {
  width: 100%;
}

.new-ssh-form .short-input {
  width: 300px;
}

</style>