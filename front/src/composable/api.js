// ts:
// import { inject, Injection } from 'vue'
// import axios, { AxiosInstance } from 'axios'
import { inject } from "vue";

// ts: export const apiKey: Injection<AxiosInstance> = Symbol('api-key')
export const apiKey = "api-key";
export function useApi() {
  const api = inject(apiKey);
  if (!api) {
    throw "api not injected";
  }
  return api;
}
