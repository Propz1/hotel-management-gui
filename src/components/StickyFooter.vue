<template>
  <div :class="$style.stickyFooter">
    <div :class="$style.bottomLinks">
      <NavigationLinkSettingsHove1
        iconId="/helpoutline1.svg"
        linkToPageTitle="О программе"
        navigationLinkSettingsHovBackgroundColor="#000"
        navigationLinkSettingsHovPosition="relative"
        navigationLinkSettingsHovCursor="pointer"
        settingsFontWeight="500"
        settingsFontFamily="Ubuntu"
        settingsOverflow="unset"
        settingsTextOverflow="unset"
        settingsWhiteSpace="unset"
        :onLinkIntegrationPageContainerClick="onLinkHelpPageContainerClick"
      /><NavigationLinkSettingsHove1
        iconId="/settings1.svg"
        linkToPageTitle="Настройки"
        navigationLinkSettingsHovBackgroundColor="#000"
        navigationLinkSettingsHovPosition="relative"
        navigationLinkSettingsHovCursor="pointer"
        settingsFontWeight="unset"
        settingsFontFamily="ABeeZee"
        settingsOverflow="hidden"
        settingsTextOverflow="ellipsis"
        settingsWhiteSpace="nowrap"
        :onLinkIntegrationPageContainerClick="onLinkSettingsPageContainerClick"
      />
    </div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import NavigationLinkSettingsHove1 from "./NavigationLinkSettingsHove1.vue";

  export default defineComponent({
    name: "StickyFooter",
    components: { NavigationLinkSettingsHove1 },
    props: {
      onLinkHelpPageContainerClick: { type: Function },
      onLinkSettingsPageContainerClick: { type: Function },
    },
  });
</script>
<style module>
  .bottomLinks {
    align-self: stretch;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
    padding: 0px 0px var(--padding-5xs);
    gap: var(--gap-9xs);
  }
  .stickyFooter {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
  }
</style>
