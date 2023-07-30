<template>
  <div
    class="w-full fixed bottom-0 left-0 h-12 border-t border-primary bg-secondary p-3 text-center"
  >
    {{ computedLogDate }}: {{ text }}
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, onUnmounted, toRefs } from "vue";
import { formatDistanceToNow } from "date-fns";
import { de } from "date-fns/locale";

interface Props {
  text: string;
  logDate: Date;
}

const computedLogDate = ref("");

const props = defineProps<Props>();
const { text, logDate } = toRefs(props);

// Funktion zum Aktualisieren des computedLogDate-Werts
function updateLogDate() {
  const distanceNow = logDate
    ? formatDistanceToNow(logDate.value, {
        addSuffix: true,
        locale: de,
        includeSeconds: true,
      })
    : "";
  computedLogDate.value = distanceNow;
}

onMounted(() => {
  console.log("mounted");
  updateLogDate();

  const timer = setInterval(updateLogDate, 1000 * 1);

  onUnmounted(() => {
    console.log("unmounted");
    clearInterval(timer);
  });
});
</script>
