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
  <button @click="onClick"><img src="@/assets/bin.svg" /></button>
</template>

<style scoped>
button{
  cursor: pointer;
}</style>