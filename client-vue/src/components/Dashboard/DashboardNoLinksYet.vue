<script setup lang="ts">
import BaseButton from '@/components/Base/BaseButton.vue';
import { defineEmits } from 'vue';

const emit = defineEmits<{
  (e: 'openAddLinkModal'): void
}>();
</script>

<template>
  <div class="box">
    <div class="wrapper">
      <div class="heading">
        <img class="mobile-img" src="@/assets/no-links-mobile.svg" />
        <h2 class="title">You haven't shorten any link yet</h2>
        <p class="description">Shorten to your heart's content! Unlimited.</p>
      </div>
      <BaseButton @click="emit('openAddLinkModal')" class="btn-create" variant="primary" shadow>Shorten your first link
      </BaseButton>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.box {
  width: 100%;
  margin: 114px auto 0;
}

.wrapper {
  margin: 0px 43px 0px;
  height: calc(100dvh - 234px);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.mobile-img {
  width: 100%;
  max-width: 600px;
  margin-bottom: 24px;
}

.heading {
  text-align: center;
}

.title {
  color: $colorGrey;
  font-family: Inter;
  font-size: 20px;
  font-style: normal;
  font-weight: 500;
  line-height: 150%;
}

.description {
  margin-top: 8px;
  color: $colorLightGrey;
  font-family: Source Sans Pro;
  font-size: 16px;
  font-style: normal;
  font-weight: 400;
  line-height: 150%;
}

@media only screen and (min-width: $widthDesktop) {
  .mobile-img {
    display: none;
  }

  .box {
    width: 818px;
    height: 565px;
    margin: 0 auto;

    border-radius: 8px;
    border: 2px solid $colorWhiteGrey;
  }

  .wrapper {
    height: calc(100% - 60px);
    background: url('@/assets/no-links.svg');
    background-size: 100%;
    background-repeat: no-repeat;

    margin: 50px 70px 10px;
  }

  .heading {
    margin-top: 66px;
    text-align: center;
  }

  .btn-create {
    margin-bottom: 48px;
  }

  .title {
    font-size: 32px;
  }

  .description {
    margin-top: 8px;
    font-size: 20px;
  }
}
</style>