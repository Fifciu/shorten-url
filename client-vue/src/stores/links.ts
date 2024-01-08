import { reactive, computed } from 'vue'
import { defineStore } from 'pinia'
import { linksService } from '@/services/links';
import type { Link, UpdateLinkDto } from '@/services/links';

export const useLinksStore = defineStore('links', () => {
  const links = reactive<Link[]>([]);
  async function fetchMyLinks () {
    try {
      const results = await linksService.getMyMany();
      links.splice(0, links.length);
      links.push(...results);
    } catch (err) {
      console.error(err);
    }
  }

  async function addNewLink(name: string, original_url: string, alias?: string) {
    try {
      const link = await linksService.add(name, original_url, alias);
      links.push(link);
    } catch (err) {
      console.error(err);
    }
  }

  async function deleteLink(linkId: number) {
    try {
      await linksService.delete(linkId);
      links.splice(links.findIndex(link => link.id === linkId), 1);
    } catch (err) {
      console.error(err);
    }
  }

  async function updateLink(linkId: number, dto: UpdateLinkDto) {
    try {
      const link = await linksService.update(linkId, dto);
      links.splice(links.findIndex(link => link.id === linkId), 1, link);
    } catch (err) {
      console.log(err);
    }
  }

  return {
    links,
    fetchMyLinks,
    addNewLink,
    deleteLink,
    updateLink
  }
});
