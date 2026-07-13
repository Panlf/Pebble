<script setup lang="ts">
import { ref } from 'vue'
import { useIssueStore } from '../stores/issue'
import { useProjectStore } from '../stores/project'

const issueStore = useIssueStore()
const projectStore = useProjectStore()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const title = ref('')
const description = ref('')
const contactPerson = ref('')
const error = ref('')

async function handleSubmit() {
  if (!title.value) {
    error.value = '请填写问题标题'
    return
  }

  if (!projectStore.currentProject) {
    error.value = '请先选择项目'
    return
  }

  try {
    await issueStore.createIssue(
      projectStore.currentProject.id,
      title.value,
      description.value,
      contactPerson.value
    )
    emit('close')
  } catch (e: any) {
    error.value = e.message || '创建失败'
  }
}
</script>

<template>
  <div class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
    <div class="bg-base-100 rounded-lg p-6 w-96">
      <h2 class="text-xl font-bold mb-4">新建问题</h2>
      
      <form @submit.prevent="handleSubmit" class="space-y-4">
        <div>
          <label class="block text-sm font-medium mb-1">问题标题</label>
          <input 
            v-model="title" 
            type="text" 
            class="input input-bordered w-full"
            placeholder="请输入问题标题"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium mb-1">问题描述</label>
          <textarea 
            v-model="description" 
            class="textarea textarea-bordered w-full"
            placeholder="请输入问题描述（可选）"
          ></textarea>
        </div>
        
        <div>
          <label class="block text-sm font-medium mb-1">对接人</label>
          <input 
            v-model="contactPerson" 
            type="text" 
            class="input input-bordered w-full"
            placeholder="请输入对接人（可选）"
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
