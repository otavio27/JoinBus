<template>
  <q-page padding>
    <div class="row q-col-gutter-lg flex-center">
      <div class="col-lg-4 col-md-4 col-xs-12 col-sm-12">
        <div class="q-pb-sm" v-for="hour in hours" :key="hour.id">
          <q-card class="no-shadow text-center" v-if="hour.hours">
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
          <q-card class="no-shadow" v-else>
            <q-card-section>
              <span class="text-h6 text-center">{{ hours.id }} {{ hours.name }}</span>
            </q-card-section>
            <q-card-section class="q-gutter-sm">
              <q-input filled readonly v-model="warning" />
            </q-card-section>
            <q-card-section class="q-gutter-sm text-center"> </q-card-section>
          </q-card>
        </div>

        <q-card class="no-shadow bg-transparent">
          <div class="grid q-pa-md q-gutter-sm">
            <q-btn v-if="!terminal" color="primary" style="width: 70%" :to="{ name: 'index' }">
              <div class="ellipsis">Voltar</div>
            </q-btn>

            <q-btn unelevated rounded v-else color="primary" label="Voltar" style="width: 70%"
              :to="{ name: 'terminal', params: { terminal: props.terminal } }" />
          </div>
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

let hours = ref([]);
const warning = "Esta linha não apresenta horarios para hoje.";

const props = defineProps({
  terminal: String,
  linha: String,
});

const api = useApi();
onMounted(async () => {
  const { data } = await api.get("/linhas/" + props.linha);
  hours.value = data;
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="css" scoped src="../css/app.css"></style>
