<template>
  <div :class="$style.dashboard">
    <div :class="$style.row">
      <b :class="$style.titletotalnumberguests">
        <p :class="$style.p">Всего гостей за месяц</p>
      </b>
      <div :class="$style.totalnumberguests">
        <div :class="$style.progressLine">
          <div :class="$style.progressLineChild" />
        </div>
        <div :class="$style.titlenumberguests">546 чел.</div>
      </div>
    </div>
    <div :class="$style.row1">
      <b :class="$style.titleprogressline"
        >Реальная заполняемость гостиницы в месяц (OCC)</b
      >
      <div :class="$style.occupancy">
        <b :class="$style.totalhourtitle">
          <span>Всего за </span>
          <span :class="$style.span">Апрель </span>
        </b>
        <div :class="$style.hourtitle">500 часов</div>
      </div>
      <div :class="$style.progressLine1">
        <div :class="$style.progressLineItem" />
        <div :class="$style.progressLineInner" />
        <div :class="$style.rectangleDiv" />
        <div :class="$style.progressLineChild1" />
        <div :class="$style.progressLineChild2" />
      </div>
    </div>
    <div :class="$style.dashboardChild" />
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "Dashboard",
  });
</script>
<style module>
  .p {
    margin: 0;
  }
  .titletotalnumberguests {
    position: absolute;
    top: 0px;
    left: 0px;
    line-height: 14px;
    display: inline-block;
    width: 237.2px;
    height: 15.6px;
  }
  .progressLineChild {
    position: absolute;
    top: 0px;
    left: 0px;
    border-radius: var(--br-3xs);
    background-color: var(--color-teal-100);
    width: 120px;
    height: 6px;
  }
  .progressLine {
    position: relative;
    border-radius: var(--br-3xs);
    background-color: var(--color-lightblue);
    width: 130px;
    height: 6px;
  }
  .titlenumberguests {
    position: relative;
    line-height: 30px;
    font-weight: 500;
  }
  .totalnumberguests {
    position: absolute;
    top: 18.9px;
    left: 0px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-3xs);
    font-size: var(--h6-size);
    color: var(--base-pure-black);
    font-family: var(--text-xs-regular);
  }
  .row {
    position: absolute;
    top: 22.3px;
    left: 38.9px;
    width: 237.2px;
    height: 48.9px;
  }
  .titleprogressline {
    position: absolute;
    top: -0.2px;
    left: 0.5px;
    line-height: 14px;
    display: inline-block;
    width: 313px;
    height: 16px;
  }
  .span {
    color: var(--base-pure-black);
  }
  .totalhourtitle {
    position: absolute;
    top: 0px;
    left: 2.6px;
    line-height: 14px;
    display: inline-block;
    width: 101.1px;
    height: 15.6px;
  }
  .hourtitle {
    position: absolute;
    top: 14.5px;
    left: 0px;
    line-height: 30px;
    color: var(--base-pure-black);
    display: inline-block;
    width: 81.7px;
    height: 33.4px;
  }
  .occupancy {
    position: absolute;
    top: 0px;
    left: 1127.7px;
    width: 103.7px;
    height: 47.9px;
    mix-blend-mode: normal;
  }
  .progressLineItem {
    position: absolute;
    top: 0px;
    left: 80px;
    border-radius: var(--br-3xs);
    background-color: var(--color-lightblue);
    width: 840px;
    height: 6px;
  }
  .progressLineInner {
    position: absolute;
    top: 0px;
    left: 420px;
    background-color: var(--color-orange-100);
    border-right: 1px solid var(--light-light-100);
    box-sizing: border-box;
    width: 197px;
    height: 6px;
  }
  .rectangleDiv {
    position: absolute;
    top: 0px;
    left: 350px;
    background-color: var(--color-red-100);
    border-right: 1px solid var(--light-light-100);
    box-sizing: border-box;
    width: 138px;
    height: 6px;
  }
  .progressLineChild1 {
    position: absolute;
    top: 0px;
    left: 220px;
    background-color: var(--color-royalblue-100);
    border-right: 1px solid var(--light-light-100);
    box-sizing: border-box;
    width: 138px;
    height: 6px;
  }
  .progressLineChild2 {
    position: absolute;
    top: 0px;
    left: 0px;
    background-color: var(--color-teal-100);
    border-right: 1px solid var(--light-light-100);
    box-sizing: border-box;
    width: 254px;
    height: 6px;
  }
  .progressLine1 {
    position: absolute;
    top: 32.3px;
    left: 0px;
    width: 1088.8px;
    height: 6.7px;
    overflow: hidden;
  }
  .row1 {
    position: absolute;
    top: 22.3px;
    left: 365.5px;
    width: 1231.4px;
    height: 47.9px;
  }
  .dashboardChild {
    position: absolute;
    top: 16.7px;
    left: 325.4px;
    background-color: var(--color-gainsboro-200);
    width: 1.3px;
    height: 49px;
  }
  .dashboard {
    position: absolute;
    top: 188px;
    left: 266px;
    background-color: var(--light-light-100);
    border: 1px solid var(--color-gainsboro-200);
    box-sizing: border-box;
    width: 1652.7px;
    height: 85.8px;
    text-align: left;
    font-size: var(--text-xs-regular-size);
    color: var(--color-dimgray-200);
    font-family: var(--font-ubuntu);
  }
</style>
