import { createRouter, createWebHistory } from 'vue-router';
import Start from '../components/Start.vue';

const routes = [
    {
        path: '/',
        redirect: '/start'
    },
    {
        path: '/start',
        name: 'Start',
        component: Start
    }
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

export default router;