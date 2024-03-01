<script setup>
import { defineProps, useSlots } from 'vue';

const props = defineProps({
  title: String,
  description: String
});

const slots = useSlots();
</script>

<template>
  <div class="cta">
    <h2 class="cta__title">{{ props.title }}</h2>
    <div class="cta__description" :class="[slots.default && 'cta__description-before-btn']" v-html="props.description"></div>
    <div class="cta__btn-wrapper">
      <slot></slot>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.cta {
  &__title {
    color: $colorBlack;
    text-align: center;
    font-family: Merriweather;
    font-size: 32px;
    font-style: normal;
    font-weight: 400;
    line-height: 150%;
    margin-bottom: 16px;
  }

  &__description {
    color: $colorDarkGrey;
    text-align: center;
    font-family: Source Sans Pro;
    font-size: 16px;
    font-style: normal;
    font-weight: 400;
    line-height: 150%;
    margin: 0 auto 0px;

    &-before-btn {
      margin: 0 auto 40px;
    }
  }

  &__btn-wrapper {
    text-align: center;
  }
}

@media only screen and (min-width: $widthDesktop) {
.cta {
  display: inline-block;

  &__title {
    font-size: 48px;
  }

  &__description {
    font-size: 20px;
  }
}
}
</style>