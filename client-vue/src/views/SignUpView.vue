<script setup lang="ts">
import LandingCTA from '@/components/Landing/LandingCTA.vue';
import SignUpForm from '@/components/Forms/SignUpForm.vue';
</script>

<template>
  <div>
    <LandingCTA title="Shorten, Share, Control"
      description="Maximum Functionality in Short Links - See How Easy You Can Shorten URLs!"
      class="hidden lg:inline lg:mt-6 lg:pl-0 xl:pl-17 lg:ml-17" />
    <img src="@/assets/mobile-path.svg" class="absolute -top-4 pointer-events-none lg:hidden" />
    <SignUpForm />
  </div>
</template>

<style lang="scss" scoped></style>