<script setup>
import { ref } from 'vue';

const props = defineProps({
  name: String,
  short_link: String,
  updated_at: String,
});

const isOpen = ref(false);
</script>

<template>
  <div :class="['link', isOpen && 'link--opened']" @click="isOpen = !isOpen">
    <button class="indicator">
      <img src="@/assets/indicator-closed.svg" v-show="!isOpen" />
      <img src="@/assets/indicator-opened.svg" v-show="isOpen" />
    </button>
    <div class="content content--closed" v-if="!isOpen">
      <div class="item-preview">
        <span class="item-label">Name</span>
        <span class="item-name">{{ props.name }}</span>
      </div>
      <button class="item-wishlist"><img src="@/assets/wishlist-empty.svg" /></button>
    </div>
    <div class="content content--opened" v-else>
      <div class="item-label">Name</div>
      <div class="item-name">{{ props.name }}</div>
      <div class="item-label">Short link</div>
      <div class="item-name">
        <span class="baselink">https://short.link/</span>
        <span class="alias">{{ props.short_link }}</span>
      </div>
      <div class="item-label">Updated at</div>
      <div class="item-name">{{ props.updated_at }}</div>
      <div class="item-label">Actions</div>
      <div class="item-actions">
        <button><img src="@/assets/copy.svg" /></button>
        <button class="mx-2"><img src="@/assets/edit.svg" /></button>
        <button><img src="@/assets/bin.svg" /></button>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.link {
  display: flex;
  align-items: center;
  padding: 16px 0;
  border-top: 2px solid $colorWhiteGrey;

  &--opened {
    padding: 24px 0;

    .indicator {
      align-self: flex-start;
    }
  }
}

.content {
  display: flex;
  width: calc(100% - 42px);
  justify-content: space-between;

  .item-label {
    margin-right: 16px;
    color: $colorBlack;
    font-family: Source Sans Pro;
    font-size: 14px;
    font-style: normal;
    font-weight: 600;
    line-height: 150%;
    text-align: left;
    /* 21px */
  }

  .item-name {
    color: $colorDarkGrey;
    font-family: Source Sans Pro;
    font-size: 14px;
    font-style: normal;
    font-weight: 400;
    line-height: 150%;
    /* 21px */
    text-transform: capitalize;
  }

  &--opened {
    display: grid;
    grid-template-columns: auto auto;
    column-gap: 32px;
    grid-template-rows: auto auto;
    row-gap: 16px;

    .item-label {
      margin-right: 32px;
    }
  }

  &--closed {
    cursor: pointer;
    .item-wishlist {
      margin-left: 16px;
    }
  }
}

.indicator {
  margin-right: 16px;
  display: flex;
  align-items: center;
}

.baselink {
  color: $colorGrey;
  font-family: Source Sans Pro;
  font-size: 14px;
  font-style: normal;
  font-weight: 400;
  line-height: 150%;
  /* 21px */
  text-transform: lowercase;
}

.alias {
  color: $colorDarkGrey;
  font-family: Source Sans Pro;
  font-size: 14px;
  font-style: normal;
  font-weight: 600;
  line-height: 150%;
  /* 21px */
  text-transform: lowercase;
}
</style>