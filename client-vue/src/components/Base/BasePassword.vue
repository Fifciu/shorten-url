<script setup>
import { defineProps, defineOptions, ref } from 'vue';

defineOptions({
  inheritAttrs: false
});

const props = defineProps({
  uniqueId: String,
  label: {
    type: String,
    required: false
  },
  placeholder: {
    type: String,
    required: false
  },
  modelValue: String,
  isError: Boolean
});

const isPasswordVisible = ref(false);
</script>

<template>
  <div class="input-wrapper">
    <label v-if="props.label" class="input-label" :for="props.uniqueId">{{ label }}</label>
    <div class="input-subwrapper" :class="[props.isError && 'input-subwrapper--error']">
      <input :type="isPasswordVisible ? 'text' : 'password'" class="input" :id="props.uniqueId" :placeholder="props.placeholder" :value="props.modelValue"
      @input="$emit('update:modelValue', $event.target.value)" v-bind="$attrs" />
      <button type="button" class="btn-toggle-password" @click="isPasswordVisible = !isPasswordVisible"><img src="@/assets/show-password.svg" v-show="!isPasswordVisible" /><img src="@/assets/hide-password.svg" v-show="isPasswordVisible" /></button>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.input {
  font-family: Inter;
  font-size: 14px;
  font-style: normal;
  font-weight: 400;
  padding: 16px;
  height: 48px;
  border-radius: 8px;
  border: 2px solid $colorLightGrey;
  border-top-right-radius: 0;
  border-bottom-right-radius: 0;
  border-right-width: 0;
  width: 100%;

  &::placeholder {
    color: $colorLightGrey;
  }

  &-label {
    color: $colorBlue;
    font-size: 14px;
    font-weight: 500;
    margin-bottom: 4px;
  }

  &-wrapper {
    // display: flex;
    // flex-wrap: wrap;
    // width: 100%;
    // overflow: hidden;

  }

  &-subwrapper {
    width: 100%;
    display: flex;
  }
}

.btn-toggle-password {
  padding: 0 16px;
  height: 48px;
  border-radius: 8px;
  border: 2px solid $colorLightGrey;
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
  border-left-width: 0;
  cursor: pointer;
}

.input-subwrapper--error {
  .input, .btn-toggle-password {
    border-color: #EA2A2A;
  }
}
</style>