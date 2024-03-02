import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { authenticationService } from '@/services/authentication';
import Cookie from 'js-cookie';
import { FetchError } from '@/utils';
import { ERR_WRONG_EMAIL_OR_PASSWORD, ERR_EMAIL_EXISTS } from '@/utils';

export const useUserStore = defineStore('user', () => {
  const sessionToken = ref('');
  function setSessionToken (token: string) {
    sessionToken.value = token;
  }

  async function login(email: string, password: string) {
    try {
      await authenticationService.login(email, password);
      const sessionToken = Cookie.get('session-token');
      setSessionToken(sessionToken);
      return true;
    } catch (err) {
      if ((err as FetchError).status === 401) {
        return {
          error: ERR_WRONG_EMAIL_OR_PASSWORD
        }
      } else {
        return {
          error: 'SERVER_ERROR'
        }
      }
    }
  }

  async function register(fullname: string, email: string, password: string) {
    try {
      await authenticationService.register(fullname, email, password);
      const sessionToken = Cookie.get('session-token');
      setSessionToken(sessionToken);
      return true;
    } catch (err) {
      if ((err as FetchError).status === 409) {
        return {
          error: ERR_EMAIL_EXISTS
        }
      } else {
        return {
          error: 'SERVER_ERROR'
        }
      }
    }
  }

  async function logout() {
    await authenticationService.logout();
    setSessionToken('');
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
    register,
    login,
    logout,
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