<template>
  <div
    :class="$style.buttonOutlinedSquareD"
    :style="buttonOutlinedSquareDStyle"
  >
    <div :class="$style.default" :style="defaultStyle">{{buttonText}}</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "ButtonOutlinedSquareD1",
    props: {
      buttonText: { type: String },
      propTop: { type: [Object, Array, String, Number, Boolean] },
      propLeft: { type: [Object, Array, String, Number, Boolean] },
      propCursor: { type: String },
      propBackgroundColor: { type: String },
      propWidth: { type: [Object, Array, String, Number, Boolean] },
      propDisplay: { type: String },
    },
    computed: {
      buttonOutlinedSquareDStyle() {
        return {
          top: this.propTop,
          left: this.propLeft,
          cursor: this.propCursor,
          backgroundColor: this.propBackgroundColor,
          width: this.propWidth,
        };
      },
      defaultStyle() {
        return {
          display: this.propDisplay,
        };
      },
    },
  });
</script>
<style module>
  .default {
    position: relative;
    line-height: 24px;
    font-weight: 500;
  }
  .buttonOutlinedSquareD {
    position: absolute;
    top: 60px;
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
