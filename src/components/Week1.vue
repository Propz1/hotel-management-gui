<template>
  <div :class="$style.week">
    <AtomDate
      carModelNumber="5"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 5px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="6"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 5px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="7"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 5px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="8"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 5px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="9"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 5px)"
      divDisplay="inline-block"
    /><StateInRangeToday
      stateModelNumber="10"
      propBorderRadius="2px"
      propBackgroundColor="#6c80e7"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#dccfcf"
      atomDateLeft="calc(50% - 9px)"
    /><AtomDate
      carModelNumber="11"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 7px)"
      divDisplay="inline-block"
    />
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import AtomDate from "./AtomDate.vue";
  import StateInRangeToday from "./StateInRangeToday.vue";

  export default defineComponent({
    name: "Week1",
    components: { AtomDate, StateInRangeToday },
  });
</script>
<style module>
  .week {
    position: absolute;
    height: 16.89%;
    width: 99.51%;
    top: 32.62%;
    right: 0.49%;
    bottom: 50.49%;
    left: 0%;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    padding: 0px var(--padding-xs);
    box-sizing: border-box;
  }
</style>
