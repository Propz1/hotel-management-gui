<template>
  <div :class="$style.occupancy">
    <button :class="$style.arrows">
      <button :class="$style.icon">
        <img :class="$style.vectorIcon" alt="" src="/vector.svg" />
      </button>
      <b :class="$style.b">1</b>
      <div :class="$style.icon1">
        <img :class="$style.vectorIcon" alt="" src="/vector1.svg" />
      </div>
    </button>
    <div :class="$style.entriesPage">
      <div :class="$style.row">
        <div :class="$style.number">
          <b :class="$style.b1">10</b>
        </div>
        <div :class="$style.arrow">
          <div :class="$style.polygonParent">
            <img :class="$style.frameChild" alt="" src="/polygon-1.svg" />
            <img :class="$style.frameItem" alt="" src="/polygon-2.svg" />
          </div>
        </div>
      </div>
      <b :class="$style.b2">Записей на странице</b>
    </div>
    <div :class="$style.entriesPage1">
      <b :class="$style.b3">Страница</b>
      <div :class="$style.row1">
        <div :class="$style.number">
          <b :class="$style.b1">1</b>
        </div>
        <div :class="$style.arrow">
          <div :class="$style.polygonParent">
            <img :class="$style.frameChild" alt="" src="/polygon-1.svg" />
            <img :class="$style.frameItem" alt="" src="/polygon-2.svg" />
          </div>
        </div>
      </div>
    </div>
    <table />
    <div :class="$style.time">
      <div :class="$style.timeChild" />
      <div :class="$style.wrapper">
        <b :class="$style.b3">13:30</b>
      </div>
    </div>
    <div :class="$style.selectedFilters">
      <div :class="$style.chips">
        <div :class="$style.content">
          <div :class="$style.contentChild" />
          <b :class="$style.b3">Полный день</b>
          <div :class="$style.div">(8:00 - 18:00)</div>
        </div>
        <img
          :class="$style.iconSetTriptileNew"
          alt=""
          src="/icon-set-triptile-new.svg"
        />
      </div>
      <div :class="$style.chips">
        <div :class="$style.content">
          <div :class="$style.contentItem" />
          <b :class="$style.b7">Свободно</b>
        </div>
        <img
          :class="$style.iconSetTriptileNew"
          alt=""
          src="/icon-set-triptile-new.svg"
        />
      </div>
      <div :class="$style.chips">
        <div :class="$style.content">
          <div :class="$style.contentInner" />
          <b :class="$style.b8">Удалено</b>
        </div>
        <img
          :class="$style.iconSetTriptileNew"
          alt=""
          src="/icon-set-triptile-new.svg"
        />
      </div>
    </div>
    <div :class="$style.category">
      <b :class="$style.b3">Category:</b>
      <b :class="$style.multiply">Multiply</b>
      <img :class="$style.icon2" alt="" src="/icon5.svg" />
    </div>
    <Dashboard />
    <div :class="$style.months">
      <nav :class="$style.monthsframeincalenar">
        <button :class="$style.januarybutton">
          <b :class="$style.starsHotel">January</b>
        </button>
        <div :class="$style.monthsframeincalenarChild" />
        <button :class="$style.januarybutton">
          <b :class="$style.starsHotel">February</b>
        </button>
        <div :class="$style.monthsframeincalenarChild" />
        <button :class="$style.januarybutton">
          <b :class="$style.starsHotel">March</b>
        </button>
        <div :class="$style.monthsframeincalenarChild" />
        <button :class="$style.aprilbutton">
          <b :class="$style.starsHotel3">April</b>
        </button>
        <div :class="$style.monthsframeincalenarChild" />
        <button :class="$style.januarybutton">
          <b :class="$style.starsHotel">May</b>
        </button>
        <div :class="$style.monthsframeincalenarChild" />
        <button :class="$style.januarybutton">
          <b :class="$style.starsHotel">June</b>
        </button>
        <div :class="$style.monthsframeincalenarChild" />
        <button :class="$style.januarybutton">
          <b :class="$style.starsHotel">July</b>
        </button>
        <div :class="$style.monthsframeincalenarChild" />
        <button :class="$style.januarybutton">
          <b :class="$style.starsHotel">August</b>
        </button>
        <div :class="$style.monthsframeincalenarChild" />
        <button :class="$style.januarybutton">
          <b :class="$style.starsHotel">September</b>
        </button>
        <div :class="$style.monthsframeincalenarChild" />
        <button :class="$style.januarybutton">
          <b :class="$style.starsHotel">October</b>
        </button>
        <div :class="$style.monthsframeincalenarChild" />
        <button :class="$style.januarybutton">
          <b :class="$style.starsHotel">November</b>
        </button>
        <div :class="$style.monthsframeincalenarChild" />
        <button :class="$style.januarybutton">
          <b :class="$style.starsHotel">December</b>
        </button>
      </nav>
    </div>
    <div :class="$style.updatesBtn">
      <div :class="$style.row2">
        <img :class="$style.icon2" alt="" src="/icon6.svg" />
        <div :class="$style.text">
          <div :class="$style.textChild" />
          <b :class="$style.updates">Updates</b>
        </div>
      </div>
    </div>
    <div :class="$style.filtersBtn" @click="onFiltersBtnContainerClick">
      <div :class="$style.row3" :onClick="openFiltersWindow">
        <img :class="$style.icon2" alt="" src="/icon7.svg" />
        <b :class="$style.filters">Filters</b>
        <img :class="$style.icon2" alt="" src="/icon8.svg" />
      </div>
    </div>
    <div :class="$style.date">
      <div :class="$style.date1">
        <b :class="$style.b9">4 Апрель</b>
        <b :class="$style.b9">2023</b>
        <button :class="$style.arrowlastday" id="ArrowLastDay">
          <img :class="$style.icon6" alt="" src="/icon9.svg" />
        </button>
        <button :class="$style.arrowlastday">
          <img :class="$style.icon6" alt="" src="/icon10.svg" />
        </button>
      </div>
    </div>
    <div :class="$style.breadcrumbs">
      <b :class="$style.filters">Календарь</b>
      <img :class="$style.icon6" alt="" src="/icon-set-triptile-new1.svg" />
      <b :class="$style.b12">Занятость номеров</b>
    </div>
    <SidebarDark6 />
  </div>
  <div v-if="isFiltersWindowOpen">
    <PortalPopup
      overlayColor="rgba(113, 113, 113, 0.3)"
      placement="Centered"
      :relativeLayerRef="$refs[rowContainer3Ref]"
      :onOutsideClick="closeFiltersWindow"
    >
      <FiltersWindow :onClose="closeFiltersWindow" />
    </PortalPopup>
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import FiltersWindow from "../components/FiltersWindow.vue";
  import PortalPopup from "../components/PortalPopup.vue";
  import Table from "../components/Table.vue";
  import Dashboard from "../components/Dashboard.vue";
  import SidebarDark6 from "../components/SidebarDark6.vue";

  export default defineComponent({
    name: "Occupancy",
    data() {
      return { isFiltersWindowOpen: false };
    },
    components: { FiltersWindow, PortalPopup, Table, Dashboard, SidebarDark6 },
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
  .vectorIcon {
    position: relative;
    width: 6.2px;
    height: 10.5px;
  }
  .icon {
    cursor: pointer;
    border: 1px solid var(--color-whitesmoke-200);
    padding: var(--padding-6xs) var(--padding-4xs);
    background-color: transparent;
    box-sizing: border-box;
    width: 30px;
    height: 30px;
    overflow: hidden;
    flex-shrink: 0;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
  }
  .b {
    position: relative;
    font-size: var(--text-xs-regular-size);
    line-height: 16px;
    font-family: var(--text-xs-regular);
    color: var(--base-pure-black);
    text-align: left;
  }
  .icon1 {
    border: 1px solid var(--color-whitesmoke-200);
    box-sizing: border-box;
    width: 30px;
    height: 30px;
    overflow: hidden;
    flex-shrink: 0;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    padding: var(--padding-6xs) var(--padding-4xs);
  }
  .arrows {
    cursor: pointer;
    border: none;
    padding: 0;
    background-color: transparent;
    position: absolute;
    top: 1047px;
    left: calc(50% + 47px);
    width: 173px;
    height: 33px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    gap: 34px;
  }
  .b1 {
    position: relative;
    line-height: 16px;
  }
  .number {
    position: absolute;
    top: 0px;
    left: 0px;
    border: 1px solid var(--color-whitesmoke-200);
    box-sizing: border-box;
    width: 30px;
    height: 30px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    padding: var(--padding-3xs);
  }
  .frameChild {
    position: relative;
    width: 10px;
    height: 6px;
  }
  .frameItem {
    position: relative;
    width: 8.7px;
    height: 4.5px;
  }
  .polygonParent {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-11xs);
  }
  .arrow {
    position: absolute;
    top: 0px;
    left: 29px;
    border: 1px solid var(--color-whitesmoke-200);
    box-sizing: border-box;
    width: 22px;
    height: 30px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    padding: var(--padding-3xs);
  }
  .row {
    position: relative;
    width: 51px;
    height: 30px;
  }
  .b2 {
    position: relative;
    line-height: 14px;
    color: var(--primary-footer);
  }
  .entriesPage {
    position: absolute;
    top: 1047px;
    left: 266px;
    width: 202px;
    height: 33px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-3xs);
  }
  .b3 {
    position: relative;
    line-height: 14px;
  }
  .row1 {
    position: relative;
    width: 51px;
    height: 30px;
    color: var(--base-pure-black);
  }
  .entriesPage1 {
    position: absolute;
    top: 1047px;
    left: 1797px;
    width: 115px;
    height: 33px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-3xs);
    color: var(--primary-footer);
  }
  .timeChild {
    position: absolute;
    top: calc(50% - 278.95px);
    left: calc(50% - 0.7px);
    background-color: #eeacac;
    width: 1.3px;
    height: 579.1px;
  }
  .wrapper {
    position: absolute;
    top: 0px;
    left: 0px;
    background-color: #ff0000;
    width: 48px;
    height: 21.2px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    padding: 0px 6px;
    box-sizing: border-box;
  }
  .time {
    position: absolute;
    top: 422.9px;
    left: 842.8px;
    width: 48px;
    height: 600.3px;
    font-size: var(--font-size-3xs);
    color: var(--light-light-100);
  }
  .contentChild {
    position: relative;
    border-radius: 50%;
    background-color: var(--color-teal-100);
    width: 8px;
    height: 8px;
  }
  .div {
    position: relative;
    line-height: 14px;
    color: rgba(108, 108, 108, 0.5);
  }
  .content {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    gap: var(--gap-9xs);
  }
  .iconSetTriptileNew {
    position: relative;
    width: 18px;
    height: 21.5px;
    overflow: hidden;
    flex-shrink: 0;
  }
  .chips {
    border-radius: var(--br-9xs);
    background-color: var(--color-whitesmoke-100);
    border: 1px solid var(--color-gainsboro-200);
    box-sizing: border-box;
    height: 30px;
    overflow: hidden;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    padding: 0px var(--padding-5xs);
    gap: var(--gap-5xs);
  }
  .contentItem {
    position: relative;
    border-radius: 50%;
    background-color: var(--color-red-100);
    width: 8px;
    height: 8px;
  }
  .b7 {
    position: relative;
    line-height: 14px;
    display: inline-block;
    width: 78px;
    flex-shrink: 0;
  }
  .contentInner {
    position: relative;
    border-radius: 50%;
    background-color: var(--color-orange-100);
    width: 8px;
    height: 8px;
  }
  .b8 {
    position: relative;
    line-height: 14px;
    display: inline-block;
    width: 66px;
    flex-shrink: 0;
  }
  .selectedFilters {
    position: absolute;
    top: 300px;
    left: 266px;
    width: 531px;
    height: 34px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-xs);
    color: var(--color-dimgray-200);
  }
  .multiply {
    position: relative;
    line-height: 14px;
    color: var(--base-pure-black);
  }
  .icon2 {
    position: relative;
    width: 18px;
    height: 18px;
    overflow: hidden;
    flex-shrink: 0;
  }
  .category {
    position: absolute;
    top: 300px;
    left: 1723px;
    background-color: var(--light-light-100);
    border: 1px solid var(--color-gainsboro-200);
    box-sizing: border-box;
    width: 196px;
    height: 34px;
    overflow: hidden;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    padding: 0px var(--padding-3xs);
    gap: var(--gap-9xs);
    color: var(--color-dimgray-200);
  }
  .starsHotel {
    position: relative;
    font-size: var(--text-xs-regular-size);
    line-height: 22px;
    font-family: var(--text-xs-regular);
    color: var(--color-dimgray-200);
    text-align: left;
  }
  .januarybutton {
    cursor: pointer;
    border: none;
    padding: 0px var(--padding-3xs);
    background-color: transparent;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
  }
  .monthsframeincalenarChild {
    position: relative;
    border-radius: var(--br-12xs);
    background-color: var(--color-gainsboro-200);
    width: 1px;
    height: 16px;
  }
  .starsHotel3 {
    position: relative;
    font-size: var(--text-xs-regular-size);
    line-height: 22px;
    font-family: var(--text-xs-regular);
    color: var(--base-pure-black);
    text-align: left;
  }
  .aprilbutton {
    cursor: pointer;
    border: none;
    padding: 0px var(--padding-3xs);
    background-color: var(--light-light-100);
    border-radius: var(--br-9xs);
    width: 104px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    box-sizing: border-box;
    mix-blend-mode: normal;
  }
  .monthsframeincalenar {
    margin: 0;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: 18px;
  }
  .months {
    position: absolute;
    top: 128px;
    left: 266px;
    background-color: var(--color-gray-200);
    border: 1px solid var(--color-gainsboro-200);
    box-sizing: border-box;
    width: 1654px;
    height: 33px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    padding: 4px var(--padding-5xs);
  }
  .textChild {
    position: absolute;
    top: 7px;
    left: 59px;
    border-radius: 50%;
    width: 6px;
    height: 6px;
  }
  .updates {
    position: absolute;
    top: 0px;
    left: 0px;
    line-height: 18px;
  }
  .text {
    position: relative;
    width: 58px;
    height: 18px;
  }
  .row2 {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-3xs);
  }
  .updatesBtn {
    position: absolute;
    top: 57px;
    left: 1733px;
    background-color: var(--color-gray-600);
    width: 181px;
    height: 44px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: var(--padding-3xs) var(--padding-xl);
    box-sizing: border-box;
    font-size: var(--label-large-label-14-size);
    color: var(--light-light-100);
  }
  .filters {
    position: relative;
    line-height: 18px;
  }
  .row3 {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-3xs);
    cursor: pointer;
  }
  .filtersBtn {
    position: absolute;
    top: 57px;
    left: 1538px;
    background-color: var(--light-light-100);
    border: 1px solid var(--color-gainsboro-200);
    box-sizing: border-box;
    width: 181px;
    height: 44px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: var(--padding-3xs) var(--padding-xl);
    cursor: pointer;
    font-size: var(--label-large-label-14-size);
  }
  .b9 {
    position: relative;
    line-height: 32px;
  }
  .icon6 {
    position: relative;
    width: 24px;
    height: 24px;
    overflow: hidden;
    flex-shrink: 0;
  }
  .arrowlastday {
    cursor: pointer;
    border: 1px solid var(--color-gainsboro-300);
    padding: var(--padding-3xs);
    background-color: transparent;
    box-sizing: border-box;
    width: 40px;
    height: 40px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
  }
  .date1 {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-base);
  }
  .date {
    position: absolute;
    top: 57px;
    left: 267px;
    width: 385px;
    height: 44px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
    font-size: var(--font-size-11xl);
  }
  .b12 {
    position: relative;
    line-height: 18px;
    color: var(--base-pure-black);
  }
  .breadcrumbs {
    position: absolute;
    top: 12px;
    left: 266px;
    width: 254px;
    height: 27px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-9xs);
    font-size: var(--label-large-label-14-size);
    color: var(--color-gainsboro-300);
  }
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
