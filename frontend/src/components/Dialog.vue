<template>
  <dialog
    ref="dialog"
    class="backdrop:backdrop-blur-sm border-2 p-3 border-primary select-none"
    v-bind="attrs"
  >
    <slot />
  </dialog>
</template>

<script lang="ts" setup>
import { onMounted, ref, watchEffect, useAttrs } from 'vue';

const dialog = ref<InstanceType<any> | undefined>(null);
const internalOpen = ref(false);

const props = defineProps({
  open: Boolean,
  inline: Boolean,
});

const attrs = useAttrs();

function openCloseDialog() {
  if (!dialog?.value) return;
  if (props.open) dialog.value.show();
  else dialog.value.close();
}

function showHideDialog() {
  if (!dialog?.value) return;
  if (props.open) dialog.value.showModal();
  else dialog.value.close();
}

onMounted(() => {
  watchEffect(() => {
    if (props.open !== internalOpen.value) {
      if (props.inline) openCloseDialog();
      else showHideDialog();
      internalOpen.value = props.open;
    }
  });
});
</script>
