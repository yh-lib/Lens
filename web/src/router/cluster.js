// /cluster/dashboard
export const clusterRoutes = {
    path: "/cluster",
    component: () => import('../view/layout/Layout.vue'),
    children: [
        {
            path: "dashboard",
            component: () => import('../view/cluster/Dashboard.vue'),
        },
    ]
}