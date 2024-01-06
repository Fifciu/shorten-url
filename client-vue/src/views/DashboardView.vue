<script setup lang="ts">
import DashboardCTA from '@/components/Dashboard/DashboardCTA.vue';
import DashboardDatatable from '@/components/Dashboard/DashboardDatatable.vue';
import { ref } from 'vue';

import { useUserStore } from '@/stores/user';

const { user } = useUserStore();
const searchValue = ref('');
</script>

<template>
  <main class="layout">
    <DashboardCTA 
      :title="`Welcome ${user!.fullname}!`"
      description="Shorten to your heart's content! Unlimited. "
      class="mb-10"
    />
    <DashboardDatatable />
  </main>
</template>

<style lang="scss" scoped>
.desktop-man-img, .desktop-zygzag-img {
  display: none;
}

.cta {
  display: none;
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
    display: block;
    margin-top: 44px;
  }
}
</style>