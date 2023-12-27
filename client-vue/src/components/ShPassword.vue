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
  modelValue: String
});

const isPasswordVisible = ref(false);
</script>

<template>
  <div class="input-wrapper">
    <label v-if="props.label" class="input-label" :for="props.uniqueId">{{ label }}</label>
    <div class="input-subwrapper">
      <input :type="isPasswordVisible ? 'text' : 'password'" class="input" :id="props.uniqueId" :placeholder="props.placeholder" :model-value="props.value"
      @input="$emit('update:modelValue', $event.target.value)" v-bind="$attrs" />
      <button type="button" class="btn-toggle-password" @click="isPasswordVisible = !isPasswordVisible"><img src="@/assets/hide-password.svg" /></button>
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
</style>