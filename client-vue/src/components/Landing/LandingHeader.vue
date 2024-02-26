<script setup lang="ts">
import BaseButton from '@/components/Base/BaseButton.vue';
import { ref } from 'vue';
import routerInstance from '@/router';

const showedMobileMenu = ref(false);
const closeMenu = () => {
  showedMobileMenu.value = false;
}
routerInstance.afterEach(closeMenu);
</script>

<template>
    <header class="header">
      <div class="header__left">
        <router-link to="/"><img src="@/assets/logo.svg" alt="Short logo" class="header__logo" /></router-link>
        <ul class="desktop-menu__list">
          <li class="desktop-menu__list-item">
            <router-link to="/" class="desktop-menu__list-item-link">Home</router-link>
          </li>
          <li class="desktop-menu__list-item">
            <router-link to="/about-us" class="desktop-menu__list-item-link">About us</router-link>
          </li>
          <li class="desktop-menu__list-item">
            <router-link to="/contact" class="desktop-menu__list-item-link">Contact</router-link>
          </li>
        </ul>
      </div>
      <div class="header__right">
        <button class="header__hamburger" @click="showedMobileMenu = true">
          <img src="@/assets/hamburger.svg" alt="Hamburger button" />
        </button>
        <ul class="desktop-menu__cta-list">
          <li class="desktop-menu__cta-list-item">
            <router-link to="/sign-up" class="desktop-menu__cta-list-item-link secondary">Sign up</router-link>
          </li>
          <li class="desktop-menu__cta-list-item">
            <router-link to="/sign-in" class="desktop-menu__cta-list-item-link primary">Sign in</router-link>
          </li>
        </ul>
      </div>
    </header>

    <div class="mobile-menu-box" v-show="showedMobileMenu">
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
        <BaseButton variant="secondary" class="mb-2" to="/sign-in">Sign in</BaseButton>
        <BaseButton variant="primary" to="/sign-up">Sign up</BaseButton>
      </div>
    </div>
</template>

<style lang="scss">
.header__logo,
.mobile-menu__logo {
  width: 89.98px;
}

.header {
  padding: 0px 16px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.desktop-menu__list,
.desktop-menu__cta-list {
  display: none;
}

.mobile-menu {
  &-box {
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

  &__heading {
    text-align: center;
    position: relative;
    width: 100%;
  }

  &__close-btn {
    position: absolute;
    right: 0; top: 0; bottom: 0;
    margin: auto;
    width: 21px;
    height: 21px;

    img {
      width: 21px;
    }
  }

  &__actions {
    width: 100%;

    button {
      width: 100%;
    }
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

        &:link {
          color: $colorBlue;
        }
      }
    }
  }
}

@media only screen and (min-width: $widthDesktop) {
  .header__hamburger {
    // todo: more mobile class name
    display: none;
  }

  .header {
    padding: 0px 138px;

    &__left {
      display: flex;
    }
  }

  .desktop-menu__list {
    display: flex;
    align-items: center;
    margin-left: 80px;

    &-item {
      margin-right: 32px;

      &:nth-last-child {
        margin-right: 0;
      }

      &-link {
        font-size: 16px;
        font-style: normal;
        font-weight: 400;
        line-height: normal;
        color: $colorBlue;

        &:link {
          color: $colorBlue;
        }

        &.router-link-active {
          position: relative;

          &:after {
            content: '';
            display: block;
            position: absolute;
            width: 32px;
            height: 3px;
            background-color: $colorOrange;
            transform: rotate(47.49deg);
            top: 0;
            bottom: 0;
            right: 0;
            left: 0;
            z-index: -1;
            border-radius: 3px;
            margin: auto;
          }
        }
      }
    }
  }

  .desktop-menu__cta-list {
    display: flex;

    &-item {
      margin-right: 16px;

      &-link {
        font-family: Inter;
        font-size: 16px;
        font-style: normal;
        font-weight: 700;
        line-height: normal;
        padding: 8px 24px;
        color: $colorBlue;

        &.primary {
          border-radius: 16px;
          border: 2px solid $colorBlue;
          box-shadow: 0px 4px 8px 0px rgba(0, 0, 0, 0.25);
        }
      }
    }
  }
}
</style>