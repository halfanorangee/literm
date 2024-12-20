import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router';
import Start from "../components/common/Start.vue";

const routes: RouteRecordRaw[] = [
    {
        path: '/',
        redirect: '/collections'
    },
    {
        path: '/',
        name: 'Start',
        component: Start,
        children: [

            {
                path: '/collections',
                name: 'Collections',
                component: () => import('../components/collection/Collections.vue'),
            },
            {
                path: '/terminals',
                name: 'Terminals',
                component: () => import('../components/terminal/Terminals.vue'),
            },
            {
                path: '/settings',
                name: 'Settings',
                component: () => import('../components/settings/Settings.vue'),
            },
        ]
    }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export default router;