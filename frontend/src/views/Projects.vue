<script setup lang="ts">
import { ref } from 'vue'
import ProjectList from '../components/ProjectList.vue'
import ProjectForm from '../components/ProjectForm.vue'

const showCreateForm = ref(false)
</script>

<template>
  <div class="p-6">
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold">项目列表</h1>
      <button class="btn btn-primary" @click="showCreateForm = true">
        新建项目
      </button>
    </div>
    
    <ProjectList />
    
    <ProjectForm v-if="showCreateForm" @close="showCreateForm = false" />
  </div>
</template>
