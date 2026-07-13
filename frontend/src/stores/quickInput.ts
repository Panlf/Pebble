import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useQuickInputStore = defineStore('quickInput', () => {
  const isVisible = ref(false)
  const selectedProjectId = ref<string | null>(null)

  function show() {
    isVisible.value = true
  }

  function hide() {
    isVisible.value = false
  }

  function toggle() {
    isVisible.value = !isVisible.value
  }

  function setSelectedProject(projectId: string) {
    selectedProjectId.value = projectId
  }

  return {
    isVisible,
    selectedProjectId,
    show,
    hide,
    toggle,
    setSelectedProject
  }
})
