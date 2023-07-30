<template>
  <textarea
    @change="handleChange"
    class="w-full resize-none h-24 rounded-lg outline-none border-2 border-primary p-4"
    spellcheck="false"
  ></textarea>
</template>
<script lang="ts" setup>
import { onMounted } from "vue";

// interface Props {
//   myProp: any;
// }

// const { myProp } = defineProps<Props>();

const handleChange = (e: Event) => {
  console.log(e);
};

onMounted(() => {
  console.log("onMounted");
});
</script>
