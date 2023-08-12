<template>
  <div>
    <header>
      <Bars3Icon
        @click="openMenu"
        class="h-5 w-5 md:h-9 md:w-9 fixed top-2 right-4 z-10 cursor-pointer hover:opacity-[.65]"
      />
      <NavMenu
        :items="navigationItems"
        :showMenu="showMenu"
        @outsideClick="handleOutsideClickMenu"
      />
    </header>
    <main class="flex w-full justify-center items-center">
      <div class="w-4/5 md:w-3/4 mt-12">
        <RouterView v-slot="{ Component, route }">
          <Transition name="page" mode="out-in">
            <div :key="route.name ?? ''">
              <component :is="Component" />
            </div>
          </Transition>
        </RouterView>
      </div>
    </main>
  </div>
</template>

<script lang="ts" setup>
import { RouterView } from 'vue-router';
import { Transition, ref } from 'vue';
import { Bars3Icon } from '@heroicons/vue/24/solid';
import NavMenu from './components/NavMenu.vue';
import router, { routes } from './router';

const handleOutsideClickMenu = () => {
  showMenu.value = false;
};

const openMenu = (e: MouseEvent) => {
  e.stopPropagation();
  showMenu.value = true;
};

router.beforeEach(async () => {
  showMenu.value = false;
});

const showMenu = ref(false);
const navigationItems = ref(routes.map(route => ({ name: route.name, href: route.path })));
</script>

<style type="postcss">
.page-enter-active,
.page-leave-active {
  transition-duration: 150ms;
  transition-property: opacity;
  will-change: opacity;
  transition-timing-function: ease;
}

.page-enter,
.page-leave-active {
  opacity: 0;
}
</style>
