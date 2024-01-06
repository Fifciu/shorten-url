<script setup lang="ts">
const props = defineProps<{
  toCopy: string
}>();

async function onClick() {
  await navigator.clipboard.writeText(props.toCopy);
}
</script>

<template>
  <button @click="onClick">
    <svg width="13" height="13" viewBox="0 0 13 13" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path
        d="M3.91034 12.0269H10.8389C11.289 12.0269 11.654 11.6462 11.654 11.1765V3.94894C11.654 3.47932 11.289 3.09863 10.8389 3.09863H3.91034C3.46015 3.09863 3.09521 3.47932 3.09521 3.94894V11.1765C3.09521 11.6462 3.46015 12.0269 3.91034 12.0269Z"
        stroke="#C7C4C4" stroke-linecap="round" stroke-linejoin="round" />
      <path
        d="M1.05737 9.47621V1.82345C1.05737 1.59793 1.14325 1.38166 1.29612 1.22219C1.44898 1.06273 1.65631 0.973145 1.87249 0.973145H9.20857"
        stroke="#C7C4C4" stroke-linecap="round" stroke-linejoin="round" />
    </svg>
  </button>
</template>

<style lang="scss" scoped>
button {
  cursor: pointer;

  &:hover {
    path {
      stroke: $colorBlue;
    }
  }
}

path {
  transition: $transitionTime;
  will-change: stroke;
}</style>