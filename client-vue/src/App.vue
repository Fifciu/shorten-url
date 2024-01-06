<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { shallowRef, provide, reactive } from 'vue';
import router from './router';
import * as layouts from '@/layouts';
import Cookie from 'js-cookie';
import { useUserStore } from '@/stores/user';

const formData = reactive({
  email: '',
  password: ''
});

const userStore = useUserStore();
const layout = shallowRef('div');

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
  <component :is="layout || 'div'">
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
</style>
