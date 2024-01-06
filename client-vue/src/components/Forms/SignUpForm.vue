<script setup>
import { reactive } from 'vue';
import BaseButton from '@/components/Base/BaseButton.vue';
import BaseInput from '@/components/Base/BaseInput.vue';
import BasePassword from '@/components/Base/BasePassword.vue';
import BaseAlternativeLink from '@/components/Base/BaseAlternativeLink.vue';
import { authenticationService } from '@/services/authentication';

const formData = reactive({
  email: '',
  fullname: '',
  password: ''
});

const onSubmit = async event => {
  try {
    const result = await authenticationService.register(formData.fullname, formData.email, formData.password);
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
      <BaseInput uniqueId="sign-up-fullname" type="text" label="Full name" placeholder="John Doe" class="mb-2" minlength="3"
        required v-model="formData.fullname" />
      <BasePassword uniqueId="sign-up-password" label="Password" placeholder="Password" class="mb-4" minlength="8" required
        v-model="formData.password" />
      <BaseButton variant="primary" class="submit-btn mb-2" shadow>Sign up</BaseButton>
      <BaseAlternativeLink question="Already have an account?" link="/sign-in" linkedText="Sign in" />
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