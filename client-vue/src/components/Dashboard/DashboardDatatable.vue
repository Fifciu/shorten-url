<script setup lang="ts">
import DashboardCTA from '@/components/Dashboard/DashboardCTA.vue';
import ShInput from '@/components/ShInput.vue';
import ShButton from '@/components/ShButton.vue';
import MobileDashboardLink from '@/components/Dashboard/MobileDashboardLink.vue';
import { ref } from 'vue';

import { useUserStore } from '@/stores/user';

const { user } = useUserStore();
const searchValue = ref('');

const records = [
  {
    id: 1,
    user_id: 3,
    name: 'My Youtube',
    original_url: 'https://fifciuu.com',
    updated_at: '2022-11-23',
    alias: 'ckdl5ndpwa'
  },
  {
    id: 2,
    user_id: 3,
    name: 'My Favourite Song',
    original_url: 'https://xdddd.com',
    updated_at: '2025-11-23',
    alias: 'dpa05987cm'
  },
  {
    id: 3,
    user_id: 3,
    name: 'Instagram',
    original_url: 'https://niewiem.com',
    updated_at: '2022-11-23',
    alias: 'smak41k2k7'
  },
  {
    id: 4,
    user_id: 3,
    name: 'My Youtube',
    original_url: 'https://fifciuu.com',
    updated_at: '2022-11-23',
    alias: 'ckdl5ndpwa'
  },
  {
    id: 5,
    user_id: 3,
    name: 'My Favourite Song',
    original_url: 'https://xdddd.com',
    updated_at: '2025-11-23',
    alias: 'dpa05987cm'
  },
  {
    id: 6,
    user_id: 3,
    name: 'Instagram',
    original_url: 'https://niewiem.com',
    updated_at: '2022-11-23',
    alias: 'smak41k2k7'
  },
  {
    id: 7,
    user_id: 3,
    name: 'My Youtube',
    original_url: 'https://fifciuu.com',
    updated_at: '2022-11-23',
    alias: 'ckdl5ndpwa'
  },
  {
    id: 8,
    user_id: 3,
    name: 'My Favourite Song',
    original_url: 'https://xdddd.com',
    updated_at: '2025-11-23',
    alias: 'dpa05987cm'
  },
  {
    id: 9,
    user_id: 3,
    name: 'Instagram',
    original_url: 'https://niewiem.com',
    updated_at: '2022-11-23',
    alias: 'smak41k2k7'
  }
]
</script>

<template>
  <div class="datatable-desktop">
    <div class="actions">
      <ShInput uniqueId="filterLinksByTextDesktop" type="search" label="Your Links" placeholder="Search"
        v-model="searchValue" />
      <button>SortBy</button>
      <ShButton variant="primary">New link</ShButton>
    </div>
    <table class="content">
      <tr class="headers">
        <th>Name</th>
        <th>Updated at</th>
        <th>Short link</th>
        <th></th>
      </tr>
      <tr class="records" v-for="record in records" :key="record.id">
        <td class="name">{{ record.name }}</td>
        <td class="date">{{ record.updated_at }}</td>
        <td class="link">https://short.link/{{ record.alias }}</td>
        <td class="actions">
          <button><img src="@/assets/copy.svg" /></button>
          <button><img src="@/assets/edit.svg" /></button>
          <button><img src="@/assets/bin.svg" /></button>
          <button><img src="@/assets/wishlist-empty.svg" /></button>
        </td>
      </tr>
    </table>
  </div>
  <div class="datatable-mobile">
    <div class="actions px-2">
      <ShButton variant="primary" class="w-100">New link</ShButton>
      <button class="sort-by--mobile">SortBy</button>
    </div>
    <div class="content px-2">
      <MobileDashboardLink 
        v-for="record in records" :key="record.id"
        :name="record.name"
        :short_link="record.alias"
        :updated_at="record.updated_at"
      />
    </div>
  </div>
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
          text-align: center;
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
}</style>