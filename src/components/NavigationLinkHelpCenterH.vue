<template>
  <div
    :class="$style.navigationLinkhelpCenterH"
    :style="navigationLinkHelpCenterHStyle"
    @click="onLinkIntegrationPageContainerClick"
  >
    <img :class="$style.helpOutlineIcon" alt="" :src="iconImageUrl" />
    <div :class="$style.helpCenter" :style="helpCenterStyle">
      {{linkToPageTitle}}
    </div>
    <div :class="$style.navigationLinkhelpCenterHChild" />
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "NavigationLinkHelpCenterH",
    props: {
      iconImageUrl: { type: String },
      linkToPageTitle: { type: String },
      navigationLinkToPagePosition: { type: String },
      navigationLinkToPageCursor: { type: String },
      navigationLinkToPageFontWeight: { type: String },
      navigationLinkToPageFontFamily: { type: String },
      navigationLinkToPageOverflow: { type: String },
      helpCenterTextOverflow: { type: String },
      helpCenterWhiteSpace: { type: String },
      onLinkIntegrationPageContainerClick: { type: Function },
    },
    computed: {
      navigationLinkHelpCenterHStyle() {
        return {
          position: this.navigationLinkToPagePosition,
          cursor: this.navigationLinkToPageCursor,
        };
      },
      helpCenterStyle() {
        return {
          fontWeight: this.navigationLinkToPageFontWeight,
          fontFamily: this.navigationLinkToPageFontFamily,
          overflow: this.navigationLinkToPageOverflow,
          textOverflow: this.helpCenterTextOverflow,
          whiteSpace: this.helpCenterWhiteSpace,
        };
      },
    },
  });
</script>
<style module>
  .helpOutlineIcon {
    position: absolute;
    top: calc(50% - 12px);
    left: 24px;
    width: 24px;
    height: 24px;
    overflow: hidden;
    opacity: 0.9;
  }
  .helpCenter {
    position: absolute;
    top: calc(50% - 8px);
    left: 71px;
    line-height: 100%;
    font-weight: 500;
  }
  .navigationLinkhelpCenterHChild {
    position: absolute;
    height: 100%;
    top: 0%;
    right: 0px;
    bottom: 0%;
    background-color: var(--primary-blue-400);
    width: 4px;
  }
  .navigationLinkhelpCenterH {
    background-color: var(--gray-800);
    width: 240px;
    height: 48px;
    text-align: left;
    font-size: var(--label-text-large-size);
    color: var(--primary-blue-50);
    font-family: var(--label-text-large);
  }
</style>
