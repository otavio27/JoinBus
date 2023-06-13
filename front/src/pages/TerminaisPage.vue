<template>
  <div class="q-pa-sm">
    <q-page class="flex flex-center">
      <q-card class="q-ma-md full-width" style="max-width: 95%">
        <q-card-section class="grid q-pa-md q-gutter-sm">
          <q-btn color="primary" style="width: 70%" to="/"> Voltar </q-btn>
        </q-card-section>

        <q-separator />

        <q-card-section
          class="grid q-pa-md q-gutter-sm"
          v-for="station in stattions.name"
          :key="station.id"
        >
          <q-btn
            color="primary"
            style="width: 70%"
            :to="{ name: 'terminal', params: { terminal: station } }"
          >
            <div class="ellipsis">
              {{ station }}
            </div>
          </q-btn>
        </q-card-section>
      </q-card>
    </q-page>
  </div>
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
