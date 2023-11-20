<template>
  <b :class="$style.shadow">shadow</b>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "ShadowText",
  });
</script>
<style module>
  .shadow {
    font-size: var(--text-xs-regular-size);
    line-height: 16px;
    font-family: var(--text-xs-regular);
    color: var(--base-pure-black);
    text-align: left;
  }
</style>
