<script setup lang="ts">
import LandingCTA from '@/components/Landing/LandingCTA.vue';
import SignInForm from '@/components/Forms/SignInForm.vue';
</script>

<template>
  <div class="lg:ml-[40%] lg:flex lg:flex-col lg:items-center lg:w-fit">
    <LandingCTA 
        title="Shorten, Share, Control"
        description="Maximum Functionality in Short Links - See How Easy You Can Shorten URLs!"
        class="mt-5 px-2"
      />
    <SignInForm />
  </div>
</template>

<style lang="scss" scoped>
.header {
  position: relative;
  z-index: 2;
}
</style>