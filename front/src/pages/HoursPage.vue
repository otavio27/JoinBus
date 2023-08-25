<template>
  <q-page class="q-pa-sm">
    <div class="row q-col-gutter-lg flex-center">
      <div class="col-lg-4 col-md-4 col-xs-12 col-sm-12">
        <q-card class="no-shadow" bordered>
          <div class="grid q-pa-md q-gutter-sm">
            <q-btn v-if="!terminal" color="primary" style="width: 70%" :to="{ name: 'index' }">
              <div class="ellipsis">Voltar</div>
            </q-btn>

            <q-btn v-else color="primary" style="width: 70%"
              :to="{ name: 'terminal', params: { terminal: props.terminal } }">
              <div class="ellipsis">Voltar</div>
            </q-btn>
          </div>
        </q-card>
        <q-separator />

        <div class="q-pb-sm" v-for="hour in hours" :key="hour.id">
          <q-card class="no-shadow text-center" bordered v-if="hour.hours">
            <q-card-section>
              <span class="text-h6">{{ hour.id }} {{ hour.name }}</span>
            </q-card-section>
            <q-card-section class="q-gutter-sm flex-center">
              <q-input filled readonly v-model="hour.station" />
              <q-input filled readonly v-model="hour.direction" />
              <q-input filled readonly v-model="hour.weekday" />
            </q-card-section>
            <q-card-section class="q-gutter-sm text-center">
              <q-chip outline v-for="val in hour.hours" :key="val" color="dark" icon="schedule">{{ val }}</q-chip>
            </q-card-section>
          </q-card>
          <q-card class="no-shadow" bordered v-else>
            <q-card-section>
              <span class="text-h6 text-center">{{ hours.id }} {{ hours.name }}</span>
            </q-card-section>
            <q-card-section class="q-gutter-sm">
              <q-input filled readonly v-model="warning" />
            </q-card-section>
            <q-card-section class="q-gutter-sm text-center"> </q-card-section>
          </q-card>
        </div>
      </div>
    </div>
  </q-page>
</template>

<script setup>
import { useApi } from "src/composable/api";
import { useQuasar } from "quasar";
import { onMounted, ref } from "vue";

const $q = useQuasar();

let hours = ref([]);
const warning = "Esta linha não apresenta horarios para hoje.";

const props = defineProps({
  terminal: String,
  linha: String,
});

const api = useApi();
onMounted(async () => {
  $q.loading.show({
    spinnerColor: "primary",
    message: "Carregando...",
    messageColor: "primary",
  });
  const { data } = await api.get("/linhas/" + props.linha);
  hours.value = data;
  $q.loading.hide();
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="css" scoped src="../css/app.css"></style>
