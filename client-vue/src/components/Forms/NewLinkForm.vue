<script setup lang="ts">
import { reactive, defineEmits } from 'vue';
import BaseInput from '@/components/Base/BaseInput.vue';
import BaseButton from '@/components/Base/BaseButton.vue';
import { useLinksStore } from '@/stores/links';
import { REDIRECT_BASE_URL } from '@/const';

const emit = defineEmits<{
  (e: 'close'): void
}>();
const formData = reactive({
  original_url: '',
  name: '',
  alias: ''
});
const { addNewLink } = useLinksStore();
async function onSubmit() {
  await addNewLink(formData.name, formData.original_url, formData.alias);
  emit('close');
}
</script>

<template>
  <h2 class="title">New link</h2>
  <hr class="title-divider" />
  <form @submit.prevent="onSubmit">
    <div class="fields">
      <BaseInput v-model="formData.original_url" class="mb-2" uniqueId="addLinkNewUrl" type="text" label="Long URL"
        placeholder="example: http://very-long-link.com/example-example" label-style="black" />
      <div class="divided-fields mb-2">
        <BaseInput uniqueId="domain" type="text" label="Domain" :value="REDIRECT_BASE_URL" disabled label-style="black" />
        <span class="span">/</span>
        <BaseInput v-model="formData.alias" uniqueId="addLinkBackHalf" type="text" label="Back-half (optional)"
          placeholder="random string if empty, or provide alias" label-style="black" />
      </div>
      <BaseInput v-model="formData.name" class="mb-3" uniqueId="addLinkBackHalf" type="text" label="Name short link"
        placeholder="example: My Favourite Song" label-style="black" />
    </div>
    <div class="buttons">
      <BaseButton variant="secondary mr-2" @click.prevent="emit('close')" type="button">Cancel</BaseButton>
      <BaseButton variant="primary" type="submit">Create</BaseButton>
    </div>
  </form>
</template>

<style lang="scss" scoped>

.buttons {
  display: flex;
  justify-content: flex-end;
}

.title {
  color: $colorBlack;
  font-family: Inter;
  font-size: 20px;
  font-style: normal;
  font-weight: 700;
  line-height: 150%;
  /* 30px */

  margin-top: 12px;
  margin-bottom: 32px;
}

.title-divider {
  margin-bottom: 32px;
  border: none;
  border-top: 2px solid $colorWhiteGrey;
}

form {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.span {
  display: block;
  padding: 8px 16px;
  width: 38px;
  align-self: flex-end;
  margin-bottom: 6px;
}

.divided-fields {
  display: grid;
  grid-template-columns: 1fr 38px 1fr;
  align-items: center;
}

@media only screen and (min-width: $widthDesktop) {
  .title-divider {
    display: none;
  }

  .title {
    margin-bottom: 16px;
    margin-top: 0;
  }
}
</style>