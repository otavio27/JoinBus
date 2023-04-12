const routes = [
  {
    path: "/",
    component: () => import("layouts/MainLayout.vue"),
    children: [
      { path: "/", component: () => import("src/pages/IndexPage.vue") },
    ],
  },
  {
    path: "/terminais",
    component: () => import("layouts/MainLayout.vue"),
    children: [
      {
        path: "/terminais",
        component: () => import("src/pages/TerminaisPage.vue"),
      },
    ],
  },
  {
    path: "/favoritos",
    component: () => import("layouts/MainLayout.vue"),
    children: [
      {
        path: "/favoritos",
        component: () => import("src/pages/FavoritosPage.vue"),
      },
    ],
  },
  {
    path: "/linhas",
    component: () => import("layouts/MainLayout.vue"),
    children: [
      {
        path: "/linhas",
        component: () => import("src/pages/LinhasPage.vue"),
      },
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: "/:catchAll(.*)*",
    component: () => import("pages/ErrorNotFound.vue"),
  },
];

export default routes;
