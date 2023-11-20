<template>
  <div :class="$style.navigationLinkdashboardHov">
    <img :class="$style.tabIcon" alt="" src="/tab1.svg" />
    <div :class="$style.navigationLinkdashboardHovChild" />
    <div :class="$style.dashboard">Dashboard</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "NavigationLinkDashboardHov",
  });
</script>
<style module>
  .tabIcon {
    position: absolute;
    top: calc(50% - 12px);
    left: 24px;
    width: 24px;
    height: 24px;
    overflow: hidden;
    opacity: 0.9;
  }
  .navigationLinkdashboardHovChild {
    position: absolute;
    height: 100%;
    top: 0%;
    right: 0px;
    bottom: 0%;
    background-color: var(--primary-blue-400);
    width: 4px;
  }
  .dashboard {
    position: absolute;
    top: calc(50% - 8px);
    left: 71px;
    line-height: 100%;
    font-weight: 500;
  }
  .navigationLinkdashboardHov {
    position: absolute;
    top: 20px;
    left: 548px;
    background-color: var(--gray-800);
    width: 240px;
    height: 48px;
    text-align: left;
    font-size: var(--label-text-large-size);
    color: var(--primary-blue-50);
    font-family: var(--font-ubuntu);
  }
</style>
