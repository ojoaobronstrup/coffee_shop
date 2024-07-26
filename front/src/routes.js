import { createRouter, createWebHistory } from "vue-router";
import IntroductionPage from "./pages/IntroductionPage.vue";
import LoginPage from "./pages/LoginPage.vue";

const routes = [
    { path: '/', component: IntroductionPage },
    { path: '/login', component: LoginPage },
    // { path: '/home', component: HomePage }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router