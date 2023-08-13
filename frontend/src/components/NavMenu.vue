<template>
  <div
    class="fixed top-0 right-0 transition-transform will-change-transform ease-in-out translate-x-full z-50 w-full md:w-96"
    @click="closeMenuOutside"
    :class="{ '!translate-x-0': showMenu }"
  >
    <nav
      class="bg-white md:bg-transparent bg-gradient-to-r from-secondary/50 to-primary/10 p-8 h-screen relative"
      ref="menuRef"
    >
      <XMarkIcon
        class="h-7 w-7 absolute top-3 right-3 cursor-pointer transition-all hover:scale-125"
        @click="emit('close')"
      />
      <ul class="list-none p-0 mt-9">
        <li
          v-for="(item, index) in items"
          :key="index"
          class="select-none mb-4 last:mb-0 w-full border-2 border-primary hover:bg-secondary font-medium uppercase"
          :class="{
            'bg-secondary underline underline-offset-4':
              router.currentRoute.value.path == item.path,
          }"
        >
          <RouterLink
            class="w-full p-4 flex justify-start items-center"
            :to="item.path"
            v-if="router.currentRoute.value.path !== item.path"
            ><component :is="item.icon" class="ml-4 w-8 h-8 mr-2" />{{ item.name }}</RouterLink
          >
          <span class="w-full p-4 flex justify-start items-center" v-else
            ><component :is="item.icon" class="ml-4 w-8 h-8 mr-2" />{{ item.name }}</span
          >
        </li>
      </ul>
    </nav>
  </div>
</template>

<script setup lang="ts">
import { toRefs, ref, onMounted, onUnmounted } from 'vue';
import { RouterLink } from 'vue-router';
import router, { RouteConfig } from '../router';
import { XMarkIcon } from '@heroicons/vue/24/solid';

interface Props {
  items: RouteConfig[];
  showMenu: boolean;
}

const menuRef = ref<HTMLElement | null>(null);

const closeMenuOutside = (event: MouseEvent) => {
  if (menuRef.value && !menuRef.value.contains(event.target as Node)) {
    emit('close');
  }
};

onMounted(() => {
  document.addEventListener('click', closeMenuOutside);
});

onUnmounted(() => {
  document.removeEventListener('click', closeMenuOutside);
});

const emit = defineEmits(['close']);
const props = defineProps<Props>();
const { items, showMenu } = toRefs(props);
</script>
