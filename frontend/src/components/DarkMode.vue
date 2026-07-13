<script lang="ts" setup>
import { Icon } from '@iconify/vue';
import { useThemeStore } from '../stores/theme';

const themeStore = useThemeStore();
</script>

<template>
  <button @click="themeStore.toggleDark()" class="text-2xl mt-1 hover:rotate-6">
    <span v-if="themeStore.isDark"><Icon icon="icon-park:dark-mode" /></span>
    <span v-else><Icon icon="icon-park:sun-one" /></span>
  </button>
</template>
