<template>
  <div :class="$style.reports"><SidebarDark4 /></div>
</template>
<script>
  import { defineComponent } from "vue";
  import SidebarDark4 from "../components/SidebarDark4.vue";

  export default defineComponent({
    name: "Reports",
    components: { SidebarDark4 },
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
        this.$router.push("/requisitions");
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
  .reports {
    position: relative;
    background-color: var(--color-whitesmoke-300);
    width: 100%;
    height: 1080px;
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    justify-content: flex-start;
  }
</style>
