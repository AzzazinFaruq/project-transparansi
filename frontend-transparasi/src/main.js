/**
 * main.js
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Plugins
import { registerPlugins } from '@/plugins'
import './styles/font.css'
import './styles/scss/main.scss'

// Components
import App from './App.vue'
import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:8000';
axios.defaults.withCredentials = true;
axios.defaults.withXSRFToken = true;
axios.defaults.headers.common['Content-Type'] = 'application/json';

// Composables
import { createApp } from 'vue'
import { createPinia } from 'pinia'

const pinia = createPinia()
const app = createApp(App)

registerPlugins(app)
app.use(pinia)

app.mount('#app')
