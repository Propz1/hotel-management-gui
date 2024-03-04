<template>
  <div :class="$style.occupancy">
    <SidebarDark6 />
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  import SidebarDark6 from "../components/SidebarDark6.vue";

  export default defineComponent({
    name: "Occupancy",
    data() {
      return { isFiltersWindowOpen: false };
    },
    components: { SidebarDark6 },
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
      onFiltersBtnContainerClick() {
        // Please sync "Requisitions_Filter" to the project
      },
      openFiltersWindow() {
        this.isFiltersWindowOpen = true;
      },
      closeFiltersWindow() {
        this.isFiltersWindowOpen = false;
      },
    },
  });
</script>
<style module>
  
  .occupancy {
    position: relative;
    background-color: var(--color-whitesmoke-300);
    width: 100%;
    height: 1080px;
    text-align: left;
    font-size: var(--text-xs-regular-size);
    color: var(--base-pure-black);
    font-family: var(--text-xs-regular);
  }
</style>
