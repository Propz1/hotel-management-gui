<template>
  <div :class="$style.buttonSolidRoundedAct">
    <div :class="$style.active">Active</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "ButtonSolidRoundedAct",
  });
</script>
<style module>
  .active {
    position: relative;
    line-height: 24px;
    font-weight: 500;
  }
  .buttonSolidRoundedAct {
    position: absolute;
    top: 120px;
    left: 201px;
    background-color: var(--primary-darker);
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
