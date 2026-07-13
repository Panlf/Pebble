<script setup lang="ts">
import { useTagStore } from '../stores/tag'
import { onMounted } from 'vue'

const tagStore = useTagStore()

onMounted(() => {
  tagStore.fetchTags()
})

function formatDate(dateString: string) {
  return new Date(dateString).toLocaleDateString('zh-CN')
}
</script>

<template>
  <div class="space-y-2">
    <div 
      v-for="tag in tagStore.tags" 
      :key="tag.id"
      class="flex items-center justify-between p-2 rounded-lg border border-base-300"
    >
      <div class="flex items-center gap-2">
        <div 
          class="w-4 h-4 rounded-full" 
          :style="{ backgroundColor: tag.color }"
        ></div>
        <span>{{ tag.name }}</span>
      </div>
      <div class="flex gap-2">
        <button class="btn btn-xs btn-ghost">编辑</button>
        <button class="btn btn-xs btn-error">删除</button>
      </div>
    </div>
    
    <div v-if="tagStore.tags.length === 0" class="text-center text-base-content/50 py-4">
      暂无标签
    </div>
  </div>
</template>
