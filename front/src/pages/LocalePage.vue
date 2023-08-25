<template>
  <q-page class="q-pa-sm">
    <div class="grid q-pa-md q-gutter-sm">
      <q-btn color="primary" style="width: 70%" to="/">
        <div class="ellipsis">Voltar</div>
      </q-btn>
    </div>
    <div class="flex justify-center q-pb-md" v-for="lna in linha" :key="lna.id">
      <q-card class="full-width" style="max-width: 95%">
        <q-card-section>
          <span class="text-h6 text-center">{{ lna.id }} {{ lna.name }}</span>
        </q-card-section>
        <q-card-section class="q-gutter-sm">
          <q-input filled readonly v-model="lna.station" />
          <q-input filled readonly v-model="lna.direction" />
          <q-input filled readonly v-model="lna.weekday" />
        </q-card-section>
        <q-card-section class="q-gutter-sm text-center">
          <q-chip outline v-for="val in lna.hours" :key="val" color="dark" icon="schedule">{{ val }}</q-chip>
        </q-card-section>
      </q-card>
    </div>
  </q-page>
</template>

<script setup>
import { useGeolocationStore } from "src/stores/geo-location";
import { useApi } from "src/composable/api";
import { useQuasar } from "quasar";
import { storeToRefs } from "pinia";
import { onMounted, ref } from "vue";

const props = defineProps({
  latitude: Number,
  longitude: Number,
});

let linha = ref([]);

const $q = useQuasar();
const api = useApi();
const geolocationStore = useGeolocationStore();
const { coords, isReady } = storeToRefs(geolocationStore);

onMounted(async () => {
  try {
    $q.loading.show({
      spinnerColor: "primary",
      message: "Carregando...",
      messageColor: "primary",
    });

    await isReady.value;

    const { data } = await api.get(
      `/geolocation/${coords.value.latitude}/${coords.value.longitude}`
    );

    linha.value = data;

    $q.loading.hide();
  } catch (error) {
    $q.loading.hide();
  }
});

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="css" scoped src="../css/app.css"></style>
