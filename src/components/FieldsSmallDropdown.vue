<template>
  <div :class="$style.fieldsSmallDropdown" :style="fieldsSmallDropdownStyle">
    <img :class="$style.inputIcon" alt="" src="/input.svg" />
    <div :class="$style.placeholderText">Placeholder Text</div>
    <div :class="$style.dropdownIcon">
      <div :class="$style.background" />
      <img :class="$style.icon" alt="" src="/icon.svg" />
    </div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "FieldsSmallDropdown",
    props: {
      fieldsSmallDropdownWidth: {
        type: [Object, Array, String, Number, Boolean],
      },
      fieldsSmallDropdownPosition: { type: String },
      fieldsSmallDropdownTop: {
        type: [Object, Array, String, Number, Boolean],
      },
      fieldsSmallDropdownLeft: {
        type: [Object, Array, String, Number, Boolean],
      },
    },
    computed: {
      fieldsSmallDropdownStyle() {
        return {
          width: this.fieldsSmallDropdownWidth,
          position: this.fieldsSmallDropdownPosition,
          top: this.fieldsSmallDropdownTop,
          left: this.fieldsSmallDropdownLeft,
        };
      },
    },
  });
</script>
<style module>
  .inputIcon {
    position: absolute;
    height: 100%;
    width: 100%;
    top: 0%;
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
    top: 20%;
    left: 3.03%;
    line-height: 150%;
  }
  .background {
    position: absolute;
    height: 100%;
    width: 100%;
    top: 0px;
    right: 0px;
    bottom: 0px;
    left: 0px;
    border-radius: var(--br-9xs);
  }
  .icon {
    position: absolute;
    height: 21.88%;
    width: 37.5%;
    top: 43.75%;
    right: 27.5%;
    bottom: 34.38%;
    left: 35%;
    max-width: 100%;
    overflow: hidden;
    max-height: 100%;
  }
  .dropdownIcon {
    position: absolute;
    height: 80%;
    width: 12.12%;
    top: 10%;
    right: 1.52%;
    bottom: 10%;
    left: 86.36%;
  }
  .fieldsSmallDropdown {
    width: 264px;
    text-align: left;
    font-size: var(--label-text-large-size);
    color: var(--dark-dark-200);
    font-family: var(--body-body-16);
  }
</style>
