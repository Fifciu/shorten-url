<script setup lang="ts">
import LandingHeader from '@/components/Landing/LandingHeader.vue';
import LandingFooter from '@/components/Landing/LandingFooter.vue';
import { useRoute } from 'vue-router';
import { computed } from 'vue';

const route = useRoute();
const perRouteClasses = computed(() => {
  if (route.name === 'sign-up') {
    return 'sign-up-background'
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
.cta {
  margin-top: 100px;
  padding: 0 16px;
  overflow: hidden;
}

.header {
  position: relative;
  z-index: 2;
}

@media only screen and (min-width: $widthDesktop) {
  .sign-up-background {
    background: url('@/assets/sign-up-man.png') right no-repeat, url('@/assets/zygzag-2.svg') bottom 60% right 20% no-repeat;
  }

  .desktop-zygzag-img {
    display: block;
    position: absolute;
    z-index: -1;
    right: 400px;
    top: 0;
  }

  .desktop-man-img {
    display: block;
    position: absolute;
    z-index: -1;
    right: 0;
    top: 0;
  }

  .cta {
    display: inline-block;
    padding-left: 138px;
    margin-top: 175px;
  }
}
</style>