<template>
  <div :class="$style.atomdayweek">
    <div :class="$style.div">ПН</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "AtomDayWeek",
  });
</script>
<style module>
  .div {
    position: absolute;
    top: calc(50% - 7px);
    left: calc(50% - 7px);
    line-height: 150%;
  }
  .atomdayweek {
    width: 24px;
    height: 24px;
    text-align: center;
    font-size: var(--font-size-3xs);
    color: var(--base-pure-black);
    font-family: var(--body-body-16);
  }
</style>
