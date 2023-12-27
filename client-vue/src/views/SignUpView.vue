<script setup lang="ts">
import LandingHeader from '@/components/Landing/LandingHeader.vue';
import LandingCTA from '@/components/Landing/LandingCTA.vue';
import SignUpForm from '@/components/SignUpForm.vue';
import ShButton from '@/components/ShButton.vue';
</script>

<template>
  <main class="layout">
    <LandingHeader class="header" />
    <div class="contet">
      <!-- <LandingCTA 
        title="Shorten, Share, Control"
        description="Maximum Functionality in Short Links - See How Easy You Can Shorten URLs!"
      /> -->
      <SignUpForm />
    </div>
  </main>
</template>

<style lang="scss" scoped>

.header {
  position: relative;
  z-index: 2;
}

.cta {
  display: none;
}

@media only screen and (min-width: $widthDesktop) {
  .cta {
    display: inline-block;
    padding-left: 138px;
    margin-top: 48px;
  }
}
</style>