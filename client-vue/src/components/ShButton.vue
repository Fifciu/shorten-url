<script setup>
import { defineProps } from 'vue';

const props = defineProps({
  variant: {
    type: String,
    validator: (val) => ['primary', 'secondary'].includes(val)
  }
});
</script>

<template>
  <button :class="`btn btn-${props.variant}`">
    <slot />
  </button>
</template>

<style lang="scss" scoped>
.btn {
  padding: 16px 0;
  border-radius: 8px;
  box-shadow: 0px 3.62px 9.05px 0px rgba(0, 0, 0, 0.25);
  font-family: $fontFamily;
  font-size: 16px;
  font-style: normal;
  font-weight: 700;
  line-height: 16px;

  &-primary {
    background-color: $colorBlue;
    color: $colorWhite;
  }

  &-secondary {
    border: 2px solid $colorBlue;
    color: $colorBlue;
  }
}
</style>