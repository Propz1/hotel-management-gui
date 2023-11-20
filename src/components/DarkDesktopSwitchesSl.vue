<template>
  <div
    :class="$style.darkDesktopSwitchesSl"
    :style="darkDesktopSwitchesSlStyle"
  >
    <div :class="$style.rectangleCopy18" />
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "DarkDesktopSwitchesSl",
    props: {
      guestsCheckBoxPosition: { type: String },
      guestsCheckBoxBorderRadius: { type: String },
      guestsCheckBoxBackgroundColor: { type: String },
    },
    computed: {
      darkDesktopSwitchesSlStyle() {
        return {
          position: this.guestsCheckBoxPosition,
          borderRadius: this.guestsCheckBoxBorderRadius,
          backgroundColor: this.guestsCheckBoxBackgroundColor,
        };
      },
    },
  });
</script>
<style module>
  .rectangleCopy18 {
    position: absolute;
    height: 100%;
    width: 100%;
    top: 0%;
    right: 0%;
    bottom: 0%;
    left: 0%;
    border-radius: var(--br-9xs);
    background-color: var(--color-gray-900);
    border: 1px solid #c8c8c8;
    box-sizing: border-box;
  }
  .darkDesktopSwitchesSl {
    width: 18px;
    height: 18px;
  }
</style>
