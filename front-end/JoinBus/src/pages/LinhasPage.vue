<template>
  <div class="q-pa-sm">
    <q-card class="ow-card">
      <div class="q-pa-sm" v-for="line in linhas" :key="line.id">
        <div>
          <p>
            <strong>{{ line.id }}: {{ line.name }} - {{ line.weekday }}</strong>
          </p>
          <p>
            <strong>{{ line.direction }}</strong>
          </p>
          <tr>
            {{
              line.hours.join(" ")
            }}
          </tr>
        </div>
      </div>

      <div class="grid q-pa-md q-gutter-sm">
        <q-btn color="primary" style="width: 70%">
          <div class="ellipsis">
            <RouterLink class="ow-router-link" to="/">Voltar</RouterLink>
          </div>
        </q-btn>
      </div>
    </q-card>
  </div>
</template>

<script setup>
import { useApi } from "src/composable/api";
import { onMounted, ref } from "vue";

let linhas = ref([]);

const api = useApi();
onMounted(async () => {
  const { data } = await api.get("/linhas/0206");
  linhas.value = data;
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="css" scoped src="../css/app.css"></style>
