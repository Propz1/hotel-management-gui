<template>
  <div :class="$style.buttonSolidSquareFocu">
    <div :class="$style.focus">Focus</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "ButtonSolidSquareFocu",
  });
</script>
<style module>
  .focus {
    position: relative;
    line-height: 24px;
    font-weight: 500;
  }
  .buttonSolidSquareFocu {
    position: absolute;
    top: -2px;
    left: 413px;
    background-color: var(--primary-darker);
    border: 2px solid var(--primary-lighter);
    box-sizing: border-box;
    height: 36px;
    overflow: hidden;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    padding: var(--padding-5xs) var(--padding-base);
    text-align: center;
    font-size: var(--label-large-label-14-size);
    color: var(--light-light-100);
    font-family: var(--body-body-16);
  }
</style>
