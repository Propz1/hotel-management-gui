<template>
  <div
    :class="$style.navigationLinksettingsHove"
    :style="navigationLinkSettingsHoveStyle"
    @click="onLinkIntegrationPageContainerClick"
  >
    <img :class="$style.settingsIcon" alt="" :src="iconId" />
    <div :class="$style.settings" :style="settingsStyle">
      {{linkToPageTitle}}
    </div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "NavigationLinkSettingsHove1",
    props: {
      iconId: { type: String },
      linkToPageTitle: { type: String },
      navigationLinkSettingsHovBackgroundColor: { type: String },
      navigationLinkSettingsHovPosition: { type: String },
      navigationLinkSettingsHovCursor: { type: String },
      settingsFontWeight: { type: String },
      settingsFontFamily: { type: String },
      settingsOverflow: { type: String },
      settingsTextOverflow: { type: String },
      settingsWhiteSpace: { type: String },
      onLinkIntegrationPageContainerClick: { type: Function },
    },
    computed: {
      navigationLinkSettingsHoveStyle() {
        return {
          backgroundColor: this.navigationLinkSettingsHovBackgroundColor,
          position: this.navigationLinkSettingsHovPosition,
          cursor: this.navigationLinkSettingsHovCursor,
        };
      },
      settingsStyle() {
        return {
          fontWeight: this.settingsFontWeight,
          fontFamily: this.settingsFontFamily,
          overflow: this.settingsOverflow,
          textOverflow: this.settingsTextOverflow,
          whiteSpace: this.settingsWhiteSpace,
        };
      },
    },
  });
</script>
<style module>
  .settingsIcon {
    position: absolute;
    top: calc(50% - 12px);
    left: 24px;
    width: 24px;
    height: 24px;
    overflow: hidden;
    opacity: 0.9;
  }
  .settings {
    position: absolute;
    top: calc(50% - 8px);
    left: 71px;
    line-height: 100%;
    font-weight: 500;
  }
  .navigationLinksettingsHove {
    background-color: var(--base-pure-black);
    width: 240px;
    height: 48px;
    text-align: left;
    font-size: var(--label-text-large-size);
    color: var(--gray-100);
    font-family: var(--label-text-large);
  }
</style>
