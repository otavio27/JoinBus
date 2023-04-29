<template>
  <q-page class="flex flex-center">
    <q-card class="q-ma-md full-width" style="max-width: 95%">
      <q-card-section class="text-center text-h6">
        <span>{{ terminal }}</span>
      </q-card-section>
      <q-separator></q-separator>
      <q-card-section>
        <div class="col col-12">
          <q-btn class="full-width" color="primary" :to="{ name: 'terminais' }">
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
            <div class="ellipsis">{{ line.id }} - {{ line.name }}</div>
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
  terminal: String,
});

const $q = useQuasar();

const api = useApi();
onMounted(async () => {
  $q.loading.show({
    spinnerColor: "primary",
    message: "Carregando...",
    messageColor: "primary",
  });
  const { data } = await api.get("/routes/" + props.terminal);
  const res = data?.[0];
  const _linhas = res.id?.map((id, index) => ({
    id,
    name: res.name[index],
  }));

  linhas.value = _linhas;
  $q.loading.hide();
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="css" scoped src="../css/app.css"></style>
