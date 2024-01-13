<template>
  <q-page padding>
    <div class="text-h3 text-center text-grey-8 q-pa-lg">
      <span>Terminais</span>
    </div>

    <div class="row q-col-gutter-lg flex-center">
      <div class="col-lg-4 col-md-4 col-xs-12 col-sm-12">
        <q-card class="no-shadow">
          <q-card-section class="grid q-pa-sm q-gutter-sm" v-for="station in  stattions.name " :key="station.id">
            <q-btn unelevated rounded color="primary" :label="station" style="width: 70%"
              :to="{ name: 'terminal', params: { terminal: station } }" />
          </q-card-section>

          <q-card-section class="grid q-pa-md q-gutter-sm">
            <q-btn unelevated rounded color="primary" style="width: 70%" to="/"> Voltar </q-btn>
          </q-card-section>
        </q-card>
      </div>
    </div>
  </q-page>
  <footer-component />
</template>
<script setup>
import FooterComponent from "src/components/global/FooterComponent.vue";
import { useApi } from "src/composable/api";
import { onMounted, ref } from "vue";

let stattions = ref([]);

const api = useApi();
onMounted(async () => {
  const { data } = await api.get("/terminais");
  stattions.value = data;
});
</script>
