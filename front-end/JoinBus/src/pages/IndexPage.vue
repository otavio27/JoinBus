<template>
  <div class="q-pa-sm">
    <q-page class="flex flex-center">
      <img src="../assets/joinbus.png" style="max-width: 50%" />

      <q-card
        class="q-ma-md full-width"
        style="max-width: 95%"
        v-if="linhas.length > 0"
      >
        <q-card-section class="row q-col-gutter-y-md">
          <div class="col col-12" v-for="line in linhas" :key="line.id">
            <q-btn
              class="full-width"
              color="primary"
              style="width: 70%"
              :to="{
                name: 'linha',
                params: { linha: line.id },
              }"
            >
              <div class="ellipsis">
                {{ line.name }}
              </div>
            </q-btn>
          </div>
        </q-card-section>
      </q-card>

      <q-card class="q-ma-md full-width" style="max-width: 95%" v-else>
        <div class="grid q-pa-md q-gutter-sm">
          <q-btn color="primary" style="margin-top: 10%; width: 70%">
            <div class="ellipsis">
              <RouterLink class="ow-router-link" :to="{ name: 'terminais' }"
                >Terminais</RouterLink
              >
            </div>
          </q-btn>
        </div>

        <div class="grid q-pa-md q-gutter-sm">
          <q-btn color="primary" style="width: 70%">
            <div class="ellipsis">
              <RouterLink class="ow-router-link" :to="{ name: 'locale' }"
                >Localização</RouterLink
              >
            </div>
          </q-btn>
        </div>

        <form
          @submit.prevent="onSearch"
          class="q-pa-md"
          style="margin-top: 15%"
        >
          <!-- a simple text field watching for the enter key release -->
          <q-input
            filled
            color="teal"
            hint="Digite o nome da linha, ou número."
            v-model="filter"
          />

          <div class="row justify-end">
            <q-btn type="submit" label="Buscar" class="q-mt-md" color="primary">
            </q-btn>
          </div>
        </form>
      </q-card>
    </q-page>
  </div>
</template>

<script setup>
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

onMounted(async () => {
  await doSearch();
});

const onReset = () => {
  filter.value = null;
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="css" scoped src="../css/app.css"></style>
