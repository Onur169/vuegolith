<template>
  <div class="flex flex-col justify-center w-full p-6">
    <Tabs :tabs="tabs" />
  </div>
</template>

<script setup lang="tsx">
import { ref } from "vue";
import Tabs from "./components/Tabs.vue";

const tabs = ref([
  {
    name: "write",
    id: 1,
    content: <p>fsd</p>,
  },
  {
    name: "upload",
    id: 2,
    content: <p>fsd</p>,
  },
]);
</script>
