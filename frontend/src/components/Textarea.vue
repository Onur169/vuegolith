<template>
  <textarea
    v-model="content"
    class="w-full resize-none h-24 rounded-none outline-none border-2 border-primary p-4 font-regular"
    spellcheck="false"
  ></textarea>
</template>
<script lang="ts" setup>
import { ref, watch, toRefs } from 'vue';

interface Props {
  value: string;
}

const props = defineProps<Props>();
const { value } = toRefs(props);

const content = ref('');

const emit = defineEmits(['content']);

watch(content, (newValue, _) => {
  emit('content', newValue);
});

watch(value, (newValue, _) => {
  content.value = newValue;
});
</script>
