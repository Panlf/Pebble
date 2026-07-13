<script setup lang="ts">
import { Icon } from '@iconify/vue';
import { ref, computed } from 'vue';

interface MenuItem {
  text: string;
  href: string;
  icon: string;
}

interface Props {
  menuItems?: MenuItem[];
  activeClass?: string;
}

const props = withDefaults(defineProps<Props>(), {
  menuItems: () => [
    { text: "首页", href: "/", icon: "material-symbols:home-outline" },
  ],
  activeClass: 'text-purple-600 dark:text-purple-400',
});

const isExpanded = ref(false);

const sidebarWidth = computed(() => isExpanded.value ? '220px' : '72px');

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
    <div class="app-header" @click="toggleExpanded">
      <div class="menu-content">
        <div class="expand-icon">
          <Icon
            icon="material-symbols:menu"
            class="w-5 h-5"
          />
        </div>
      </div>
    </div>

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
  </nav>
</template>

<style scoped>
.sidebar {
  display: flex;
  flex-direction: column;
  user-select: none;
  z-index: 20;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  background: rgba(255, 255, 255, 0.6);
  backdrop-filter: blur(20px) saturate(1.8);
  border-right: 1px solid rgba(255, 255, 255, 0.2);
  box-shadow:
    1px 0 3px rgba(0, 0, 0, 0.05),
    inset -1px 0 0 rgba(255, 255, 255, 0.5);
}

.dark .sidebar {
  background: rgba(30, 41, 59, 0.6);
  backdrop-filter: blur(20px) saturate(1.8);
  border-right: 1px solid rgba(71, 85, 105, 0.3);
  box-shadow:
    1px 0 3px rgba(0, 0, 0, 0.3),
    inset -1px 0 0 rgba(255, 255, 255, 0.1);
}

.app-header {
  padding: 16px 8px;
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
  background: rgba(102, 126, 234, 0.05);
}

.expand-icon {
  @apply flex items-center justify-center;
  width: 56px;
  height: 32px;
  color: #6b7280;
  transition: all 0.3s ease;
}

.expand-icon:hover {
  color: #667eea;
  transform: scale(1.1);
}

.dark .expand-icon {
  color: #9ca3af;
}

.dark .expand-icon:hover {
  color: #a5b4fc;
}

.main-menu {
  flex: 1;
  padding: 0 8px 16px 0;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.menu-item {
  display: block;
  margin: 0 8px;
  border-radius: 12px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  text-decoration: none;
}

.menu-item:not(.router-link-active) {
  background: transparent;
}

.menu-item:not(.router-link-active):hover {
  background: rgba(102, 126, 234, 0.08);
  transform: translateX(4px);
}

.menu-item.router-link-active {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.15), rgba(118, 75, 162, 0.1));
  box-shadow:
    0 4px 12px rgba(102, 126, 234, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 0.3);
  border: 1px solid rgba(102, 126, 234, 0.2);
}

.dark .menu-item.router-link-active {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.2), rgba(118, 75, 162, 0.15));
  box-shadow:
    0 4px 12px rgba(102, 126, 234, 0.25),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(102, 126, 234, 0.3);
}

.menu-content {
  display: flex;
  align-items: center;
  padding: 12px 16px;
}

.menu-icon {
  font-size: 22px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
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

.menu-item:not(.router-link-active) .menu-icon {
  opacity: 0.7;
}

.menu-item:not(.router-link-active):hover .menu-icon {
  opacity: 1;
  color: #667eea;
  transform: scale(1.1);
}

.dark .menu-item:not(.router-link-active):hover .menu-icon {
  color: #a5b4fc;
}

.menu-item.router-link-active .menu-icon {
  color: #667eea;
  opacity: 1;
  transform: scale(1.15);
  filter: drop-shadow(0 2px 4px rgba(102, 126, 234, 0.3));
}

.dark .menu-item.router-link-active .menu-icon {
  color: #a5b4fc;
  filter: drop-shadow(0 2px 4px rgba(165, 180, 252, 0.4));
}

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

.dark .menu-item:hover {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

.sidebar[style*="72px"] .menu-content {
  justify-content: center;
  padding-left: 0;
  padding-right: 0;
}

.sidebar[style*="72px"] .menu-text {
  display: none;
}
</style>