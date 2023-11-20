<template>
  <div
    :class="$style.typemacHorizontalfalsePo"
    :style="typeMacHorizontalFalsePoStyle"
  >
    <div :class="$style.bar" :style="barStyle" />
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "TypeMacHorizontalFalsePo",
    props: {
      propLeft: { type: [Object, Array, String, Number, Boolean] },
      propTop: { type: [Object, Array, String, Number, Boolean] },
      propBorderRight: { type: String },
      propBorderLeft: { type: String },
      propWidth: { type: [Object, Array, String, Number, Boolean] },
      propHeight: { type: [Object, Array, String, Number, Boolean] },
      propBorderTop: { type: String },
      propBorderBottom: { type: String },
      propBottom: { type: [Object, Array, String, Number, Boolean] },
      propTop1: { type: [Object, Array, String, Number, Boolean] },
      propLeft1: { type: [Object, Array, String, Number, Boolean] },
      propTransform: { type: String },
      propTransformOrigin: { type: String },
      propRight: { type: [Object, Array, String, Number, Boolean] },
    },
    computed: {
      typeMacHorizontalFalsePoStyle() {
        return {
          left: this.propLeft,
          top: this.propTop,
          borderRight: this.propBorderRight,
          borderLeft: this.propBorderLeft,
          width: this.propWidth,
          height: this.propHeight,
          borderTop: this.propBorderTop,
          borderBottom: this.propBorderBottom,
        };
      },
      barStyle() {
        return {
          bottom: this.propBottom,
          top: this.propTop1,
          left: this.propLeft1,
          transform: this.propTransform,
          transformOrigin: this.propTransformOrigin,
          right: this.propRight,
        };
      },
    },
  });
</script>
<style module>
  .bar {
    position: absolute;
    bottom: 2px;
    left: calc(50% - 3px);
    border-radius: var(--br-81xl);
    background-color: var(--color-gray-500);
    width: 6px;
    height: 25px;
  }
  .typemacHorizontalfalsePo {
    position: absolute;
    top: 16px;
    left: 337px;
    background-color: var(--color-gray-100);
    border-right: 1px solid var(--color-gainsboro-100);
    border-left: 1px solid var(--color-gainsboro-100);
    box-sizing: border-box;
    width: 12px;
    height: 150px;
  }
</style>
