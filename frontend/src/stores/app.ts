import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export interface AppInfo {
  name: string
  version: string
  description: string
  author: string
  repository: string
  license: string
}

export interface TechStack {
  name: string
  version: string
  description: string
  icon: string
  color: string
  url: string
}

export const useAppStore = defineStore('app', () => {
  // 应用信息
  const appInfo = ref<AppInfo>({
    name: 'Wails Template Vue3',
    version: '1.0.0',
    description: '基于 Wails + Vue3 + TypeScript + Pinia + TailwindCSS + DaisyUI 的现代桌面应用模板',
    author: 'onewinner',
    repository: 'https://github.com/onewinner/wails-template-vue3',
    license: 'MIT'
  })
  
  // 技术栈信息
  const techStack = ref<TechStack[]>([
    {
      name: 'Wails',
      version: 'v2.x',
      description: '使用 Go 和 Web 技术构建桌面应用',
      icon: 'simple-icons:go',
      color: '#00ADD8',
      url: 'https://wails.io'
    },
    {
      name: 'Vue 3',
      version: '3.4.x',
      description: '渐进式 JavaScript 框架',
      icon: 'simple-icons:vuedotjs',
      color: '#4FC08D',
      url: 'https://vuejs.org'
    },
    {
      name: 'TypeScript',
      version: '5.x',
      description: 'JavaScript 的超集，添加了静态类型',
      icon: 'simple-icons:typescript',
      color: '#3178C6',
      url: 'https://www.typescriptlang.org'
    },
    {
      name: 'Vite',
      version: '5.x',
      description: '下一代前端构建工具',
      icon: 'simple-icons:vite',
      color: '#646CFF',
      url: 'https://vitejs.dev'
    },
    {
      name: 'Pinia',
      version: '2.x',
      description: 'Vue 的状态管理库',
      icon: 'simple-icons:pinia',
      color: '#FFD859',
      url: 'https://pinia.vuejs.org'
    },
    {
      name: 'TailwindCSS',
      version: '3.x',
      description: '实用优先的 CSS 框架',
      icon: 'simple-icons:tailwindcss',
      color: '#06B6D4',
      url: 'https://tailwindcss.com'
    },
    {
      name: 'DaisyUI',
      version: '4.x',
      description: 'TailwindCSS 的组件库',
      icon: 'simple-icons:daisyui',
      color: '#5A67D8',
      url: 'https://daisyui.com'
    }
  ])
  
  // 示例数据
  const sampleData = ref({
    users: [
      { id: 1, name: '张三', email: 'zhangsan@example.com', role: 'admin' },
      { id: 2, name: '李四', email: 'lisi@example.com', role: 'user' },
      { id: 3, name: '王五', email: 'wangwu@example.com', role: 'user' }
    ],
    stats: {
      totalUsers: 1234,
      activeUsers: 856,
      totalProjects: 42,
      completedTasks: 789
    }
  })
  
  // 计算属性
  const techStackCount = computed(() => techStack.value.length)
  
  // 方法
  const updateAppInfo = (info: Partial<AppInfo>) => {
    appInfo.value = { ...appInfo.value, ...info }
  }
  
  const addTechStack = (tech: TechStack) => {
    techStack.value.push(tech)
  }
  
  return {
    appInfo,
    techStack,
    sampleData,
    techStackCount,
    updateAppInfo,
    addTechStack
  }
})
