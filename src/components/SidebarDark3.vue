<template>
  <div :class="$style.sidebarDark">
    <TopHeader />
    <div :class="$style.mainLinks">
      <NavigationLinkHomeHoverNo
        pageTitle="/home.svg"
        linkToPageTitle="Гостиницы"
        propPosition="relative"
        propTop="unset"
        propLeft="unset"
        propWidth="unset"
        propAlignSelf="stretch"
        propCursor="pointer"
        propBackgroundColor="#000"
        propBorder="unset"
        propPadding="unset"
        propTop1="calc(50% - 12px)"
        propWidth1="24px"
        propFontWeight="unset"
        propColor="#f1f1f1"
        propDisplay="inline-block"
        :onLinkMainPageContainerClick="onLinkMainPageContainerClick"
      /><NavigationLinkHomeHoverNo
        pageTitle="/calendar.svg"
        linkToPageTitle="Календарь"
        propPosition="relative"
        propTop="unset"
        propLeft="unset"
        propWidth="240px"
        propAlignSelf="unset"
        propCursor="pointer"
        propBackgroundColor="unset"
        propBorder="unset"
        propPadding="unset"
        propTop1="calc(50% - 11px)"
        propWidth1="24px"
        propFontWeight="unset"
        propColor="#f1f1f1"
        propDisplay="inline-block"
        :onLinkMainPageContainerClick="onLinkCalendarPageContainerClick"
      /><NavigationLinkHomeHoverNo
        pageTitle="/notificationsnone1.svg"
        linkToPageTitle="Заявки"
        propPosition="relative"
        propTop="unset"
        propLeft="unset"
        propWidth="240px"
        propAlignSelf="unset"
        propCursor="pointer"
        propBackgroundColor="unset"
        propBorder="unset"
        propPadding="unset"
        propTop1="calc(50% - 12px)"
        propWidth1="24px"
        propFontWeight="unset"
        propColor="#f1f1f1"
        propDisplay="inline-block"
        :onLinkMainPageContainerClick="onLinkRequisitionsPageContainerClick"
      /><NavigationLinkHomeHoverNo
        pageTitle="/occupancyoutline.svg"
        linkToPageTitle="Заполненность"
        propPosition="relative"
        propTop="unset"
        propLeft="unset"
        propWidth="240px"
        propAlignSelf="unset"
        propCursor="pointer"
        propBackgroundColor="unset"
        propBorder="unset"
        propPadding="unset"
        propTop1="calc(50% - 12px)"
        propWidth1="24px"
        propFontWeight="500"
        propColor="#f1f1f1"
        propDisplay="inline-block"
        :onLinkMainPageContainerClick="onLinkOccupancyPageContainerClick"
      /><NavigationLinkHomeHoverNo
        pageTitle="/chatoutline.svg"
        linkToPageTitle="Чат поддержки"
        propPosition="relative"
        propTop="unset"
        propLeft="unset"
        propWidth="240px"
        propAlignSelf="unset"
        propCursor="pointer"
        propBackgroundColor="unset"
        propBorder="unset"
        propPadding="unset"
        propTop1="calc(50% - 12px)"
        propWidth1="24px"
        propFontWeight="500"
        propColor="#f1f1f1"
        propDisplay="inline-block"
        :onLinkMainPageContainerClick="onLinkChatPageContainerClick"
      /><NavigationLinkHomeHoverNo
        pageTitle="/report.svg"
        linkToPageTitle="Отчеты"
        propPosition="relative"
        propTop="unset"
        propLeft="unset"
        propWidth="240px"
        propAlignSelf="unset"
        propCursor="pointer"
        propBackgroundColor="unset"
        propBorder="unset"
        propPadding="unset"
        propTop1="calc(50% - 12px)"
        propWidth1="34px"
        propFontWeight="unset"
        propColor="#f1f1f1"
        propDisplay="inline-block"
        :onLinkMainPageContainerClick="onLinkReportsPageContainerClick"
      /><NavigationLinkHelpCenterH
        iconImageUrl="/dashboard.svg"
        linkToPageTitle="Интеграция"
        navigationLinkToPagePosition="relative"
        navigationLinkToPageCursor="pointer"
        navigationLinkToPageFontWeight="unset"
        navigationLinkToPageFontFamily="Ubuntu"
        navigationLinkToPageOverflow="unset"
        helpCenterTextOverflow="unset"
        helpCenterWhiteSpace="unset"
        :onLinkIntegrationPageContainerClick="onLinkIntegrationPageContainerClick"
      />
    </div>
    <StickyFooter
      :onLinkHelpPageContainerClick="onLinkHelpPageContainerClick"
      :onLinkSettingsPageContainerClick="onLinkSettingsPageContainerClick"
    />
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import TopHeader from "./TopHeader.vue";
  import NavigationLinkHomeHoverNo from "./NavigationLinkHomeHoverNo.vue";
  import NavigationLinkHelpCenterH from "./NavigationLinkHelpCenterH.vue";
  import StickyFooter from "./StickyFooter.vue";

  export default defineComponent({
    name: "SidebarDark3",
    components: {
      TopHeader,
      NavigationLinkHomeHoverNo,
      NavigationLinkHelpCenterH,
      StickyFooter,
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
  .mainLinks {
    align-self: stretch;
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
    gap: var(--gap-5xs);
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
