import { sendToApi, sendToApiExpectBody } from "./base";

const MODULE_BASE_PATH = 'links';

export type Link = {
  id: number,
  user_id: number,
  name: string,
  original_url: string,
  updated_at: string, // TODO: To date
  alias: string,
};

export type PaginatedLinks = {
  count: number,
  per_page: number,
  page: number,
  page_amount: number,
  links: Link[],
}

export type UpdateLinkDto = {
  name?: string,
  original_url?: string,
  alias?: string,
};

export const linksService = {
  async add(name: string, original_url: string, alias?: string) {
    return sendToApiExpectBody<Link>(`${MODULE_BASE_PATH}/`, {
      method: 'POST',
      body: JSON.stringify({
        name,
        original_url,
        alias
      })
    });
  },

  async delete(id: number) {
    await sendToApi(`${MODULE_BASE_PATH}/${id}`, {
      method: 'DELETE'
    });
  },

  async getMyMany(page = 1, sortBy?: string) {
    let url = `${MODULE_BASE_PATH}/?page=${page}`;
    if (sortBy) {
      url += `&sortBy=${sortBy}`
    }
    return await sendToApiExpectBody<PaginatedLinks>(url, {
      method: 'GET'
    });
  },

  async update(id: number, updateLinkDto: UpdateLinkDto) {
    return await sendToApiExpectBody<Link>(`${MODULE_BASE_PATH}/${id}`, {
      method: 'PATCH',
      body: JSON.stringify(updateLinkDto)
    });
  },
};
