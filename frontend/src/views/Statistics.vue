<script setup lang="ts">
import { ref, onMounted } from 'vue'

const stats = ref<any>({})
const isLoading = ref(true)

onMounted(async () => {
  try {
    // @ts-ignore
    stats.value = await window.go.main.ActivityHandler.GetStatistics()
  } catch (error) {
    console.error('Failed to load statistics:', error)
  } finally {
    isLoading.value = false
  }
})
</script>

<template>
  <div class="p-6">
    <h1 class="text-3xl font-bold mb-6">统计报表</h1>
    
    <div v-if="isLoading" class="text-center py-8">
      <span class="loading loading-spinner loading-lg"></span>
    </div>
    
    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
      <div class="stat bg-base-200 rounded-lg">
        <div class="stat-title">项目总数</div>
        <div class="stat-value text-primary">{{ stats.projects || 0 }}</div>
      </div>
      
      <div class="stat bg-base-200 rounded-lg">
        <div class="stat-title">问题总数</div>
        <div class="stat-value text-secondary">{{ stats.issues || 0 }}</div>
      </div>
      
      <div class="stat bg-base-200 rounded-lg">
        <div class="stat-title">文档总数</div>
        <div class="stat-value text-accent">{{ stats.documents || 0 }}</div>
      </div>
      
      <div class="stat bg-base-200 rounded-lg">
        <div class="stat-title">标签总数</div>
        <div class="stat-value text-info">{{ stats.tags || 0 }}</div>
      </div>
    </div>
    
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div class="card bg-base-200">
        <div class="card-body">
          <h2 class="card-title">问题状态分布</h2>
          <div class="space-y-2">
            <div class="flex justify-between items-center">
              <span>待解决</span>
              <span class="badge badge-warning">{{ stats.issue_status?.pending || 0 }}</span>
            </div>
            <div class="flex justify-between items-center">
              <span>进行中</span>
              <span class="badge badge-info">{{ stats.issue_status?.in_progress || 0 }}</span>
            </div>
            <div class="flex justify-between items-center">
              <span>已解决</span>
              <span class="badge badge-success">{{ stats.issue_status?.resolved || 0 }}</span>
            </div>
          </div>
        </div>
      </div>
      
      <div class="card bg-base-200">
        <div class="card-body">
          <h2 class="card-title">最近活动</h2>
          <div class="stat">
            <div class="stat-title">最近7天活动数</div>
            <div class="stat-value text-primary">{{ stats.recent_activities || 0 }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
