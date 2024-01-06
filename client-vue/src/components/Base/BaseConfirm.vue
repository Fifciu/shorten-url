<script setup lang="ts">
import { defineEmits } from 'vue';

const emit = defineEmits<{
  (e: 'close'): void
}>();
</script>

<template>
  <div class="modal" @click="emit('close')">
    <div class="dialog" @click.stop>
      <div class="wrapper">
        <button class="dialog-close-btn" @click="emit('close')">
          <img src="@/assets/close.svg" alt="Close button" />
        </button>
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.modal {
  position: fixed;
  width: 100%;
  height: 100vh;
  top: 0;
  left: 0;
  background: rgba(0, 0, 0, .3);
  display: flex;
  align-items: center;
  justify-content: center;
}

//TODO: only desktop
.wrapper {
  height: 100%;
  display: flex;
  flex-direction: column;
}
//end

.dialog {
  height: 100%;
  width: 100%;
  max-width: 800px;
  background: $colorWhite;
  padding: 32px;

  .wrapper {
    position: relative;
  }

  &-close-btn {
    cursor: pointer;
    position: absolute;
    top: 0;
    right: 0;
    transform: translate(100%, -100%);

    img {
      width: 12px;
    }
  }
}

@media screen and (min-width: $widthDesktop) {
  .dialog {
    height: initial;
    border-radius: 8px;
    border: 2px solid $colorLightGrey;
  }

  .dialog-close-btn {
    cursor: pointer;
  }
}
</style>
