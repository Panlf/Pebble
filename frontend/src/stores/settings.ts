import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useSettingsStore = defineStore('settings', () => {
  const settings = ref<Record<string, string>>({})
  const theme = ref<'light' | 'dark'>('light')

  async function loadSettings() {
    try {
      // @ts-ignore
      const result = await window.go.main.SettingsHandler.GetAllSettings()
      settings.value = result || {}
      
      // Load theme
      if (settings.value.theme) {
        theme.value = settings.value.theme as 'light' | 'dark'
        applyTheme()
      }
    } catch (error) {
      console.error('Failed to load settings:', error)
    }
  }

  async function saveSetting(key: string, value: string) {
    try {
      // @ts-ignore
      await window.go.main.SettingsHandler.SetSetting(key, value)
      settings.value[key] = value
    } catch (error) {
      console.error('Failed to save setting:', error)
      throw error
    }
  }

  async function setTheme(newTheme: 'light' | 'dark') {
    theme.value = newTheme
    await saveSetting('theme', newTheme)
    applyTheme()
  }

  function applyTheme() {
    if (theme.value === 'dark') {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
  }

  function getStoragePath(): string {
    return settings.value.storage_path || ''
  }

  async function setStoragePath(path: string) {
    await saveSetting('storage_path', path)
  }

  return {
    settings,
    theme,
    loadSettings,
    saveSetting,
    setTheme,
    applyTheme,
    getStoragePath,
    setStoragePath
  }
})
