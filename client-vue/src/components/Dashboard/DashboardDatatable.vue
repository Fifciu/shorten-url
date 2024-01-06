<script setup lang="ts">
import DashboardCTA from '@/components/Dashboard/DashboardCTA.vue';
import ShInput from '@/components/ShInput.vue';
import ShButton from '@/components/ShButton.vue';
import MobileDashboardLink from '@/components/Dashboard/MobileDashboardLink.vue';
import BaseModal from '@/components/BaseModal.vue';
import { ref } from 'vue';

import { useUserStore } from '@/stores/user';

const { user } = useUserStore();
const searchValue = ref('');
const isAddLinkOpen = ref(true);

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
];


</script>

<template>
  <div class="datatable-desktop">
    <div class="actions">
      <ShInput uniqueId="filterLinksByTextDesktop" type="search" label="Your Links" placeholder="Search"
        v-model="searchValue" />
      <button>Sort By</button>
      <ShButton variant="primary" @click="isAddLinkOpen = !isAddLinkOpen">New link</ShButton>
      <Teleport to="body">
        <!-- It will be component -->
        <BaseModal v-if="isAddLinkOpen" @close="isAddLinkOpen = false">
          <h2 class="title">New link</h2>
          <hr class="title-divider" />
          <form>
            <div class="fields">
              <ShInput class="mb-2" uniqueId="addLinkNewUrl" type="text" label="Long URL"
                placeholder="example: http://very-long-link.com/example-example" hasBlackLabel />
              <div class="divided-fields mb-2">
                <ShInput uniqueId="domain" type="text" label="Domain" value="https://short.link" disabled hasBlackLabel />
                <span class="span">/</span>
                <ShInput uniqueId="addLinkBackHalf" type="text" label="Back-half (optional)"
                  placeholder="random string if empty, or provide alias" hasBlackLabel />
              </div>
              <ShInput class="mb-3" uniqueId="addLinkBackHalf" type="text" label="Name short link"
                placeholder="example: My Favourite Song" hasBlackLabel />
            </div>
            <div class="buttons">
              <ShButton variant="secondary mr-2" @click.prevent="isAddLinkOpen = false">Cancel</ShButton>
              <ShButton variant="primary">Create</ShButton>
            </div>
          </form>
        </BaseModal>
      </Teleport>
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
      <button class="sort-by--mobile">Sort By</button>
    </div>
    <div class="content px-2">
      <MobileDashboardLink v-for="record in records" :key="record.id" :name="record.name" :short_link="record.alias"
        :updated_at="record.updated_at" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.buttons {
  display: flex;
  justify-content: flex-end;
}

.title {
  color: $colorBlack;
  font-family: Inter;
  font-size: 20px;
  font-style: normal;
  font-weight: 700;
  line-height: 150%;
  /* 30px */

  margin-top: 12px;
  margin-bottom: 32px;
}

form {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.title-divider {
  margin-bottom: 32px;
  border: none;
  border-top: 2px solid $colorWhiteGrey;
}

.span {
  display: block;
  padding: 8px 16px;
  width: 38px;
  align-self: flex-end;
  margin-bottom: 6px;
}

.divided-fields {
  display: grid;
  grid-template-columns: 1fr 38px 1fr;
  align-items: center;
}

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
  .title {
    margin-bottom: 16px;
    margin-top: 0;
  }

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
