import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Vue from 'vue'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import './assets/scss/index.scss'
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

createApp(App).use(store).use(router).mount('#app')
