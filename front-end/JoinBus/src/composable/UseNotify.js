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
      actions: [{ icon: "close", color: "white" }],
      message: msg,
    });
  };

  const notifySuccsses = (msg) => {
    $q.notify({
      position: "top",
      progress: true,
      type: "positive",
      actions: [{ icon: "check", color: "white" }],
      message: msg,
    });
  };

  return {
    notifyWarning,
    notifyDanger,
    notifySuccsses,
  };
}
