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
  onClick: {
    type: Function,
    required: false
  }
});

const isPasswordVisible = ref(false);
</script>

<template>
  <!-- TODO: POSSIBLE TO MERGE WITH BASE PASSWORD AND MODAL -->
  <!-- TODO: POSSIBLE TO PUT IN ON BOTH SIDE OF THE INPUT -->
  <div class="input-wrapper">
    <label v-if="props.label" class="input-label" :for="props.uniqueId">{{ label }}</label>
    <div class="input-subwrapper">
      <div class="search-icon-wrapper" v-if="!props.onClick">
        <img src="@/assets/search.svg" />
      </div>
      <button class="search-icon-wrapper" type="button" v-else>
        <img src="@/assets/search.svg" />
      </button>
      <input type="search" class="input" :id="props.uniqueId" :placeholder="props.placeholder" :value="props.modelValue"
      @input="$emit('update:modelValue', $event.target.value)" v-bind="$attrs" />
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
  padding-left: 0;
  height: 48px;
  border-radius: 8px;
  border: 2px solid $colorWhiteGrey2;
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
  border-left-width: 0;
  width: 100%;
  transition: all $transitionTime;

  &:focus {
    border: 2px solid $colorLightGrey;
    border-left-width: 0;
  }

  &::placeholder {
    color: $colorLightGrey;
  }

  &-label {
    color: $colorDarkGrey;
    font-size: 14px;
    font-weight: 500;
    font-family: Source Sans Pro;
    font-size: 14px;
    font-style: normal;
    font-weight: 600;
    line-height: 150%; /* 21px */
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
    margin-top: 4px;
  }
}

.search-icon-wrapper {
  padding: 0 16px;
  height: 48px;
  border-radius: 8px;
  border: 2px solid $colorWhiteGrey2;
  border-top-right-radius: 0;
  border-bottom-right-radius: 0;
  border-right-width: 0;
  cursor: pointer;
  display: flex;
  align-items: center;
  transition: all $transitionTime;

  img {
    opacity: .3;
    transition: all $transitionTime;
  }
}

.search-icon-wrapper:has(+ .input:focus) {
  border: 2px solid $colorLightGrey;
  border-right-width: 0;

  img {
    opacity: 1;
  }
}
</style>