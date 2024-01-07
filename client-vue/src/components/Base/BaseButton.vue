<script setup>
import { defineProps } from 'vue';

const props = defineProps({
  variant: {
    type: String,
    validator: (val) => ['primary', 'secondary'].includes(val)
  },
  to: String,
  roundy: Boolean,
  shadow: Boolean,
  disabled: Boolean
});
</script>

<template>
  <router-link v-if="props.to" :to="props.to" :class="['btn', 'btn-link', `btn-${props.variant}`, props.roundy && 'btn-roundy', props.shadow && 'btn-shadow']">
    <slot></slot>
  </router-link>
  <button v-else :class="['btn', `btn-${props.variant}`, props.roundy && 'btn-roundy', props.shadow && 'btn-shadow']"
    :disabled="props.disabled">
    <slot></slot>
  </button>
</template>

<style lang="scss" scoped>
.btn {
  text-align: center;
  cursor: pointer;
  padding: 16px 32px;
  border-radius: 8px;
  font-family: $fontFamily;
  font-size: 16px;
  font-style: normal;
  font-weight: 700;
  line-height: 16px;

  &-link {
    display: block;
  }

  &-primary,
  &-primary>a {
    background-color: $colorBlue;
    color: $colorWhite;
  }

  &-primary>a {
    color: $colorWhite;
  }

  &-secondary {
    border: 2px solid $colorBlue;
    background-color: $colorWhite;
    color: $colorBlue;
  }

  &-secondary>a {
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

  &-shadow {
    box-shadow: 0px 3.62px 9.05px 0px rgba(0, 0, 0, 0.25);
  }
}
</style>