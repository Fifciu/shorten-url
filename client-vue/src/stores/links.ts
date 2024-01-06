import { reactive, computed } from 'vue'
import { defineStore } from 'pinia'
import { linksService } from '@/services/links';
import type { Link } from '@/services/links';

export const useLinksStore = defineStore('links', () => {
  const links = reactive<Link[]>([]);
  async function fetchMyLinks () {
    try {
      const results = await linksService.getMyMany();
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

  return {
    links,
    fetchMyLinks,
    addNewLink,
    deleteLink
  }
})

function parseJwt (token: string): Record<string, any> {
  const base64Url = token.split('.')[1];
  const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
  const jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
      return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
  }).join(''));

  return JSON.parse(jsonPayload);
}