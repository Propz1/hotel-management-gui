<template>
  <router-view />
 <!-- <Toast position="bottom-right" group="br" /> -->
  <Toast />
</template>

<script>
import Toast from "primevue/toast";

export default {
  components: {
    Toast: Toast,
  },
}
</script>
