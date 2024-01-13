<template>
  <q-page padding>
    <div class="text-h3 text-center text-grey-8 q-pa-xl">
      <span>{{ terminal }}</span>
    </div>

    <div class="row q-col-gutter-lg flex-center">
      <div class="col-lg-4 col-md-4 col-xs-12 col-sm-12">
        <q-card class="no-shadow">
          <q-card-section class="grid q-pa-sm q-gutter-sm" v-for="line in linhas" :key="line.id">
            <q-btn unelevated rounded color="primary" style="width: 70%" :to="{
              name: 'linha',
              params: { terminal: props.terminal, linha: line.id },
            }">
              <div class="ellipsis">{{ line.id }} - {{ line.name }}</div>
            </q-btn>
          </q-card-section>

          <q-card-section class="grid q-pa-sm q-gutter-sm">
            <q-btn unelevated rounded color="primary" label="Voltar" style="width: 70%" :to="{ name: 'terminais' }" />
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

let linhas = ref([]);
const props = defineProps({
  terminal: String,
});

const api = useApi();
onMounted(async () => {
  const { data } = await api.get("/routes/" + props.terminal);
  const res = data?.[0];
  const _linhas = res.id?.map((id, index) => ({
    id,
    name: res.name[index],
  }));
  linhas.value = _linhas;
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="css" scoped src="../css/app.css"></style>
