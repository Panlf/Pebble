<script setup lang="ts">
import { ref, computed } from 'vue'
import { useQuickInputStore } from '../stores/quickInput'
import { useProjectStore } from '../stores/project'
import { useIssueStore } from '../stores/issue'

const quickInputStore = useQuickInputStore()
const projectStore = useProjectStore()
const issueStore = useIssueStore()

const title = ref('')
const description = ref('')
const error = ref('')
const isSubmitting = ref(false)

const selectedProject = computed(() => 
  projectStore.projects.find(p => p.id === quickInputStore.selectedProjectId)
)

async function handleSubmit() {
  if (!title.value) {
    error.value = '请输入问题标题'
    return
  }

  if (!quickInputStore.selectedProjectId) {
    error.value = '请选择项目'
    return
  }

  isSubmitting.value = true
  error.value = ''

  try {
    await issueStore.createIssue(
      quickInputStore.selectedProjectId,
      title.value,
      description.value,
      ''
    )
    
    // Reset form
    title.value = ''
    description.value = ''
    quickInputStore.hide()
  } catch (e: any) {
    error.value = e.message || '创建失败'
  } finally {
    isSubmitting.value = false
  }
}

function handleCancel() {
  title.value = ''
  description.value = ''
  error.value = ''
  quickInputStore.hide()
}
</script>

<template>
  <Teleport to="body">
    <div 
      v-if="quickInputStore.isVisible"
      class="fixed bottom-4 right-4 w-96 bg-base-100 rounded-lg shadow-2xl border border-base-300 z-50"
    >
      <div class="p-4 border-b border-base-300">
        <div class="flex justify-between items-center">
          <h3 class="font-bold">快速记录问题</h3>
          <button class="btn btn-ghost btn-sm" @click="handleCancel">×</button>
        </div>
      </div>
      
      <form @submit.prevent="handleSubmit" class="p-4 space-y-3">
        <div>
          <label class="label">
            <span class="label-text">选择项目</span>
          </label>
          <select 
            v-model="quickInputStore.selectedProjectId"
            class="select select-bordered w-full"
          >
            <option value="">请选择项目</option>
            <option 
              v-for="project in projectStore.projects" 
              :key="project.id" 
              :value="project.id"
            >
              {{ project.name }}
            </option>
          </select>
        </div>
        
        <div>
          <label class="label">
            <span class="label-text">问题标题</span>
          </label>
          <input 
            v-model="title"
            type="text" 
            class="input input-bordered w-full"
            placeholder="请输入问题标题"
          />
        </div>
        
        <div>
          <label class="label">
            <span class="label-text">问题描述</span>
          </label>
          <textarea 
            v-model="description"
            class="textarea textarea-bordered w-full h-20"
            placeholder="请输入问题描述（可选）"
          ></textarea>
        </div>
        
        <div v-if="error" class="text-error text-sm">{{ error }}</div>
        
        <div class="flex gap-2">
          <button type="submit" class="btn btn-primary flex-1" :disabled="isSubmitting">
            {{ isSubmitting ? '提交中...' : '提交' }}
          </button>
          <button type="button" class="btn btn-ghost flex-1" @click="handleCancel">取消</button>
        </div>
      </form>
    </div>
  </Teleport>
</template>
