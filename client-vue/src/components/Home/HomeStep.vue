<script setup lang="ts">
const props = defineProps<{
  title: string;
  description: string;
  special: boolean;
}>();
</script>

<template>
  <div class="relative">
    <div class="w-[40px] h-[40px] bg-blue-50 rounded-full border-blue-500 border-8 border-solid absolute left-[-67px]" :class="[props.special && 'bg-blue-500']"></div>
    <h3 class="text-2xl text-gray-750 mb-2 font-medium" :class="[props.special && 'text-blue-500 font-bold']">{{ props.title }}</h3>
    <p class="text-base text-gray-500 font-normal">{{ props.description }}</p>
  </div>
</template>