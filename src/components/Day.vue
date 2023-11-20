<template>
  <div :class="$style.day" :style="dayStyle">
    <div :class="$style.separator" />
    <div :class="$style.day1" :style="day1Style">
      <div :class="$style.div" :style="divStyle">
        {{numberDaySchedulerOccupancyLabel}}
      </div>
      <div :class="$style.mon" :style="monStyle">
        {{dayWeekSchedulerOccupancyLabel}}
      </div>
    </div>
    <div :class="$style.separator" />
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "Day",
    props: {
      numberDaySchedulerOccupancyLabel: { type: String },
      dayWeekSchedulerOccupancyLabel: { type: String },
      dayScheduleTableOccupancyWidth: {
        type: [Object, Array, String, Number, Boolean],
      },
      dayScheduleTableOccupancyShrink: { type: String },
      dayScheduleTableOccupancyMarginLeft: {
        type: [Object, Array, String, Number, Boolean],
      },
      dayScheduleTableOccupancyBorderRadius: { type: String },
      dayScheduleTableOccupancyBackgroundColor: { type: String },
      divColor: { type: String },
      monColor: { type: String },
    },
    computed: {
      dayStyle() {
        return {
          width: this.dayScheduleTableOccupancyWidth,
          flexShrink: this.dayScheduleTableOccupancyShrink,
          marginLeft: this.dayScheduleTableOccupancyMarginLeft,
        };
      },
      day1Style() {
        return {
          borderRadius: this.dayScheduleTableOccupancyBorderRadius,
          backgroundColor: this.dayScheduleTableOccupancyBackgroundColor,
        };
      },
      divStyle() {
        return {
          color: this.divColor,
        };
      },
      monStyle() {
        return {
          color: this.monColor,
        };
      },
    },
  });
</script>
<style module>
  .separator {
    position: relative;
    border-radius: var(--br-12xs);
    background-color: var(--color-whitesmoke-200);
    width: 1px;
    height: 40px;
  }
  .div {
    position: relative;
    line-height: 28px;
    font-weight: 500;
  }
  .mon {
    position: relative;
    font-size: var(--text-xs-regular-size);
    line-height: 18px;
    color: var(--color-darkgray-200);
    text-align: center;
  }
  .day1 {
    flex: 1;
    height: 54px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }
  .day {
    background-color: var(--light-light-100);
    width: 55px;
    height: 60px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    padding: var(--padding-10xs) 0px;
    box-sizing: border-box;
    gap: var(--gap-11xs);
    text-align: left;
    font-size: var(--text-ig-semibold-size);
    color: var(--primary-footer);
    font-family: var(--text-xs-regular);
  }
</style>
