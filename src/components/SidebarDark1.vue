<template>
  <div :class="$style.sidebarDark">
    <TopHeader /><MainLinks
      :onLinkMainPageContainerClick="onLinkMainPageContainerClick"
      :onLinkCalendarPageContainerClick="onLinkCalendarPageContainerClick"
      :onLinkRequisitionsPageContainerClick="onLinkRequisitionsPageContainerClick"
      :onLinkOccupancyPageContainerClick="onLinkOccupancyPageContainerClick"
      :onLinkChatPageContainerClick="onLinkChatPageContainerClick"
      :onLinkReportsPageContainerClick="onLinkReportsPageContainerClick"
      :onLinkIntegrationPageContainerClick="onLinkIntegrationPageContainerClick"
    />
    <div :class="$style.stickyFooter">
      <div :class="$style.bottomLinks">
        <NavigationLinkHelpCenterH
          iconImageUrl="/helpoutline.svg"
          linkToPageTitle="О программе"
          navigationLinkToPagePosition="relative"
          navigationLinkToPageCursor="pointer"
          navigationLinkToPageFontWeight="500"
          navigationLinkToPageFontFamily="Ubuntu"
          navigationLinkToPageOverflow="unset"
          helpCenterTextOverflow="unset"
          helpCenterWhiteSpace="unset"
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
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import TopHeader from "./TopHeader.vue";
  import MainLinks from "./MainLinks.vue";
  import NavigationLinkHelpCenterH from "./NavigationLinkHelpCenterH.vue";
  import NavigationLinkSettingsHove1 from "./NavigationLinkSettingsHove1.vue";

  export default defineComponent({
    name: "SidebarDark1",
    components: {
      TopHeader,
      MainLinks,
      NavigationLinkHelpCenterH,
      NavigationLinkSettingsHove1,
    },
    methods: {
      onLinkMainPageContainerClick() {
        this.$router.push("/main");
      },
      onLinkCalendarPageContainerClick() {
        this.$router.push("/calendar");
      },
      onLinkRequisitionsPageContainerClick() {
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
