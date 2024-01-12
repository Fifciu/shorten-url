import { reactive, computed, ref } from 'vue'
import { defineStore } from 'pinia'
import { linksService } from '@/services/links';
import type { Link, UpdateLinkDto } from '@/services/links';

const SORT_OPTIONS = [
  {
    label: 'Updated At',
    value: 'updatedAt',
  },
  {
    label: 'Updated At Descending',
    value: 'updatedAt:desc',
  },
  {
    label: 'Name',
    value: 'name',
  },
  {
    label: 'Name Descending',
    value: 'name:desc',
  },
  {
    label: 'Favourites',
    value: 'favourites',
  },
  {
    label: 'Favourites Descending',
    value: 'favourites:desc',
  },
];

export const useUiStore = defineStore('ui', () => {
  const searchQuery = ref('');
  const sortBy = ref(SORT_OPTIONS[0].value);
  const sortObject = computed(() => SORT_OPTIONS.find(opt => opt.value === sortBy.value));

  return {
    searchQuery,
    sortBy,
    sortObject,
    SORT_OPTIONS
  }
});
