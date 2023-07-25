<template>
  <TabGroup>
    <TabList class="flex space-x-1 rounded-xl bg-blue-900/20 p-1">
      <Tab
        v-for="tab in tabs"
        :key="tab.id"
        v-slot="{ selected }"
        as="template"
      >
        <button
          :class="[
            'w-full rounded-lg py-2.5 text-sm font-medium leading-5 focus:outline-none',
            selected
              ? 'bg-white shadow'
              : 'text-blue-100 hover:bg-white/[0.12] hover:text-white',
          ]"
          @click="handleTabBtnClick(tab)"
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
        <div v-if="activeTab === tab.name">{{ tab.content }}</div>
      </TabPanel>
    </TabPanels>
  </TabGroup>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { TabGroup, TabList, Tab, TabPanels, TabPanel } from "@headlessui/vue";

export interface TabItem {
  name: string;
  id: number;
  content: string;
}

const { tabs } = defineProps<{
  tabs: TabItem[];
}>();

const handleTabBtnClick = (tab: TabItem) => (activeTab.value = tab.name);

const activeTab = ref(tabs.length > 0 ? tabs[0].name : null);
</script>
