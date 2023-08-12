<template>
  <Dialog :open="openDeleteDialog">
    <p class="m-3">
      Datei <strong>{{ fileToDelete }}</strong> löschen?
    </p>
    <div class="flex items-center justify-evenly gap-x-3">
      <Button @clicked="handleFileToDeleteDialog(true)" text="Ja, bitte löschen!" />
      <Button @clicked="handleFileToDeleteDialog(false)" text="Schließen" />
    </div>
  </Dialog>
  <Progressbar :percentage="progressVal" class="fixed top-0 left-0" v-show="showProgress" />
  <Tabs :tabs="tabs" @changed="handleTabChange">
    <template v-slot:default="{ tab, activeTab }">
      <template v-if="'upload' === activeTab">
        <Filechooser :reset="shouldResetFileChooser" @files="handleFilesSelected" />
        <hr class="my-6 bg-primary border border-primary" v-if="fetchedUploadsList?.length > 0" />
        <ul class="flex flex-col" v-if="fetchedUploadsList?.length > 0">
          <li
            class="flex flex-row justify-between items-center font-regular select-none betterhover:hover:bg-secondary w-12- h-12 p-3 mb-6 outline hover:outline-primary hover:shadow-md"
            v-for="(uploadedFile, index) in fetchedUploadsList"
            :key="index"
          >
            <a
              :href="`${apiBaseUrl}uploads/${uploadedFile.name}`"
              :data-text="uploadedFile.name"
              target="_blank"
              class="overflow-hidden text-ellipsis whitespace-nowrap"
              >{{ uploadedFile.name }} ({{ formatBytes(uploadedFile.size) }})</a
            >
            <div class="flex flex-row gap-x-3">
              <IconHover>
                <template #default="{ hover, hovered }">
                  <div
                    @mouseenter="hover(true)"
                    @mouseleave="hover(false)"
                    @click="
                      handleDownload(`${apiBaseUrl}uploads/${uploadedFile.name}`, uploadFile.name)
                    "
                  >
                    <ArrowDownOnSquareIcon class="h-6 w-6 cursor-pointer" v-show="hovered" />
                    <ArrowDownOnSquareIconSolid class="h-6 w-6 cursor-pointer" v-show="!hovered" />
                  </div>
                </template>
              </IconHover>

              <IconHover>
                <template #default="{ hover, hovered }">
                  <div
                    @mouseenter="hover(true)"
                    @mouseleave="hover(false)"
                    @click="handleDelete(uploadedFile.name)"
                  >
                    <TrashIcon class="h-6 w-6 cursor-pointer" v-show="hovered" />
                    <TrashIconSolid class="h-6 w-6 cursor-pointer" v-show="!hovered" />
                  </div>
                </template>
              </IconHover>
            </div>
          </li>
        </ul>
      </template>
      <template v-if="'log' === activeTab">
        <Textarea class="mb-4" @content="handleLogContent" :value="logContent" />
        <Button
          class="mb-4 w-full font-medium"
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
</template>

<script setup lang="tsx">
import { ref } from 'vue';
import Tabs, { TabItem } from './../components/Tabs.vue';
import Textarea from './../components/Textarea.vue';
import Dialog from './../components/Dialog.vue';
import Button from './../components/Button.vue';
import Filechooser from './../components/Filechooser.vue';
import IconHover from './../components/IconHover.vue';
import { apiBaseUrl } from './../api/api';
import {
  logPost,
  logGet,
  uploadFile,
  uploadsGet,
  UploadFile,
  uploadFileWithProgress,
  uploadsDelete,
} from './../api/calls';
import StatusBar from './../components/StatusBar.vue';
import Progressbar from './../components/Progressbar.vue';
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
import CodePreview from './../components/CodePreview.vue';
import { formatBytes } from './../helper';

const showProgress = ref(false);
const progressVal = ref(0);

const logContent = ref('');
const fetchedLogContent = ref('');
const fetchedUploadsList = ref([] as UploadFile[]);

const isLoading = ref(false);

const lastActionDate = ref<Date>(new Date());

const shouldResetFileChooser = ref(false);

const openDeleteDialog = ref(false);
const fileToDelete = ref('');

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
    showProgress.value = true;

    uploadFileWithProgress(file, (percentage: number) => (progressVal.value = percentage))
      .then(async () => {
        setStatus('Upload erfolgreich');
        handleUploadGet();
      })
      .catch(() => setStatus('Upload nicht erfolgreich'))
      .finally(() => {
        shouldResetFileChooser.value = true;
        setTimeout(() => {
          showProgress.value = false;
          progressVal.value = 0;
        }, 1000);
      });
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
      setStatus(
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

    // Hyperlink wieder entfernen
    a.remove();

    setStatus(`${path} erfolgreich heruntergeladen`);
  } catch (error) {
    setStatus(`Fehler beim Herunterladen von ${path}`);
  }
};

const handleDelete = (fileName: string) => {
  openDeleteDialog.value = true;
  fileToDelete.value = fileName;
};

const setStatus = (msg: string) => {
  lastActionDate.value = new Date();
  statusText.value = msg;
};

const handleFileToDeleteDialog = (deleteFile: boolean) => {
  openDeleteDialog.value = false;

  if (deleteFile) {
    uploadsDelete(fileToDelete.value)
      .then(res => {
        if (res.ack === 'success') {
          setStatus(`${fileToDelete.value} erfolgreich gelöscht`);
          handleUploadGet();
        }
      })
      .catch(() => setStatus(`${fileToDelete.value} konnte nicht gelöscht werden`))
      .finally(() => (fileToDelete.value = ''));
  }
};
</script>
