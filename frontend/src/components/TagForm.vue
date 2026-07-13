<script setup lang="ts">
import { ref } from 'vue'
import { useTagStore } from '../stores/tag'

const tagStore = useTagStore()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const name = ref('')
const color = ref('#3b82f6')
const error = ref('')

const colorOptions = [
  '#3b82f6', '#10b981', '#f59e0b', '#ef4444', '#8b5cf6', '#ec4899',
  '#06b6d4', '#84cc16', '#f97316', '#6366f1', '#14b8a6', '#e11d48'
]

async function handleSubmit() {
  if (!name.value) {
    error.value = '请输入标签名称'
    return
  }

  try {
    await tagStore.createTag(name.value, color.value)
    emit('close')
  } catch (e: any) {
    error.value = e.message || '创建失败'
  }
}
</script>

<template>
  <div class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
    <div class="bg-base-100 rounded-lg p-6 w-80">
      <h2 class="text-xl font-bold mb-4">新建标签</h2>
      
      <form @submit.prevent="handleSubmit" class="space-y-4">
        <div>
          <label class="block text-sm font-medium mb-1">标签名称</label>
          <input 
            v-model="name" 
            type="text" 
            class="input input-bordered w-full"
            placeholder="请输入标签名称"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium mb-1">标签颜色</label>
          <div class="flex flex-wrap gap-2">
            <div 
              v-for="c in colorOptions" 
              :key="c"
              class="w-8 h-8 rounded-full cursor-pointer border-2 transition-transform hover:scale-110"
              :class="{ 'border-primary': color === c }"
              :style="{ backgroundColor: c }"
              @click="color = c"
            ></div>
          </div>
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
