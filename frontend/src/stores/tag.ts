import { defineStore } from 'pinia'
import { ref } from 'vue'

interface Tag {
  id: string
  name: string
  color: string
  created_at: string
}

export const useTagStore = defineStore('tag', () => {
  const tags = ref<Tag[]>([])

  async function fetchTags() {
    try {
      // @ts-ignore
      const result = await window.go.main.TagHandler.ListTags()
      tags.value = result || []
    } catch (error) {
      console.error('Failed to fetch tags:', error)
    }
  }

  async function createTag(name: string, color: string) {
    try {
      // @ts-ignore
      const tag = await window.go.main.TagHandler.CreateTag(name, color)
      tags.value.push(tag)
      return tag
    } catch (error) {
      console.error('Failed to create tag:', error)
      throw error
    }
  }

  async function updateTag(id: string, name: string, color: string) {
    try {
      // @ts-ignore
      await window.go.main.TagHandler.UpdateTag(id, name, color)
      const index = tags.value.findIndex(t => t.id === id)
      if (index !== -1) {
        tags.value[index] = { ...tags.value[index], name, color }
      }
    } catch (error) {
      console.error('Failed to update tag:', error)
      throw error
    }
  }

  async function deleteTag(id: string) {
    try {
      // @ts-ignore
      await window.go.main.TagHandler.DeleteTag(id)
      tags.value = tags.value.filter(t => t.id !== id)
    } catch (error) {
      console.error('Failed to delete tag:', error)
      throw error
    }
  }

  async function addTagToDocument(documentId: string, tagId: string) {
    try {
      // @ts-ignore
      await window.go.main.TagHandler.AddTagToDocument(documentId, tagId)
    } catch (error) {
      console.error('Failed to add tag to document:', error)
      throw error
    }
  }

  async function removeTagFromDocument(documentId: string, tagId: string) {
    try {
      // @ts-ignore
      await window.go.main.TagHandler.RemoveTagFromDocument(documentId, tagId)
    } catch (error) {
      console.error('Failed to remove tag from document:', error)
      throw error
    }
  }

  async function getDocumentTags(documentId: string) {
    try {
      // @ts-ignore
      return await window.go.main.TagHandler.GetDocumentTags(documentId)
    } catch (error) {
      console.error('Failed to get document tags:', error)
      return []
    }
  }

  return {
    tags,
    fetchTags,
    createTag,
    updateTag,
    deleteTag,
    addTagToDocument,
    removeTagFromDocument,
    getDocumentTags
  }
})
