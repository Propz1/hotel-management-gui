<template>
  <div :class="$style.sidebardark">
    <TopHeader />
    <div :class="$style.mainLinks">
      <NavigationLinkHomeHoverNo
        pageTitle="/home.svg"
        linkToPageTitle="Главная"
        propPosition="relative"
        propTop="unset"
        propLeft="unset"
        propWidth="unset"
        propAlignSelf="stretch"
        propCursor="pointer"
        propBackgroundColor="#000"
        propBorder="none"
        propPadding="0"
        propTop1="calc(50% - 12px)"
        propWidth1="24px"
        propFontWeight="unset"
        propColor="#f1f1f1"
        propDisplay="inline-block"
        :onLinkMainPageContainerClick="onLinkMainPageClick"
      /><NavigationLinkHomeHoverNo
        pageTitle="/calendar.svg"
        linkToPageTitle="Календарь"
        propPosition="relative"
        propTop="unset"
        propLeft="unset"
        propWidth="240px"
        propAlignSelf="unset"
        propCursor="pointer"
        propBackgroundColor="transparent"
        propBorder="none"
        propPadding="0"
        propTop1="calc(50% - 11px)"
        propWidth1="24px"
        propFontWeight="unset"
        propColor="#f1f1f1"
        propDisplay="inline-block"
        :onLinkMainPageContainerClick="onLinkCalendarPageClick"
      /><NavigationLinkHomeHoverNo1
        iconImageUrl="/notificationsnone.svg"
        pageTitle="Заявки"
        linkToPageTitle="Заявки"
        propTop="unset"
        propLeft="unset"
        propPosition="relative"
        propBackgroundColor="#242d35"
        propCursor="pointer"
        propWidth="240px"
        propBorder="none"
        propPadding="0"
        propAlignSelf="unset"
        propFontWeight="unset"
        propDisplay="inline-block"
        :onLinkChatPageContainerClick="onLinkRequisitionPageClick"
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
      /><NavigationLinkSettingsHove1
        iconId="/dashboard2.svg"
        linkToPageTitle="Интеграция"
        navigationLinkSettingsHovBackgroundColor="unset"
        navigationLinkSettingsHovPosition="relative"
        navigationLinkSettingsHovCursor="pointer"
        settingsFontWeight="unset"
        settingsFontFamily="Ubuntu"
        settingsOverflow="unset"
        settingsTextOverflow="unset"
        settingsWhiteSpace="unset"
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
  import NavigationLinkHomeHoverNo1 from "./NavigationLinkHomeHoverNo1.vue";
  import NavigationLinkSettingsHove1 from "./NavigationLinkSettingsHove1.vue";
  import StickyFooter from "./StickyFooter.vue";

  export default defineComponent({
    name: "SidebarDark",
    components: {
      TopHeader,
      NavigationLinkHomeHoverNo,
      NavigationLinkHomeHoverNo1,
      NavigationLinkSettingsHove1,
      StickyFooter,
    },
    methods: {
      onLinkMainPageClick() {
        this.$router.push("/main");
      },
      onLinkCalendarPageClick() {
        this.$router.push("/calendar");
      },
      onLinkRequisitionPageClick() {
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
  .sidebardark {
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
    text-align: left;
    font-size: var(--font-size-9xl);
    color: var(--primary-blue-900);
    font-family: var(--font-ubuntu);
  }
</style>
