<script setup lang="ts">
import LandingHeader from '@/components/Landing/LandingHeader.vue';
import LandingFooter from '@/components/Landing/LandingFooter.vue';
import { useRoute } from 'vue-router';
import { computed } from 'vue';

const route = useRoute();
const perRouteClasses = computed(() => {
  if (route.name === 'sign-up') {
    return 'sign-up-background'
  } else if (route.name === 'sign-in') {
    return 'sign-in-background'
  }
  return ''
})

</script>

<template>
  <div class="min-h-screen flex flex-col" :class="perRouteClasses">
    <main class="flex-1 mt-5 lg:mt-4">
      <LandingHeader class="header" />
      <slot></slot>
    </main>
    <LandingFooter />
  </div>
</template>

<style lang="scss" scoped>
.header {
  position: relative;
  z-index: 2;
}

@media only screen and (min-width: $widthDesktop) {
  .sign-up-background {
    background: url('@/assets/sign-up-man.png') right no-repeat, url('@/assets/zygzag-2.svg') bottom 60% right 20% no-repeat;
  }

  .sign-in-background {
    background: url('@/assets/sign-in-man.png') left no-repeat;
  }
}
</style>