import { defineStore } from 'pinia'
import { ref } from 'vue'

interface Document {
  id: string
  issue_id: string | null
  project_id: string
  title: string
  file_path: string
  encrypted_name: string
  file_type: string
  file_size: number
  created_at: string
  updated_at: string
}

export const useDocumentStore = defineStore('document', () => {
  const documents = ref<Document[]>([])

  async function fetchDocuments(projectId: string, issueId?: string) {
    try {
      // @ts-ignore
      const result = await window.go.main.DocumentHandler.ListDocuments(projectId, issueId || null)
      documents.value = result || []
    } catch (error) {
      console.error('Failed to fetch documents:', error)
    }
  }

  async function uploadDocument(projectId: string, issueId: string | null, title: string, fileName: string, fileContent: number[], password: string) {
    try {
      // @ts-ignore
      const doc = await window.go.main.DocumentHandler.UploadDocument(
        projectId, issueId, title, fileName, fileContent, password
      )
      documents.value.unshift(doc)
      return doc
    } catch (error) {
      console.error('Failed to upload document:', error)
      throw error
    }
  }

  async function deleteDocument(id: string) {
    try {
      // @ts-ignore
      await window.go.main.DocumentHandler.DeleteDocument(id)
      documents.value = documents.value.filter(d => d.id !== id)
    } catch (error) {
      console.error('Failed to delete document:', error)
      throw error
    }
  }

  async function exportDocument(id: string, outputPath: string, password: string) {
    try {
      // @ts-ignore
      await window.go.main.DocumentHandler.ExportDocument(id, outputPath, password)
    } catch (error) {
      console.error('Failed to export document:', error)
      throw error
    }
  }

  return {
    documents,
    fetchDocuments,
    uploadDocument,
    deleteDocument,
    exportDocument
  }
})
