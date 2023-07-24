<template>
  <div class="flex flex-col justify-center w-full p-6">
    <TabGroup>
      <TabList class="flex space-x-1 rounded-xl bg-blue-900/20 p-1">
        <Tab
          v-for="tab in tabs"
          as="template"
          :key="tab.id"
          v-slot="{ selected }"
        >
          <button
            :class="[
              'w-full rounded-lg py-2.5 text-sm font-medium leading-5 text-blue-700 focus:outline-none',
              selected
                ? 'bg-white shadow'
                : 'text-blue-100 hover:bg-white/[0.12] hover:text-white',
            ]"
          >
            {{ tab.name }}
          </button>
        </Tab>
      </TabList>

      <TabPanels class="mt-2">
        <TabPanel
          v-for="(tab, idx) in tabs"
          :key="idx"
          :class="['rounded-xl bg-white p-3']"
        >
          <div v-if="tab.name == 'write'">write!</div>
          <div v-if="tab.name == 'upload'">upload!</div>
        </TabPanel>
      </TabPanels>
    </TabGroup>
  </div>
</template>

<script setup lang="tsx">
import { ref } from "vue";
import { TabGroup, TabList, Tab, TabPanels, TabPanel } from "@headlessui/vue";

const tabs = ref([
  {
    name: "write",
    id: 1,
  },
  {
    name: "upload",
    id: 2,
  },
]);
</script>
