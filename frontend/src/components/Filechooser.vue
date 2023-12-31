<template>
  <div class="font-regular select-none">
    <input
      type="file"
      ref="fileInput"
      class="cursor-pointer font-regular text-md outline-primary outline-dashed outline-2 p-3 w-full outline-offset-[-3px]"
      @change="onFileChange"
      :accept="accept"
      :multiple="multiple"
      :disabled="disabled"
    />
    <button v-if="!disabled" @click="openFileChooser">{{ buttonText }}</button>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, toRefs } from 'vue';

interface FileChooserProps {
  accept?: string;
  multiple?: boolean;
  disabled?: boolean;
  buttonText?: string;
  reset?: boolean;
}

const props = defineProps<FileChooserProps>();

const { accept, multiple, disabled, buttonText, reset } = toRefs(props);

const emit = defineEmits(['files']);

const fileInput = ref<HTMLInputElement | null>(null);

const openFileChooser = () => {
  if (fileInput.value) {
    fileInput.value.click();
  }
};

const resetFileChooser = () => {
  if (fileInput.value) {
    fileInput.value.value = '';
  }
};

const onFileChange = (event: Event) => {
  const inputElement = event.target as HTMLInputElement;
  if (!inputElement.files) return;

  const selectedFiles: FileList = inputElement.files;
  emit('files', selectedFiles); // Emit the selectedFiles as an event
};

watch(reset, (newValue, _) => {
  if (newValue) {
    resetFileChooser();
  }
});
</script>

<style scoped lang="postcss">
input[type='file']::file-selector-button {
  @apply mr-4 px-4 py-2 border border-primary bg-transparent shadow cursor-pointer transition duration-150 rounded-none;
}

@media (hover: hover) {
  input[type='file']::file-selector-button:hover {
    @apply hover:bg-secondary;
  }
}
</style>
