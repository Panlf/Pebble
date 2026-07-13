import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import { pinia } from './stores'
import { initShortcuts } from './utils/shortcuts'

const app = createApp(App)
app.use(pinia)
app.use(router)

initShortcuts()

app.mount('#app')
