<template>
  <div :class="$style.buttonSolidSquareDefa">
    <div :class="$style.default">Default</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "ButtonSolidSquareDefa",
  });
</script>
<style module>
  .default {
    position: relative;
    line-height: 24px;
    font-weight: 500;
  }
  .buttonSolidSquareDefa {
    position: absolute;
    top: 0px;
    left: 0px;
    background-color: var(--primary-main-color);
    height: 32px;
    overflow: hidden;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    padding: var(--padding-5xs) var(--padding-base);
    box-sizing: border-box;
    text-align: center;
    font-size: var(--label-large-label-14-size);
    color: var(--light-light-100);
    font-family: var(--body-body-16);
  }
</style>
