<script setup lang="ts">


</script>

<template>
  <button>
    <svg width="13" height="13" viewBox="0 0 13 13" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path
        d="M4.47673 11.1441L0.653809 12.2608L1.72423 8.27283L8.72443 1.0059C8.80352 0.921489 8.89807 0.854421 9.00231 0.808633C9.10664 0.762844 9.21869 0.739258 9.33185 0.739258C9.44501 0.739258 9.55706 0.762844 9.66138 0.808633C9.76562 0.854421 9.86018 0.921489 9.93927 1.0059L11.443 2.58335C11.5226 2.66573 11.5858 2.76376 11.6289 2.87174C11.6721 2.97974 11.6942 3.09557 11.6942 3.21256C11.6942 3.32955 11.6721 3.44539 11.6289 3.55338C11.5858 3.66138 11.5226 3.75939 11.443 3.84177L4.47673 11.1441Z"
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
}
</style>