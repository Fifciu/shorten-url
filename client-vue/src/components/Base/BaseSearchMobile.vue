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
  },
  wrapperClasses: {
    type: String,
    required: false
  }
});

const isPasswordVisible = ref(false);
</script>

<template>
  <!-- TODO: POSSIBLE TO MERGE WITH BASE PASSWORD AND MODAL -->
  <!-- TODO: POSSIBLE TO PUT IN ON BOTH SIDE OF THE INPUT -->
  <!-- TODO: I NEED TO ADD CLASS TO MAIN DIV AND OTHER ATTRS TO INPUT -->
  <div :class="`input-wrapper${props.wrapperClasses ? ' ' + props.wrapperClasses : ''}`">
    <label v-if="props.label" class="input-label" :for="props.uniqueId">{{ label }}</label>
    <div class="input-subwrapper">
      <input type="search" class="input" :id="props.uniqueId" :placeholder="props.placeholder" :value="props.modelValue"
      @input="$emit('update:modelValue', $event.target.value)" v-bind="$attrs" />
      <button class="search-icon-wrapper" type="button" @click="props.onClick">
        <img src="@/assets/search.svg" />
      </button>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.input {
  font-family: Inter;
  font-size: 14px;
  font-style: normal;
  font-weight: 400;
  padding: 8px;
  padding-left: 0;
  height: 38px;
  border: none;
  border-bottom: 2px solid $colorLightGrey;
  width: 100%;

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
    margin-top: -1px; // TODO: Fix bug with unwanted extend of div by 1px
  }

  &-subwrapper {
    width: 100%;
    display: flex;
  }
}

.search-icon-wrapper {
  padding-left: 8px;
  height: 38px;
  border: none;
  border-bottom: 2px solid $colorLightGrey;
  cursor: pointer;
  display: flex;
  align-items: center;
}
</style>