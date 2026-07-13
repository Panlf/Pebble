<script setup lang="ts">
import { ref } from 'vue'
import { useDocumentStore } from '../stores/document'
import { useProjectStore } from '../stores/project'
import { useIssueStore } from '../stores/issue'

const documentStore = useDocumentStore()
const projectStore = useProjectStore()
const issueStore = useIssueStore()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const title = ref('')
const file = ref<File | null>(null)
const password = ref('')
const error = ref('')
const isUploading = ref(false)

async function handleFileSelect(event: Event) {
  const target = event.target as HTMLInputElement
  if (target.files && target.files.length > 0) {
    file.value = target.files[0]
    if (!title.value) {
      title.value = file.value.name.replace(/\.[^/.]+$/, '')
    }
  }
}

async function handleDrop(event: DragEvent) {
  event.preventDefault()
  if (event.dataTransfer?.files && event.dataTransfer.files.length > 0) {
    file.value = event.dataTransfer.files[0]
    if (!title.value) {
      title.value = file.value.name.replace(/\.[^/.]+$/, '')
    }
  }
}

async function handleDragOver(event: DragEvent) {
  event.preventDefault()
}

async function handleSubmit() {
  if (!title.value || !file.value || !password.value) {
    error.value = '请填写标题、选择文件并输入密码'
    return
  }

  if (!projectStore.currentProject) {
    error.value = '请先选择项目'
    return
  }

  isUploading.value = true
  error.value = ''

  try {
    const arrayBuffer = await file.value.arrayBuffer()
    const fileContent = Array.from(new Uint8Array(arrayBuffer))

    await documentStore.uploadDocument(
      projectStore.currentProject.id,
      issueStore.currentIssue?.id || null,
      title.value,
      file.value.name,
      fileContent,
      password.value
    )
    emit('close')
  } catch (e: any) {
    error.value = e.message || '上传失败'
  } finally {
    isUploading.value = false
  }
}
</script>

<template>
  <div class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
    <div class="bg-base-100 rounded-lg p-6 w-96">
      <h2 class="text-xl font-bold mb-4">上传文档</h2>
      
      <form @submit.prevent="handleSubmit" class="space-y-4">
        <div>
          <label class="block text-sm font-medium mb-1">文档标题</label>
          <input 
            v-model="title" 
            type="text" 
            class="input input-bordered w-full"
            placeholder="请输入文档标题"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium mb-1">选择文件</label>
          <div 
            class="border-2 border-dashed border-base-300 rounded-lg p-4 text-center cursor-pointer hover:border-primary transition-colors"
            @drop="handleDrop"
            @dragover="handleDragOver"
            @click="($refs.fileInput as HTMLInputElement).click()"
          >
            <input 
              ref="fileInput"
              type="file" 
              class="hidden" 
              @change="handleFileSelect"
            />
            <div v-if="file" class="text-sm">
              <div class="font-medium">{{ file.name }}</div>
              <div class="text-base-content/70">{{ (file.size / 1024).toFixed(2) }} KB</div>
            </div>
            <div v-else class="text-base-content/50">
              点击或拖拽文件到此处
            </div>
          </div>
        </div>
        
        <div>
          <label class="block text-sm font-medium mb-1">加密密码</label>
          <input 
            v-model="password" 
            type="password" 
            class="input input-bordered w-full"
            placeholder="请输入加密密码"
          />
        </div>
        
        <div v-if="error" class="text-error text-sm">{{ error }}</div>
        
        <div class="flex gap-2">
          <button type="submit" class="btn btn-primary flex-1" :disabled="isUploading">
            {{ isUploading ? '上传中...' : '上传' }}
          </button>
          <button type="button" class="btn btn-ghost flex-1" @click="emit('close')">取消</button>
        </div>
      </form>
    </div>
  </div>
</template>
