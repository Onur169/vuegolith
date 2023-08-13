<template>
  <div class="flex justify-center">
    <div class="w-4/5 md:w-3/4">
      <header
        class="fixed flex flex-row items-center justify-between p-3 w-4/5 md:w-3/4 bg-white/40 drop-shadow-sm backdrop-blur-sm shadow-lg"
      >
        <div
          class="antialiased gap-x-3 text-center font-regular uppercase select-none flex flex-row items-center justify-center w-auto"
        >
          <h1 class="text-xl md:text-3xl font-bold white-text-shadow">
            <RouterLink to="/" class="focus-style">Vuegolith</RouterLink>
          </h1>
          <p class="font-bold italic white-text-shadow">Vue + Go ⚡</p>
        </div>
        <button class="focus-style" @focus="openMenu">
          <Bars3Icon
            @click="openMenu"
            class="h-7 w-7 md:h-9 md:w-9 cursor-pointer hover:opacity-[.65] betterhover:hover:shadow-lg betterhover:hover:rounded-full betterhover:hover:p-1 betterhover:hover:scale-125 transition-all"
            v-if="!showMenu"
          />
        </button>
      </header>
      <NavMenu
        :items="navigationItems"
        :showMenu="showMenu"
        @outsideClick="handleOutsideClickMenu"
      />
      <main class="flex w-full justify-center items-center mt-20 md:mt-24">
        <div class="w-full">
          <RouterView v-slot="{ Component, route }">
            <Transition name="fade" mode="out-in">
              <div :key="route.name ?? ''">
                <component :is="Component" />
              </div>
            </Transition>
          </RouterView>
        </div>
      </main>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { RouterView, RouterLink } from 'vue-router';
import { Transition, ref } from 'vue';
import { Bars3Icon } from '@heroicons/vue/24/solid';
import NavMenu from './components/NavMenu.vue';
import router, { routes } from './router';

const handleOutsideClickMenu = () => {
  showMenu.value = false;
};

const openMenu = (e: MouseEvent | FocusEvent) => {
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
