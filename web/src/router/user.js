export const userRoutes = {
    path: "/user",
    component: () => import('../view/layout/Layout.vue'),
    children: [
        {
            path: "list",
            component: () => import('../view/user/List.vue'),
        },
    ]
}