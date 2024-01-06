<script setup>
import { defineProps, defineOptions } from 'vue';

defineOptions({
  inheritAttrs: false
});

const props = defineProps({
  uniqueId: String,
  label: {
    type: String,
    required: false
  },
  placeholder: {
    type: String,
    required: false
  },
  hasBlackLabel: Boolean,
  modelValue: String
});
</script>

<template>
  <div class="input-wrapper">
    <label v-if="props.label" :class="['input-label', props.hasBlackLabel && 'input-label--black']" :for="props.uniqueId">{{ label }}</label>
    <input class="input" :id="props.uniqueId" :placeholder="props.placeholder" :model-value="props.value" @input="$emit('update:modelValue', $event.target.value)" v-bind="$attrs" />
  </div>
</template>

<style lang="scss" scoped>
  .input {
    font-family: Inter;
    font-size: 14px;
    font-style: normal;
    font-weight: 400;
    width: 100%;
    padding: 16px;
    height: 48px;
    border-radius: 8px;
    border: 2px solid $colorLightGrey;

    &::placeholder {
      color: $colorLightGrey;
    }
    &:disabled{
      background: $colorWhiteGrey;
      border: 1px solid $colorLightGrey;
    }

    &-label {
      color: $colorBlue;
      font-size: 14px;
      font-weight: 500;
      margin-bottom: 4px;

      &--black {
        color: $colorBlack;
      }
    }

    &-wrapper {
      display: flex;
      flex-wrap: wrap;
    }
  }
</style>