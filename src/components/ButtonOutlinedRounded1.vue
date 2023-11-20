<template>
  <div :class="$style.buttonOutlinedRounded">
    <div :class="$style.disabled">Disabled</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "ButtonOutlinedRounded1",
  });
</script>
<style module>
  .disabled {
    position: relative;
    line-height: 24px;
    font-weight: 500;
  }
  .buttonOutlinedRounded {
    position: absolute;
    top: 180px;
    left: 300px;
    border: 2px solid var(--light-light-500);
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
    color: var(--light-light-500);
    font-family: var(--body-body-16);
  }
</style>
