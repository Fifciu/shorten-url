<script setup>
import { defineProps } from 'vue';

const props = defineProps({
  question: String,
  link: String,
  linkedText: String
});
</script>

<template>
  <p class="alt-link">
    {{ question }} <router-link class="link" :to="props.link">{{ props.linkedText }}</router-link>
  </p>
</template>

<style lang="scss" scoped>
  .alt-link {
    color: $colorBlue;
    text-shadow: 0px 3.62px 9.05px rgba(0, 0, 0, 0.25);
    font-size: 16px;
    font-weight: 500;
  }

  .link {
    font-weight: 700;
    text-decoration-line: underline;

    &:visited {
    color: $colorBlue;
    }
  }
</style>