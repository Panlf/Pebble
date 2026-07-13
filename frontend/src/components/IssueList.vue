<script setup lang="ts">
import { useIssueStore } from '../stores/issue'
import { useProjectStore } from '../stores/project'
import { watch } from 'vue'

const issueStore = useIssueStore()
const projectStore = useProjectStore()

watch(() => projectStore.currentProject, (newProject) => {
  if (newProject) {
    issueStore.fetchIssues(newProject.id)
  }
}, { immediate: true })

function formatDate(dateString: string) {
  return new Date(dateString).toLocaleDateString('zh-CN')
}

function getStatusText(status: string) {
  const statusMap: Record<string, string> = {
    'pending': '待解决',
    'in_progress': '进行中',
    'resolved': '已解决'
  }
  return statusMap[status] || status
}

function getStatusClass(status: string) {
  const classMap: Record<string, string> = {
    'pending': 'badge-warning',
    'in_progress': 'badge-info',
    'resolved': 'badge-success'
  }
  return classMap[status] || 'badge-ghost'
}
</script>

<template>
  <div class="space-y-2">
    <div 
      v-for="issue in issueStore.issues" 
      :key="issue.id"
      class="p-3 rounded-lg border border-base-300 hover:border-primary cursor-pointer transition-colors"
      @click="issueStore.setCurrentIssue(issue)"
    >
      <div class="flex justify-between items-start">
        <div class="font-medium">{{ issue.title }}</div>
        <span class="badge badge-sm" :class="getStatusClass(issue.status)">
          {{ getStatusText(issue.status) }}
        </span>
      </div>
      <div class="text-sm text-base-content/70 mt-1">{{ issue.description || '无描述' }}</div>
      <div class="flex justify-between text-xs text-base-content/50 mt-2">
        <span>对接人: {{ issue.contact_person || '未指定' }}</span>
        <span>{{ formatDate(issue.created_at) }}</span>
      </div>
    </div>
    
    <div v-if="issueStore.issues.length === 0" class="text-center text-base-content/50 py-8">
      暂无问题记录
    </div>
  </div>
</template>
