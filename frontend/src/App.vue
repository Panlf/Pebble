<script setup lang="ts">
import TitleBar from './components/TitleBar.vue';
import Sidebar from './components/Sidebar.vue';
import { useThemeStore } from './stores/theme';
import { onMounted } from 'vue';

const themeStore = useThemeStore();

const menuItems = [
  { text: "首页", href: "/", icon: "material-symbols:home-outline" },
  { text: "项目管理", href: "/projects", icon: "material-symbols:folder-open-outline" },
  { text: "统计", href: "/statistics", icon: "material-symbols:bar-chart-outline" },
  { text: "设置", href: "/settings", icon: "material-symbols:settings-outline" },
];

onMounted(() => {
  themeStore.initTheme();
});
</script>

<template>
  <div class="flex flex-col h-screen bg-gradient-to-br from-slate-50 to-slate-100 dark:from-slate-900 dark:to-slate-800">
    <!-- 顶部标题栏 -->
    <TitleBar title="砾石" />

    <!-- 主要内容区域 -->
    <main class="flex flex-1 overflow-hidden">
      <!-- 侧边栏导航 -->
      <Sidebar :menu-items="menuItems" />

      <!-- 内容面板 -->
      <div class="flex-1 overflow-hidden">
        <!-- 页面内容 -->
        <router-view class="h-full overflow-y-auto p-6" style="--wails-draggable:none;" />
      </div>
    </main>
  </div>
</template>
