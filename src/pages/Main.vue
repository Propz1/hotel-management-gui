<template>
  <div :class="$style.main"><SidebarDark8 /></div>
</template>
<script>
  import { defineComponent } from "vue";
  import SidebarDark8 from "../components/SidebarDark8.vue";

  export default defineComponent({
    name: "Main",
    components: { SidebarDark8 },
    methods: {
      onLinkCalendarPageClick() {
        this.$router.push("/calendar");
      },
      onLinkRequisitionPageClick() {
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
  .main {
    position: relative;
    background-color: var(--color-whitesmoke-300);
    width: 100%;
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    justify-content: flex-start;
  }
</style>
