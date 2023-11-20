<template>
  <div :class="$style.fieldsSmallLabel" :style="fieldsSmallLabelStyle">
    <img :class="$style.inputIcon" alt="" src="/input.svg" />
    <div :class="$style.placeholderText">Placeholder Text</div>
    <div :class="$style.labelTitle">Label Title</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "FieldsSmallLabel",
    props: {
      fieldsSmallLabelWidth: { type: [Object, Array, String, Number, Boolean] },
      fieldsSmallLabelBorder: { type: String },
      fieldsSmallLabelBackgroundColor: { type: String },
      fieldsSmallLabelPosition: { type: String },
      fieldsSmallLabelTop: { type: [Object, Array, String, Number, Boolean] },
      fieldsSmallLabelLeft: { type: [Object, Array, String, Number, Boolean] },
    },
    computed: {
      fieldsSmallLabelStyle() {
        return {
          width: this.fieldsSmallLabelWidth,
          border: this.fieldsSmallLabelBorder,
          backgroundColor: this.fieldsSmallLabelBackgroundColor,
          position: this.fieldsSmallLabelPosition,
          top: this.fieldsSmallLabelTop,
          left: this.fieldsSmallLabelLeft,
        };
      },
    },
  });
</script>
<style module>
  .inputIcon {
    position: absolute;
    height: 66.67%;
    width: 100%;
    top: 33.33%;
    right: 0%;
    bottom: 0%;
    left: 0%;
    border-radius: var(--br-5xs);
    max-width: 100%;
    overflow: hidden;
    max-height: 100%;
    opacity: 0.16;
  }
  .placeholderText {
    position: absolute;
    top: 46.67%;
    left: 3.03%;
    line-height: 150%;
  }
  .labelTitle {
    position: absolute;
    top: 0%;
    left: 0%;
    font-size: var(--text-xs-regular-size);
    letter-spacing: 0.2px;
    line-height: 150%;
    font-weight: 500;
    color: var(--dark-dark-300);
  }
  .fieldsSmallLabel {
    width: 264px;
    height: 60px;
    text-align: left;
    font-size: var(--label-text-large-size);
    color: var(--dark-dark-200);
    font-family: var(--body-body-16);
  }
</style>
