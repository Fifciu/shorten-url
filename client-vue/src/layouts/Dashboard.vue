<script setup lang="ts">
import DashboardHeader from '@/components/Dashboard/DashboardHeader.vue';
</script>

<template>
  <main class="mt-5 lg:mt-4">
    <DashboardHeader class="header" />
    <slot></slot>
  </main>
</template>

<style lang="scss" scoped>
.desktop-man-img, .desktop-zygzag-img {
  display: none;
}

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