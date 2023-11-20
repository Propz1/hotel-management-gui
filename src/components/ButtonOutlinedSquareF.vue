<template>
  <div :class="$style.buttonOutlinedSquareF">
    <div :class="$style.focus">Focus</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "ButtonOutlinedSquareF",
  });
</script>
<style module>
  .focus {
    position: relative;
    line-height: 24px;
    font-weight: 500;
  }
  .buttonOutlinedSquareF {
    position: absolute;
    top: 58px;
    left: 413px;
    border: 2px solid var(--primary-darker);
    box-sizing: border-box;
    height: 36px;
    overflow: hidden;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    padding: var(--padding-5xs) var(--padding-base);
    opacity: 0.64;
    text-align: center;
    font-size: var(--label-large-label-14-size);
    color: var(--primary-darker);
    font-family: var(--body-body-16);
  }
</style>
