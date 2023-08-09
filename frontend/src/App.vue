<template>
  <div>
    <header>
      <Bars3Icon
        @click="toggleMenu"
        class="h-5 w-5 md:h-9 md:w-9 fixed top-2 right-4 cursor-pointer hover:opacity-[.65]"
      />
      <div v-if="showMenu">
        <NavMenu :items="navigationItems" @outsideClick="showMenu = false" />
      </div>
    </header>
    <main>
      <RouterView />
    </main>
  </div>
</template>

<script lang="ts" setup>
import { RouterView } from 'vue-router';
import { ref } from 'vue';
import { Bars3Icon } from '@heroicons/vue/24/solid';
import NavMenu from './components/NavMenu.vue';
import router, { routes } from './router';

const toggleMenu = (e: MouseEvent) => {
  e.stopPropagation();
  showMenu.value = !showMenu.value;
};

router.beforeEach(async (to, from) => {
  console.log(to, from);
  showMenu.value = false;
});

const showMenu = ref(false);
const navigationItems = ref(routes.map(route => ({ name: route.name, href: route.path })));
</script>
