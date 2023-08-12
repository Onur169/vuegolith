<template>
  <div
    class="fixed top-0 right-0 transition-transform will-change-transform ease-in-out translate-x-full z-50 w-full md:w-96"
    @click="closeMenuOutside"
    :class="{ '!translate-x-0': showMenu }"
  >
    <nav
      class="bg-white md:bg-transparent bg-gradient-to-r from-secondary/50 to-primary/10 p-8 h-screen"
      ref="menuRef"
    >
      <ul class="list-none p-0">
        <li
          v-for="(item, index) in items"
          :key="index"
          class="select-none mb-4 last:mb-0 w-full border-2 border-primary hover:bg-secondary font-medium uppercase flex items-stretch p-4"
          :class="{
            'bg-secondary underline underline-offset-4':
              router.currentRoute.value.path == item.href,
          }"
        >
          <RouterLink
            class="w-full"
            :to="item.href"
            v-if="router.currentRoute.value.path !== item.href"
            >{{ item.name }}</RouterLink
          >
          <span v-else>{{ item.name }}</span>
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
  showMenu: boolean;
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
const { items, showMenu } = toRefs(props);
</script>
