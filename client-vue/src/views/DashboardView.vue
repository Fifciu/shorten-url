<script setup lang="ts">
import DashboardCTA from '@/components/Dashboard/DashboardCTA.vue';
import DashboardDatatable from '@/components/Dashboard/DashboardDatatable.vue';
import DashboardNoLinksYet from '@/components/Dashboard/DashboardNoLinksYet.vue';
import BaseModal from '@/components/Base/BaseModal.vue';
import NewLinkForm from '@/components/Forms/NewLinkForm.vue';
import EditLinkForm from '@/components/Forms/EditLinkForm.vue';
import { ref, onBeforeMount } from 'vue';

import { useUserStore } from '@/stores/user';
import { useLinksStore } from '@/stores/links';

const searchValue = ref('');
const isAddLinkOpen = ref(false);
const editLinkData = ref(null);
const { user } = useUserStore();
const { links, fetchMyLinks } = useLinksStore();
onBeforeMount(() => {
  fetchMyLinks();
});
</script>

<template>
  <main class="layout">
    <DashboardCTA :title="`Welcome ${user!.fullname}!`" description="Shorten to your heart's content! Unlimited. "
      class="mb-10" />
    <DashboardDatatable v-if="links?.length" @openAddLinkModal="isAddLinkOpen = true" @openEditLinkModal="editLinkData = $event" />
    <DashboardNoLinksYet v-else @openAddLinkModal="isAddLinkOpen = true" />
    <Teleport to="body">
      <BaseModal v-if="isAddLinkOpen" @close="isAddLinkOpen = false">
        <NewLinkForm @close="isAddLinkOpen = false" />
      </BaseModal>
      <BaseModal v-if="editLinkData" @close="editLinkData = null">
        <EditLinkForm :link="editLinkData" @close="editLinkData = null" />
      </BaseModal>
    </Teleport>
  </main>
</template>

<style lang="scss" scoped>
.desktop-man-img,
.desktop-zygzag-img {
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