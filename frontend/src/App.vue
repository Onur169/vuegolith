<template>
  <div class="flex flex-col justify-center w-full p-6">
    <Tabs :tabs="tabs" />
  </div>
</template>

<script setup lang="tsx">
import { ref } from "vue";
import Tabs from "./components/Tabs.vue";
import Textarea from "./components/Textarea.vue";
import Button from "./components/Button.vue";

const tabs = ref([
  {
    name: "write",
    id: 1,
    content: (
      <div>
        <Textarea class="mb-4" />
        <Button class="mb-4 w-full" text="Speichern" isLoading={false} />
      </div>
    ),
  },
  {
    name: "upload",
    id: 2,
    content: <p>fsd</p>,
  },
]);
</script>
