<script setup lang="ts">
import { useLinksStore } from '@/stores/links';

const props = defineProps<{
  linkId: number
}>();

const { deleteLink } = useLinksStore();

async function onClick() {
  // TODO: Add are you sure question
  await deleteLink(props.linkId);
}
</script>

<template>
  <button @click="onClick">
    <svg width="12" height="13" viewBox="0 0 12 13" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path
        d="M3.76294 3.29878V2.84147C3.76294 2.23504 3.99387 1.65345 4.40493 1.22464C4.816 0.795834 5.37353 0.554932 5.95486 0.554932C6.53619 0.554932 7.09372 0.795834 7.50478 1.22464C7.91585 1.65345 8.14678 2.23504 8.14678 2.84147V3.29878"
        stroke="#C7C4C4" stroke-linecap="round" stroke-linejoin="round" />
      <path d="M0.694336 3.29858H11.2156" stroke="#C7C4C4" stroke-linecap="round" stroke-linejoin="round" />
      <path
        d="M2.00928 3.29858H9.90019V11.5301C9.90019 11.7727 9.80778 12.0054 9.64338 12.1769C9.47899 12.3483 9.25594 12.4447 9.02342 12.4447H2.88605C2.65351 12.4447 2.4305 12.3483 2.26607 12.1769C2.10165 12.0054 2.00928 11.7727 2.00928 11.5301V3.29858Z"
        stroke="#C7C4C4" stroke-linecap="round" stroke-linejoin="round" />
      <path d="M4.63965 6.0437V9.70357" stroke="#C7C4C4" stroke-linecap="round" stroke-linejoin="round" />
      <path d="M7.27002 6.0437V9.70357" stroke="#C7C4C4" stroke-linecap="round" stroke-linejoin="round" />
    </svg>
  </button>
</template>

<style lang="scss" scoped>
button {
  cursor: pointer;

  &:hover {
    path {
      stroke: $colorBlue;
    }
  }
}

path {
  transition: $transitionTime;
  will-change: stroke;
}
</style>