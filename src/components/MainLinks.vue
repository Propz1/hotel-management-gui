<template>
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
</template>
<script>
  import { defineComponent } from "vue";
  import NavigationLinkHomeHoverNo from "./NavigationLinkHomeHoverNo.vue";
  import NavigationLinkSettingsHove1 from "./NavigationLinkSettingsHove1.vue";

  export default defineComponent({
    name: "MainLinks",
    components: { NavigationLinkHomeHoverNo, NavigationLinkSettingsHove1 },
    props: {
      onLinkMainPageContainerClick: { type: Function },
      onLinkCalendarPageContainerClick: { type: Function },
      onLinkRequisitionsPageContainerClick: { type: Function },
      onLinkOccupancyPageContainerClick: { type: Function },
      onLinkChatPageContainerClick: { type: Function },
      onLinkReportsPageContainerClick: { type: Function },
      onLinkIntegrationPageContainerClick: { type: Function },
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
</style>
