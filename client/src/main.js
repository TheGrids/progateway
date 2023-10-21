import { createApp } from 'vue'
import { createPinia } from 'pinia'
import Notifications from '@kyvg/vue3-notification'
import App from './App.vue'
import router from './router'

const pinia = createPinia();

createApp(App).use(pinia).use(router).use(Notifications).mount('#app')
