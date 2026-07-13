import { defineStore } from 'pinia'
import { ref } from 'vue'

interface Project {
  id: string
  name: string
  description: string
  created_at: string
  updated_at: string
  storage_path: string
}

export const useProjectStore = defineStore('project', () => {
  const projects = ref<Project[]>([])
  const currentProject = ref<Project | null>(null)
  const isAuthenticated = ref<Record<string, boolean>>({})

  async function fetchProjects() {
    try {
      // @ts-ignore
      const result = await window.go.main.ProjectHandler.ListProjects()
      projects.value = result || []
    } catch (error) {
      console.error('Failed to fetch projects:', error)
    }
  }

  async function createProject(name: string, description: string, password: string) {
    try {
      // @ts-ignore
      const project = await window.go.main.ProjectHandler.CreateProject(name, description, password)
      projects.value.unshift(project)
      return project
    } catch (error) {
      console.error('Failed to create project:', error)
      throw error
    }
  }

  async function updateProject(id: string, name: string, description: string) {
    try {
      // @ts-ignore
      await window.go.main.ProjectHandler.UpdateProject(id, name, description)
      const index = projects.value.findIndex(p => p.id === id)
      if (index !== -1) {
        projects.value[index] = { ...projects.value[index], name, description }
      }
    } catch (error) {
      console.error('Failed to update project:', error)
      throw error
    }
  }

  async function deleteProject(id: string) {
    try {
      // @ts-ignore
      await window.go.main.ProjectHandler.DeleteProject(id)
      projects.value = projects.value.filter(p => p.id !== id)
      if (currentProject.value?.id === id) {
        currentProject.value = null
      }
    } catch (error) {
      console.error('Failed to delete project:', error)
      throw error
    }
  }

  async function verifyPassword(id: string, password: string): Promise<boolean> {
    try {
      // @ts-ignore
      const valid = await window.go.main.ProjectHandler.VerifyPassword(id, password)
      if (valid) {
        isAuthenticated.value[id] = true
      }
      return valid
    } catch (error) {
      console.error('Failed to verify password:', error)
      return false
    }
  }

  function setCurrentProject(project: Project) {
    currentProject.value = project
  }

  function isProjectAuthenticated(id: string): boolean {
    return isAuthenticated.value[id] === true
  }

  return {
    projects,
    currentProject,
    isAuthenticated,
    fetchProjects,
    createProject,
    updateProject,
    deleteProject,
    verifyPassword,
    setCurrentProject,
    isProjectAuthenticated
  }
})
