<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps({
  level: {
    type: Number,
    validator (val: number) {
      return val >= 0 && val <= 3;
    }
  }
});

const DEFAULT_COLOR_CLASS = 'bg-gray-200';
const LEVEL_1_COLOR_CLASS = 'bg-red';
const LEVEL_2_COLOR_CLASS = 'bg-yellow';
const LEVEL_3_COLOR_CLASS = 'bg-green';
const colorClasses = computed(() => {
  return [
    props.level === 1 ? LEVEL_1_COLOR_CLASS : DEFAULT_COLOR_CLASS,
    props.level === 2 ? LEVEL_2_COLOR_CLASS : DEFAULT_COLOR_CLASS,
    props.level === 3 ? LEVEL_3_COLOR_CLASS : DEFAULT_COLOR_CLASS,
  ]
})
</script>

<template>
  <div class="flex gap-x-[4px]">
    <div class="h-[4px] flex-1 rounded-[2px]" :class="colorClass" v-for="(colorClass, index) in colorClasses" :key="`${colorClass}-${index}`"></div>
  </div>
</template>