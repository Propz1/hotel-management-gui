<template>
  <div :class="$style.buttonOutlinedSquareH">
    <div :class="$style.hover">Hover</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "ButtonOutlinedSquareH",
  });
</script>
<style module>
  .hover {
    position: relative;
    line-height: 24px;
    font-weight: 500;
  }
  .buttonOutlinedSquareH {
    position: absolute;
    top: 60px;
    left: 105px;
    border: 2px solid var(--primary-lighter);
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
    color: var(--primary-lighter);
    font-family: var(--body-body-16);
  }
</style>
