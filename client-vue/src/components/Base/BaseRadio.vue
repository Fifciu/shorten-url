<script setup>
import { computed } from 'vue';

const props = defineProps({
  uniqueId: String,
  modelValue: String,
  fields: {
    type: Array,
    validator(fields) {
      return fields.every(field => field?.label && field?.value)
    }
  }
});

const emit = defineEmits(['update:modelValue']);
const model = defineModel();
const uniqueIds = computed(() => {
  return props.fields.map(({ value }) => `radio_unique_${props.uniqueId}_${value}`)
});
</script>

<template>
  <label 
    class="pl-1 radio-label"
    v-for="(field, index) in props.fields" :key="field.value" :for="uniqueIds[index]"
    :class="{ 'pb-2': index !== props.fields.length - 1 }"
    >
    <input class="radio-field" type="radio" :id="uniqueIds[index]" :value="field.value" v-model="model" /> {{ field.label }}
  </label>
</template>

<style lang="scss">
  .radio-label {
    display: flex;
    align-items: center;
    gap: 8px;
    color: $colorDarkGrey;
    font-family: Source Sans Pro;
    font-size: 14px;
    font-style: normal;
    font-weight: 600;
    line-height: 150%; /* 21px */
  }

  .radio-field {
    color: $colorLightGrey;
    width: 18px;
    height: 18px;
    border: 0.15em solid $colorLightGrey;
    border-radius: 50%;

    display: grid;
    place-content: center;

    &::before {
      content: '';
      width: 8px;
      height: 8px;
      border-radius: 50%;
      transform: scale(0);
      transition: 120ms transform ease-in-out;
      box-shadow: inset 1em 1em $colorOrange;
    }

    &:checked {
      border: 0.15em solid $colorOrange;
    }

    &:checked::before{
      transform: scale(1);
    }
  }
</style>