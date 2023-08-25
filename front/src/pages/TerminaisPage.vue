<template>
  <q-page class="q-pa-sm">
    <div class="row q-col-gutter-lg flex-center">
      <div class="col-lg-4 col-md-4 col-xs-12 col-sm-12">
        <q-card class="no-shadow" bordered>
          <q-card-section class="grid q-pa-md q-gutter-sm">
            <q-btn color="primary" style="width: 70%" to="/"> Voltar </q-btn>
          </q-card-section>

          <q-separator />

          <q-card-section class="grid q-pa-sm q-gutter-sm" v-for="station in stattions.name" :key="station.id">
            <q-btn color="primary" style="width: 70%" :to="{ name: 'terminal', params: { terminal: station } }">
              <div class="ellipsis">
                {{ station }}
              </div>
            </q-btn>
          </q-card-section>
        </q-card>
      </div>
    </div>
  </q-page>
</template>
<script setup>
import { useApi } from "src/composable/api";
import { useQuasar } from "quasar";
import { onMounted, ref } from "vue";

const $q = useQuasar();

let stattions = ref([]);

const api = useApi();
onMounted(async () => {
  $q.loading.show({
    spinnerColor: "primary",
    message: "Carregando...",
    messageColor: "primary",
  });
  const { data } = await api.get("/terminais");
  stattions.value = data;
  $q.loading.hide();
});
</script>
