<template>
  <img
    :class="$style.typewindowsHorizontalfalseIcon"
    alt=""
    :src="dimensionCode"
    :style="typeWindowsHorizontalFalseIconStyle"
  />
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "TypeWindowsHorizontalFalse1",
    props: {
      dimensionCode: { type: String },
      propLeft: { type: [Object, Array, String, Number, Boolean] },
      propWidth: { type: [Object, Array, String, Number, Boolean] },
      propHeight: { type: [Object, Array, String, Number, Boolean] },
      propTop: { type: [Object, Array, String, Number, Boolean] },
    },
    computed: {
      typeWindowsHorizontalFalseIconStyle() {
        return {
          left: this.propLeft,
          width: this.propWidth,
          height: this.propHeight,
          top: this.propTop,
        };
      },
    },
  });
</script>
<style module>
  .typewindowsHorizontalfalseIcon {
    position: absolute;
    top: 16px;
    left: 41px;
    width: 17px;
    height: 150px;
  }
</style>
