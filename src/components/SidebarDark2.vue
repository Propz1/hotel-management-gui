<template>
  <div :class="$style.sidebarDark">
    <TopHeader /><MainLinks
      :onLinkMainPageContainerClick="onLinkMainPageContainerClick"
      :onLinkCalendarPageContainerClick="onLinkCalendarPageContainerClick"
      :onLinkRequisitionsPageContainerClick="onLinkRequisitionPageContainerClick"
      :onLinkOccupancyPageContainerClick="onLinkOccupancyPageContainerClick"
      :onLinkChatPageContainerClick="onLinkChatPageContainerClick"
      :onLinkReportsPageContainerClick="onLinkReportsPageContainerClick"
      :onLinkIntegrationPageContainerClick="onLinkIntegrationPageContainerClick"
    />
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
        /><NavigationLinkHelpCenterH
          iconImageUrl="/settings.svg"
          linkToPageTitle="Настройки"
          navigationLinkToPagePosition="relative"
          navigationLinkToPageCursor="pointer"
          navigationLinkToPageFontWeight="unset"
          navigationLinkToPageFontFamily="ABeeZee"
          navigationLinkToPageOverflow="hidden"
          helpCenterTextOverflow="ellipsis"
          helpCenterWhiteSpace="nowrap"
          :onLinkIntegrationPageContainerClick="onLinkSettingsPageContainerClick"
        />
      </div>
    </div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import TopHeader from "./TopHeader.vue";
  import MainLinks from "./MainLinks.vue";
  import NavigationLinkSettingsHove1 from "./NavigationLinkSettingsHove1.vue";
  import NavigationLinkHelpCenterH from "./NavigationLinkHelpCenterH.vue";

  export default defineComponent({
    name: "SidebarDark2",
    components: {
      TopHeader,
      MainLinks,
      NavigationLinkSettingsHove1,
      NavigationLinkHelpCenterH,
    },
    methods: {
      onLinkMainPageContainerClick() {
        this.$router.push("/main");
      },
      onLinkCalendarPageContainerClick() {
        this.$router.push("/calendar");
      },
      onLinkRequisitionPageContainerClick() {
        this.$router.push("/");
      },
      onLinkOccupancyPageContainerClick() {
        this.$router.push("/occupancy");
      },
      onLinkChatPageContainerClick() {
        this.$router.push("/chat");
      },
      onLinkReportsPageContainerClick() {
        this.$router.push("/reports");
      },
      onLinkIntegrationPageContainerClick() {
        this.$router.push("/integrations");
      },
      onLinkHelpPageContainerClick() {
        this.$router.push("/help");
      },
      onLinkSettingsPageContainerClick() {
        this.$router.push("/settings");
      },
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
  .sidebarDark {
    background-color: var(--base-pure-black);
    box-shadow: -1px 0px 0px #3e3e3e inset;
    width: 240px;
    height: 1080px;
    overflow: hidden;
    flex-shrink: 0;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
    gap: var(--gap-base);
  }
</style>
