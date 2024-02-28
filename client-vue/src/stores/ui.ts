import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useUiStore = defineStore('ui', () => {
  const isBodyScrollLocked = ref(false);

  function setBodyScroll (value: boolean) {
    isBodyScrollLocked.value = value;
  }

  return {
    isBodyScrollLocked: computed(() => isBodyScrollLocked.value),
    setBodyScroll
  }
});
