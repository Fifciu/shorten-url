<script setup lang="ts">
import { reactive, defineEmits } from 'vue';
import ShInput from '@/components/ShInput.vue';
import ShButton from '@/components/ShButton.vue';
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
const { links, addNewLink } = useLinksStore();
async function onSubmit() {
  return await addNewLink(formData.name, formData.original_url, formData.alias);
}
</script>

<template>
  <h2 class="title">New link</h2>
  <hr class="title-divider" />
  <form @submit.prevent="onSubmit">
    <div class="fields">
      <ShInput v-model="formData.original_url" class="mb-2" uniqueId="addLinkNewUrl" type="text" label="Long URL"
        placeholder="example: http://very-long-link.com/example-example" hasBlackLabel />
      <div class="divided-fields mb-2">
        <ShInput uniqueId="domain" type="text" label="Domain" :value="REDIRECT_BASE_URL" disabled hasBlackLabel />
        <span class="span">/</span>
        <ShInput v-model="formData.alias" uniqueId="addLinkBackHalf" type="text" label="Back-half (optional)"
          placeholder="random string if empty, or provide alias" hasBlackLabel />
      </div>
      <ShInput v-model="formData.name" class="mb-3" uniqueId="addLinkBackHalf" type="text" label="Name short link"
        placeholder="example: My Favourite Song" hasBlackLabel />
    </div>
    <div class="buttons">
      <ShButton variant="secondary mr-2" @click.prevent="emit('close')" type="button">Cancel</ShButton>
      <ShButton variant="primary" type="submit">Create</ShButton>
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