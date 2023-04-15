<template>
  <q-page class="flex flex-center">
    <q-card class="q-ma-md full-width" style="max-width: 95%">
      <q-card-section class="text-center text-h6">
        <span>{{ terminal }}</span>
      </q-card-section>
      <q-separator></q-separator>
      <q-card-section>
        <div class="col col-12">
          <q-btn class="full-width" color="primary" to="/">
            <div class="ellipsis">Voltar</div>
          </q-btn>
        </div>
      </q-card-section>
      <q-separator></q-separator>
      <q-card-section class="row q-col-gutter-y-md">
        <div class="col col-12" v-for="line in linhas" :key="line.id">
          <q-btn
            class="full-width"
            color="primary"
            :to="{
              name: 'linha',
              params: { terminal: props.terminal, linha: line.id },
            }"
          >
            <div class="ellipsis">
              {{ line.name }}
            </div>
          </q-btn>
        </div>
      </q-card-section>
      <q-separator></q-separator>
    </q-card>
  </q-page>
</template>

<script setup>
import { useApi } from "src/composable/api";
import { useQuasar } from "quasar";
import { onMounted, ref } from "vue";

let linhas = ref([]);
const props = defineProps({
  index: String,
  texto: {
    type: String,
    required: true,
  },
});

console.log("Search text received: ", props.texto);

const $q = useQuasar();

const api = useApi();
onMounted(async () => {
  $q.loading.show({
    spinnerColor: "primary",
    message: "Carregando...",
    messageColor: "amber",
  });
  const { data } = await api.get(`/search/${props.texto}`);
  linhas.value = data;
  $q.loading.hide();
});
</script>
