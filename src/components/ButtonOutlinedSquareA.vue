<template>
  <div :class="$style.buttonOutlinedSquareA">
    <div :class="$style.active">Active</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "ButtonOutlinedSquareA",
  });
</script>
<style module>
  .active {
    position: relative;
    line-height: 24px;
    font-weight: 500;
  }
  .buttonOutlinedSquareA {
    position: absolute;
    top: 60px;
    left: 201px;
    border: 2px solid var(--primary-darker);
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
    color: var(--primary-darker);
    font-family: var(--body-body-16);
  }
</style>
