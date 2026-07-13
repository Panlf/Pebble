import { defineStore } from 'pinia'
import { ref } from 'vue'

interface Issue {
  id: string
  project_id: string
  title: string
  description: string
  status: string
  contact_person: string
  created_at: string
  updated_at: string
}

export const useIssueStore = defineStore('issue', () => {
  const issues = ref<Issue[]>([])
  const currentIssue = ref<Issue | null>(null)

  async function fetchIssues(projectId: string) {
    try {
      // @ts-ignore
      const result = await window.go.main.IssueHandler.ListIssues(projectId)
      issues.value = result || []
    } catch (error) {
      console.error('Failed to fetch issues:', error)
    }
  }

  async function createIssue(projectId: string, title: string, description: string, contactPerson: string) {
    try {
      // @ts-ignore
      const issue = await window.go.main.IssueHandler.CreateIssue(projectId, title, description, contactPerson)
      issues.value.unshift(issue)
      return issue
    } catch (error) {
      console.error('Failed to create issue:', error)
      throw error
    }
  }

  async function updateIssue(id: string, title: string, description: string, status: string, contactPerson: string) {
    try {
      // @ts-ignore
      await window.go.main.IssueHandler.UpdateIssue(id, title, description, status, contactPerson)
      const index = issues.value.findIndex(i => i.id === id)
      if (index !== -1) {
        issues.value[index] = { ...issues.value[index], title, description, status, contact_person: contactPerson }
      }
    } catch (error) {
      console.error('Failed to update issue:', error)
      throw error
    }
  }

  async function deleteIssue(id: string) {
    try {
      // @ts-ignore
      await window.go.main.IssueHandler.DeleteIssue(id)
      issues.value = issues.value.filter(i => i.id !== id)
      if (currentIssue.value?.id === id) {
        currentIssue.value = null
      }
    } catch (error) {
      console.error('Failed to delete issue:', error)
      throw error
    }
  }

  async function updateStatus(id: string, status: string) {
    try {
      // @ts-ignore
      await window.go.main.IssueHandler.UpdateStatus(id, status)
      const index = issues.value.findIndex(i => i.id === id)
      if (index !== -1) {
        issues.value[index] = { ...issues.value[index], status }
      }
    } catch (error) {
      console.error('Failed to update status:', error)
      throw error
    }
  }

  function setCurrentIssue(issue: Issue) {
    currentIssue.value = issue
  }

  return {
    issues,
    currentIssue,
    fetchIssues,
    createIssue,
    updateIssue,
    deleteIssue,
    updateStatus,
    setCurrentIssue
  }
})
