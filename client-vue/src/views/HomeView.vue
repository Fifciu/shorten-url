<script setup lang="ts">
import LandingHeader from '@/components/Landing/LandingHeader.vue';
import LandingCTA from '@/components/Landing/LandingCTA.vue';
import BaseButton from '@/components/Base/BaseButton.vue';
</script>

<template>
  <LandingCTA class="cta" title="Simplify Your Links!"
    :description="`Transform lengthy URLs into short,<br> shareable ones with a single click.`">
    <BaseButton variant="primary" roundy>Discover</BaseButton>
  </LandingCTA>
  <img src="@/assets/zygzag-2.svg" class="desktop-zygzag-img" />
  <img src="@/assets/home-man.png" class="desktop-man-img" />
</template>

<style lang="scss" scoped>
.desktop-man-img,
.desktop-zygzag-img {
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