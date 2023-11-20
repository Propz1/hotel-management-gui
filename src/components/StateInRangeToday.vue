<template>
  <div :class="$style.stateinrangetoday" :style="stateInRangeTodayStyle">
    <AtomDate
      carModelNumber="01"
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

  export default defineComponent({
    name: "StateInRangeToday",
    components: { AtomDate },
    props: {
      stateModelNumber: { type: String },
      propBorderRadius: { type: String },
      propBackgroundColor: { type: String },
      propCursor: { type: String },
      propBorder: { type: String },
      propPadding: { type: [Object, Array, String, Number, Boolean] },
      atomDateColor: { type: String },
      atomDateLeft: { type: [Object, Array, String, Number, Boolean] },
    },
    computed: {
      stateInRangeTodayStyle() {
        return {
          borderRadius: this.propBorderRadius,
          backgroundColor: this.propBackgroundColor,
          cursor: this.propCursor,
          border: this.propBorder,
          padding: this.propPadding,
        };
      },
      div2Style() {
        return {
          color: this.atomDateColor,
          left: this.atomDateLeft,
        };
      },
    },
  });
</script>
<style module>
  .stateinrangetoday {
    border-radius: var(--br-11xl);
    background-color: var(--fuschia-100);
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
  }
</style>
