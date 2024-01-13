<template>
  <q-page padding>
    <q-card class="no-shadow q-mt-md">
      <q-card-section>
        <q-img src="../assets/Joinville.jpg" />
      </q-card-section>
      <q-separator />

      <div v-if="linhas.length > 0">
        <q-card-section class="grid q-pa-sm q-gutter-sm">
          <q-btn unelevated rounded color="primary" style="width: 70%" @click="refreshPage">
            <div class="ellipsis">Voltar</div>
          </q-btn>
        </q-card-section>
        <q-separator />
        <div v-for="line in linhas" :key="line.id">
          <q-card-section class="grid q-pa-sm q-gutter-sm" v-if="line.name !== ''">
            <q-btn unelevated rounded color="primary" style="width: 70%" :to="{
              name: 'linha',
              params: { terminal: props.terminal, linha: line.id },
            }">
              <div class="ellipsis">{{ line.id }} - {{ line.name }}</div>
            </q-btn>
          </q-card-section>

          <q-card-section class="grid q-pa-sm q-gutter-sm" v-else>
            <q-icon name="warning" color="warning" size="4rem" />
            <span class="text-h6 text-center text-grey-7">{{ warning }}</span>
          </q-card-section>
        </div>
      </div>

      <q-card-section v-else>
        <div class="grid q-pa-md q-gutter-sm">
          <q-btn unelevated rounded color="primary" style="margin-top: 10%; width: 70%" :to="{ name: 'terminais' }">
            Terminais
          </q-btn>
        </div>

        <div class="grid q-pa-md q-gutter-sm">
          <q-btn unelevated rounded color="primary" style="width: 70%" :to="{ name: 'locale' }">
            Localização
          </q-btn>
        </div>

        <form @submit.prevent="onSearch" class="q-pa-md" style="margin-top: 15%">
          <!-- a simple text field watching for the enter key release -->
          <q-input filled color="teal" hint="Digite o nome da linha, ou número." v-model="filter" />

          <div class="row justify-end">
            <q-btn unelevated rounded type="submit" label="Buscar" class="q-mt-md" color="primary" />
          </div>
        </form>
      </q-card-section>
    </q-card>
  </q-page>
  <footer-component />
</template>

<script setup>
import FooterComponent from "src/components/global/FooterComponent.vue"
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useRoute } from "vue-router";
import { useApi } from "src/composable/api";
import useNotify from "src/composable/UseNotify";

const { notifyWarning } = useNotify();

const router = useRouter();
const route = useRoute();
const api = useApi();
const initial = route.query.filter || "";

const props = defineProps({
  name: {
    type: String,
    default: "index",
  },
  linhas: String,
});

const linhas = ref([]);
const filter = ref(initial);
const warning = "A Linha solicitada não foi encontrada!";

const updateRoute = () => {
  const { name, query: curQuery, params } = route;
  const query = {
    ...curQuery,
    filter: filter.value,
  };
  router.replace({
    name,
    query,
    params,
  });
};

const redirectToHours = () => {
  const ID = filter.value.replace(/[^0-9]/g);

  if (ID.length === 4) {
    router.push({
      name: "linha",
      params: { terminal: props.terminal, linha: filter.value },
    });
    return;
  }
};

const doSearch = async () => {
  redirectToHours();
  if (filter.value) {
    const { data } = await api.get("search/" + filter.value);
    linhas.value = data;
  } else {
    linhas.value = [];
  }
  onReset();
};

const onSearch = async () => {
  if (filter.value) {
    await doSearch();
    updateRoute();
  } else {
    notifyWarning("O campo de busca deve ser preenchido!");
  }
};

const refreshPage = () => {
  linhas.value = [];
  filter.value = initial;
};

onMounted(async () => {
  await doSearch();
});

const onReset = () => {
  filter.value = null;
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="css" scoped src="../css/app.css"></style>
