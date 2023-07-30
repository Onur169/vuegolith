<template>
  <div>
    <input
      type="file"
      ref="fileInput"
      @change="onFileChange"
      :accept="accept"
      :multiple="multiple"
      :disabled="disabled"
    />
    <button v-if="!disabled" @click="openFileChooser">{{ buttonText }}</button>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

interface FileChooserProps {
  accept?: string;
  multiple?: boolean;
  disabled?: boolean;
  buttonText?: string;
}

const { accept, multiple, disabled, buttonText }: FileChooserProps =
  defineProps<FileChooserProps>();
const emit = defineEmits(["files"]);

const fileInput = ref<HTMLInputElement | null>(null);

const openFileChooser = () => {
  if (fileInput.value) {
    fileInput.value.click();
  }
};

const onFileChange = (event: Event) => {
  const inputElement = event.target as HTMLInputElement;
  if (!inputElement.files) return;

  const selectedFiles: FileList = inputElement.files;
  emit("files", selectedFiles); // Emit the selectedFiles as an event
};
</script>
