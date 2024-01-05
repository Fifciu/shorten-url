<script setup lang="ts">
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import ShButton from '@/components/ShButton.vue';
import ShInput from '@/components/ShInput.vue';
import ShPassword from '@/components/ShPassword.vue';
import ShAlternativeLink from '@/components/ShAlternativeLink.vue';
import { authenticationService } from '@/services/authentication';
import { useUserStore } from '@/stores/user';
import Cookie from 'js-cookie';

const formData = reactive({
  email: '',
  password: ''
});

const router = useRouter();
const userStore = useUserStore();

const onSubmit = async event => {
  // TODO: Validators
  try {
    await authenticationService.login(formData.email, formData.password);
    const sessionToken = Cookie.get('session-token');
    userStore.setSessionToken(sessionToken);
    router.push('/dashboard');
  } catch (err) {
    console.error(err);
  }
};
</script>

<template>
  <div class="box">
    <!-- <h2 class="desktop-heading">
      One account,
      Unlimited Possibilities!
    </h2> -->
    <div class="img-wrapper">
      <img src="@/assets/ludzik-maly.svg" class="img-hi" />
    </div>
    <form class="sign-up" @submit.prevent="onSubmit">
      <ShInput uniqueId="sign-up-email" type="email" label="Email" placeholder="example@example.com" validator="email"
        class="mb-2" required v-model="formData.email" />
      <ShPassword uniqueId="sign-up-password" label="Password" placeholder="Password" class="mb-4" minlength="8" required
        v-model="formData.password" />
      <ShButton variant="primary" class="submit-btn mb-2">Sign in</ShButton>
      <ShAlternativeLink question="Already have an account?" link="/sign-up" linkedText="Sign up" />
    </form>
  </div>
</template>

<style lang="scss" scoped>
.img-wrapper {
  text-align: center;
}

.img-hi {
  margin: 80px auto 30px;
}

.sign-up {
  padding: 0 16px;
}

.submit-btn {
  width: 100%;
}
</style>