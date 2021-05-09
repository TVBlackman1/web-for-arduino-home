import { createApp } from 'vue'
import {createRouter, createWebHistory} from "vue-router";
import App from './App.vue'

import Devices from './components/Device/Devices'
import DevicePage from './components/Device/DevicePage'
import LoginZone from "./components/Login/LoginZone";
import Profile from "./components/Login/Profile";
import News from "./components/News/News";


const routes = [
    {
        path: "/",
        name: "Devices",
        component: Devices
    },
    {
        path: "/device/:id",
        name: "current-device",
        component: DevicePage
    },
    {
        path: "/register",
        name: "register",
        component: LoginZone
    },
    {
        path: "/profile",
        name: "profile",
        component: Profile
    },
    {
        path: "/news",
        name: "news",
        component: News
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes: routes,
})

const app = createApp(App)
app.use(router)
router.isReady().then(() => app.mount('#app'))
// app.mount('#app')
