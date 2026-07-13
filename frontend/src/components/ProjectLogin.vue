<script setup lang="ts">
import { ref } from 'vue'
import { useProjectStore } from '../stores/project'
import { useAuthStore } from '../stores/auth'

const projectStore = useProjectStore()
const authStore = useAuthStore()

const emit = defineEmits<{
  (e: 'close'): void
  (e: 'success'): void
}>()

const password = ref('')
const error = ref('')
const isLoading = ref(false)

async function handleSubmit() {
  if (!password.value) {
    error.value = '请输入密码'
    return
  }

  if (!projectStore.currentProject) {
    error.value = '请先选择项目'
    return
  }

  isLoading.value = true
  error.value = ''

  try {
    const valid = await projectStore.verifyPassword(projectStore.currentProject.id, password.value)
    if (valid) {
      authStore.authenticateProject(projectStore.currentProject.id)
      emit('success')
    } else {
      error.value = '密码错误'
    }
  } catch (e: any) {
    error.value = e.message || '验证失败'
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <div class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
    <div class="bg-base-100 rounded-lg p-6 w-80">
      <h2 class="text-xl font-bold mb-4">项目验证</h2>
      <p class="text-sm text-base-content/70 mb-4">
        请输入项目密码以访问「{{ projectStore.currentProject?.name }}」
      </p>
      
      <form @submit.prevent="handleSubmit" class="space-y-4">
        <div>
          <input 
            v-model="password" 
            type="password" 
            class="input input-bordered w-full"
            placeholder="请输入密码"
            autofocus
          />
        </div>
        
        <div v-if="error" class="text-error text-sm">{{ error }}</div>
        
        <div class="flex gap-2">
          <button type="submit" class="btn btn-primary flex-1" :disabled="isLoading">
            {{ isLoading ? '验证中...' : '验证' }}
          </button>
          <button type="button" class="btn btn-ghost flex-1" @click="emit('close')">取消</button>
        </div>
      </form>
    </div>
  </div>
</template>
