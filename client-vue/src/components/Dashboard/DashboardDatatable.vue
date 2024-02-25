<script setup lang="ts">
import BaseSelect from '@/components/Base/BaseSelect.vue';
import BaseSearch from '@/components/Base/BaseSearch.vue';
import BaseButton from '@/components/Base/BaseButton.vue';
import MobileDashboardLink from '@/components/Dashboard/MobileDashboardLink.vue';
import MobileSortBy from '@/components/Dashboard/MobileSortBy.vue';
import DashboardNoLinksMatchingCriteria from '@/components/Dashboard/DashboardNoLinksMatchingCriteria.vue';
import CopyButton from '@/components/Actions/CopyButton.vue';
import EditButton from '@/components/Actions/EditButton.vue';
import WishlistButton from '@/components/Actions/WishlistButton.vue';
import BinButton from '@/components/Actions/BinButton.vue';
import DashboardLoadMore from '@/components/Dashboard/DashboardLoadMore.vue';
import { computed, defineEmits } from 'vue';

import { useLinksStore } from '@/stores/links';

import { REDIRECT_BASE_URL } from '@/const';
import type { Link } from '@/services/links';
import { storeToRefs } from 'pinia';

const emit = defineEmits<{
  (e: 'openAddLinkModal'): void,
  (e: 'openEditLinkModal', link: Link): void, // Is it proper type? TODO
}>();

const { links, SORT_OPTIONS } = useLinksStore();
const { sortBy, searchQuery } = storeToRefs(useLinksStore())
const filteredLinks = computed(() => {
  if (!searchQuery) {
    return links;
  }
  const searchQueryLowerCased = (''+searchQuery).toLocaleLowerCase();
  return links.filter(link => 
    link.name.toLocaleLowerCase().match(searchQueryLowerCased)
    || link.alias.toLocaleLowerCase().match(searchQueryLowerCased)
    || link.original_url.toLowerCase().match(searchQueryLowerCased)
  );
});
</script>

<template>
  <div class="datatable-desktop">
    <div class="actions">
      <BaseSearch uniqueId="filterLinksByTextDesktop" type="search" label="Your Links" label-style="dark-grey" placeholder="Search"
        v-model="searchQuery" />
      <BaseSelect v-model="sortBy" :fields="SORT_OPTIONS" />
      <BaseButton variant="primary" @click="emit('openAddLinkModal')">New link</BaseButton>
    </div>
    <table class="content" v-if="filteredLinks.length">
      <tr class="headers">
        <th>Name</th>
        <th>Updated at</th>
        <th>Short link</th>
        <th></th>
      </tr>
      <tr class="records" v-for="record in filteredLinks" :key="record.id">
        <td class="name">{{ record.name }}</td>
        <td class="date">{{ record.updated_at }}</td>
        <td class="link">{{ REDIRECT_BASE_URL }}{{ record.alias }}</td>
        <td class="actions">
          <CopyButton :toCopy="REDIRECT_BASE_URL + record.alias" />
          <EditButton @click="emit('openEditLinkModal', record)" />
          <BinButton :linkId="record.id" />
          <WishlistButton :linkId="record.id" />
        </td>
      </tr>
    </table>
  </div>
  <div class="datatable-mobile">
    <div class="actions px-2">
      <BaseButton variant="primary" class="w-100" @click="emit('openAddLinkModal')">New link</BaseButton>
      <MobileSortBy />
    </div>
    <div class="content px-2" v-if="filteredLinks.length">
      <MobileDashboardLink v-for="record in filteredLinks" :key="record.id" :id="record.id" :name="record.name" :short_link="record.alias"
        :updated_at="record.updated_at" @edit="emit('openEditLinkModal', record)" />
    </div>
  </div>
  <DashboardLoadMore />
  <DashboardNoLinksMatchingCriteria v-if="!filteredLinks.length" class="mt-9 md:mt-2" />
</template>

<style lang="scss" scoped>
.datatable-desktop {
  display: none;
}

.sort-by--mobile {
  color: $colorDarkGrey;
  text-align: center;
  font-family: Source Sans Pro;
  font-size: 12px;
  font-style: normal;
  font-weight: 600;
  line-height: 150%;
  /* 18px */
  text-transform: uppercase;
  margin: 16px 0;
}

@media only screen and (min-width: $widthDesktop) {
  .datatable-desktop {
    display: block;
    max-width: 752px;
    margin: 0 auto;

    .actions {
      display: flex;
      align-items: flex-end;
      justify-content: space-between;
    }

    .content {
      padding: 16px 32px;
      margin-top: 16px;
      margin-bottom: 52px;
      width: 100%;
      border-radius: 8px;
      border: 2px solid $colorWhiteBlue;
      background: $colorWhite;

      th,
      td {
        text-align: left;
      }

      .headers {
        color: $colorGrey;
        font-family: Source Sans Pro;
        font-size: 10px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%;

        th {
          text-align: left;
          padding: 10px;
        }
      }

      .records {
        td {
          padding: 10px;
        }

        .name {
          color: $colorBlack;
          font-family: Source Sans Pro;
          font-size: 14px;
          font-style: normal;
          font-weight: 600;
          line-height: 150%;
        }

        .date,
        .link {
          color: $colorBlack;
          font-family: Source Sans Pro;
          font-size: 14px;
          font-style: normal;
          font-weight: 400;
          line-height: 150%;
        }
      }
    }
  }

  .datatable-mobile {
    display: none;
  }
}
</style>
