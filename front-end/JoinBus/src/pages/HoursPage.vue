<template>
  <q-page class="q-pa-sm">
    <div class="grid q-pa-md q-gutter-sm">
      <q-btn
        color="primary"
        style="width: 220px"
        :to="{ name: 'terminal', params: { terminal: props.terminal } }"
      >
        <div class="ellipsis">Voltar</div>
      </q-btn>
    </div>
    <div
      class="flex justify-center q-pb-md"
      v-for="hour in hours"
      :key="hour.id"
    >
      <q-card class="full-width" style="max-width: 720px">
        <q-card-section>
          <span class="text-h6 text-center">{{ hour.name }}</span>
        </q-card-section>
        <q-card-section class="q-gutter-sm">
          <q-input filled readonly label="Estação" v-model="hour.station" />
          <q-input filled readonly label="Direceção" v-model="hour.direction" />
          <q-input
            filled
            readonly
            label="Dia da Semana"
            v-model="hour.weekday"
          />
        </q-card-section>
        <q-card-section class="q-gutter-sm">
          <q-chip
            outline
            v-for="val in hour.hours"
            :key="val"
            color="primary"
            icon="schedule"
            >{{ val }}</q-chip
          >
        </q-card-section>
      </q-card>
    </div>
  </q-page>
</template>

<script setup>
import { useApi } from "src/composable/api";
import { onMounted, ref } from "vue";

let hours = ref([]);
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
