<template>
  <div :class="$style.dayofcalendarinrangedisable">
    <AtomDate
      carModelNumber="01"
      atomDatePosition="relative"
      divColor="#9d9ca3"
      divLeft="calc(50% - 9px)"
      divDisplay="inline-block"
    />
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import AtomDate from "./AtomDate.vue";

  export default defineComponent({
    name: "DayOfCalendarInRangeDisable",
    components: { AtomDate },
    props: {
      carNumber: { type: String },
      atomDateLeft: { type: [Object, Array, String, Number, Boolean] },
    },
    computed: {
      div1Style() {
        return {
          left: this.atomDateLeft,
        };
      },
    },
  });
</script>
<style module>
  .dayofcalendarinrangedisable {
    border-radius: var(--br-5xs);
    background-color: var(--enddate);
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
  }
</style>
