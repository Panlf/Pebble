<script lang="ts" setup>
import { Icon } from '@iconify/vue';
import { useThemeStore } from '../stores/theme';
import { useAppStore } from '../stores/app';
import { ref } from 'vue';

const themeStore = useThemeStore();
const appStore = useAppStore();

// 设置选项
const settings = ref({
  autoStart: false,
  minimizeToTray: true,
  notifications: true,
  autoUpdate: false,
  language: 'zh-CN'
});

// 语言选项
const languages = [
  { value: 'zh-CN', label: '简体中文' },
  { value: 'en-US', label: 'English' },
  { value: 'ja-JP', label: '日本語' }
];

// 保存设置
const saveSettings = () => {
  localStorage.setItem('app-settings', JSON.stringify(settings.value));
  alert('设置已保存！');
};

// 重置设置
const resetSettings = () => {
  if (confirm('确定要重置所有设置吗？')) {
    settings.value = {
      autoStart: false,
      minimizeToTray: true,
      notifications: true,
      autoUpdate: false,
      language: 'zh-CN'
    };
    themeStore.setTheme('light');
  }
};
</script>

<template>
  <div class="min-h-full bg-base-100 p-6">
    <!-- 页面标题 -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-base-content flex items-center gap-3">
        <Icon icon="material-symbols:settings" class="text-primary" />
        应用设置
      </h1>
      <p class="text-base-content/70 mt-2">配置您的应用偏好设置</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- 主题设置 -->
      <div class="card bg-base-200 shadow-xl">
        <div class="card-body">
          <h2 class="card-title">
            <Icon icon="material-symbols:palette" class="text-secondary" />
            主题设置
          </h2>

          <div class="form-control">
            <label class="label">
              <span class="label-text">当前主题</span>
            </label>
            <select
              v-model="themeStore.currentTheme"
              @change="themeStore.setTheme(themeStore.currentTheme)"
              class="select select-bordered w-full"
            >
              <option v-for="theme in themeStore.themes" :key="theme.name" :value="theme.name">
                {{ theme.label }}
              </option>
            </select>
          </div>

          <div class="grid grid-cols-2 md:grid-cols-3 gap-2 mt-4">
            <button
              v-for="theme in themeStore.themes.slice(0, 6)"
              :key="theme.name"
              @click="themeStore.setTheme(theme.name)"
              class="btn btn-sm"
              :class="{ 'btn-primary': themeStore.currentTheme === theme.name }"
            >
              <Icon :icon="theme.icon" class="mr-1" />
              {{ theme.label }}
            </button>
          </div>
        </div>
      </div>

      <!-- 应用设置 -->
      <div class="card bg-base-200 shadow-xl">
        <div class="card-body">
          <h2 class="card-title">
            <Icon icon="material-symbols:tune" class="text-accent" />
            应用设置
          </h2>

          <div class="space-y-4">
            <div class="form-control">
              <label class="label cursor-pointer">
                <span class="label-text">开机自启动</span>
                <input v-model="settings.autoStart" type="checkbox" class="toggle toggle-primary" />
              </label>
            </div>

            <div class="form-control">
              <label class="label cursor-pointer">
                <span class="label-text">最小化到系统托盘</span>
                <input v-model="settings.minimizeToTray" type="checkbox" class="toggle toggle-primary" />
              </label>
            </div>

            <div class="form-control">
              <label class="label cursor-pointer">
                <span class="label-text">桌面通知</span>
                <input v-model="settings.notifications" type="checkbox" class="toggle toggle-primary" />
              </label>
            </div>

            <div class="form-control">
              <label class="label cursor-pointer">
                <span class="label-text">自动更新</span>
                <input v-model="settings.autoUpdate" type="checkbox" class="toggle toggle-primary" />
              </label>
            </div>

            <div class="form-control">
              <label class="label">
                <span class="label-text">语言</span>
              </label>
              <select v-model="settings.language" class="select select-bordered">
                <option v-for="lang in languages" :key="lang.value" :value="lang.value">
                  {{ lang.label }}
                </option>
              </select>
            </div>
          </div>
        </div>
      </div>

      <!-- 应用信息 -->
      <div class="card bg-base-200 shadow-xl">
        <div class="card-body">
          <h2 class="card-title">
            <Icon icon="material-symbols:info" class="text-info" />
            应用信息
          </h2>

          <div class="space-y-3">
            <div class="flex justify-between">
              <span class="text-base-content/70">应用名称</span>
              <span class="font-medium">{{ appStore.appInfo.name }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-base-content/70">版本</span>
              <div class="flex items-center gap-2">
                <span class="font-medium">{{ appStore.appInfo.version }}</span>
                <div class="badge badge-success badge-sm">最新</div>
              </div>
            </div>
            <div class="flex justify-between">
              <span class="text-base-content/70">作者</span>
              <span class="font-medium">{{ appStore.appInfo.author }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-base-content/70">许可证</span>
              <span class="font-medium">{{ appStore.appInfo.license }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-base-content/70">仓库</span>
              <a :href="appStore.appInfo.repository" target="_blank" class="link link-primary">
                GitHub
              </a>
            </div>
          </div>
        </div>
      </div>

      <!-- 操作按钮 -->
      <div class="card bg-base-200 shadow-xl">
        <div class="card-body">
          <h2 class="card-title">
            <Icon icon="material-symbols:build" class="text-warning" />
            操作
          </h2>

          <div class="space-y-3">
            <button @click="saveSettings" class="btn btn-primary w-full">
              <Icon icon="material-symbols:save" class="mr-2" />
              保存设置
            </button>

            <button @click="resetSettings" class="btn btn-warning w-full">
              <Icon icon="material-symbols:refresh" class="mr-2" />
              重置设置
            </button>

            <div class="divider">危险操作</div>

            <button class="btn btn-error w-full" onclick="confirm('确定要清除所有数据吗？') && localStorage.clear()">
              <Icon icon="material-symbols:delete-forever" class="mr-2" />
              清除所有数据
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- 技术栈信息 -->
    <div class="mt-8">
      <h2 class="text-2xl font-bold mb-4 flex items-center gap-2">
        <Icon icon="material-symbols:code" class="text-primary" />
        技术栈
      </h2>
      <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-7 gap-4">
        <div
          v-for="tech in appStore.techStack"
          :key="tech.name"
          class="card bg-base-200 shadow-lg hover:shadow-xl transition-all duration-300 hover:scale-105"
        >
          <div class="card-body items-center text-center p-4">
            <Icon
              :icon="tech.icon"
              class="text-3xl mb-2"
              :style="{ color: tech.color }"
            />
            <h3 class="font-semibold text-sm">{{ tech.name }}</h3>
            <p class="text-xs opacity-70">{{ tech.version }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
