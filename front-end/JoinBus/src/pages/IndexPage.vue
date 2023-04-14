<template>
  <div class="q-pa-sm">
    <q-page class="flex flex-center">
      <img src="../assets/favicon-196.png" />
      <q-card class="q-ma-md full-width" style="max-width: 100%">
        <div class="grid q-pa-md q-gutter-sm">
          <q-btn color="primary" style="width: 70%">
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
          @submit.prevent="sendSubmit"
          class="q-pa-md"
          style="margin-top: 15%"
        >
          <!-- a simple text field watching for the enter key release -->
          <q-input
            filled
            color="teal"
            hint="Digite o nome da linha, ou o número."
            v-model="text"
          />

          <div class="row justify-end">
            <q-btn type="submit" label="Buscar" class="q-mt-md" color="primary">
              <template v-slot:loading>
                <q-spinner-facebook />
              </template>
            </q-btn>
          </div>
        </form>
      </q-card>
    </q-page>
  </div>
</template>

<script setup>
import { useApi } from "src/composable/api";
import { ref } from "vue";

const props = defineProps({
  name: {
    type: String,
    default: "index",
  },
  terminal: String,
  linhas: String,
});

const api = useApi();
const text = ref("");

const sendSubmit = async () => {
  const { data } = await api.get(`/search/${text.value}`);
  onReset();
};

const onReset = () => {
  text.value = null;
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="css" scoped src="../css/app.css"></style>
