<template>
  <div
    :class="$style.navigationLinkhomeHoverno"
    :style="navigationLinkHomeHoverNo1Style"
    @click="onLinkMainPageContainerClick"
  >
    <img
      :class="$style.homeIcon"
      alt=""
      :src="pageTitle"
      :style="homeIconStyle"
    />
    <div :class="$style.div" :style="div5Style">{{linkToPageTitle}}</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "NavigationLinkHomeHoverNo",
    props: {
      pageTitle: { type: String },
      linkToPageTitle: { type: String },
      propPosition: { type: String },
      propTop: { type: [Object, Array, String, Number, Boolean] },
      propLeft: { type: [Object, Array, String, Number, Boolean] },
      propWidth: { type: [Object, Array, String, Number, Boolean] },
      propAlignSelf: { type: String },
      propCursor: { type: String },
      propBackgroundColor: { type: String },
      propBorder: { type: String },
      propPadding: { type: [Object, Array, String, Number, Boolean] },
      propTop1: { type: [Object, Array, String, Number, Boolean] },
      propWidth1: { type: [Object, Array, String, Number, Boolean] },
      propFontWeight: { type: String },
      propColor: { type: String },
      propDisplay: { type: String },
      onLinkMainPageContainerClick: { type: Function },
    },
    computed: {
      navigationLinkHomeHoverNo1Style() {
        return {
          position: this.propPosition,
          top: this.propTop,
          left: this.propLeft,
          width: this.propWidth,
          alignSelf: this.propAlignSelf,
          cursor: this.propCursor,
          backgroundColor: this.propBackgroundColor,
          border: this.propBorder,
          padding: this.propPadding,
        };
      },
      homeIconStyle() {
        return {
          top: this.propTop1,
          width: this.propWidth1,
        };
      },
      div5Style() {
        return {
          fontWeight: this.propFontWeight,
          color: this.propColor,
          display: this.propDisplay,
        };
      },
    },
  });
</script>
<style module>
  .homeIcon {
    position: absolute;
    top: calc(50% - 12px);
    left: 24px;
    width: 24px;
    height: 24px;
    overflow: hidden;
    opacity: 0.9;
  }
  .div {
    position: absolute;
    top: calc(50% - 8px);
    left: 71px;
    line-height: 100%;
  }
  .navigationLinkhomeHoverno {
    position: absolute;
    top: 664px;
    left: 20px;
    background-color: var(--base-pure-black);
    width: 240px;
    height: 48px;
    text-align: left;
    font-size: var(--label-text-large-size);
    color: var(--gray-100);
    font-family: var(--font-ubuntu);
  }
</style>
