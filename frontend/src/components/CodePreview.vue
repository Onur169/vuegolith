<template>
  <div class="relative" v-if="content">
    <pre
      ref="preElement"
      @click="selectText()"
      class="max-h-96 border-2 border-dashed border-primary p-3 overflow-y-scroll font-regular text-md"
      >{{ splitAndReverse(content, false).join('\n') }}</pre
    >
    <template v-if="isHovering && !hasCopied">
      <ClipboardIconSolid
        class="absolute h-6 w-6 top-3 right-3 cursor-pointer"
        @click="handleClick"
        @mouseleave="isHovering = false"
      />
    </template>
    <template v-if="!isHovering && !hasCopied">
      <ClipboardIcon
        class="absolute h-6 w-6 top-3 right-3 cursor-pointer"
        @click="handleClick"
        @mouseover="isHovering = true"
      />
    </template>
    <template v-if="hasCopied">
      <ClipboardDocumentCheckIcon class="absolute h-6 w-6 top-3 right-3 cursor-pointer" />
    </template>
  </div>
</template>
<script lang="ts" setup>
import { toRefs, ref, watch } from 'vue';
import { ClipboardIcon, ClipboardDocumentCheckIcon } from '@heroicons/vue/24/outline';
import { ClipboardIcon as ClipboardIconSolid } from '@heroicons/vue/24/solid';

interface Props {
  content: string;
}

const props = defineProps<Props>();
const { content } = toRefs(props);

const emit = defineEmits(['clipboardSuccess', 'clipboardFail']);

const isHovering = ref(false);

const hasCopied = ref(false);

const preElement = ref<HTMLPreElement | null>(null);

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
    hasCopied.value = true;
  } catch (err) {
    emit('clipboardFail');
    hasCopied.value = false;
  }
};

const selectText = () => {
  if (preElement.value) {
    const range = document.createRange();
    range.selectNode(preElement.value);
    const selection = window.getSelection();
    if (selection) {
      selection.removeAllRanges();
      selection.addRange(range);
    }
  }
};

watch(content, (newValue, _) => {
  if (newValue) {
    hasCopied.value = false;
    isHovering.value = false;
  }
});
</script>
