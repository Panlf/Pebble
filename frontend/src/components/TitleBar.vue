<script setup lang="ts">
import WindowControls from './WindowControls.vue';
import DarkMode from './DarkMode.vue';
import { computed } from 'vue';

interface Props {
  title?: string;
  showIcon?: boolean;
  iconSrc?: string;
}

withDefaults(defineProps<Props>(), {
  title: '砾石',
  showIcon: true,
  iconSrc: '/src/assets/pebble-icon.svg'
});

const isMac = navigator.userAgent.toUpperCase().indexOf('MAC') >= 0;

const titleBarClass = computed(() => {
  return [
    'title-bar',
    isMac ? 'title-bar-mac' : 'title-bar-windows'
  ];
});
</script>

<template>
  <div :class="titleBarClass" style="--wails-draggable: drag">
    <div class="title-section">
      <div class="app-icon-container">
        <svg class="app-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <circle cx="8" cy="12" r="3" fill="url(#gradient1)" opacity="0.8"/>
          <circle cx="14" cy="8" r="2" fill="url(#gradient2)" opacity="0.6"/>
          <circle cx="16" cy="14" r="2.5" fill="url(#gradient1)" opacity="0.7"/>
          <circle cx="10" cy="16" r="1.5" fill="url(#gradient2)" opacity="0.5"/>
          <defs>
            <linearGradient id="gradient1" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stop-color="#667eea"/>
              <stop offset="100%" stop-color="#764ba2"/>
            </linearGradient>
            <linearGradient id="gradient2" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stop-color="#f093fb"/>
              <stop offset="100%" stop-color="#f5576c"/>
            </linearGradient>
          </defs>
        </svg>
      </div>
      <h1 class="app-title">{{ title }}</h1>
    </div>
    
    <div class="drag-region"></div>
    
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
  @apply flex items-center justify-between h-12 px-4 select-none;
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(20px) saturate(1.8);
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  position: relative;
  z-index: 1000;
  box-shadow:
    0 1px 3px rgba(0, 0, 0, 0.05),
    inset 0 1px 0 rgba(255, 255, 255, 0.5);
}

.dark .title-bar {
  background: rgba(30, 41, 59, 0.7);
  backdrop-filter: blur(20px) saturate(1.8);
  border-bottom: 1px solid rgba(71, 85, 105, 0.3);
  box-shadow:
    0 1px 3px rgba(0, 0, 0, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
}

.title-bar-mac {
  @apply pl-20;
}

.title-bar-windows {
  @apply pl-4;
}

.title-section {
  @apply flex items-center gap-3 flex-shrink-0;
}

.app-icon-container {
  @apply flex items-center justify-center;
  width: 32px;
  height: 32px;
}

.app-icon {
  width: 28px;
  height: 28px;
  filter: drop-shadow(0 2px 4px rgba(102, 126, 234, 0.3));
}

.app-title {
  @apply text-lg font-semibold;
  background: linear-gradient(135deg, #667eea, #764ba2);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  text-shadow: none;
}

.dark .app-title {
  background: linear-gradient(135deg, #a5b4fc, #c084fc);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
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
  border-radius: 8px;
  overflow: hidden;
  gap: 4px;
}

.theme-toggle {
  @apply flex items-center justify-center;
  width: 36px;
  height: 36px;
  margin-right: 4px;
  border-radius: 8px;
  transition: all 0.2s ease-in-out;
  color: #6b7280;
}

.dark .theme-toggle {
  color: #9ca3af;
}

.theme-toggle:hover {
  background: rgba(102, 126, 234, 0.1);
  color: #667eea;
}

.dark .theme-toggle:hover {
  background: rgba(165, 180, 252, 0.15);
  color: #a5b4fc;
}

.theme-toggle :deep(button) {
  font-size: 16px !important;
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

.title-bar-mac .controls-section {
  @apply hidden;
}

.title-bar::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg,
    transparent,
    rgba(102, 126, 234, 0.5),
    transparent
  );
  opacity: 0.6;
}

.dark .title-bar::before {
  background: linear-gradient(90deg,
    transparent,
    rgba(165, 180, 252, 0.3),
    transparent
  );
  opacity: 0.4;
}
</style>
