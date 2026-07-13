<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useSettingsStore } from '../stores/settings'

const settingsStore = useSettingsStore()

const storagePath = ref('')
const theme = ref<'light' | 'dark'>('light')

onMounted(async () => {
  await settingsStore.loadSettings()
  storagePath.value = settingsStore.getStoragePath()
  theme.value = settingsStore.theme
})

async function handleSave() {
  await settingsStore.setStoragePath(storagePath.value)
  await settingsStore.setTheme(theme.value)
}

async function handleSelectPath() {
  // @ts-ignore
  const path = await window.go.runtime.OpenDirectoryDialog({
    title: '选择存储路径',
    defaultPath: storagePath.value
  })
  if (path) {
    storagePath.value = path
  }
}
</script>

<template>
  <div class="p-6 max-w-2xl mx-auto">
    <h1 class="text-3xl font-bold mb-6">设置</h1>
    
    <div class="space-y-6">
      <div class="card bg-base-200">
        <div class="card-body">
          <h2 class="card-title">存储设置</h2>
          
          <div class="form-control">
            <label class="label">
              <span class="label-text">文档存储路径</span>
            </label>
            <div class="input-group">
              <input 
                v-model="storagePath"
                type="text" 
                class="input input-bordered flex-1"
                placeholder="请选择文档存储路径"
              />
              <button class="btn" @click="handleSelectPath">选择</button>
            </div>
            <label class="label">
              <span class="label-text-alt">文档将加密存储在此路径下</span>
            </label>
          </div>
        </div>
      </div>
      
      <div class="card bg-base-200">
        <div class="card-body">
          <h2 class="card-title">外观设置</h2>
          
          <div class="form-control">
            <label class="label">
              <span class="label-text">主题</span>
            </label>
            <div class="flex gap-4">
              <label class="cursor-pointer label gap-2">
                <input 
                  type="radio" 
                  name="theme"
                  class="radio radio-primary" 
                  :checked="theme === 'light'"
                  @change="theme = 'light'"
                />
                <span class="label-text">浅色</span>
              </label>
              <label class="cursor-pointer label gap-2">
                <input 
                  type="radio" 
                  name="theme"
                  class="radio radio-primary" 
                  :checked="theme === 'dark'"
                  @change="theme = 'dark'"
                />
                <span class="label-text">深色</span>
              </label>
            </div>
          </div>
        </div>
      </div>
      
      <div class="flex justify-end">
        <button class="btn btn-primary" @click="handleSave">保存设置</button>
      </div>
    </div>
  </div>
</template>
