<template>
  <div :class="$style.buttonOutlinedRounded">
    <div :class="$style.default">Default</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "ButtonOutlinedRounded4",
  });
</script>
<style module>
  .default {
    position: relative;
    line-height: 24px;
    font-weight: 500;
  }
  .buttonOutlinedRounded {
    position: absolute;
    top: 180px;
    left: 0px;
    border: 2px solid var(--primary-main-color);
    box-sizing: border-box;
    height: 32px;
    overflow: hidden;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    padding: var(--padding-5xs) var(--padding-base);
    text-align: center;
    font-size: var(--label-large-label-14-size);
    color: var(--primary-main-color);
    font-family: var(--body-body-16);
  }
</style>
