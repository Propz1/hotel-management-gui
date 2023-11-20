<template>
  <div :class="$style.navigationLinkdashboardHov">
    <img :class="$style.tabIcon" alt="" src="/tab.svg" />
    <img
      :class="$style.navLinksDarknavigationLink"
      alt=""
      src="/nav-links-darknavigation-linkdashboard-hoverno-activenotabvector1.svg"
    />
    <div :class="$style.dashboard">Dashboard</div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "NavigationLinkDashboardHov1",
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
  }
  .navLinksDarknavigationLink {
    position: absolute;
    top: calc(50% - 12px);
    left: 24px;
    width: 22px;
    height: 18px;
    overflow: hidden;
    display: none;
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
    left: 20px;
    background-color: var(--base-pure-black);
    width: 240px;
    height: 48px;
    text-align: left;
    font-size: var(--label-text-large-size);
    color: var(--gray-100);
    font-family: var(--font-ubuntu);
  }
</style>
