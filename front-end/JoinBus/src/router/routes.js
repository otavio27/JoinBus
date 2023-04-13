const routes = [
  {
    path: "/",
    component: () => import("layouts/MainLayout.vue"),
    children: [
      { path: "", component: () => import("src/pages/IndexPage.vue") },
      {
        path: "favoritos",
        component: () => import("src/pages/FavoritosPage.vue"),
      },
      {
        name: "terminais",
        path: "terminais",
        component: () => import("src/pages/TerminaisPage.vue"),
      },
      {
        name: "terminal",
        path: "terminal/:terminal/linhas",
        component: () => import("src/pages/LinhasPage.vue"),
        props: true,
      },
      {
        name: "linha",
        path: "terminal/:terminal/linha/:linha/horarios",
        component: () => import("src/pages/HoursPage.vue"),
        props: true,
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
