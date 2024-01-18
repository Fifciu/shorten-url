<script setup>
const links = [
  {
    to: '/',
    label: 'Privacy Policy'
  },
  {
    to: '/',
    label: 'Cookie Policy'
  },
  {
    to: '/',
    label: 'Contact'
  },
  {
    to: '/',
    label: 'About Us'
  }
]
</script>

<template>
  <footer class="p-2 flex flex-wrap md:flex-nowrap md:py-3 md:px-[138px] md:justify-between">
    <ul class="flex justify-between order-2 w-[100%] md:justify-normal md:gap-x-2 md:w-auto">
      <li v-for="link in links" :key="link.label">
        <router-link :to="link.to" class="link text-[12px]">{{ link.label }}</router-link>
      </li>
    </ul>
    <div class="flex justify-between order-1 w-[100%] mb-4 md:order-3 md:mb-0 md:w-[295px]">
        <div class="flex items-center gap-x-2 md:gap-x-1">
            <img src="@/assets/logo_white.svg" alt="Logo footer" class="h-[19px]">
            <span class="text-white text-[12px] font-normal">© 2023 Short Company</span>
        </div>
        <div class="flex items-center justify-between w-[82px]">
          <a href="https://x.com" target="_blank">
            <img src="@/assets/social_twitter.svg" alt="Twitter">
          </a>
          <a href="https://instagram.com" target="_blank">
            <img src="@/assets/social_instagram.svg" alt="Instagram">
          </a>
          <a href="https://facebook.com" target="_blank">
            <img src="@/assets/social_facebook.svg" alt="Facebook">
          </a>
        </div>
    </div>
  </footer>
</template>

<style lang="scss" scoped>
  footer {
    background-color: $colorBlue;
  }

  .link {
    color: $colorWhite;
  }
</style>