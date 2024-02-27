<script setup lang="ts">
import LandingHeader from '@/components/Landing/LandingHeader.vue';
import LandingCTA from '@/components/Landing/LandingCTA.vue';
import BaseButton from '@/components/Base/BaseButton.vue';
import HomeStep from '@/components/Home/HomeStep.vue';

const steps = [
  {
    title: '1. Paste Your Link',
    description: 'Begin by pasting your long URL into the text box. Our system is ready to handle any URL, no matter how lengthy.'
  },
  {
    title: '2. Choose Your Shortened Version',
    description: 'Decide whether to use an automatically generated short link or create your own customized one.'
  },
  {
    title: "3. Click 'Shorten'",
    description: 'After selecting your preferred shortcut, click the ‘Shorten’ button and wait for your link to be magically transformed.'
  },
  {
    title: '4. Copy and Share',
    description: 'Once your link is shortened, simply copy the new, concise URL and share it anywhere you want - through social media, email, or text messages. It’s that easy!',
    special: true
  },
]
</script>

<template>
  <LandingCTA class="cta lg:min-h-screen" title="Simplify Your Links!"
    :description="`Transform lengthy URLs into short,<br> shareable ones with a single click.`">
    <BaseButton variant="primary" roundy>Discover</BaseButton>
  </LandingCTA>
  <img src="@/assets/zygzag-2.svg" class="desktop-zygzag-img" />
  <img src="@/assets/home-man.png" class="desktop-man-img" />
  <img src="@/assets/ludzik-mobile.png" class="w-full mt-3 lg:hidden" />
  <section class="bg-blue-50 bg-[url('@/assets/home-mobile-path.png')] bg-no-repeat bg-[right_top_40px] bg-[length:87%] pb-7 pt-29 pr-2 -mt-17 flex justify-end">
    <div class="flex flex-col gap-y-7 w-3/4">
      <HomeStep v-for="step in steps" :key="step.title" v-bind="step" />
    </div>
  </section>
  <section class="px-2 py-6">
    <h2 class="text-center text-2xl mb-2 font-medium uppercase">Shortening and editing links has never been easier</h2>
    <p class="text-center font-medium mb-5 text-gray-400 text-xl">- discover the possibilities with your own account today!</p>
    <BaseButton variant="primary" to="/sign-up" roundy class="inline">Start for Free</BaseButton>
  </section>
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