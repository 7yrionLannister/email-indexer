<script setup>
import HeaderAndContent from "./HeaderAndContentItem.vue";
defineProps({
  subject: {
    type: String,
    required: true,
  },
  body: {
    type: String,
    required: true,
  },
});
</script>

<template>
  <HeaderAndContent>
    <template #heading>{{ subject }}</template>

    {{ body }}
  </HeaderAndContent>
</template>
