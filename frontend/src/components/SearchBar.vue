<script setup lang="ts">
import { ref, watch } from 'vue'
import { useSearchStore } from '../stores/search'
import { useProjectStore } from '../stores/project'

const searchStore = useSearchStore()
const projectStore = useProjectStore()

const query = ref('')
const showResults = ref(false)

watch(query, (newQuery) => {
  if (newQuery.trim()) {
    searchStore.search(newQuery, projectStore.currentProject?.id)
    showResults.value = true
  } else {
    searchStore.clearResults()
    showResults.value = false
  }
})

function handleResultClick(result: any) {
  showResults.value = false
  query.value = ''
  
  if (result.type === 'project') {
    // Navigate to project
  } else if (result.type === 'issue') {
    // Navigate to issue
  } else if (result.type === 'document') {
    // Navigate to document
  }
}

function handleBlur() {
  setTimeout(() => {
    showResults.value = false
  }, 200)
}

function getTypeIcon(type: string) {
  const iconMap: Record<string, string> = {
    'project': '📁',
    'issue': '❓',
    'document': '📄',
    'tag': '🏷️'
  }
  return iconMap[type] || '📎'
}

function getTypeName(type: string) {
  const nameMap: Record<string, string> = {
    'project': '项目',
    'issue': '问题',
    'document': '文档',
    'tag': '标签'
  }
  return nameMap[type] || type
}
</script>

<template>
  <div class="relative">
    <div class="input-group">
      <input 
        v-model="query"
        type="text" 
        placeholder="搜索项目、问题、文档、标签..." 
        class="input input-bordered w-full"
        @focus="showResults = !!query.trim()"
        @blur="handleBlur"
      />
      <button class="btn btn-square">
        <span class="material-symbols">search</span>
      </button>
    </div>
    
    <div 
      v-if="showResults && (searchStore.results.length > 0 || searchStore.isLoading)"
      class="absolute top-full left-0 right-0 mt-2 bg-base-100 border border-base-300 rounded-lg shadow-lg z-50 max-h-96 overflow-y-auto"
    >
      <div v-if="searchStore.isLoading" class="p-4 text-center">
        <span class="loading loading-spinner loading-sm"></span>
      </div>
      
      <div v-else-if="searchStore.results.length === 0" class="p-4 text-center text-base-content/50">
        未找到相关结果
      </div>
      
      <div v-else>
        <div 
          v-for="result in searchStore.results" 
          :key="`${result.type}-${result.id}`"
          class="p-3 hover:bg-base-200 cursor-pointer border-b border-base-300 last:border-b-0"
          @mousedown="handleResultClick(result)"
        >
          <div class="flex items-center gap-2">
            <span>{{ getTypeIcon(result.type) }}</span>
            <span class="text-xs text-base-content/50">{{ getTypeName(result.type) }}</span>
          </div>
          <div class="font-medium">{{ result.title }}</div>
          <div class="text-sm text-base-content/70">{{ result.snippet }}</div>
          <div class="text-xs text-base-content/50 mt-1">{{ result.created_at }}</div>
        </div>
      </div>
    </div>
  </div>
</template>
