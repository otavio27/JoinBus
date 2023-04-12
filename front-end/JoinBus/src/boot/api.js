import { boot } from "quasar/wrappers";
import { apiKey } from "src/composable/api";
import axios from "axios";
// import { useAuthStore } from 'src/stores/auth'
// import { Notify } from 'quasar';

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
      /*
    const res = error.response;
    switch (res.status) {
      case 401:
        if (authStore.token) {
          await authStore.refresh();
          if (authStore.token) {
            return api.request(error.config)
          }
        }
        if (res.data?.path != '/api/auth/refresh') {
          Notify.create({ message: 'Usuario não autenticado', color: 'warning' })
          if (router.currentRoute.value.name !== 'login') {
            router.push('/login')
          }
        }
        break;
      case 403:
        Notify.create({ message: 'Você não tem permissão', color: 'warning' })
        if (router.currentRoute.value.name !== 'home') {
          router.push('/home')
        }
        break;
    }
    */
      return Promise.reject(error);
    }
  );
});
