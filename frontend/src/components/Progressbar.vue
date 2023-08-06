<template>
  <div class="w-full bg-gray-200 h-3">
    <div class="h-3 bg-secondary transition-[width]" :style="{ width: `${percentage}%` }"></div>
  </div>
</template>

<script lang="ts" setup>
import { toRefs } from 'vue';

interface Props {
  percentage: number;
}

const props = defineProps<Props>();
const { percentage } = toRefs(props);
</script>
