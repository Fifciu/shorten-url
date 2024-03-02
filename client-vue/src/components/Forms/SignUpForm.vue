<script setup lang="ts">
import { reactive, ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import BaseButton from '@/components/Base/BaseButton.vue';
import BaseInput from '@/components/Base/BaseInput.vue';
import BasePassword from '@/components/Base/BasePassword.vue';
import BasePasswordIndicator from '@/components/Base/BasePasswordIndicator.vue';
import BaseTooltipBox from '@/components/Base/BaseTooltipBox.vue';
import BaseRulesList from '@/components/Base/BaseRulesList.vue';
import type { ListElement } from '@/components/Base/BaseRulesList.vue';
import BaseAlternativeLink from '@/components/Base/BaseAlternativeLink.vue';
import BaseSingleSignOn from '@/components/Base/BaseSingleSignOn.vue';
import { useUserStore } from '@/stores/user';
import { useForm } from 'vee-validate';
import { toTypedSchema } from '@vee-validate/zod';
import { z } from 'zod';
import { isUpperCase, ERROR_MESSAGES, ERR_EMAIL_EXISTS } from '@/utils';

const ERR_MSG_EMAIL_ALREADY_USED = ERROR_MESSAGES[ERR_EMAIL_EXISTS];
const ERR_MSG_EMAIL_TOO_SHORT = 'Email must be at least 1 character long';
const ERR_MSG_FULL_NAME_TOO_SHORT = 'Full name must be at least 2 characters long';
const ERR_MSG_PASSWORD_TOO_SHORT = 'Password must be at least 8 characters long';
const ERR_MSG_PASSWORD_MUST_CONTAIN_AT_LEAST_1_NUMBER = 'Password must contain at least 1 number';
const ERR_MSG_PASSWORD_MUST_CONTAIN_AT_LEAST_1_UPPERCASE_LETTER = 'Password must contain at least 1 uppercase letter';
const schema = toTypedSchema(
  z.object({
    email: z.string().min(1, ERR_MSG_EMAIL_TOO_SHORT).email().refine(val => !existingEmails.includes(val + ''), {
      message: ERR_MSG_EMAIL_ALREADY_USED
    }),
    fullname: z.string().min(2, ERR_MSG_FULL_NAME_TOO_SHORT),
    password: z.string().min(8, ERR_MSG_PASSWORD_TOO_SHORT)
      .refine(val => val.split('').filter(sign => !isNaN(sign as any)).length > 0, {
        message: ERR_MSG_PASSWORD_MUST_CONTAIN_AT_LEAST_1_NUMBER,
      })
      .refine(val => val.split('').filter(isUpperCase).length > 0, {
        message: ERR_MSG_PASSWORD_MUST_CONTAIN_AT_LEAST_1_UPPERCASE_LETTER,
      }),
  })
);

const { errors, meta, defineField, handleSubmit, setFieldError, errorBag } = useForm({
  validationSchema: schema,
});

const [email, emailAttrs] = defineField('email');
const [fullname, fullnameAttrs] = defineField('fullname');
const [password, passwordAttrs] = defineField('password');
const receivedServerError = ref(false);

const existingEmails = reactive<string[]>([]);
// If user still tries to use already existing email even though error message 
// informing it already exists has been showed
const tryingToUseExistingEmailAfterError = computed(() => existingEmails.includes(email.value + ''));

const router = useRouter();
const { register } = useUserStore();

const onSubmit = handleSubmit(async formData => {
  const result = await register(formData.fullname, formData.email, formData.password);
  if (result === true) {
    return router.push('/dashboard');
  }

  if (result.error === ERR_EMAIL_EXISTS) {
    existingEmails.push(formData.email);
    return setFieldError('email', ERR_MSG_EMAIL_ALREADY_USED);
  }
  receivedServerError.value = true;
})

const rulesList = computed<ListElement[]>(() => {
  return [
    {
      description: 'At least 1 uppercase',
      status: Boolean(password.value?.length && password.value?.length > 7 && !errorBag.value.password?.includes(ERR_MSG_PASSWORD_MUST_CONTAIN_AT_LEAST_1_UPPERCASE_LETTER))
    },
    {
      description: 'At least 1 number',
      status: Boolean(password.value?.length && password.value?.length > 7 && !errorBag.value.password?.includes(ERR_MSG_PASSWORD_MUST_CONTAIN_AT_LEAST_1_NUMBER))
    },
    {
      description: 'At least 8 characters',
      status: Boolean(password.value?.length && password.value?.length > 7 && !errorBag.value.password?.includes(ERR_MSG_PASSWORD_TOO_SHORT))
    }
  ]
});
const passwordLevel = computed(() => {
  return rulesList.value.reduce((total, curr) => {
    if (curr.status) {
      return total + 1
    }
    return total
  }, 0)
})

const readableErrors = computed(() => {
  const errMessages = Object.values(errors.value);
  if (!errMessages.length) {
    return
  }
  return errMessages.join('. ') + '.';
});
</script>

<template>
  <div class="px-2 md:max-w-[700px] md:mx-auto lg:max-w-[434px] lg:px-0 lg:ml-17 lg:mb-5">
    <div class="relative mt-9 mb-6 lg:mb-4">
      <h2 class="text-blue-500 text-center text-2xl leading-9 font-ternary font-semibold">Create Account</h2>
      <BaseTooltipBox class="mx-2 absolute top-0 w-[calc(100%-32px)]" v-show="readableErrors">
        {{ readableErrors }}
      </BaseTooltipBox>
    </div>
    <form @submit.prevent="onSubmit">
      <BaseInput uniqueId="sign-up-email" type="email" label="Email" placeholder="example@example.com" validator="email"
        class="mb-2" required v-model.trim="email" v-bind="emailAttrs" :isError="Boolean(errors.email) || tryingToUseExistingEmailAfterError"/>
      <BaseInput uniqueId="sign-up-fullname" type="text" label="Full name" placeholder="John Doe" class="mb-2" minlength="3"
        required v-model="fullname" v-bind="fullnameAttrs" :isError="Boolean(errors.fullname)"/>
      <BasePassword uniqueId="sign-up-password" label="Password" placeholder="Password" minlength="8" required
        v-model.trim="password" v-bind="passwordAttrs" :isError="Boolean(errors.password)"/>
      <BasePasswordIndicator :level="passwordLevel" class="mb-3 mt-1" />
      <BaseRulesList title="Weak password. It must contain:" :list="rulesList" class="mb-2"/>
      <BaseButton variant="primary" class="submit-btn mb-2" shadow :disabled="!meta.valid || tryingToUseExistingEmailAfterError">Sign up</BaseButton>
      <BaseAlternativeLink question="Already have an account?" link="/sign-in" linkedText="Sign in" class="mb-2" />
    </form>
    <BaseSingleSignOn class="my-4"/>
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