<template>
  <div class="flex justify-center">
    <div class="flex flex-col justify-center w-full p-6 md:w-3/4">
      <Tabs :tabs="tabs">
        <template v-slot:default="{ tab, activeTab }">
          <template v-if="'upload' === activeTab">
            <Filechooser
              :reset="shouldResetFileChooser"
              :onFiles="handleFilesSelected"
            />
          </template>
          <template v-if="'log' === activeTab">
            <Textarea class="mb-4" />
            <Button
              class="mb-4 w-full"
              text="Speichern"
              :isLoading="isLoading"
              :onClicked="handleLogButton"
            />
          </template>
        </template>
      </Tabs>
      <StatusBar :text="statusText" :logDate="lastActionDate" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import Tabs from "./components/Tabs.vue";
import Textarea from "./components/Textarea.vue";
import Button from "./components/Button.vue";
import Filechooser from "./components/Filechooser.vue";
import { log, uploadFile } from "./api/api";
import StatusBar from "./components/StatusBar.vue";

const handleFilesSelected = (files: FileList) => {
  if (files.length > 0) {
    const file = files[0];

    uploadFile(file)
      .then(() => setStatus("Upload erfolgreich"))
      .catch(() => setStatus("Upload nicht erfolgreich"))
      .finally(() => (shouldResetFileChooser.value = true));
  }
};

const handleLogButton = () => {
  const logData = {
    message: "Dies ist eine Testmeldung",
    timestamp: new Date().toISOString(),
  };

  isLoading.value = true;

  log(logData)
    .then(() => setStatus("Log erfolgreich"))
    .catch(() => setStatus("Log nicht erfolgreich"))
    .finally(() =>
      setTimeout(() => {
        isLoading.value = false;
      }, 1000)
    );
};

const setStatus = (msg: string) => {
  lastActionDate.value = new Date();
  statusText.value = msg;
};

const isLoading = ref(false);

const lastActionDate = ref<Date>(new Date());

const shouldResetFileChooser = ref(false);

const statusText = ref("Bisher keine Aktion durchgeführt");

const tabs = ref([
  {
    name: "log",
    id: 1,
  },
  {
    name: "upload",
    id: 2,
  },
]);
</script>
