import { ref, reactive, computed, watch } from 'vue'
import { defineStore } from 'pinia'
import { linksService } from '@/services/links';
import type { PaginatedLinks, UpdateLinkDto } from '@/services/links';

const SORT_OPTIONS = [
  {
    label: 'Updated At',
    value: 'updated_at',
  },
  {
    label: 'Updated At Descending',
    value: 'updated_at:desc',
  },
  {
    label: 'Name',
    value: 'name',
  },
  {
    label: 'Name Descending',
    value: 'name:desc',
  },
  // {
  //   label: 'Favourites',
  //   value: 'favourites',
  // },
  // {
  //   label: 'Favourites Descending',
  //   value: 'favourites:desc',
  // },
] as const;

type SortOptions = typeof SORT_OPTIONS[number]["value"];

export const useLinksStore = defineStore('links', () => {
  const searchQuery = ref('');
  const sortBy = ref<SortOptions>('updated_at:desc');
  const sortObject = computed(() => SORT_OPTIONS.find(opt => opt.value === sortBy.value));
  const paginatedLinks = reactive<PaginatedLinks>({
    count: 0,
    per_page: 10,
    page: 1,
    page_amount: 1,
    links: []
  });

  watch(sortBy, () => {
    fetchMyLinks(sortBy.value)
  })

  async function fetchMyLinks (sortBy?: string) {
    try {
      const results = await linksService.getMyMany(1, sortBy);
      paginatedLinks.count = results.count;
      paginatedLinks.per_page = results.per_page;
      paginatedLinks.page = results.page;
      paginatedLinks.page_amount = results.page_amount;
      paginatedLinks.links.splice(0, paginatedLinks.links.length);
      paginatedLinks.links.push(...results.links);
    } catch (err) {
      console.error(err);
    }
  }

  async function addNewLink(name: string, original_url: string, alias?: string) {
    try {
      const link = await linksService.add(name, original_url, alias);
      if (sortBy.value === 'updatedAt:desc') {
        paginatedLinks.links.unshift(link);
      } else {
        await fetchMyLinks();
      }
    } catch (err) {
      console.error(err);
    }
  }

  async function deleteLink(linkId: number) {
    try {
      await linksService.delete(linkId);
      await fetchMyLinks();
    } catch (err) {
      console.error(err);
    }
  }

  async function updateLink(linkId: number, dto: UpdateLinkDto) {
    try {
      const link = await linksService.update(linkId, dto);
      paginatedLinks.links.splice(paginatedLinks.links.findIndex(link => link.id === linkId), 1, link);
    } catch (err) {
      console.log(err);
    }
  }

  return {
    links: computed(() => paginatedLinks.links),
    pagination: computed(() => ({
      count: paginatedLinks.count,
      per_page: paginatedLinks.per_page,
      page: paginatedLinks.page,
      page_amount: paginatedLinks.page_amount,
    })),
    fetchMyLinks,
    addNewLink,
    deleteLink,
    updateLink,

    searchQuery,
    sortBy,
    sortObject,
    SORT_OPTIONS
  }
});
