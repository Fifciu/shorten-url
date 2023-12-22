<script setup lang="ts">
import ShButton from '@/components/ShButton.vue';
import { ref } from 'vue';

const showedMobileMenu = ref(false);
</script>

<template>
  <main class="layout">
    <header class="header--mobile">
      <img src="@/assets/logo.svg" alt="Short logo" class="header__logo" />
      <button class="header__hamburger" @click="showedMobileMenu = true">
        <img src="@/assets/hamburger.svg" alt="Hamburger button" />
      </button>
    </header>

    <div class="header--mobile-menu-box" v-show="showedMobileMenu">
      <div class="mobile-menu__top-wrapper">
        <div class="mobile-menu__heading">
          <img src="@/assets/logo.svg" alt="Short logo" class="mobile-menu__logo" />
          <button class="mobile-menu__close-btn" @click="showedMobileMenu = false">
            <img src="@/assets/close.svg" alt="Close button" />
          </button>
        </div>

        <ul class="mobile-menu__list">
          <li class="mobile-menu__list-item">
            <router-link to="/" class="mobile-menu__list-item-link">Home</router-link>
          </li>
          <li class="mobile-menu__list-item">
            <router-link to="/about-us" class="mobile-menu__list-item-link">About us</router-link>
          </li>
          <li class="mobile-menu__list-item">
            <router-link to="/contact" class="mobile-menu__list-item-link">Contact</router-link>
          </li>
          <li class="mobile-menu__list-item">
            <router-link to="/sign-in" class="mobile-menu__list-item-link">Sign in</router-link>
          </li>
        </ul>
      </div>

      <div class="mobile-menu__actions">
        <ShButton variant="secondary" class="mb-2">Sign in</ShButton>
        <ShButton variant="primary">Sign up</ShButton>
      </div>
    </div>
  </main>
</template>

<style lang="scss">
.layout {
  margin-top: 44px;
}

.header {
  &--mobile {
    padding: 0px 16px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    &-menu-box {
      box-sizing: border-box;
      width: 100vw;
      height: 100dvh;
      background: white url('@/assets/zygzag-1.svg') no-repeat 0px 270px;
      background-size: cover;
      position: absolute;
      left: 0;
      top: 0;
      padding: 44px 16px 32px;

      display: flex;
      flex-direction: column;
      justify-content: space-between;
    }
  }

  &__logo {
    width: 89.98px;
  }
}

.mobile-menu {
  &__heading {
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;
    width: 100%;
  }

  &__logo {
    width: 89.98px;
  }

  &__close-btn {
    align-self: flex-start;
    position: absolute;
    right: 0;
    transform: translateY(-100%);
    width: 14px;
    height: 14px;
  }

  &__actions {
    width: 100%;
  }

  &__list {
    margin-top: 100px;
    &-item {
      margin-bottom: 40px;
      text-align: center;

      &:nth-last-child {
        margin-bottom: 0;
      }

      &-link {
        color: $colorBlue;
        font-family: $fontFamily;
        font-size: 20px;
        font-style: normal;
        font-weight: 500;
        line-height: normal;
      }
    }
  }
}

.btn {
  padding: 16px 0;
  line-height: 24px;
  width: 100%;

  &-primary {
    background: blue;
    color: #fff;
  }

  &-secondary {
    background: white;
    border: 1px solid blue;
  }
}

.mb-2 {
  margin-bottom: 16px;
}
</style>