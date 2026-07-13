<script setup lang="ts">
import { Icon } from '@iconify/vue';
import { ref, computed } from 'vue';

// 定义菜单项类型
interface MenuItem {
  text: string;
  href: string;
  icon: string;
}

// 定义props
interface Props {
  menuItems?: MenuItem[];
  activeClass?: string;
  appName?: string;
  appIcon?: string;
}

const props = withDefaults(defineProps<Props>(), {
  menuItems: () => [
    { text: "首页", href: "/home", icon: "material-symbols:home-app-logo" },
  ],
  activeClass: 'text-blue-600 dark:text-blue-400',
  appName: 'Wails Template Vue3',
  appIcon: '/src/assets/vue.svg'
});

// 侧边栏展开状态
const isExpanded = ref(false);

// 底部固定菜单项
const bottomMenuItems = [
  { text: "设置", href: "/setup", icon: "material-symbols:settings" }
];

// 计算侧边栏宽度
const sidebarWidth = computed(() => isExpanded.value ? '200px' : '70px');

// 切换展开状态
const toggleExpanded = () => {
  isExpanded.value = !isExpanded.value;
};
</script>

<template>
  <nav
    class="sidebar"
    :style="{
      width: sidebarWidth,
      '--wails-draggable': 'drag'
    }"
  >
    <!-- 应用头部 -->
    <div class="app-header" @click="toggleExpanded">
      <div class="menu-content">
        <div class="app-logo">
          <img :src="appIcon" alt="App Icon" class="app-icon" />
        </div>
        <span v-if="isExpanded" class="app-name">Wails</span>
      </div>
    </div>

    <!-- 主要导航菜单 -->
    <div class="main-menu">
      <router-link
        v-for="item in props.menuItems"
        :key="item.text"
        :to="item.href"
        v-slot="{ isActive }"
        class="menu-item"
        :title="!isExpanded ? item.text : ''"
      >
        <div class="menu-content">
          <Icon
            :icon="item.icon"
            :class="[
              'menu-icon',
              isActive && props.activeClass
            ]"
          />
          <span v-if="isExpanded" class="menu-text" :class="isActive && props.activeClass">
            {{ item.text }}
          </span>
        </div>
      </router-link>
    </div>

    <!-- 底部菜单 -->
    <div class="bottom-menu">
      <router-link
        v-for="item in bottomMenuItems"
        :key="item.text"
        :to="item.href"
        v-slot="{ isActive }"
        class="menu-item"
        :title="!isExpanded ? item.text : ''"
      >
        <div class="menu-content">
          <Icon
            :icon="item.icon"
            :class="[
              'menu-icon',
              isActive && props.activeClass
            ]"
          />
          <span v-if="isExpanded" class="menu-text" :class="isActive && props.activeClass">
            {{ item.text }}
          </span>
        </div>
      </router-link>
    </div>
  </nav>
</template>

<style scoped>
.sidebar {
  display: flex;
  flex-direction: column;
  user-select: none;
  z-index: 20;
  transition: all 0.3s ease-in-out;
  /* 与TitleBar保持一致的亚克力毛玻璃效果 */
  background: rgba(255, 255, 255, 0.5);
  backdrop-filter: blur(25px) saturate(2.0);
  border-right: 1px solid rgba(255, 255, 255, 0.15);
  /* 增强毛玻璃质感 */
  box-shadow:
    1px 0 3px rgba(0, 0, 0, 0.08),
    inset -1px 0 0 rgba(255, 255, 255, 0.4);
}

.dark .sidebar {
  background: rgba(30, 41, 59, 0.6);
  backdrop-filter: blur(25px) saturate(2.0);
  border-right: 1px solid rgba(71, 85, 105, 0.2);
  box-shadow:
    1px 0 3px rgba(0, 0, 0, 0.4),
    inset -1px 0 0 rgba(255, 255, 255, 0.08);
}

.app-header {
  padding: 12px 8px 12px 8px;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
  background: transparent;
  border-bottom: 1px solid rgba(0, 0, 0, 0.05);
  margin-bottom: 8px;
}

.dark .app-header {
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
}

.app-header:hover {
  background: rgba(59, 130, 246, 0.05);
}

.app-logo {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 54px;
  height: 24px;
}

.app-icon {
  width: 24px;
  height: 24px;
  object-fit: contain;
}

.app-name {
  margin-left: 12px;
  font-size: 16px;
  font-weight: 700;
  color: #1f2937;
  transition: colors 0.3s;
  letter-spacing: -0.025em;
}

.dark .app-name {
  color: #f9fafb;
}

.main-menu {
  flex: 1;
  padding: 0 0 16px 0;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.bottom-menu {
  padding: 16px 0;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.menu-item {
  display: block;
  margin: 0 8px;
  border-radius: 8px;
  transition: all 0.3s ease-in-out;
  position: relative;
  text-decoration: none;
}

/* 未选中状态 */
.menu-item:not(.router-link-active) {
  background: transparent;
}

.menu-item:not(.router-link-active):hover {
  background: rgba(59, 130, 246, 0.08);
  transform: scale(1.02);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

/* 选中状态 */
.menu-item.router-link-active {
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.15), rgba(37, 99, 235, 0.1));
  box-shadow:
    0 4px 12px rgba(59, 130, 246, 0.2),
    inset 0 1px 0 rgba(255, 255, 255, 0.2);
  border: 1px solid rgba(59, 130, 246, 0.2);
}

.dark .menu-item.router-link-active {
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.2), rgba(37, 99, 235, 0.15));
  box-shadow:
    0 4px 12px rgba(59, 130, 246, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(59, 130, 246, 0.3);
}

.menu-content {
  display: flex;
  align-items: center;
  padding: 12px;
}

/* 图标样式 */
.menu-icon {
  font-size: 20px;
  transition: all 0.3s;
  flex-shrink: 0;
  color: #6b7280;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.dark .menu-icon {
  color: #9ca3af;
}

/* 未选中状态的图标 */
.menu-item:not(.router-link-active) .menu-icon {
  opacity: 0.7;
}

.menu-item:not(.router-link-active):hover .menu-icon {
  opacity: 1;
  color: #3b82f6;
  transform: scale(1.1);
}

.dark .menu-item:not(.router-link-active):hover .menu-icon {
  color: #60a5fa;
}

/* 选中状态的图标 */
.menu-item.router-link-active .menu-icon {
  color: #3b82f6;
  opacity: 1;
  transform: scale(1.15);
  filter: drop-shadow(0 2px 4px rgba(59, 130, 246, 0.3));
}

.dark .menu-item.router-link-active .menu-icon {
  color: #60a5fa;
  filter: drop-shadow(0 2px 4px rgba(96, 165, 250, 0.4));
}

/* 文本样式 */
.menu-text {
  margin-left: 12px;
  font-size: 14px;
  font-weight: 500;
  color: #374151;
  transition: colors 0.3s;
}

.dark .menu-text {
  color: #d1d5db;
}

/* 悬停时的阴影效果 */
.dark .menu-item:hover {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
}

/* 收缩状态下的图标居中 */
.sidebar[style*="70px"] .menu-content {
  justify-content: center;
  padding-left: 0;
  padding-right: 0;
}

.sidebar[style*="70px"] .menu-text {
  display: none;
}
</style>