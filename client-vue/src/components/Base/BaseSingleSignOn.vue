<script setup lang="ts">
const props = defineProps({
  disabled: {
    type: Boolean,
    default: true
  }
});
</script>

<template>
  <div>
    <p class="mb-2 text-center font-ternary font-normal text-sm text-gray-200">or</p>
    <div class="flex gap-x-4 justify-between w-[154px] mx-auto">
      <button type="button" class="cursor-not-allowed"><img src="@/assets/sso-fb.svg"></button>
      <button type="button" class="cursor-not-allowed"><img src="@/assets/sso-google.svg"></button>
      <button type="button" class="cursor-not-allowed"><img src="@/assets/sso-instagram.svg"></button>
    </div>
  </div>
</template>