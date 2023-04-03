import { useQuasar } from "quasar";

export default function useNotify() {
  const $q = useQuasar();

  const notifyWarning = (msg) => {
    $q.notify({
      position: "top",
      progress: true,
      type: "warning",
      message: msg,
    });
  };

  const notifyDanger = (msg) => {
    $q.notify({
      position: "top",
      progress: true,
      type: "negative",
      message: msg,
    });
  };

  const notifySuccsses = (msg) => {
    $q.notify({
      position: "top",
      progress: true,
      type: "positive",
      message: msg,
    });
  };

  return {
    notifyWarning,
    notifyDanger,
    notifySuccsses,
  };
}
