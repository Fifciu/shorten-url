<script setup>
import { ref } from 'vue';
import BaseRadio from '@/components/Base/BaseRadio.vue';
import { useUiStore } from '@/stores/ui';

const uiStore = useUiStore();
const isOpen = ref(false);
</script>

<template>
  <div class="sort-by mt-2" :class="{ opened: isOpen }">
    <div class="opener py-1" @click="isOpen = !isOpen">
      <span>Sort By</span>
      <img src="@/assets/dropdown.svg" />
    </div>
    <div class="py-2" v-if="isOpen">
      <BaseRadio uniqueId="mobile-dashboard-sort" :fields="uiStore.SORT_OPTIONS" v-model="uiStore.sortBy" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.opener {
  cursor: pointer;
  display: flex;
  justify-content: space-between;
  color: $colorDarkGrey;
  font-family: Source Sans Pro;
  font-size: 12px;
  font-style: normal;
  font-weight: 600;
  line-height: 150%;
  /* 18px */
  text-transform: uppercase;
  border-bottom: none;
}

.sort-by {
  &.opened {
    .opener {
      border-bottom: 2px solid $colorWhiteGrey;

      img {
        transform: rotate(180deg);
      }
    }
  }
}
</style>