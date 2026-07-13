<script setup lang="ts">
import { useProjectStore } from '../stores/project'
import { onMounted } from 'vue'

const projectStore = useProjectStore()

onMounted(() => {
  projectStore.fetchProjects()
})

function formatDate(dateString: string) {
  return new Date(dateString).toLocaleDateString('zh-CN')
}
</script>

<template>
  <div class="space-y-2">
    <div 
      v-for="project in projectStore.projects" 
      :key="project.id"
      class="p-3 rounded-lg border border-base-300 hover:border-primary cursor-pointer transition-colors"
      @click="projectStore.setCurrentProject(project)"
    >
      <div class="font-medium">{{ project.name }}</div>
      <div class="text-sm text-base-content/70">{{ project.description || '无描述' }}</div>
      <div class="text-xs text-base-content/50 mt-1">
        更新于: {{ formatDate(project.updated_at) }}
      </div>
    </div>
  </div>
</template>
