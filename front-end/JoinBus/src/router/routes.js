const routes = [
  {
    path: "/",
    component: () => import("layouts/MainLayout.vue"),
    children: [
      {
        name: "index",
        path: "",
        component: () => import("src/pages/IndexPage.vue"),
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
        name: "locale",
        path: "locale",
        component: () => import("src/pages/LocalePage.vue"),
        props: true,
      },
      {
        name: "linha",
        path: "terminal/:terminal?/linha/:linha/horarios",
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
