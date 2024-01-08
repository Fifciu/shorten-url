<script setup lang="ts">
import { reactive, defineEmits, defineProps, computed } from 'vue';
import BaseInput from '@/components/Base/BaseInput.vue';
import BaseButton from '@/components/Base/BaseButton.vue';
import { useLinksStore } from '@/stores/links';
import { REDIRECT_BASE_URL } from '@/const';

const props = defineProps({
  link: {
    type: Object,
    required: true,
    validator(val: Record<string, any>) {
      const objKeys = Object.keys(val);
      return ['id', 'name',  'original_url', 'alias'].every(key => objKeys.includes(key));
    }
  }
});
const emit = defineEmits<{
  (e: 'close'): void
}>();
const formData = reactive({
  original_url: props.link.original_url,
  name: props.link.name,
  alias: props.link.alias
});
const { updateLink } = useLinksStore();
async function onSubmit() {
  await updateLink(props.link.id, {
    name: props.link.name !== formData.name ? formData.name : undefined,
    alias: props.link.alias !== formData.alias ? formData.alias : undefined,
    original_url: props.link.original_url !== formData.original_url ? formData.original_url : undefined,
  });
  emit('close');
}

const modifiedData = computed(() => {
  return props.link.name !== formData.name
    || props.link.alias !== formData.alias
    || props.link.original_url !== formData.original_url;
});
</script>

<template>
  <h2 class="title">Edit your link</h2>
  <hr class="title-divider" />
  <form @submit.prevent="onSubmit">
    <div class="fields">
      <BaseInput v-model.trim="formData.original_url" class="mb-2" uniqueId="addLinkNewUrl" type="text" label="Long URL"
        placeholder="example: http://very-long-link.com/example-example" label-style="black" />
      <div class="divided-fields mb-2">
        <BaseInput uniqueId="domain" type="text" label="Domain" :value="REDIRECT_BASE_URL" disabled label-style="black" />
        <span class="span">/</span>
        <BaseInput v-model.trim="formData.alias" uniqueId="addLinkBackHalf" type="text" label="Back-half (optional)"
          placeholder="random string if empty, or provide alias" label-style="black" />
      </div>
      <BaseInput v-model.trim="formData.name" class="mb-3" uniqueId="addLinkBackHalf" type="text" label="Name short link"
        placeholder="example: My Favourite Song" label-style="black" />
    </div>
    <div class="buttons">
      <BaseButton variant="secondary mr-2" @click.prevent="emit('close')" type="button">Cancel</BaseButton>
      <BaseButton variant="primary" type="submit" :disabled="!modifiedData">Save</BaseButton>
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