<template>
  <pre class="max-h-96 border border-primary p-3 overflow-y-scroll" @click="handleClick">{{
    splitAndReverse(content, false).join('\n')
  }}</pre>
</template>
<script lang="ts" setup>
import { onMounted, toRefs } from 'vue';

interface Props {
  content: string;
}

const props = defineProps<Props>();
const { content } = toRefs(props);

const emit = defineEmits(['clipboardSuccess', 'clipboardFail']);

// reverse macht Sinn bei Logs die appended werden
// BA muss O_APPEND Flag nutzen beim Log-Post-Call
const splitAndReverse = (inputString: string, reverse: boolean = true): string[] => {
  const linesArray = inputString.split('\n');
  const reversed = reverse ? linesArray.reverse() : linesArray;
  return reversed.filter(item => item.trim() !== '');
};

const handleClick = async () => {
  try {
    await navigator.clipboard.writeText(content.value);
    emit('clipboardSuccess');
  } catch (err) {
    emit('clipboardFail');
  }
};
</script>
