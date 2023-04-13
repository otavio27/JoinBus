<template>
  <div class="q-pa-sm">
    <q-card class="ow-card">
      <div class="grid q-pa-md q-gutter-sm">
        <q-btn color="primary" style="width: 70%">
          <div class="ellipsis">
            <RouterLink class="ow-router-link" to="/linhas">Linhas</RouterLink>
          </div>
        </q-btn>
      </div>

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
            <RouterLink class="ow-router-link" to="/favoritos"
              >Favoritos</RouterLink
            >
          </div>
        </q-btn>
      </div>

      <div class="grid q-pa-md q-gutter-sm">
        <q-btn color="primary" style="width: 70%">
          <div class="ellipsis">Localização</div>
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
          <q-btn
            type="submit"
            label="Buscar"
            class="q-mt-md"
            color="primary"
            @click="onReset()"
          >
            <template v-slot:loading>
              <q-spinner-facebook />
            </template>
          </q-btn>
        </div>
      </form>
    </q-card>
  </div>
</template>

<script setup>
import { ref } from "vue";

const props = defineProps({
  name: {
    type: String,
    default: "Index",
  },
});

const text = ref("");

const sendSubmit = () => {};

const onReset = () => {
  text.value = null;
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="css" scoped src="../css/app.css"></style>
