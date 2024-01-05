import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', () => {
  const sessionToken = ref('');
  function setSessionToken (token: string) {
    sessionToken.value = token;
  }

  const parsedToken = computed(() => {
    if (sessionToken.value) {
      return parseJwt(sessionToken.value);
    }
    return undefined
  });

  const user = computed(() => {
    if (!parsedToken.value) {
      return undefined
    }
    const { exp, iat, ...user } = parsedToken.value
    return user;
  });

  const isAuth = computed(() => Boolean(user.value?.id));

  return {
    setSessionToken,
    user,
    isAuth
  }
})

function parseJwt (token: string): Record<string, any> {
  const base64Url = token.split('.')[1];
  const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
  const jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
      return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
  }).join(''));

  return JSON.parse(jsonPayload);
}