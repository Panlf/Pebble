<script lang="ts" setup>
import { Icon } from '@iconify/vue';
import { useAppStore } from '../stores/app';
import { useThemeStore } from '../stores/theme';
import { ref } from 'vue';

const appStore = useAppStore();
const themeStore = useThemeStore();

// 示例计数器
const counter = ref(0);

// 示例表单数据
const formData = ref({
  name: '',
  email: '',
  message: ''
});

// 提交表单
const submitForm = () => {
  console.log('表单提交:', formData.value);
  formData.value = { name: '', email: '', message: '' };
  alert('表单提交成功！');
};
</script>

<template>
  <div class="min-h-full bg-base-100">
    <!-- 欢迎区域 -->
    <div class="hero min-h-96 bg-gradient-to-r from-primary to-secondary">
      <div class="hero-content text-center text-primary-content">
        <div class="max-w-md">
          <h1 class="mb-5 text-5xl font-bold">{{ appStore.appInfo.name }}</h1>
          <p class="mb-5">{{ appStore.appInfo.description }}</p>
          <div class="flex gap-2 justify-center">
            <div class="badge badge-outline">v{{ appStore.appInfo.version }}</div>
            <div class="badge badge-outline">{{ appStore.appInfo.license }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 技术栈展示 -->
    <div class="container mx-auto px-4 py-12">
      <h2 class="text-3xl font-bold text-center mb-8">技术栈</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        <div 
          v-for="tech in appStore.techStack" 
          :key="tech.name"
          class="card bg-base-200 shadow-xl hover:shadow-2xl transition-all duration-300 hover:scale-105"
        >
          <div class="card-body items-center text-center">
            <Icon 
              :icon="tech.icon" 
              class="text-4xl mb-2" 
              :style="{ color: tech.color }"
            />
            <h3 class="card-title text-lg">{{ tech.name }}</h3>
            <p class="text-sm opacity-70">{{ tech.version }}</p>
            <p class="text-xs">{{ tech.description }}</p>
            <div class="card-actions justify-end">
              <a 
                :href="tech.url" 
                target="_blank" 
                class="btn btn-primary btn-sm"
              >
                了解更多
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 统计数据展示 -->
    <div class="bg-base-200 py-12">
      <div class="container mx-auto px-4">
        <h2 class="text-3xl font-bold text-center mb-8">项目统计</h2>
        <div class="stats stats-vertical lg:stats-horizontal shadow w-full">
          <div class="stat">
            <div class="stat-figure text-primary">
              <Icon icon="icon-park:user" class="text-3xl" />
            </div>
            <div class="stat-title">总用户数</div>
            <div class="stat-value text-primary">{{ appStore.sampleData.stats.totalUsers }}</div>
            <div class="stat-desc">21% 比上月增长</div>
          </div>
          
          <div class="stat">
            <div class="stat-figure text-secondary">
              <Icon icon="icon-park:check-one" class="text-3xl" />
            </div>
            <div class="stat-title">活跃用户</div>
            <div class="stat-value text-secondary">{{ appStore.sampleData.stats.activeUsers }}</div>
            <div class="stat-desc">↗︎ 400 (22%)</div>
          </div>
          
          <div class="stat">
            <div class="stat-figure text-accent">
              <Icon icon="icon-park:folder" class="text-3xl" />
            </div>
            <div class="stat-title">项目数量</div>
            <div class="stat-value text-accent">{{ appStore.sampleData.stats.totalProjects }}</div>
            <div class="stat-desc">↘︎ 90 (14%)</div>
          </div>
          
          <div class="stat">
            <div class="stat-figure text-info">
              <Icon icon="icon-park:task" class="text-3xl" />
            </div>
            <div class="stat-title">完成任务</div>
            <div class="stat-value text-info">{{ appStore.sampleData.stats.completedTasks }}</div>
            <div class="stat-desc">↗︎ 90 (14%)</div>
          </div>
        </div>
      </div>
    </div>

    <!-- DaisyUI 组件展示 -->
    <div class="container mx-auto px-4 py-12">
      <h2 class="text-3xl font-bold text-center mb-8">DaisyUI 组件展示</h2>
      
      <!-- 按钮展示 -->
      <div class="mb-8">
        <h3 class="text-xl font-semibold mb-4">按钮组件</h3>
        <div class="flex flex-wrap gap-2">
          <button class="btn">默认</button>
          <button class="btn btn-primary">主要</button>
          <button class="btn btn-secondary">次要</button>
          <button class="btn btn-accent">强调</button>
          <button class="btn btn-info">信息</button>
          <button class="btn btn-success">成功</button>
          <button class="btn btn-warning">警告</button>
          <button class="btn btn-error">错误</button>
          <button class="btn btn-outline">轮廓</button>
          <button class="btn btn-ghost">幽灵</button>
          <button class="btn btn-link">链接</button>
        </div>
      </div>

      <!-- 计数器示例 -->
      <div class="mb-8">
        <h3 class="text-xl font-semibold mb-4">交互示例</h3>
        <div class="card bg-base-200 shadow-xl">
          <div class="card-body">
            <h4 class="card-title">计数器</h4>
            <div class="flex items-center gap-4">
              <button class="btn btn-circle btn-outline" @click="counter--">
                <Icon icon="icon-park:minus" />
              </button>
              <span class="text-2xl font-bold">{{ counter }}</span>
              <button class="btn btn-circle btn-outline" @click="counter++">
                <Icon icon="icon-park:plus" />
              </button>
            </div>
            <div class="card-actions justify-end">
              <button class="btn btn-primary" @click="counter = 0">重置</button>
            </div>
          </div>
        </div>
      </div>

      <!-- 主题切换 -->
      <div class="mb-8">
        <h3 class="text-xl font-semibold mb-4">主题切换</h3>
        <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-2">
          <button
            v-for="theme in themeStore.themes.slice(0, 12)"
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

      <!-- 表单示例 -->
      <div class="mb-8">
        <h3 class="text-xl font-semibold mb-4">表单组件</h3>
        <div class="card bg-base-200 shadow-xl">
          <div class="card-body">
            <h4 class="card-title">联系表单</h4>
            <form @submit.prevent="submitForm" class="space-y-4">
              <div class="form-control">
                <label class="label">
                  <span class="label-text">姓名</span>
                </label>
                <input
                  v-model="formData.name"
                  type="text"
                  placeholder="请输入您的姓名"
                  class="input input-bordered"
                  required
                />
              </div>

              <div class="form-control">
                <label class="label">
                  <span class="label-text">邮箱</span>
                </label>
                <input
                  v-model="formData.email"
                  type="email"
                  placeholder="请输入您的邮箱"
                  class="input input-bordered"
                  required
                />
              </div>

              <div class="form-control">
                <label class="label">
                  <span class="label-text">消息</span>
                </label>
                <textarea
                  v-model="formData.message"
                  class="textarea textarea-bordered"
                  placeholder="请输入您的消息"
                  rows="3"
                  required
                ></textarea>
              </div>

              <div class="form-control mt-6">
                <button type="submit" class="btn btn-primary">
                  <Icon icon="icon-park:send" class="mr-2" />
                  提交
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>

      <!-- 用户列表示例 -->
      <div class="mb-8">
        <h3 class="text-xl font-semibold mb-4">数据表格</h3>
        <div class="overflow-x-auto">
          <table class="table table-zebra">
            <thead>
              <tr>
                <th>ID</th>
                <th>姓名</th>
                <th>邮箱</th>
                <th>角色</th>
                <th>操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="user in appStore.sampleData.users" :key="user.id">
                <td>{{ user.id }}</td>
                <td>{{ user.name }}</td>
                <td>{{ user.email }}</td>
                <td>
                  <div class="badge" :class="user.role === 'admin' ? 'badge-primary' : 'badge-secondary'">
                    {{ user.role }}
                  </div>
                </td>
                <td>
                  <div class="flex gap-2">
                    <button class="btn btn-ghost btn-xs">编辑</button>
                    <button class="btn btn-ghost btn-xs text-error">删除</button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
