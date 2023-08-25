<template>
  <q-page class="q-pa-sm">
    <div class="row q-col-gutter-lg flex-center">
      <div class="col-lg-4 col-md-4 col-xs-12 col-sm-12">
        <q-card class="no-shadow" bordered>
          <q-card-section class="text-center text-h6">
            <span>{{ terminal }}</span>
          </q-card-section>
          <q-separator></q-separator>
          <q-card-section class="grid q-pa-md q-gutter-sm">
            <q-btn color="primary" style="width: 70%" :to="{ name: 'terminais' }">
              <div class="ellipsis">Voltar</div>
            </q-btn>
          </q-card-section>
          <q-separator></q-separator>
          <q-card-section class="grid q-pa-md q-gutter-sm" v-for="line in linhas" :key="line.id">
            <q-btn color="primary" style="width: 70%" :to="{
              name: 'linha',
              params: { terminal: props.terminal, linha: line.id },
            }">
              <div class="ellipsis">{{ line.id }} - {{ line.name }}</div>
            </q-btn>
          </q-card-section>
          <q-separator></q-separator>
        </q-card>
      </div>
    </div>
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
