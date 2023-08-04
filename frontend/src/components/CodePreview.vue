<template>
  <div class="relative">
    <pre class="max-h-96 border-2 border-dashed border-primary p-3 overflow-y-scroll">{{
      splitAndReverse(content, false).join('\n')
    }}</pre>
    <template v-if="!isHovering">
      <ClipboardIcon
        class="absolute h-6 w-6 top-3 right-3 cursor-pointer"
        @click="handleClick"
        @mouseover="isHovering = true"
      />
    </template>
    <template v-else>
      <ClipboardIconSolid
        class="absolute h-6 w-6 top-3 right-3 cursor-pointer"
        @click="handleClick"
        @mouseleave="isHovering = false"
      />
    </template>
  </div>
</template>
<script lang="ts" setup>
import { toRefs, ref } from 'vue';
import { ClipboardIcon } from '@heroicons/vue/24/outline';
import { ClipboardIcon as ClipboardIconSolid } from '@heroicons/vue/24/solid';

interface Props {
  content: string;
}

const props = defineProps<Props>();
const { content } = toRefs(props);

const emit = defineEmits(['clipboardSuccess', 'clipboardFail']);

const isHovering = ref(false);

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
