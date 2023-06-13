import { defineStore } from "pinia";
import { useGeolocation } from "@vueuse/core";
import { ref, watch } from "vue";

export const useGeolocationStore = defineStore("geolocation", () => {
  const { coords, locatedAt, error, resume, pause } = useGeolocation();

  let resolve;
  const isReady = ref(new Promise((_resolve) => (resolve = _resolve)));

  const now = new Date().getTime();
  const unwatch = watch(
    () => coords.value,
    () => {
      if (locatedAt.value > now) {
        unwatch();
        resolve();
      }
    },
    { deep: true, immediate: true }
  );

  return {
    isReady,
    coords,
    locatedAt,
    error,
    resume,
    pause,
  };
});
