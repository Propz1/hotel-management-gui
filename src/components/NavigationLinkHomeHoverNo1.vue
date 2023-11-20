<template>
  <div
    :class="$style.navigationLinkhomeHoverno"
    :style="navigationLinkHomeHoverNoStyle"
    @click="onLinkChatPageContainerClick"
  >
    <img :class="$style.homeIcon" alt="" :src="iconImageUrl" />
    <div :class="$style.div" :style="div4Style">{{pageTitle}}</div>
    <div :class="$style.navigationLinkhomeHovernoChild" />
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "NavigationLinkHomeHoverNo1",
    props: {
      iconImageUrl: { type: String },
      pageTitle: { type: String },
      propTop: { type: [Object, Array, String, Number, Boolean] },
      propLeft: { type: [Object, Array, String, Number, Boolean] },
      propPosition: { type: String },
      propBackgroundColor: { type: String },
      propCursor: { type: String },
      propWidth: { type: [Object, Array, String, Number, Boolean] },
      propBorder: { type: String },
      propPadding: { type: [Object, Array, String, Number, Boolean] },
      propAlignSelf: { type: String },
      propFontWeight: { type: String },
      propDisplay: { type: String },
      onLinkChatPageContainerClick: { type: Function },
    },
    computed: {
      navigationLinkHomeHoverNoStyle() {
        return {
          top: this.propTop,
          left: this.propLeft,
          position: this.propPosition,
          backgroundColor: this.propBackgroundColor,
          cursor: this.propCursor,
          width: this.propWidth,
          border: this.propBorder,
          padding: this.propPadding,
          alignSelf: this.propAlignSelf,
        };
      },
      div4Style() {
        return {
          fontWeight: this.propFontWeight,
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
  .navigationLinkhomeHovernoChild {
    position: absolute;
    height: 100%;
    top: 0%;
    right: 0px;
    bottom: 0%;
    background-color: var(--primary-blue-400);
    width: 4px;
  }
  .navigationLinkhomeHoverno {
    position: absolute;
    top: 664px;
    left: 548px;
    background-color: var(--gray-800);
    width: 240px;
    height: 48px;
    text-align: left;
    font-size: var(--label-text-large-size);
    color: var(--primary-blue-50);
    font-family: var(--font-ubuntu);
  }
</style>
