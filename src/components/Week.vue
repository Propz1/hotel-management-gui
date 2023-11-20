<template>
  <div :class="$style.week">
    <StateInRangeToday
      stateModelNumber="28"
      propBorderRadius="unset"
      propBackgroundColor="rgba(255, 255, 255, 0.2)"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#9d9ca3"
      atomDateLeft="calc(50% - 10px)"
    /><StateInRangeToday
      stateModelNumber="29"
      propBorderRadius="unset"
      propBackgroundColor="rgba(255, 255, 255, 0.2)"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#9d9ca3"
      atomDateLeft="calc(50% - 10px)"
    /><StateInRangeToday
      stateModelNumber="30"
      propBorderRadius="unset"
      propBackgroundColor="rgba(255, 255, 255, 0.2)"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#9d9ca3"
      atomDateLeft="calc(50% - 10px)"
    /><AtomDate
      carModelNumber="1"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 4px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="2"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 5px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="3"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 5px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="4"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 5px)"
      divDisplay="inline-block"
    />
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import StateInRangeToday from "./StateInRangeToday.vue";
  import AtomDate from "./AtomDate.vue";

  export default defineComponent({
    name: "Week",
    components: { StateInRangeToday, AtomDate },
  });
</script>
<style module>
  .week {
    position: absolute;
    height: 16.89%;
    width: 99.51%;
    top: 15.73%;
    right: 0%;
    bottom: 67.38%;
    left: 0.49%;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    padding: 0px var(--padding-xs);
    box-sizing: border-box;
  }
</style>
