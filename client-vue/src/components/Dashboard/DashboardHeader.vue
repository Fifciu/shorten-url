<script setup lang="ts">
import BaseButton from '@/components/Base/BaseButton.vue';
import BaseSearchMobile from '@/components/Base/BaseSearchMobile.vue';
import { ref } from 'vue';
import { useLinksStore } from '@/stores/links';
import { useUserStore } from '@/stores/user';
import { useRouter } from 'vue-router';
import routerInstance from '@/router';

const showedMobileMenu = ref(false);
const showedDesktopSubmenu = ref(false);
const showedMobileSearchBar = ref(false);
const linkStore = useLinksStore();
const { user, logout } = useUserStore();
const router = useRouter();

const closeMenu = () => {
  showedMobileMenu.value = false;
  showedDesktopSubmenu.value = false;
}
routerInstance.afterEach(closeMenu);

const onClickLogout = async () => {
  await logout();
  router.push('/');
}
</script>

<template>
  <header class="header">
    <div class="header__left">
      <router-link to="/dashboard"><img src="@/assets/logo.svg" alt="Short logo" class="header__logo" /></router-link>
    </div>
    <div class="header__right">
      <button class="header__search pointer mr-2" v-show="!showedMobileSearchBar"
        @click="showedMobileSearchBar = !showedMobileSearchBar">
        <img src="@/assets/search.svg" alt="Search button" />
      </button>
      <BaseSearchMobile wrapperClasses="header__search-input mr-2" v-model="linkStore.searchQuery"
        @click="showedMobileSearchBar = !showedMobileSearchBar" v-show="showedMobileSearchBar" />
      <button class="header__hamburger" @click="showedMobileMenu = true">
        <img src="@/assets/hamburger.svg" alt="Hamburger button" />
      </button>
      <div class="submenu-wrapper">
        <ul class="desktop-menu__cta-list">
          <li class="desktop-menu__cta-list-item">
            {{ user?.fullname }}
          </li>
          <li class="desktop-menu__cta-list-item">
            <!-- TODO: Make it button? -->
            <span to="/account" class="desktop-menu__cta-list-item-link pointer"><img src="@/assets/settings.svg"
                @click="showedDesktopSubmenu = !showedDesktopSubmenu" />
            </span>
          </li>
        </ul>
        <ul class="submenu" v-show="showedDesktopSubmenu">
          <li class="submenu-item">
            <router-link class="submenu-text" to="/account">
              <img src="@/assets/account.svg" />
              Account
            </router-link>
          </li>
          <li class="submenu-item">
            <button class="submenu-text" type="button" @click="onClickLogout">
              <img src="@/assets/logout.svg" />
              Log out
            </button>
          </li>
        </ul>
      </div>
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
      <BaseButton variant="secondary" class="mb-2" @click="onClickLogout">Log out</BaseButton>
    </div>
  </div>
</template>

<style lang="scss">
.pointer {
  cursor: pointer;
}

.submenu {
  display: none;
}

.layout {
  margin-top: 44px;
}

.header__logo,
.mobile-menu__logo {
  width: 89.98px;
}

.header {
  padding: 0px 16px;
  height: 38px;
  display: flex;
  align-items: center;
  justify-content: space-between;

  &__right {
    display: flex;
  }
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
    right: 0;
    top: 0;
    bottom: 0;
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

  .header__search-input {
    display: none;
  }

  .submenu-wrapper {
    position: relative;
  }

  .submenu {
    display: block;
    position: absolute;
    bottom: -13.5px;
    transform: translateY(100%);
    padding: 8px;
    display: flex;
    flex-direction: column;
    border-radius: 8px;
    border: 2px solid $colorWhiteGrey;
    background: $colorWhite;
  }

  .submenu-item {
    padding: 8px 16px;

    img {
      margin-right: 8px;
    }

    .submenu-text {
      display: inline-flex;
      align-items: center;
      cursor: pointer;
      line-height: 100%;
      color: $colorBlack;
      font-family: Inter;
      font-size: 14px;
      font-style: normal;
      font-weight: 400;
      line-height: 150%;
      /* 21px */
    }
  }

  // end submenu

  .header__hamburger,
  .header__search {
    // todo: more mobile class name
    display: none;
  }

  .layout {
    margin-top: 32px;
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
}</style>
