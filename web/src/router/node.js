// /node/dashboard
export const nodeRoutes = {
    path: "/node",
    component: () => import('../view/layout/Layout.vue'),
    children: [
        {
            path: "list",
            component: () => import('../view/node/Node.vue'),
        }
    ]
}