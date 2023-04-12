import { defineStore } from "pinia";
import { ref } from "vue";
import { useGeolocation } from "@vueuse/core";

export const useJoinBusStore = defineStore("JoinBus", () => {
  const location = ref([]);
  const terminals = ref([]);
  const lines = ref([]);
  const { coords, error } = useGeolocation();

  async function fetchLocation() {
    // não há necessidade de catch, tratamento generico de erro deve ser feito no interceptor.
    if (!coords.value || error.value) {
      // throw 'location not avaliable'
      return;
    }
    // não há necessidade de catch, tratamento generico de erro deve ser feito no interceptor.
    const { latitude: lat, longitude: lng } = coords.value;
    const { data } = await this.api.get(`/geolocation/${lat}/${lng}`);
    this.location = data;
  }

  async function fetchTerminals() {
    const { data } = await this.api.get("/terminais");
    this.terminais = data;
  }

  async function fetchLines() {
    // não há necessidade de catch, tratamento generico de erro deve ser feito no interceptor.
    const { data } = await this.api.get("/linhas/0300");
    this.lines = data;
  }

  async function fetch() {
    await fetchLines();
    await fetchTerminals();
    await fetchLocation();

    // or in parallel:
    /*
    await Promise.all([
      fetchLines(),
      fetchTerminals(),
      fetchLocation(),
    ])
    */
  }

  return {
    fetch,
    location,
    terminals,
    lines,
    fetchLocation,
    fetchTerminals,
    fetchLines,
  };
});
