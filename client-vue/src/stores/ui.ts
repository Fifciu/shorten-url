import { reactive, computed, ref } from 'vue'
import { defineStore } from 'pinia'
import { linksService } from '@/services/links';
import type { Link, UpdateLinkDto } from '@/services/links';

export const useUiStore = defineStore('ui', () => {
  const searchQuery = ref('');

  return {
    searchQuery,
  }
});
