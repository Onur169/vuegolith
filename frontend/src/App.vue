<template>
  <div class="flex justify-center">
    <div class="flex flex-col justify-center w-full p-6 md:w-3/4">
      <Tabs :tabs="tabs" @changed="handleTabChange">
        <template v-slot:default="{ tab, activeTab }">
          <template v-if="'upload' === activeTab">
            <Filechooser :reset="shouldResetFileChooser" @files="handleFilesSelected" />
            <hr class="my-6 bg-primary border border-primary" />
            <ul class="flex flex-col" v-if="fetchedUploadsList?.length > 0">
              <li
                class="flex flex-row justify-between my-3"
                v-for="(uploadedFile, index) in fetchedUploadsList"
                :key="index"
              >
                <a :href="`${baseUrl}uploads/${uploadedFile.name}`" target="_blank">{{
                  uploadedFile.name
                }}</a>
                <div class="flex flex-row gap-x-3">
                  <template v-if="!isHoveringDownloadIcon[index]">
                    <ArrowDownOnSquareIconSolid
                      class="h-6 w-6 cursor-pointer"
                      @mouseenter="handleIsHoverDownloadIcon(index, true)"
                      @mouseleave="handleIsHoverDownloadIcon(index, false)"
                    />
                  </template>
                  <template v-else>
                    <ArrowDownOnSquareIcon
                      class="h-6 w-6 cursor-pointer"
                      @click="
                        handleDownload(`${baseUrl}uploads/${uploadedFile.name}`, uploadFile.name)
                      "
                      @mouseenter="handleIsHoverDownloadIcon(index, true)"
                      @mouseleave="handleIsHoverDownloadIcon(index, false)"
                    />
                  </template>

                  <template v-if="!isHoveringTrashIcon[index]">
                    <TrashIconSolid
                      class="h-6 w-6 cursor-pointer"
                      @mouseenter="handleIsHoverTrashIcon(index, true)"
                      @mouseleave="handleIsHoverTrashIcon(index, false)"
                    />
                  </template>
                  <template v-else>
                    <TrashIcon
                      class="h-6 w-6 cursor-pointer"
                      @click="handleDelete(uploadedFile.name)"
                      @mouseenter="handleIsHoverTrashIcon(index, true)"
                      @mouseleave="handleIsHoverTrashIcon(index, false)"
                    />
                  </template>

                  <!-- <TrashIcon class="h-6 w-6" /> -->
                </div>
              </li>
            </ul>
          </template>
          <template v-if="'log' === activeTab">
            <Textarea class="mb-4" @content="handleLogContent" :value="logContent" />
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
import { ref } from 'vue';
import Tabs, { TabItem } from './components/Tabs.vue';
import Textarea from './components/Textarea.vue';
import Button from './components/Button.vue';
import Filechooser from './components/Filechooser.vue';
import {
  logPost,
  logGet,
  uploadFile,
  uploadsGet,
  UploadFile,
  baseUrl,
  uploadsDelete,
  UploadsPayload,
} from './api/api';
import StatusBar from './components/StatusBar.vue';
import {
  PencilSquareIcon,
  CloudArrowUpIcon,
  InformationCircleIcon,
  ArrowDownOnSquareIcon,
  TrashIcon,
} from '@heroicons/vue/24/solid';
import {
  ArrowDownOnSquareIcon as ArrowDownOnSquareIconSolid,
  TrashIcon as TrashIconSolid,
} from '@heroicons/vue/24/outline';
import CodePreview from './components/CodePreview.vue';

const handleClipboardSuccess = () => setStatus('Erfolgreich kopiert');
const handleClipboardFail = () => setStatus('Kopieren hat fehlgeschlagen');

const handleLogGet = () => {
  logGet()
    .then(logData => (fetchedLogContent.value = logData.data))
    .catch(() => {
      setStatus('Log-Read fail');
    });
};

