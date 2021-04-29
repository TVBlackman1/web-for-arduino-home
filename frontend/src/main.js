import { createApp } from 'vue'
import {createRouter, createWebHistory} from "vue-router";
import App from './App.vue'

import Devices from './components/Device/Devices'
import DevicePage from './components/Device/DevicePage'


const routes = [
    {
        path: "/device/:id",
        name: "current-device",
        component: DevicePage
    },
    {
        path: "/",
        name: "Devices",
        component: Devices
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes: routes,
})

const app = createApp(App)
app.use(router)
router.isReady().then(() => app.mount('#app'))
// app.mount('#app')
