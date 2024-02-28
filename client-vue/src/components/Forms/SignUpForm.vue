<script setup lang="ts">
import { reactive, computed } from 'vue';
import { useRouter } from 'vue-router';
import BaseButton from '@/components/Base/BaseButton.vue';
import BaseInput from '@/components/Base/BaseInput.vue';
import BasePassword from '@/components/Base/BasePassword.vue';
import BasePasswordIndicator from '@/components/Base/BasePasswordIndicator.vue';
import BaseTooltipBox from '@/components/Base/BaseTooltipBox.vue';
import BaseRulesList from '@/components/Base/BaseRulesList.vue';
import type { ListElement } from '@/components/Base/BaseRulesList.vue';
import BaseAlternativeLink from '@/components/Base/BaseAlternativeLink.vue';
import { useUserStore } from '@/stores/user';

const formData = reactive({
  email: '',
  fullname: '',
  password: ''
});

// values should be always trimmed
const validation = computed(() => {
  return {
    email: formData.email.includes('@') || 'Provide proper email',
    fullname: formData.fullname.includes(' ') || 'Provide proper full name',
    password: {
      uppercase: /[A-Z]+/.test(formData.password),
      number: /[0-9]+/.test(formData.password) || 'At least 1 number',
      length: formData.password.length >= 8 || 'At least 8 characters'
    }
  }
})
const router = useRouter();
const { register } = useUserStore();

const onSubmit = async (event: any) => {
  try {
    await register(formData.fullname, formData.email, formData.password);
    router.push('/dashboard');
  } catch (err) {
    console.error(err);
  }
};

const rulesList = computed<ListElement[]>(() => {
  return [
    {
      description: 'At least 1 uppercase',
      status: validation.value.password.uppercase === true,
    },
    {
      description: 'At least 1 number',
      status: validation.value.password.number === true,
    },
    {
      description: 'At least 8 characters',
      status: validation.value.password.length === true,
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
</script>

<template>
  <div class="box">
    <!-- <h2 class="desktop-heading">
      One account,
      Unlimited Possibilities!
    </h2> -->
    <div class="relative mt-9 mb-6">
      <h2 class="text-blue-500 text-center text-2xl leading-9 font-ternary font-semibold">Create Account</h2>
      <!-- <BaseTooltipBox class="mx-2">
        Oops! Something went wrong!<br>
Make sure your login and password are correct.
      </BaseTooltipBox> -->
    </div>
    <form class="sign-up" @submit.prevent="onSubmit">
      <BaseInput uniqueId="sign-up-email" type="email" label="Email" placeholder="example@example.com" validator="email"
        class="mb-2" required v-model.trim="formData.email" />
      <BaseInput uniqueId="sign-up-fullname" type="text" label="Full name" placeholder="John Doe" class="mb-2" minlength="3"
        required v-model.trim="formData.fullname" />
      <BasePassword uniqueId="sign-up-password" label="Password" placeholder="Password" minlength="8" required
        v-model.trim="formData.password" />
      <BasePasswordIndicator :level="passwordLevel" class="mb-3 mt-1" />
      <BaseRulesList title="Weak password. It must contain:" :list="rulesList" class="mb-2"/>
      <BaseButton variant="primary" class="submit-btn mb-2" shadow>Sign up</BaseButton>
      <BaseAlternativeLink question="Already have an account?" link="/sign-in" linkedText="Sign in" class="mb-2" />
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