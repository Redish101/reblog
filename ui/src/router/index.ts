import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
    {
        path: '/',
        name: 'a',
        component: () => import("@/view/a.vue"),
        meta: {
            title: 'A'
        }
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

router.beforeEach(async (to, _) => {
    document.title = `${to.meta["title"]} | reblog admin`
})

export default router;