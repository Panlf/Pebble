<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useProjectStore } from '../stores/project'

const projectStore = useProjectStore()
const stats = ref({
  projects: 0,
  issues: 0,
  documents: 0
})

onMounted(async () => {
  await projectStore.fetchProjects()
  stats.value.projects = projectStore.projects.length
})

const quickActions = [
  {
    title: '创建项目',
    description: '开始一个新的项目',
    icon: 'material-symbols:add-circle-outline',
    color: 'from-purple-500 to-indigo-500',
    href: '/projects'
  },
  {
    title: '最近项目',
    description: '查看最近的项目',
    icon: 'material-symbols:history',
    color: 'from-pink-500 to-rose-500',
    href: '/projects'
  },
  {
    title: '快速搜索',
    description: '搜索文档和问题',
    icon: 'material-symbols:search',
    color: 'from-cyan-500 to-blue-500',
    href: '/projects'
  }
]
</script>

<template>
  <div class="min-h-full">
    <!-- 欢迎区域 -->
    <div class="welcome-section mb-8">
      <div class="welcome-content">
        <div class="welcome-text">
          <h1 class="welcome-title">欢迎使用砾石</h1>
          <p class="welcome-subtitle">您的项目文档管理助手</p>
        </div>
        <div class="welcome-decoration">
          <div class="decoration-circle circle-1"></div>
          <div class="decoration-circle circle-2"></div>
          <div class="decoration-circle circle-3"></div>
        </div>
      </div>
    </div>

    <!-- 快速操作卡片 -->
    <div class="mb-8">
      <h2 class="section-title">快速操作</h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <router-link
          v-for="action in quickActions"
          :key="action.title"
          :to="action.href"
          class="action-card"
        >
          <div class="action-icon" :class="`bg-gradient-to-br ${action.color}`">
            <Icon :icon="action.icon" class="w-6 h-6 text-white" />
          </div>
          <div class="action-content">
            <h3 class="action-title">{{ action.title }}</h3>
            <p class="action-description">{{ action.description }}</p>
          </div>
          <div class="action-arrow">
            <Icon icon="material-symbols:arrow-forward" class="w-5 h-5" />
          </div>
        </router-link>
      </div>
    </div>

    <!-- 系统统计 -->
    <div class="mb-8">
      <h2 class="section-title">系统概览</h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="stat-card">
          <div class="stat-icon bg-gradient-to-br from-purple-500 to-indigo-500">
            <Icon icon="material-symbols:folder-open-outline" class="w-6 h-6 text-white" />
          </div>
          <div class="stat-content">
            <div class="stat-value">{{ stats.projects }}</div>
            <div class="stat-label">项目总数</div>
          </div>
        </div>
        
        <div class="stat-card">
          <div class="stat-icon bg-gradient-to-br from-pink-500 to-rose-500">
            <Icon icon="material-symbols:bug-outline" class="w-6 h-6 text-white" />
          </div>
          <div class="stat-content">
            <div class="stat-value">{{ stats.issues }}</div>
            <div class="stat-label">问题总数</div>
          </div>
        </div>
        
        <div class="stat-card">
          <div class="stat-icon bg-gradient-to-br from-cyan-500 to-blue-500">
            <Icon icon="material-symbols:description-outline" class="w-6 h-6 text-white" />
          </div>
          <div class="stat-content">
            <div class="stat-value">{{ stats.documents }}</div>
            <div class="stat-label">文档总数</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 最近项目 -->
    <div v-if="projectStore.projects.length > 0">
      <h2 class="section-title">最近项目</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="project in projectStore.projects.slice(0, 6)"
          :key="project.id"
          class="project-card"
        >
          <div class="project-header">
            <div class="project-icon">
              <Icon icon="material-symbols:folder-outline" class="w-5 h-5" />
            </div>
            <div class="project-info">
              <h3 class="project-name">{{ project.name }}</h3>
              <p class="project-date">{{ new Date(project.updated_at).toLocaleDateString('zh-CN') }}</p>
            </div>
          </div>
          <p class="project-description">{{ project.description || '暂无描述' }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.welcome-section {
  @apply relative overflow-hidden rounded-2xl p-8;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.welcome-content {
  @apply relative z-10;
}

.welcome-text {
  @apply text-white;
}

.welcome-title {
  @apply text-3xl font-bold mb-2;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.welcome-subtitle {
  @apply text-lg opacity-90;
}

.welcome-decoration {
  @apply absolute top-0 right-0 w-64 h-64;
}

.decoration-circle {
  @apply absolute rounded-full;
  background: rgba(255, 255, 255, 0.1);
}

.circle-1 {
  @apply w-32 h-32;
  top: -20px;
  right: -20px;
}

.circle-2 {
  @apply w-24 h-24;
  top: 40px;
  right: 60px;
}

.circle-3 {
  @apply w-16 h-16;
  top: 10px;
  right: 120px;
}

.section-title {
  @apply text-xl font-semibold text-gray-800 dark:text-gray-200 mb-4;
}

.action-card {
  @apply flex items-center p-4 rounded-xl bg-white dark:bg-gray-800 shadow-sm hover:shadow-md transition-all duration-300 cursor-pointer border border-gray-100 dark:border-gray-700;
}

.action-card:hover {
  @apply transform translate-y-[-2px];
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.1);
}

.action-icon {
  @apply flex items-center justify-center w-12 h-12 rounded-xl flex-shrink-0;
}

.action-content {
  @apply flex-1 ml-4;
}

.action-title {
  @apply text-sm font-semibold text-gray-800 dark:text-gray-200 mb-1;
}

.action-description {
  @apply text-xs text-gray-500 dark:text-gray-400;
}

.action-arrow {
  @apply text-gray-400 dark:text-gray-500;
}

.stat-card {
  @apply flex items-center p-4 rounded-xl bg-white dark:bg-gray-800 shadow-sm border border-gray-100 dark:border-gray-700;
}

.stat-icon {
  @apply flex items-center justify-center w-12 h-12 rounded-xl flex-shrink-0;
}

.stat-content {
  @apply ml-4;
}

.stat-value {
  @apply text-2xl font-bold text-gray-800 dark:text-gray-200;
}

.stat-label {
  @apply text-sm text-gray-500 dark:text-gray-400;
}

.project-card {
  @apply p-4 rounded-xl bg-white dark:bg-gray-800 shadow-sm hover:shadow-md transition-all duration-300 border border-gray-100 dark:border-gray-700;
}

.project-card:hover {
  @apply transform translate-y-[-2px];
}

.project-header {
  @apply flex items-center mb-3;
}

.project-icon {
  @apply flex items-center justify-center w-10 h-10 rounded-lg bg-gradient-to-br from-purple-500 to-indigo-500 text-white;
}

.project-info {
  @apply ml-3;
}

.project-name {
  @apply text-sm font-semibold text-gray-800 dark:text-gray-200;
}

.project-date {
  @apply text-xs text-gray-500 dark:text-gray-400;
}

.project-description {
  @apply text-sm text-gray-600 dark:text-gray-300 line-clamp-2;
}
</style>
