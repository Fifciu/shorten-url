<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { shallowRef, provide, reactive, computed } from 'vue';
import { storeToRefs } from 'pinia';
import router from './router';
import * as layouts from '@/layouts';
import Cookie from 'js-cookie';
import { useUserStore } from '@/stores/user';
import { useUiStore } from '@/stores/ui';

const formData = reactive({
  email: '',
  password: ''
});

//TODO: Token expired handling

const userStore = useUserStore();
const uiStore = useUiStore();
const { isBodyScrollLocked } = storeToRefs(uiStore);
const layout = shallowRef('div');
const wrapperClasses = computed(() => {
  return isBodyScrollLocked.value 
    ? 'h-screen overflow-hidden'
    : ''
})

router.afterEach((to) => {
  layout.value = layouts[to.meta.layout] || 'div';
});

router.beforeEach((to) => {
  // TODO: Improve check
  // TODO: Move away this check
  const sessionToken = Cookie.get('session-token');
  userStore.setSessionToken(sessionToken);
  if (to.meta?.isAuth && !sessionToken) {
    return { name: 'sign-in' }
  } else if (sessionToken && ['sign-in', 'sign-up'].includes(to.name)) {
    return { name: 'dashboard'}
  }
})

provide('app:layout', layout);
</script>

<template>
  <component :is="layout || 'div'" :class="wrapperClasses">
    <RouterView></RouterView>
  </component>
</template>

<style lang="scss">
*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  outline: none;
  list-style: none;
  font-weight: normal;
}

body {
  // min-height: 100vh;
  transition:
    color 0.5s,
    background-color 0.5s;
  font-family:
    Inter,
    -apple-system,
    BlinkMacSystemFont,
    'Segoe UI',
    Roboto,
    Oxygen,
    Ubuntu,
    Cantarell,
    'Fira Sans',
    'Droid Sans',
    'Helvetica Neue',
    sans-serif;
  text-rendering: optimizeLegibility;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

a {
  text-decoration: none;
}

button {
  border: 0;
  background: transparent;
  -webkit-font-smoothing: inherit;
  -moz-osx-font-smoothing: inherit;
  -webkit-appearance: none;
}

input[type="radio"] {
  /* Add if not using autoprefixer */
  -webkit-appearance: none;
  appearance: none;
  /* For iOS < 15 to remove gradient background */
  background-color: #fff;
}

</style>
