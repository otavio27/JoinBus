import { boot } from "quasar/wrappers";
import { apiKey } from "src/composable/api";
import axios from "axios";
// import { useAuthStore } from 'src/stores/auth'
import { Notify } from "quasar";

export default boot(({ app, store, router }) => {
  const api = axios.create({
    baseURL: process.env.API_URL,
  });

  app.provide(apiKey, api);
  store.use(() => ({ api }));

  // const authStore = useAuthStore(store)
  api.interceptors.request.use(
    function (config) {
      // aplique aqui qual quer alteração no request, como um token ou lang
      /*
    if (authStore.token) {
      config.headers.Authorization = `Bearer ${authStore.token}`;
    }
    */
      return config;
    },
    function (error) {
      return Promise.reject(error);
    }
  );

  api.interceptors.response.use(
    function (response) {
      return response;
    },
    function (error) {
      const res = error.response;
      switch (res.status) {
        case 403:
          Notify.create({
            message: "forbiden!",
            position: "top",
            progress: true,
            type: "warning",
            actions: [{ icon: "close", color: "white" }],
          });
          router.push({ name: "index" });
          break;
        case 404:
          Notify.create({
            message: "not found!",
            position: "top",
            progress: true,
            type: "warning",
            actions: [{ icon: "close", color: "white" }],
          });
          router.push({ name: "index" });
          break;
        case 500:
          Notify.create({
            message: "internal error!",
            position: "top",
            progress: true,
            type: "negative",
            actions: [{ icon: "close", color: "white" }],
          });
          router.push({ name: "index" });
          break;
        case 502:
          Notify.create({
            message: "bad resquest!",
            position: "top",
            progress: true,
            type: "negative",
            actions: [{ icon: "close", color: "white" }],
          });
          router.push({ name: "index" });
          break;
      }

      return Promise.reject(error);
    }
  );
});
