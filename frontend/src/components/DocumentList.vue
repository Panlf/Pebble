<script setup lang="ts">
import { useDocumentStore } from '../stores/document'
import { useProjectStore } from '../stores/project'
import { useIssueStore } from '../stores/issue'
import { watch } from 'vue'

const documentStore = useDocumentStore()
const projectStore = useProjectStore()
const issueStore = useIssueStore()

watch(() => [projectStore.currentProject, issueStore.currentIssue], ([newProject, newIssue]) => {
  if (newProject) {
    documentStore.fetchDocuments(newProject.id, newIssue?.id)
  }
}, { immediate: true })

function formatDate(dateString: string) {
  return new Date(dateString).toLocaleDateString('zh-CN')
}

function formatFileSize(bytes: number) {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

function getFileIcon(fileType: string) {
  const iconMap: Record<string, string> = {
    '.pdf': '📄',
    '.doc': '📝',
    '.docx': '📝',
    '.xls': '📊',
    '.xlsx': '📊',
    '.ppt': '📊',
    '.pptx': '📊',
    '.jpg': '🖼️',
    '.jpeg': '🖼️',
    '.png': '🖼️',
    '.gif': '🖼️',
    '.txt': '📄',
    '.md': '📄',
  }
  return iconMap[fileType.toLowerCase()] || '📎'
}
</script>

<template>
  <div class="space-y-2">
    <div 
      v-for="doc in documentStore.documents" 
      :key="doc.id"
      class="p-3 rounded-lg border border-base-300 hover:border-primary transition-colors"
    >
      <div class="flex items-center gap-3">
        <span class="text-2xl">{{ getFileIcon(doc.file_type) }}</span>
        <div class="flex-1">
          <div class="font-medium">{{ doc.title }}</div>
          <div class="text-sm text-base-content/70">
            {{ formatFileSize(doc.file_size) }} • {{ formatDate(doc.created_at) }}
          </div>
        </div>
        <div class="flex gap-2">
          <button class="btn btn-sm btn-ghost">下载</button>
          <button class="btn btn-sm btn-error">删除</button>
        </div>
      </div>
    </div>
    
    <div v-if="documentStore.documents.length === 0" class="text-center text-base-content/50 py-8">
      暂无文档
    </div>
  </div>
</template>
