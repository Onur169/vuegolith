<template>
  <div class="flex justify-center">
    <div class="flex flex-col justify-center w-full p-6 md:w-3/4">
      <Tabs :tabs="tabs">
        <template v-slot:default="{ tab, activeTab }">
          <component :is="tab.content" />
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
import { log } from "./api/api";

const handleFilesSelected = (files: FileList) => {
  console.log(files);
};

const isLoading = ref(false);

const handleLogButton = () => {
  log("hallo welt")
    .then(() => {
      console.log("Log erfolgreich");
    })
    .catch(() => {
      console.log("Log nicht erfolgreich");
    });
};

const tabs = ref([
  {
    name: "log",
    id: 1,
    content: (
      <>
        <Textarea class="mb-4" />
        <Button
          class="mb-4 w-full"
          text="Speichern"
          isLoading={isLoading.value}
          onClicked={handleLogButton}
        />
      </>
    ),
  },
  {
    name: "upload",
    id: 2,
    content: <Filechooser onFiles={handleFilesSelected} />,
  },
]);
</script>
