<template>
  <div class="flex justify-center">
    <div class="flex flex-col justify-center w-full p-6 md:w-3/4">
      <Tabs :tabs="tabs">
        <template v-slot:default="{ tab, activeTab }">
          <component
            :is="tab.content"
            v-if="'upload' === activeTab"
            @files="handleFilesSelected"
          />
          <component :is="tab.content" v-else />
        </template>
      </Tabs>
    </div>
  </div>
</template>

<script setup lang="tsx">
import { ref } from "vue";
import Tabs from "./components/Tabs.vue";
import Textarea from "./components/Textarea.vue";
import Button from "./components/Button.vue";
import Filechooser from "./components/Filechooser.vue";

const handleFilesSelected = (files: FileList) => {
  console.log(files);
};

const tabs = ref([
  {
    name: "log",
    id: 1,
    content: (
      <>
        <Textarea class="mb-4" />
        <Button class="mb-4 w-full" text="Speichern" isLoading={false} />
      </>
    ),
  },
  {
    name: "upload",
    id: 2,
    content: (
      <Filechooser v-on:files="handleFilesSelected" v-on:foo="handleFoo" />
    ),
  },
]);
</script>
