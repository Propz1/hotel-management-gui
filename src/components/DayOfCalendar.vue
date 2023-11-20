<template>
  <div :class="$style.dayofcalendar">
    <StateInRangeToday
      stateModelNumber="01"
      propBorderRadius="unset"
      propBackgroundColor="unset"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#000"
      atomDateLeft="calc(50% - 9px)"
    /><StateInRangeToday
      stateModelNumber="01"
      propBorderRadius="2px"
      propBackgroundColor="#ef5da8"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#000"
      atomDateLeft="calc(50% - 9px)"
    /><StateInRangeToday
      stateModelNumber="01"
      propBorderRadius="2px"
      propBackgroundColor="#6c80e7"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#dccfcf"
      atomDateLeft="calc(50% - 9px)"
    /><StateInRangeToday
      stateModelNumber="01"
      propBorderRadius="unset"
      propBackgroundColor="rgba(255, 255, 255, 0.2)"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#9d9ca3"
      atomDateLeft="calc(50% - 9px)"
    /><StateInRangeToday
      stateModelNumber="01"
      propBorderRadius="unset"
      propBackgroundColor="unset"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#ec1717"
      atomDateLeft="calc(50% - 9px)"
    /><StateInRangeToday
      stateModelNumber="01"
      propBorderRadius="unset"
      propBackgroundColor="unset"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#0e8521"
      atomDateLeft="calc(50% - 9px)"
    /><StateInRangeToday
      stateModelNumber="01"
      propBorderRadius="30px"
      propBackgroundColor="rgba(68, 153, 238, 0.75)"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#000"
      atomDateLeft="calc(50% - 9px)"
    /><StateInRangeToday
      stateModelNumber="01"
      propBorderRadius="30px"
      propBackgroundColor="rgba(68, 153, 238, 0.75)"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#000"
      atomDateLeft="calc(50% - 9px)"
    /><StateInRangeToday stateModelNumber="01" />
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import StateInRangeToday from "./StateInRangeToday.vue";

  export default defineComponent({
    name: "DayOfCalendar",
    components: { StateInRangeToday },
  });
</script>
<style module>
  .dayofcalendar {
    border-radius: var(--br-8xs);
    border: 1px dashed #9747ff;
    box-sizing: border-box;
    width: 280px;
    height: 46px;
    overflow: hidden;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    padding: var(--padding-xl);
    gap: 3px;
  }
</style>
