<script setup lang="ts">
import { Icon } from '@iconify/vue';
import { WindowMinimise, WindowMaximise, WindowUnmaximise, Quit } from "../../wailsjs/runtime/runtime";
import { ref, onMounted } from 'vue';

// 是否最大化状态
const isMaximised = ref(false);

// 检查是否为Mac平台
const isMac = navigator.userAgent.toUpperCase().indexOf('MAC') >= 0;

// 最小化窗口
const minimizeWindow = () => {
  WindowMinimise();
};

// 切换最大化状态
const toggleMaximize = async () => {
  if (isMaximised.value) {
    await WindowUnmaximise();
    isMaximised.value = false;
  } else {
    await WindowMaximise();
    isMaximised.value = true;
  }
};

// 关闭应用
const closeWindow = () => {
  Quit();
};

// 组件挂载时检查窗口状态
onMounted(async () => {
  // 这里可以添加检查窗口状态的逻辑
  // 由于Wails的API限制，我们暂时使用默认状态
});
</script>

<template>
  <!-- 只在非Mac平台显示窗口控制按钮 -->
  <div v-if="!isMac" class="flex items-center">
    <!-- 最小化按钮 -->
    <button
      @click="minimizeWindow"
      class="window-control-btn hover:bg-gray-200 dark:hover:bg-gray-700"
      title="最小化"
    >
      <Icon icon="fluent:subtract-16-regular" class="text-sm" />
    </button>

    <!-- 最大化/还原按钮 -->
    <button
      @click="toggleMaximize"
      class="window-control-btn hover:bg-gray-200 dark:hover:bg-gray-700"
      :title="isMaximised ? '还原' : '最大化'"
    >
      <Icon
        :icon="isMaximised ? 'fluent:square-multiple-16-regular' : 'fluent:maximize-16-regular'"
        class="text-sm"
      />
    </button>

    <!-- 关闭按钮 -->
    <button
      @click="closeWindow"
      class="window-control-btn hover:bg-red-500 hover:text-white"
      title="关闭"
    >
      <Icon icon="fluent:dismiss-16-regular" class="text-sm" />
    </button>
  </div>
</template>

<style scoped>
.window-control-btn {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6b7280;
  transition: all 0.2s ease-in-out;
  border-radius: 4px;
  border: none;
  background: transparent;
  font-size: 14px;
}

.dark .window-control-btn {
  color: #9ca3af;
}

.window-control-btn:hover {
  color: #374151;
  transform: scale(1.05);
}

.dark .window-control-btn:hover {
  color: #d1d5db;
}

.window-control-btn:active {
  transform: scale(0.95);
  transition-duration: 75ms;
}

/* 最小化和最大化按钮的悬停效果 */
.window-control-btn:not(:last-child):hover {
  background-color: rgba(156, 163, 175, 0.1);
}

.dark .window-control-btn:not(:last-child):hover {
  background-color: rgba(156, 163, 175, 0.15);
}

/* 关闭按钮的特殊悬停效果 */
.window-control-btn:last-child:hover {
  background-color: #ef4444;
  color: white;
}

/* 添加微妙的阴影效果 */
.window-control-btn:hover {
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.08);
}

/* 按钮组的整体样式 */
.window-control-btn + .window-control-btn {
  margin-left: 1px;
}
</style>
