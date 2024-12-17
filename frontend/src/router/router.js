import { createRouter, createWebHistory } from 'vue-router';
import NewConnection from '../components/NewConnection.vue';

const routes = [
    {
        path: '/',
        name: 'NewConnection',
        component: NewConnection,
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

export default router;