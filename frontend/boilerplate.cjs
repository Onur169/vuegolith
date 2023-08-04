const fs = require("fs");
const path = require("path");

const componentName = capitalizeFirstLetter(process.argv[2]);

function capitalizeFirstLetter(str) {
  return str.charAt(0).toUpperCase() + str.slice(1);
}

if (!componentName) {
  console.error("Error: Please provide a component name.");
  process.exit(1);
}

const template = `
<template>
  <div></div>
</template>
<script lang="ts" setup>
import { toRefs } from "vue";

interface Props {
  value: string;
}

const props = defineProps<Props>();
const { value } = toRefs(props);

</script>
`;

const componentsDir = path.join(__dirname, "src/components");
const componentPath = path.join(componentsDir, `${componentName}.vue`);

if (!fs.existsSync(componentsDir)) {
  fs.mkdirSync(componentsDir, { recursive: true });
}

fs.writeFileSync(componentPath, template);

console.log(
  `Component "${componentName}" created successfully at ${componentPath}.`
);
