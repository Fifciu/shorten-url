<script setup>
import { defineProps } from 'vue';

const props = defineProps({
  variant: {
    type: String,
    validator: (val) => ['primary', 'secondary'].includes(val)
  },
  to: String,
  roundy: Boolean,
  disabled: Boolean
});
</script>

<template>
  <button :class="`btn btn-${props.variant}${props.roundy ? ' btn-roundy' : ''}`" :disabled="props.disabled">
    <router-link v-if="props.to" :to="props.to">
      <slot></slot>
    </router-link>
    <slot v-else></slot>
  </button>
</template>

<style lang="scss" scoped>
.btn {
  padding: 16px 32px;
  border-radius: 8px;
  box-shadow: 0px 3.62px 9.05px 0px rgba(0, 0, 0, 0.25);
  font-family: $fontFamily;
  font-size: 16px;
  font-style: normal;
  font-weight: 700;
  line-height: 16px;

  &-primary, &-primary > a {
    background-color: $colorBlue;
    color: $colorWhite;
  }

  &-primary > a {
    color: $colorWhite;
  }

  &-secondary {
    border: 2px solid $colorBlue;
    background-color: $colorWhite;
    color: $colorBlue;
  }

  &-secondary > a {
    color: $colorBlue;
  }

  &:disabled {
    border: none;
    background-color: $colorGrey;
    color: $colorWhite;
    cursor: not-allowed;
  }

  &-roundy {
    border-radius: 24px;
  }
}
</style>