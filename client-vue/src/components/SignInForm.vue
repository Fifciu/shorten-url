<script setup>
import { reactive } from 'vue';
import ShButton from '@/components/ShButton.vue';
import ShInput from '@/components/ShInput.vue';
import ShPassword from '@/components/ShPassword.vue';
import ShAlternativeLink from '@/components/ShAlternativeLink.vue';

const formData = reactive({
  email: '',
  password: ''
});

const onSubmit = event => {
  // TODO: Validators
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