<template>
  <div :class="$style.week">
    <AtomDate
      carModelNumber="26"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 10px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="27"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 9px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="28"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 10px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="29"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 10px)"
      divDisplay="inline-block"
    /><StateInRangeToday
      stateModelNumber="30"
      propBorderRadius="30px"
      propBackgroundColor="rgba(68, 153, 238, 0.75)"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#000"
      atomDateLeft="calc(50% - 10px)"
    /><StateInRangeToday
      stateModelNumber="31"
      propBorderRadius="30px"
      propBackgroundColor="rgba(68, 153, 238, 0.75)"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#000"
      atomDateLeft="calc(50% - 9px)"
    /><DayOfCalendarInRangeDisable
      carNumber="1"
      atomDateLeft="calc(50% - 4px)"
    />
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import AtomDate from "./AtomDate.vue";
  import StateInRangeToday from "./StateInRangeToday.vue";
  import DayOfCalendarInRangeDisable from "./DayOfCalendarInRangeDisable.vue";

  export default defineComponent({
    name: "Week3",
    components: { AtomDate, StateInRangeToday, DayOfCalendarInRangeDisable },
  });
</script>
<style module>
  .week {
    position: absolute;
    height: 17.04%;
    width: 99.51%;
    top: 82.96%;
    right: 0%;
    bottom: 0%;
    left: 0.49%;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    padding: 0px var(--padding-xs);
    box-sizing: border-box;
  }
</style>
