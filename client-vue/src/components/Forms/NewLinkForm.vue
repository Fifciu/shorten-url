<script setup lang="ts">
import { defineEmits } from 'vue';

const emit = defineEmits<{
  (e: 'close'): void
}>();
</script>

<template>
  <h2 class="title">New link</h2>
  <hr class="title-divider" />
  <form>
    <div class="fields">
      <ShInput class="mb-2" uniqueId="addLinkNewUrl" type="text" label="Long URL"
        placeholder="example: http://very-long-link.com/example-example" hasBlackLabel />
      <div class="divided-fields mb-2">
        <ShInput uniqueId="domain" type="text" label="Domain" value="https://short.link" disabled hasBlackLabel />
        <span class="span">/</span>
        <ShInput uniqueId="addLinkBackHalf" type="text" label="Back-half (optional)"
          placeholder="random string if empty, or provide alias" hasBlackLabel />
      </div>
      <ShInput class="mb-3" uniqueId="addLinkBackHalf" type="text" label="Name short link"
        placeholder="example: My Favourite Song" hasBlackLabel />
    </div>
    <div class="buttons">
      <ShButton variant="secondary mr-2" @click.prevent="emit('close')">Cancel</ShButton>
      <ShButton variant="primary">Create</ShButton>
    </div>
  </form>
</template>

<style lang="scss" scoped>

</style>