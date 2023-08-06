<template>
  <slot :hovered="hovered" :hover="hover" />
</template>

<script lang="ts" setup>
import { ref } from 'vue';
const hovered = ref(false);
const hover = (value: boolean) => (hovered.value = value);
</script>