const handleUploadGet = () => {
  uploadsGet()
    .then(uploadsList => {
      const { data } = uploadsList;
      fetchedUploadsList.value = data;
    })
    .catch(() => {
      setStatus('Log-Read fail');
    });
};

const handleTabChange = (name: string) => {
  console.log(name);
  if (name === 'log') {
    handleLogGet();
  }
  if (name === 'upload') {
    handleUploadGet();
  }
};

const handleLogContent = (content: string) => {
  logContent.value = content;
};

const handleFilesSelected = (files: FileList) => {
  if (files.length > 0) {
    const file = files[0];

    shouldResetFileChooser.value = false;

    uploadFile(file)
      .then(async () => {
        setStatus('Upload erfolgreich');
        handleUploadGet();
      })
      .catch(() => setStatus('Upload nicht erfolgreich'))
      .finally(() => (shouldResetFileChooser.value = true));
  }
};

const handleLogButton = () => {
  isLoading.value = true;

  logPost(logContent.value)
    .then(() => {
      setStatus('Log erfolgreich');
      logContent.value = '';

      handleLogGet();
    })
    .catch(() => setStatus('Log nicht erfolgreich'))
    .finally(() =>
      setTimeout(() => {
        isLoading.value = false;
      }, 300),
    );
};

const handleDownload = async (path: string, fileName: string) => {
  try {
    const response = await fetch(path);

    if (!response.ok) {
      throw new Error(
        `Fehler beim Herunterladen der Datei. Status: ${response.status} ${response.statusText}`,
      );
    }

    const blob = await response.blob();
    const url = URL.createObjectURL(blob);

    const a = document.createElement('a');
    a.href = url;
    a.download = fileName;
    document.body.appendChild(a);
    a.click();

    // Nach dem Download die URL des Blobs wieder freigeben
    URL.revokeObjectURL(url);
  } catch (error) {
    console.error('Fehler beim Herunterladen der Datei:', error);
  }
};

const handleDelete = (fileName: string) => {
  const res = confirm(`${fileName} wirklich löschen?`);
  if (res) {
    uploadsDelete({ file: fileName } as UploadsPayload)
      .then(res => {
        if (res.ack === 'success') {
          setStatus(`${fileName} erfolgreich gelöscht`);
          handleUploadGet();
        }
      })
      .catch(() => setStatus(`${fileName} konnte nicht gelöscht werden`));
  }
};

const setStatus = (msg: string) => {
  lastActionDate.value = new Date();
  statusText.value = msg;
};

const logContent = ref('');
const fetchedLogContent = ref('');
const fetchedUploadsList = ref([] as UploadFile[]);

// Array(fetchedUploadsList.value.length).fill(false)... lieberi onMounted
// length check durchführen, da in prod bei keinen uploads die seite nicht korrekt gerendert wird

const isHoveringDownloadIcon = ref(Array(fetchedUploadsList.value.length).fill(false));
const handleIsHoverDownloadIcon = (index: number, hovering: boolean) => {
  const arr = isHoveringDownloadIcon.value;
  arr[index] = hovering;
  isHoveringDownloadIcon.value = arr;
};

const isHoveringTrashIcon = ref(Array(fetchedUploadsList.value.length).fill(false));
const handleIsHoverTrashIcon = (index: number, hovering: boolean) => {
  const arr = isHoveringTrashIcon.value;
  arr[index] = hovering;
  isHoveringTrashIcon.value = arr;
};

const isLoading = ref(false);

const lastActionDate = ref<Date>(new Date());

const shouldResetFileChooser = ref(false);

const statusText = ref('Bisher keine Aktion durchgeführt');

const tabs = ref([
  {
    icon: <PencilSquareIcon class="h-6 w-6 mr-1.5" />,
    name: 'log',
    id: 1,
  },
  {
    icon: <CloudArrowUpIcon class="h-6 w-6 mr-1.5" />,
    name: 'upload',
    id: 2,
  },
] as TabItem[]);
</script>
