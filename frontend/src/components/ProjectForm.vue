<script setup lang="ts">
import { ref } from 'vue'
import { useProjectStore } from '../stores/project'

const projectStore = useProjectStore()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const name = ref('')
const description = ref('')
const password = ref('')
const confirmPassword = ref('')
const error = ref('')

async function handleSubmit() {
  if (!name.value || !password.value) {
    error.value = '请填写项目名称和密码'
    return
  }

  if (password.value !== confirmPassword.value) {
    error.value = '两次输入的密码不一致'
    return
  }

  try {
    await projectStore.createProject(name.value, description.value, password.value)
    emit('close')
  } catch (e: any) {
    error.value = e.message || '创建失败'
  }
}
</script>

<template>
  <div class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
    <div class="bg-base-100 rounded-lg p-6 w-96">
      <h2 class="text-xl font-bold mb-4">创建新项目</h2>
      
      <form @submit.prevent="handleSubmit" class="space-y-4">
        <div>
          <label class="block text-sm font-medium mb-1">项目名称</label>
          <input 
            v-model="name" 
            type="text" 
            class="input input-bordered w-full"
            placeholder="请输入项目名称"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium mb-1">项目描述</label>
          <textarea 
            v-model="description" 
            class="textarea textarea-bordered w-full"
            placeholder="请输入项目描述（可选）"
          ></textarea>
        </div>
        
        <div>
          <label class="block text-sm font-medium mb-1">密码</label>
          <input 
            v-model="password" 
            type="password" 
            class="input input-bordered w-full"
            placeholder="请输入密码"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium mb-1">确认密码</label>
          <input 
            v-model="confirmPassword" 
            type="password" 
            class="input input-bordered w-full"
            placeholder="请再次输入密码"
          />
        </div>
        
        <div v-if="error" class="text-error text-sm">{{ error }}</div>
        
        <div class="flex gap-2">
          <button type="submit" class="btn btn-primary flex-1">创建</button>
          <button type="button" class="btn btn-ghost flex-1" @click="emit('close')">取消</button>
        </div>
      </form>
    </div>
  </div>
</template>
