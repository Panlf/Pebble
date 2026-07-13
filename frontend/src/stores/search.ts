import { defineStore } from 'pinia'
import { ref } from 'vue'

interface SearchResult {
  type: string
  id: string
  title: string
  snippet: string
  project_id: string
  created_at: string
}

export const useSearchStore = defineStore('search', () => {
  const results = ref<SearchResult[]>([])
  const isLoading = ref(false)
  const searchHistory = ref<string[]>([])

  async function search(query: string, projectId?: string) {
    if (!query.trim()) {
      results.value = []
      return
    }

    isLoading.value = true
    try {
      // @ts-ignore
      const searchResults = await window.go.main.SearchHandler.Search(query, projectId || '')
      results.value = searchResults || []

      // Add to history
      if (!searchHistory.value.includes(query)) {
        searchHistory.value.unshift(query)
        if (searchHistory.value.length > 20) {
          searchHistory.value.pop()
        }
      }
    } catch (error) {
      console.error('Search failed:', error)
      results.value = []
    } finally {
      isLoading.value = false
    }
  }

  function clearResults() {
    results.value = []
  }

  function clearHistory() {
    searchHistory.value = []
  }

  return {
    results,
    isLoading,
    searchHistory,
    search,
    clearResults,
    clearHistory
  }
})
