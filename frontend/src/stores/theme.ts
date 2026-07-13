import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useThemeStore = defineStore('theme', () => {
  // 当前主题
  const currentTheme = ref<string>('light')
  
  // 是否为深色模式
  const isDark = computed(() => currentTheme.value === 'dark')
  
  // DaisyUI主题列表
  const themes = [
    { name: 'light', label: '浅色', icon: 'icon-park:sun-one' },
    { name: 'dark', label: '深色', icon: 'icon-park:dark-mode' },
    { name: 'cupcake', label: '纸杯蛋糕', icon: 'icon-park:cake' },
    { name: 'bumblebee', label: '大黄蜂', icon: 'icon-park:bee' },
    { name: 'emerald', label: '翡翠', icon: 'icon-park:diamond' },
    { name: 'corporate', label: '企业', icon: 'icon-park:office-building' },
    { name: 'synthwave', label: '合成波', icon: 'icon-park:music' },
    { name: 'retro', label: '复古', icon: 'icon-park:retro' },
    { name: 'cyberpunk', label: '赛博朋克', icon: 'icon-park:robot' },
    { name: 'valentine', label: '情人节', icon: 'icon-park:heart' },
    { name: 'halloween', label: '万圣节', icon: 'icon-park:pumpkin' },
    { name: 'garden', label: '花园', icon: 'icon-park:leaves' },
    { name: 'forest', label: '森林', icon: 'icon-park:tree' },
    { name: 'aqua', label: '水蓝', icon: 'icon-park:water' },
    { name: 'lofi', label: 'Lo-Fi', icon: 'icon-park:headphones' },
    { name: 'pastel', label: '粉彩', icon: 'icon-park:palette' },
    { name: 'fantasy', label: '幻想', icon: 'icon-park:magic-wand' },
    { name: 'wireframe', label: '线框', icon: 'icon-park:wireframe' },
    { name: 'black', label: '黑色', icon: 'icon-park:moon' },
    { name: 'luxury', label: '奢华', icon: 'icon-park:diamond-ring' },
    { name: 'dracula', label: '德古拉', icon: 'icon-park:vampire' },
    { name: 'cmyk', label: 'CMYK', icon: 'icon-park:printer' },
    { name: 'autumn', label: '秋天', icon: 'icon-park:maple-leaf' },
    { name: 'business', label: '商务', icon: 'icon-park:briefcase' },
    { name: 'acid', label: '酸性', icon: 'icon-park:chemistry' },
    { name: 'lemonade', label: '柠檬水', icon: 'icon-park:lemon' },
    { name: 'night', label: '夜晚', icon: 'icon-park:moon-stars' },
    { name: 'coffee', label: '咖啡', icon: 'icon-park:coffee' },
    { name: 'winter', label: '冬天', icon: 'icon-park:snowflake' },
  ]
  
  // 初始化主题
  const initTheme = () => {
    const savedTheme = localStorage.getItem('theme') || 'light'
    setTheme(savedTheme)
  }
  
  // 设置主题
  const setTheme = (theme: string) => {
    currentTheme.value = theme

    // 设置DaisyUI主题
    document.documentElement.setAttribute('data-theme', theme)

    // 设置dark class用于兼容自定义样式
    const isDarkTheme = ['dark', 'synthwave', 'halloween', 'forest', 'black', 'luxury', 'dracula', 'night', 'coffee'].includes(theme)
    if (isDarkTheme) {
      document.documentElement.classList.add('dark')
      document.body.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
      document.body.classList.remove('dark')
    }

    // 强制更新CSS变量
    document.documentElement.style.setProperty('color-scheme', isDarkTheme ? 'dark' : 'light')

    localStorage.setItem('theme', theme)
  }
  
  // 切换深色模式
  const toggleDark = () => {
    setTheme(isDark.value ? 'light' : 'dark')
  }
  
  return {
    currentTheme,
    isDark,
    themes,
    initTheme,
    setTheme,
    toggleDark
  }
})
