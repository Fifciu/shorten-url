<script setup>
import { computed, defineProps, defineEmits, ref } from 'vue';

const isOpened = ref(false);

const props = defineProps({
  modelValue: String,
  fields: {
    type: Array,
    validator(fields) {
      return fields.every(field => field?.label && field?.value)
    }
  }
});

const emit = defineEmits(['update:modelValue']);

const currentLabel = computed(() => {
  const selectedField = props.fields.find(field => field.value === props.modelValue);
  if (!selectedField) {
    return '---';
  }
  return selectedField.label;
});

const updateModelAndCloseSelect = (val) => {
  emit('update:modelValue', val);
  isOpened.value = false;
};
</script>

<template>
  <div :class="['select', isOpened && 'select--open']" v-click-outside="() => isOpened = false">
    <div class="main-item" @click="isOpened = !isOpened">
      <span>{{ currentLabel }}</span>
      <img src="@/assets/dropdown.svg" />
    </div>
    <div class="select-fields" v-show="isOpened">
      <div class="item" v-for="item in props.fields" :key="item.value" @click="updateModelAndCloseSelect(item.value)">
        {{ item.label }}
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.select {
  border-radius: 8px;
  border: 2px solid $colorWhiteGrey2;
  position: relative;
  width: 140px;

  &--open {
    border-color: $colorLightGrey;
    border-bottom-left-radius: 0;
    border-bottom-right-radius: 0;
  }
}

.main-item {
  padding: 8px 24px;
  height: 49px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  color: $colorDarkGrey;
  text-align: center;
  font-family: Source Sans Pro;
  font-size: 14px;
  font-style: normal;
  font-weight: 600;
  line-height: 150%; /* 21px */

  img {
    margin-left: auto;
  }
}

.select--open {
  .main-item > img {
    transform: rotate(180deg);
  }
}

.select-fields {
  position: absolute;
  padding: 8px 0;
  width: calc(100% + 4px);
  bottom: 0;
  left: -2px;
  transform: translateY(100%);
  background: $colorWhite;
  border-bottom-left-radius: 8px;
  border-bottom-right-radius: 8px;
  border: 2px solid $colorWhiteGrey2;
}

.item {
  width: calc(100% + 4px);
  margin-left: -2px;
  padding: 8px 16px;
  color: $colorDarkGrey;
  font-family: Source Sans Pro;
  font-size: 14px;
  font-style: normal;
  font-weight: 600;
  line-height: 150%; /* 21px */
  border-right: 2px solid $colorLightGrey;
  border-left: 2px solid $colorLightGrey;
  cursor: pointer;

  &:hover {
    border-right: 2px solid $colorBlue;
    border-left: 2px solid $colorBlue;
    background: $colorWhiteGrey2;
  }
}

.select--open {
  .select-fields {
    border-color: $colorLightGrey;
  }
}
</style>