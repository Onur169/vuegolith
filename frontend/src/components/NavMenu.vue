<template>
  <div
    class="fixed inset-0 flex justify-center items-center bg-black bg-opacity-50 z-50"
    @click="closeMenuOutside"
  >
    <nav
      class="bg-white p-8 shadow-md w-full sm:w-[calc(50vw-1rem)] m-3 sm:m-0 h-auto"
      ref="menuRef"
    >
      <ul class="list-none p-0">
        <li
          v-for="(item, index) in items"
          :key="index"
          class="select-none mb-4 last:mb-0 w-full border-2 border-primary hover:bg-secondary cursor-pointer font-medium uppercase"
          :class="{ 'bg-secondary': router.currentRoute.value.path == item.href }"
        >
          <RouterLink class="text-black font-bold w-full items-stretch flex p-3" :to="item.href">{{
            item.name
          }}</RouterLink>
        </li>
      </ul>
    </nav>
  </div>
</template>

<script setup lang="ts">
import { toRefs, ref, onMounted, onUnmounted } from 'vue';
import { RouterLink } from 'vue-router';
import router from '../router';

export interface NavItem {
  name: string;
  href: string;
}

interface Props {
  items: NavItem[];
}

const menuRef = ref<HTMLElement | null>(null);

const closeMenuOutside = (event: MouseEvent) => {
  if (menuRef.value && !menuRef.value.contains(event.target as Node)) {
    emit('outsideClick');
  }
};

onMounted(() => {
  document.addEventListener('click', closeMenuOutside);
});

onUnmounted(() => {
  document.removeEventListener('click', closeMenuOutside);
});

const emit = defineEmits(['outsideClick']);
const props = defineProps<Props>();
const { items } = toRefs(props);
</script>
