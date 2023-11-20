<template>
  <div :class="$style.buttonSolidRoundedHov">
    <div :class="$style.hover">Hover</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "ButtonSolidRoundedHov",
  });
</script>
<style module>
  .hover {
    position: relative;
    line-height: 24px;
    font-weight: 500;
  }
  .buttonSolidRoundedHov {
    position: absolute;
    top: 120px;
    left: 105px;
    background-color: var(--primary-lighter);
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
