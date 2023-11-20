<template>
  <div :class="$style.week">
    <AtomDate
      carModelNumber="12"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 9px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="13"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 9px)"
      divDisplay="inline-block"
    /><StateInRangeToday
      stateModelNumber="14"
      propBorderRadius="2px"
      propBackgroundColor="#ef5da8"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#000"
      atomDateLeft="calc(50% - 9px)"
    /><AtomDate
      carModelNumber="15"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 9px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="16"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 9px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="17"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 8px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="18"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 9px)"
      divDisplay="inline-block"
    />
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import AtomDate from "./AtomDate.vue";
  import StateInRangeToday from "./StateInRangeToday.vue";

  export default defineComponent({
    name: "Week2",
    components: { AtomDate, StateInRangeToday },
  });
</script>
<style module>
  .week {
    position: absolute;
    height: 17.04%;
    width: 99.51%;
    top: 49.51%;
    right: 0.49%;
    bottom: 33.45%;
    left: 0%;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    padding: 0px var(--padding-xs);
    box-sizing: border-box;
  }
</style>
