<template>
  <div class="q-pa-sm">
    <q-card class="ow-card">
      <div class="grid q-pa-md q-gutter-sm">
        <q-btn color="primary" style="width: 70%" @click="getLines">
          <div class="ellipsis">Linhas</div>
        </q-btn>
      </div>

      <div class="grid q-pa-md q-gutter-sm">
        <q-btn color="primary" style="width: 70%" @click="getTerminals">
          <div class="ellipsis">Terminais</div>
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
        <q-btn color="primary" style="width: 70%" @click="getGeolocation">
          <div class="ellipsis">Localização</div>
        </q-btn>
      </div>

      <div class="row justify-center">
        <div class="col-10 text-center">
          <p>Latitude: {{ latitude }}</p>
          <p>Longitude: {{ longitude }}</p>
        </div>
      </div>

      <form
        @submit.prevent="simulateSubmit"
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
            :loading="submitting"
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
import { useQuasar } from "quasar";
import axios from "axios";
import useNotify from "src/composable/UseNotify";

const { notifyDanger, notifySuccsses } = useNotify();
const $q = useQuasar();
const text = ref("");
const submitting = ref(false);

let latitude = ref(null);
let longitude = ref(null);

const getGeolocation = () => {
  if (navigator.geolocation) {
    $q.loading.show();
    navigator.geolocation.getCurrentPosition(
      apiSearchByLocation,
      errorPosition
    );
  } else {
    errorPosition();
  }
};

const apiSearchByLocation = (position) => {
  var lat = position.coords.latitude;
  var lng = position.coords.longitude;
  axios
    .get(`https://owtechsystems.com/api/geolocation?lat=${lat}&lng=${lng}`, {
      headers: {
        "Access-Control-Allow-Origin": "*",
      },
    })
    .then(function (response) {
      $q.loading.hide();
      console.log(response.data);
    })
    .catch(function (error) {
      $q.loading.hide();
      errorPosition(error);
    });
};

const getTerminals = () => {
  axios
    .get("https://owtechsystems.com/api/terminais")
    .then(function (response) {
      $q.loading.hide();
      console.log(response.data);
    })
    .catch(function (error) {
      $q.loading.hide();
      errorPosition(error);
    });
};

const getLines = () => {
  axios
    .get("https://owtechsystems.com/api/linhas/0300")
    .then(function (response) {
      $q.loading.hide();
      console.log(response.data);
    })
    .catch(function (error) {
      $q.loading.hide();
      errorPosition(error);
    });
};

const errorPosition = (error) => {
  notifyDanger(`Não foi possível enviar sua localização! ${error}`),
    $q.loading.hide();
};

const simulateSubmit = () => {
  submitting.value = true;

  setTimeout(() => {
    submitting.value = false;
    notifyDanger("Não foi possível listar a requsição!");
  }, 3000);
};

const onReset = () => {
  text.value = null;
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="css" scoped src="../css/app.css"></style>
