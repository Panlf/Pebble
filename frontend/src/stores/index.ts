import { createPinia } from 'pinia'

export const pinia = createPinia()

// 导出所有store
export * from './theme'
export * from './app'
