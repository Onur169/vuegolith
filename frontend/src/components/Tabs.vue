<template>
  <TabGroup>
    <TabList class="flex space-x-1 rounded-none bg-primary p-1 mb-4">
      <Tab
        v-for="tab in tabs"
        :key="tab.id"
        v-slot="{ selected }"
        as="template"
      >
        <button
          :class="[
            'capitalize w-full select-none rounded-none py-2.5 text-sm font-medium leading-5 focus:outline-none',
            selected
              ? 'bg-secondary shadow'
              : 'text-blue-100 hover:bg-white/[0.12] hover:text-white',
          ]"
          @click="handleTabBtnClick(tab)"
        >
          {{ tab.name }}
        </button>
      </Tab>
    </TabList>
    <TabPanels>
      <TabPanel
        v-for="(tab, idx) in tabs"
        :key="idx"
        :class="['rounded-none bg-white p-1']"
      >
        <slot :tab="tab" :activeTab="activeTab"></slot>
        <!-- <component :is="tab.content" v-if="activeTab === tab.name" /> -->
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
}

const { tabs } = defineProps<{
  tabs: TabItem[];
}>();

const handleTabBtnClick = (tab: TabItem) => (activeTab.value = tab.name);

const activeTab = ref(tabs.length > 0 ? tabs[0].name : null);
</script>
