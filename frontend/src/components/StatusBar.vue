<template>
  <div
    class="select-none flex flex-row items-center justify-center font-regular w-full fixed bottom-0 left-0 h-12 border-t border-primary bg-secondary md:p-3"
  >
    <div
      class="overflow-hidden text-ellipsis whitespace-nowrap w-3/4 md:w-auto"
      @click="handleClick"
      :style="{ '--custom-transition-duration': '300ms' }"
    >
      <Transition name="fade" mode="out-in">
        <div class="flex flex-row" :key="new Date().getTime()">
          <slot v-if="computedLogDate"></slot
          >{{ computedLogDate ? `${computedLogDate}: ${text}` : `` }}
        </div></Transition
      >
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, toRefs, Transition } from 'vue';
import { formatDistanceToNow } from 'date-fns';
import { de } from 'date-fns/locale';

interface Props {
  text: string;
  logDate: Date;
}

const props = defineProps<Props>();
const { text, logDate } = toRefs(props);

const isTouchDevice = matchMedia('(pointer: coarse)').matches;

const computedLogDate = ref<string>('');

const updateLogDate = () => {
  const distanceNow = logDate.value
    ? formatDistanceToNow(logDate.value, {
        addSuffix: true,
        locale: de,
        includeSeconds: true,
      })
    : '';
  computedLogDate.value = distanceNow;
};

const handleClick = () => {
  if (isTouchDevice) {
    alert(`${computedLogDate.value}: ${text.value}`);
  }
};

onMounted(() => {
  updateLogDate();
  const timer = setInterval(updateLogDate, 1000 * 10);

  onUnmounted(() => {
    clearInterval(timer);
    emit('destroy');
  });
});

const emit = defineEmits();
</script>
