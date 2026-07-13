import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const authenticatedProjects = ref<Set<string>>(new Set())

  function authenticateProject(projectId: string) {
    authenticatedProjects.value.add(projectId)
  }

  function isProjectAuthenticated(projectId: string): boolean {
    return authenticatedProjects.value.has(projectId)
  }

  function logoutProject(projectId: string) {
    authenticatedProjects.value.delete(projectId)
  }

  function logoutAll() {
    authenticatedProjects.value.clear()
  }

  return {
    authenticatedProjects,
    authenticateProject,
    isProjectAuthenticated,
    logoutProject,
    logoutAll
  }
})
