<template>
  <div :class="$style.atomdate" :style="atomDateStyle">
    <div :class="$style.div" :style="div2Style">{{carModelNumber}}</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "AtomDate",
    props: {
      carModelNumber: { type: String },
      atomDatePosition: { type: String },
      divColor: { type: String },
      divLeft: { type: [Object, Array, String, Number, Boolean] },
      divDisplay: { type: String },
    },
    computed: {
      atomDateStyle() {
        return {
          position: this.atomDatePosition,
        };
      },
      div2Style() {
        return {
          color: this.divColor,
          left: this.divLeft,
          display: this.divDisplay,
        };
      },
    },
  });
</script>
<style module>
  .div {
    position: absolute;
    top: calc(50% - 12px);
    left: calc(50% - 9px);
    line-height: 150%;
  }
  .atomdate {
    width: 24px;
    height: 24px;
    text-align: center;
    font-size: var(--label-text-large-size);
    color: var(--base-pure-black);
    font-family: var(--body-body-16);
  }
</style>
