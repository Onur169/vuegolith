// vite.config.js
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import vueJsx from '@vitejs/plugin-vue-jsx'; // Hinzufügen des Plugins

export default defineConfig({
  plugins: [
    vue(),
    vueJsx(), // Fügen Sie das Plugin zu Ihren Plugins hinzu
  ],
  optimizeDeps: {
    include: ['@vue/runtime-dom'], // Fügen Sie @vue/runtime-dom zu optimizeDeps hinzu
  },
});
