import { useApi } from "./api";
import { useQuasar } from "quasar";
import { useGeolocation } from "@vueuse/core";

function useGeo() {
  const api = useApi();
  const { coords } = useGeolocation();
  const quasar = useQuasar();

  async function sendCoords() {
    try {
      quasar.loading.show();
      const { data } = await api.post("/geolocation", {
        latitude: coords.value.latitude,
        longitude: coords.value.longitude,
      });
      quasar.notify({
        message: "coordinates successfully sent.",
        color: "positive",
      });
    } catch (err) {
      console.error(err);
      quasar.notify({
        message: "something goes wrong...",
        color: "negative",
      });
    } finally {
      quasar.loading.hide();
    }
  }
}
