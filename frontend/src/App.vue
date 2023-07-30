<template>
  <div class="flex justify-center">
    <div class="flex flex-col justify-center w-full p-6 md:w-3/4">
      <Tabs :tabs="tabs">
        <template v-slot:default="{ tab, activeTab }">
          <component :is="tab.content" />
        </template>
      </Tabs>
      <StatusBar :text="statusText" :logDate="lastActionDate" />
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
import StatusBar from "./components/StatusBar.vue";

const handleFilesSelected = (files: FileList) => {
  console.log(files);
};

const handleLogButton = () => {
  const logData = {
    message: "Dies ist eine Testmeldung",
    timestamp: new Date().toISOString(),
  };

  log(logData)
    .then(() => setStatus("Log erfolgreich"))
    .catch(() => setStatus("Log nicht erfolgreich"));
};

const setStatus = (msg: string) => {
  lastActionDate.value = new Date();
  statusText.value = msg;
};

const isLoading = ref(false);

const lastActionDate = ref<Date>(new Date());

const statusText = ref("Bisher keine Aktion durchgeführt");

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
