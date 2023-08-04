<template>
  <div class="flex justify-center">
    <div class="flex flex-col justify-center w-full p-6 md:w-3/4">
      <Tabs :tabs="tabs">
        <template v-slot:default="{ tab, activeTab }">
          <template v-if="'upload' === activeTab">
            <Filechooser
              :reset="shouldResetFileChooser"
              @files="handleFilesSelected"
            />
            <hr class="my-6 bg-primary border border-primary" />
            <ul v-if="fetchedUploadsList.length > 0">
              <li class="my-3" v-for="uploadedFileName in fetchedUploadsList">
                <!-- todo: port austauschen mit ba port -->
                <a :href="`uploads/${uploadedFileName}`" target="_blank">{{
                  uploadedFileName
                }}</a>
              </li>
            </ul>
          </template>
          <template v-if="'log' === activeTab">
            <Textarea
              class="mb-4"
              @content="handleLogContent"
              :value="logContent"
            />
            <Button
              class="mb-4 w-full"
              text="Speichern"
              :isLoading="isLoading"
              @clicked="handleLogButton"
            />
            <CodePreview
              @clipboardSuccess="handleClipboardSuccess"
              @clipboardFail="handleClipboardFail"
              :content="fetchedLogContent"
            />
          </template>
        </template>
      </Tabs>
      <StatusBar :text="statusText" :logDate="lastActionDate">
        <template v-slot>
          <InformationCircleIcon class="h-6 w-6 mr-1.5" />
        </template>
      </StatusBar>
    </div>
  </div>
</template>

<script setup lang="tsx">
import { onMounted, ref } from "vue";
import Tabs, { TabItem } from "./components/Tabs.vue";
import Textarea from "./components/Textarea.vue";
import Button from "./components/Button.vue";
import Filechooser from "./components/Filechooser.vue";
import { logPost, logGet, uploadFile, uploadsGet } from "./api/api";
import StatusBar from "./components/StatusBar.vue";
import {
  PencilSquareIcon,
  CloudArrowUpIcon,
  InformationCircleIcon,
} from "@heroicons/vue/24/solid";
import CodePreview from "./components/CodePreview.vue";

const handleClipboardSuccess = () => setStatus("Erfolgreich kopiert");
const handleClipboardFail = () => setStatus("Kopieren hat fehlgeschlagen");

const handleLogContent = (content: string) => {
  logContent.value = content;
};

const handleFilesSelected = (files: FileList) => {
  if (files.length > 0) {
    const file = files[0];

    shouldResetFileChooser.value = false;

    uploadFile(file)
      .then(async () => {
        setStatus("Upload erfolgreich");
      })
      .catch(() => setStatus("Upload nicht erfolgreich"))
      .finally(() => (shouldResetFileChooser.value = true));
  }
};

const handleLogButton = () => {
  isLoading.value = true;

  logPost(logContent.value)
    .then(() => {
      setStatus("Log erfolgreich");
      logContent.value = "";

      logGet()
        .then((logData) => (fetchedLogContent.value = logData.data))
        .catch(() => {
          console.log("Log-Read fail");
        });
    })
    .catch(() => setStatus("Log nicht erfolgreich"))
    .finally(() =>
      setTimeout(() => {
        isLoading.value = false;
      }, 300)
    );
};

const setStatus = (msg: string) => {
  lastActionDate.value = new Date();
  statusText.value = msg;
};

const logContent = ref("");
const fetchedLogContent = ref("");
const fetchedUploadsList = ref([] as string[]);

const isLoading = ref(false);

const lastActionDate = ref<Date>(new Date());

const shouldResetFileChooser = ref(false);

const statusText = ref("Bisher keine Aktion durchgeführt");

const tabs = ref([
  {
    icon: <PencilSquareIcon class="h-6 w-6 mr-1.5" />,
    name: "log",
    id: 1,
  },
  {
    icon: <CloudArrowUpIcon class="h-6 w-6 mr-1.5" />,
    name: "upload",
    id: 2,
  },
] as TabItem[]);

onMounted(() => {
  logGet()
    .then((logData) => (fetchedLogContent.value = logData.data))
    .catch(() => {
      console.log("Log-Read fail");
    });

  uploadsGet()
    .then((uploadsList) => {
      const { data } = uploadsList;
      fetchedUploadsList.value = data;
    })
    .catch(() => {
      console.log("Log-Read fail");
    });
});
</script>
