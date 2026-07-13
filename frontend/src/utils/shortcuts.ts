import { useQuickInputStore } from '../stores/quickInput'

export function initShortcuts() {
  const quickInputStore = useQuickInputStore()

  document.addEventListener('keydown', (event) => {
    // Ctrl+N: Quick input
    if (event.ctrlKey && event.key === 'n') {
      event.preventDefault()
      quickInputStore.toggle()
    }
    
    // Escape: Close quick input
    if (event.key === 'Escape') {
      if (quickInputStore.isVisible) {
        quickInputStore.hide()
      }
    }
  })
}
