<script setup lang="ts">
import TitleBar from './components/TitleBar.vue';
import Sidebar from './components/Sidebar.vue';
import { useThemeStore } from './stores/theme';
import { onMounted } from 'vue';

const themeStore = useThemeStore();

const menuItems = [
  { text: "首页", href: "/", icon: "material-symbols:home-app-logo" },
  { text: "项目管理", href: "/projects", icon: "material-symbols:folder" },
  { text: "设置", href: "/settings", icon: "material-symbols:settings" },
];

onMounted(() => {
  themeStore.initTheme();
});
</script>

<template>
  <div class="flex flex-col h-screen">
    <!-- 顶部标题栏 -->
    <TitleBar title="项目文档管理系统" />

    <!-- 主要内容区域 -->
    <main class="flex flex-1 overflow-hidden">
      <!-- 侧边栏导航 -->
      <Sidebar :menu-items="menuItems" />

    <!-- 内容面板 -->
    <div class="flex-1 bg-white dark:bg-gray-900/90 overflow-hidden">
      <!-- 页面内容 -->
      <router-view class="h-full overflow-y-auto p-4" style="--wails-draggable:none;" />
    </div>
    </main>
  </div>
</template>
