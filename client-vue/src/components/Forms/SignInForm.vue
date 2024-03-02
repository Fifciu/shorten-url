<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import BaseButton from '@/components/Base/BaseButton.vue';
import BaseInput from '@/components/Base/BaseInput.vue';
import BasePassword from '@/components/Base/BasePassword.vue';
import BaseAlternativeLink from '@/components/Base/BaseAlternativeLink.vue';
import BaseSingleSignOn from '@/components/Base/BaseSingleSignOn.vue';
import { useUserStore } from '@/stores/user';
import { useForm } from 'vee-validate';
import { ERROR_MESSAGES, ERR_WRONG_EMAIL_OR_PASSWORD } from '@/utils';
import { z } from 'zod';
import { toTypedSchema } from '@vee-validate/zod';

const ERR_MSG_EMAIL_TOO_SHORT = 'Email must be at least 1 character long';
const ERR_MSG_PASSWORD_TOO_SHORT = 'Password must be at least 1 characters long';
const schema = toTypedSchema(
  z.object({
    email: z.string().min(1, ERR_MSG_EMAIL_TOO_SHORT).email(),
    password: z.string().min(1, ERR_MSG_PASSWORD_TOO_SHORT)
  })
);

const { errors, meta, defineField, handleSubmit, setFieldError } = useForm({
  validationSchema: schema,
});

const [email, emailAttrs] = defineField('email');
const [password, passwordAttrs] = defineField('password');
const receivedServerError = ref(false);

const router = useRouter();
const { login } = useUserStore();

const onSubmit = handleSubmit(async formData => {
  const result = await login(formData.email, formData.password);
  if (result === true) {
    return router.push('/dashboard');
  }

  if (result.error === ERR_WRONG_EMAIL_OR_PASSWORD) {
    return setFieldError('email', ERROR_MESSAGES[ERR_WRONG_EMAIL_OR_PASSWORD]);
  }
  receivedServerError.value = true;
});

const readableErrors = computed(() => {
  const errMessages = Object.values(errors.value);
  if (!errMessages.length) {
    return
  }
  return errMessages.join('. ') + '.';
});
</script>

<template>
  <div class="px-2 md:max-w-[700px] md:mx-auto lg:w-[434px] lg:px-0 lg:mb-5">
    <div class="text-center my-3 relative">
      <img src="@/assets/ludzik-maly.svg" />
      <BaseTooltipBox class="mx-2 absolute top-0 w-[calc(100%-32px)]" v-show="readableErrors">
        {{ readableErrors }}
      </BaseTooltipBox>
    </div>
    <form @submit.prevent="onSubmit">
      <BaseInput uniqueId="sign-up-email" type="email" label="Email" placeholder="example@example.com" validator="email"
        class="mb-2" required v-model="email" v-bind="emailAttrs" :isError="Boolean(errors.email)" />
      <BasePassword uniqueId="sign-up-password" label="Password" placeholder="Password" class="mb-4" minlength="8" required
        v-model="password" v-bind="passwordAttrs" :isError="Boolean(errors.password)" />
      <BaseButton variant="primary" class="w-full mb-2" shadow :disabled="!meta.valid">Sign in</BaseButton>
      <BaseAlternativeLink class="mb-2" question="You don't have an account yet?" link="/sign-up" linkedText="Sign up" />
    </form>
    <BaseSingleSignOn class="my-4"/>
  </div>
</template>

<style lang="scss" scoped>
</style>