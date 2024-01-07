<script setup lang="ts">
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import BaseButton from '@/components/Base/BaseButton.vue';
import BaseInput from '@/components/Base/BaseInput.vue';
import BasePassword from '@/components/Base/BasePassword.vue';
import BaseAlternativeLink from '@/components/Base/BaseAlternativeLink.vue';
import { useUserStore } from '@/stores/user';

const formData = reactive({
  email: '',
  password: ''
});

const router = useRouter();
const { login } = useUserStore();

const onSubmit = async () => {
  // TODO: Validators
  try {
    await login(formData.email, formData.password);
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
      <BaseInput uniqueId="sign-up-email" type="email" label="Email" placeholder="example@example.com" validator="email"
        class="mb-2" required v-model="formData.email" />
      <BasePassword uniqueId="sign-up-password" label="Password" placeholder="Password" class="mb-4" minlength="8" required
        v-model="formData.password" />
      <BaseButton variant="primary" class="submit-btn mb-2" shadow>Sign in</BaseButton>
      <BaseAlternativeLink question="Already have an account?" link="/sign-up" linkedText="Sign up" />
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