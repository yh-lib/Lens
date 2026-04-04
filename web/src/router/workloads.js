// /node/dashboard
export const workloadsRoutes = {
    path: "/workloads",
    component: () => import('../view/layout/Layout.vue'),
    children: [
        {
            path: "pod",
            component: () => import('../view/pod/pod.vue'),
        }
    ]
}