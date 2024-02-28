<script setup lang="ts">
export interface ListElement {
  description: string,
  status: boolean
};

const props = defineProps({
  title: String,
  list: {
    type: Array,
    validator (value: ListElement[]) {
      return value.every(({ description, status }: ListElement) => {
        return typeof description === 'string' && typeof status === 'boolean'
      })
    }
  }
});
</script>

<template>
  <div class="flex flex-col gap-1 text-gray-400 font-normal text-xs font-ternary">
    <span>{{ props.title }}</span>
    <ul class="flex flex-col gap-1">
      <li v-for="element in (props.list as ListElement[])" :key="element.description" class="flex align-center gap-x-1">
        <img v-show="element.status" src="@/assets/list-plus.svg" />
        <img v-show="!element.status" src="@/assets/list-minus.svg" />
        <span>{{ element.description }}</span>
      </li>
    </ul>
  </div>
</template> 