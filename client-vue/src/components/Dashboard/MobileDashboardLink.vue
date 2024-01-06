<script setup>
import { ref } from 'vue';
import { REDIRECT_BASE_URL } from '@/const';
import BinButton from '@/components/Actions/BinButton.vue';

const props = defineProps({
  id: Number,
  name: String,
  short_link: String,
  updated_at: String,
});

const isOpen = ref(false);
function openIfClosed() {
  if (!isOpen.value) {
    isOpen.value = true;
  }
}

function toggle() {
  isOpen.value = !isOpen.value;
}
</script>

<template>
  <div :class="['link', isOpen && 'link--opened']" @click="openIfClosed">
    <button class="indicator" @click.stop="toggle">
      <img src="@/assets/indicator-closed.svg" v-show="!isOpen" />
      <img src="@/assets/indicator-opened.svg" v-show="isOpen" />
    </button>

    <div class="content content--closed" v-if="!isOpen">
      <div class="item-preview">
        <span class="item-label">Name</span>
        <span class="item-value">{{ props.name }}</span>
      </div>
      <button class="item-wishlist"><img src="@/assets/wishlist-empty.svg" /></button>
    </div>

    <div class="content content--opened" v-else>
      <div class="item-label">Name</div>
      <div class="item-value">{{ props.name }}</div>
      <div class="item-label">Short link</div>
      <div class="item-value">
        <span class="baselink">{{ REDIRECT_BASE_URL }}</span>
        <span class="alias">{{ props.short_link }}</span>
      </div>
      <div class="item-label">Updated at</div>
      <div class="item-value">{{ props.updated_at }}</div>
      <div class="item-label">Actions</div>
      <div class="item-actions">
        <button><img src="@/assets/copy.svg" /></button>
        <button class="mx-2"><img src="@/assets/edit.svg" /></button>
        <BinButton :linkId="props.id" />
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

.item-label,
.item-value,
.baselink,
.alias {
  font-family: Source Sans Pro;
  font-size: 14px;
  font-style: normal;
  font-weight: 400;
  line-height: 150%;
}

.item-label {
  margin-right: 16px;
  color: $colorBlack;
  font-weight: 600;
  text-align: left;
}

.item-value {
  color: $colorDarkGrey;
  text-transform: capitalize;
}

.baselink {
  color: $colorGrey;
  text-transform: lowercase;
}

.alias {
  color: $colorDarkGrey;
  font-weight: 600;
  text-transform: none;
}</style>