<script setup lang="ts">
import WindowControls from './WindowControls.vue';
import DarkMode from './DarkMode.vue';
import { computed } from 'vue';

// 定义props
interface Props {
  title?: string;
  showIcon?: boolean;
  iconSrc?: string;
}

withDefaults(defineProps<Props>(), {
  title: 'Wails Template Vue3',
  showIcon: true,
  iconSrc: '/src/assets/vue.svg'
});

// 检查是否为Mac平台
const isMac = navigator.userAgent.toUpperCase().indexOf('MAC') >= 0;

// 计算标题栏样式
const titleBarClass = computed(() => {
  return [
    'title-bar',
    isMac ? 'title-bar-mac' : 'title-bar-windows'
  ];
});
</script>

<template>
  <div :class="titleBarClass" style="--wails-draggable: drag">
    <!-- 左侧：应用图标和标题 -->
    <div class="title-section">
      <img 
        v-if="showIcon" 
        :src="iconSrc" 
        alt="App Icon" 
        class="app-icon"
      />
      <h1 class="app-title">{{ title }}</h1>
    </div>
    
    <!-- 中间：可拖拽区域 -->
    <div class="drag-region"></div>
    
    <!-- 右侧：主题切换和窗口控制按钮 -->
    <div class="controls-section" style="--wails-draggable: none">
      <div class="theme-toggle">
        <DarkMode />
      </div>
      <WindowControls />
    </div>
  </div>
</template>

<style scoped>
.title-bar {
  @apply flex items-center justify-between h-8 px-3 select-none;
  /* 增强亚克力毛玻璃透明效果 */
  background: rgba(255, 255, 255, 0.5);
  backdrop-filter: blur(25px) saturate(2.0);
  border-bottom: 1px solid rgba(255, 255, 255, 0.15);
  position: relative;
  z-index: 1000;
  /* 增强毛玻璃质感 */
  box-shadow:
    0 1px 3px rgba(0, 0, 0, 0.08),
    inset 0 1px 0 rgba(255, 255, 255, 0.4);
}

.dark .title-bar {
  background: rgba(30, 41, 59, 0.6);
  backdrop-filter: blur(25px) saturate(2.0);
  border-bottom: 1px solid rgba(71, 85, 105, 0.2);
  box-shadow:
    0 1px 3px rgba(0, 0, 0, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.08);
}

.title-bar-mac {
  @apply pl-20; /* 为Mac的红绿灯按钮留出空间 */
}

.title-bar-windows {
  @apply pl-3;
}

.title-section {
  @apply flex items-center gap-2 flex-shrink-0;
}

.app-icon {
  @apply w-4 h-4 object-contain drop-shadow-sm;
}

.app-title {
  @apply text-sm font-semibold text-gray-700 dark:text-gray-200 truncate;
  max-width: 200px;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

.drag-region {
  @apply flex-1 h-full;
  cursor: default;
}

.drag-region:active {
  cursor: move;
}

.controls-section {
  @apply flex items-center flex-shrink-0;
  border-radius: 4px;
  overflow: hidden;
  gap: 4px;
}

.theme-toggle {
  @apply flex items-center justify-center;
  width: 32px;
  height: 32px;
  margin-right: 2px;
  border-radius: 4px;
  transition: all 0.2s ease-in-out;
  color: #6b7280;
}

.dark .theme-toggle {
  color: #9ca3af;
}

.theme-toggle:hover {
  background-color: rgba(156, 163, 175, 0.1);
  color: #374151;
}

.dark .theme-toggle:hover {
  background-color: rgba(156, 163, 175, 0.15);
  color: #d1d5db;
}

/* 让主题切换按钮内的图标更小一些 */
.theme-toggle :deep(button) {
  font-size: 14px !important;
  margin: 0 !important;
  padding: 0 !important;
  background: none !important;
  border: none !important;
  color: inherit !important;
  transition: none !important;
}

.theme-toggle :deep(button):hover {
  transform: none !important;
  background: none !important;
}

/* 在Mac上隐藏窗口控制按钮，因为系统会提供 */
.title-bar-mac .controls-section {
  @apply hidden;
}

/* 增强亚克力毛玻璃光泽效果 */
.title-bar::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg,
    transparent,
    rgba(255, 255, 255, 0.8),
    transparent
  );
  opacity: 0.6;
}

.dark .title-bar::before {
  background: linear-gradient(90deg,
    transparent,
    rgba(255, 255, 255, 0.2),
    transparent
  );
  opacity: 0.4;
}

/* 添加亚克力材质的噪点纹理效果 */
.title-bar::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-image:
    radial-gradient(circle at 1px 1px, rgba(255,255,255,0.15) 1px, transparent 0);
  background-size: 20px 20px;
  opacity: 0.3;
  pointer-events: none;
}
</style>
