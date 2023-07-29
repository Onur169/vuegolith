<template>
  <textarea @change="handleChange"></textarea>
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
