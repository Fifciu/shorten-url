<script setup lang="ts">
import { onMounted, onUnmounted, ref } from 'vue';
import { useLinksStore } from '@/stores/links';
import { storeToRefs } from 'pinia';

const loaderRef = ref<Element>();
const { fetchMoreMyLinks } = useLinksStore();
const { isLoading, loadedAllPages } = storeToRefs(useLinksStore());
let observer: IntersectionObserver;
onMounted(() => {
  observer = new IntersectionObserver((entries) => {
    entries.forEach((entry: IntersectionObserverEntry) => {
      if (entry.isIntersecting) {
        if (!isLoading.value && !loadedAllPages.value) {
          fetchMoreMyLinks();
        }
      }
    })
  }, {
    // root: loaderRef.value,
    rootMargin: '0px',
    threshold: 1.0
  });
  observer.observe(loaderRef.value!);
});

onUnmounted(() => {
  observer.unobserve(loaderRef.value!);
});
</script>

<template>
  <div ref="loaderRef"></div>
</template>