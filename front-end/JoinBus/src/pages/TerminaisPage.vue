<template>
  <div class="q-pa-sm">
    <q-card class="ow-card">
      <div class="grid q-pa-md q-gutter-sm">
        <q-btn color="primary" style="width: 70%">
          <div class="ellipsis">
            <RouterLink class="ow-router-link" to="/">Voltar</RouterLink>
          </div>
        </q-btn>
      </div>
      <div
        class="grid q-pa-md q-gutter-sm"
        v-for="station in stattions.name"
        :key="station.id"
      >
        <q-btn
          color="primary"
          style="width: 70%"
          :to="{ name: 'terminal', params: { terminal: station } }"
        >
          <div class="ellipsis">
            {{ station }}
          </div>
        </q-btn>
      </div>
    </q-card>
  </div>
</template>
<script setup>
import { useApi } from "src/composable/api";
import { onMounted, ref } from "vue";

let stattions = ref([]);

const api = useApi();
onMounted(async () => {
  const { data } = await api.get("/terminais");
  stattions.value = data;
});
</script>
