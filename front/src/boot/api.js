import { boot } from "quasar/wrappers";
import { apiKey } from "src/composable/api";
import axios from "axios";
// import { useAuthStore } from 'src/stores/auth'
import { Notify } from "quasar";

export default boot(({ app, store /*, router*/ }) => {
  const api = axios.create({
    baseURL: process.env.BASE_URL, // vai ler o 'BASE_URL' que está no json, ou seja 'http://localhost:8000/api'
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
      // faça aqui qual quer logica em caso de erro, como uma notificação, log ou refresh.

      const res = error.response;
      switch (res.status) {
        case 403:
          Notify.create({
            message: "Você não tem permissão!",
            position: "top",
            color: "warning",
          });
          break;
        case 500:
          Notify.create({
            message: "Erro interno.",
            position: "top",
            color: "danger",
          });
          break;
      }

      return Promise.reject(error);
    }
  );
});
