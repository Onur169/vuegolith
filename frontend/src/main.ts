import { createApp } from 'vue';
import { createRippleDirective } from 'vue-create-ripple';
import './index.css';
import App from './App.vue';
import router from './router';

createApp(App)
  .use(router)
  .directive(
    'Ripple',
    createRippleDirective({
      class: 'bg-primary opacity-30',
      // Set other defaults
    }),
  )
  .mount('#app');
