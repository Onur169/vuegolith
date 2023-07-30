<template>
  <div
    class="w-full fixed bottom-0 left-0 h-12 border-t border-primary bg-secondary p-3 text-center"
  >
    {{ text }}
  </div>
</template>
<script lang="ts" setup>
interface Props {
  text: string;
}

const { text } = defineProps<Props>();
</script>
