<script setup lang="ts">
import BaseButton from '@/components/Base/BaseButton.vue';
import { defineEmits } from 'vue';

const emit = defineEmits<{
  (e: 'openAddLinkModal'): void
}>();
</script>

<template>
  <div class="flex flex-col gap-[34px] w-[274px] mx-auto md:w-[752px] md:h-[495px] md:bg-[url('@/assets/no-links-with-filters-desktop.svg')] md:bg-center md:bg-[length:80%] md:bg-no-repeat md:grid md:place-items-center">
    <img class="md:hidden" src="@/assets/no-links-with-filters-mobile.svg" />
    <h2 class="title text-center text-[20px] non-italic font-medium leading-[150%] md:text-[32px] md:w-[70%]">You don't have links matching specified criteria</h2>
  </div>
</template>

<style lang="scss" scoped>
.title {
  color: $colorGrey;
}
</style>