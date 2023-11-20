<template>
  <div :class="$style.hederdaysofweek" :style="hederDaysOfWeekStyle">
    <StateInRangeToday
      stateModelNumber="ПН"
      propBorderRadius="unset"
      propBackgroundColor="unset"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#0e8521"
      atomDateLeft="calc(50% - 12px)"
    /><StateInRangeToday
      stateModelNumber="ВТ"
      propBorderRadius="unset"
      propBackgroundColor="unset"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#0e8521"
      atomDateLeft="calc(50% - 10px)"
    /><StateInRangeToday
      stateModelNumber="СР"
      propBorderRadius="unset"
      propBackgroundColor="unset"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#0e8521"
      atomDateLeft="calc(50% - 11px)"
    /><StateInRangeToday
      stateModelNumber="ЧТ"
      propBorderRadius="unset"
      propBackgroundColor="unset"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#0e8521"
      atomDateLeft="calc(50% - 11px)"
    /><StateInRangeToday
      stateModelNumber="ПТ"
      propBorderRadius="unset"
      propBackgroundColor="unset"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#0e8521"
      atomDateLeft="calc(50% - 11px)"
    /><StateInRangeToday
      stateModelNumber="СБ"
      propBorderRadius="unset"
      propBackgroundColor="unset"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#ec1717"
      atomDateLeft="calc(50% - 11px)"
    /><StateInRangeToday
      stateModelNumber="ВС"
      propBorderRadius="unset"
      propBackgroundColor="unset"
      propCursor="unset"
      propBorder="unset"
      propPadding="unset"
      atomDateColor="#ec1717"
      atomDateLeft="calc(50% - 11px)"
    />
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import StateInRangeToday from "./StateInRangeToday.vue";

  export default defineComponent({
    name: "HederDaysOfWeek",
    components: { StateInRangeToday },
    props: {
      monday: { type: String },
      tuesday: { type: String },
      wednesday: { type: String },
      thursday: { type: String },
      friday: { type: String },
      saturday: { type: String },
      sunday: { type: String },
      propHeight: { type: [Object, Array, String, Number, Boolean] },
      propWidth: { type: [Object, Array, String, Number, Boolean] },
      propTop: { type: [Object, Array, String, Number, Boolean] },
      propRight: { type: [Object, Array, String, Number, Boolean] },
      propBottom: { type: [Object, Array, String, Number, Boolean] },
      propBackgroundColor: { type: String },
      propLeft: { type: [Object, Array, String, Number, Boolean] },
      propBackgroundColor1: { type: String },
      propCursor: { type: String },
      propBorder: { type: String },
      propPadding: { type: [Object, Array, String, Number, Boolean] },
      propBorderRadius: { type: String },
      propLeft1: { type: [Object, Array, String, Number, Boolean] },
      propColor: { type: String },
      propBackgroundColor2: { type: String },
      propBorderRadius1: { type: String },
      dayLeft: { type: [Object, Array, String, Number, Boolean] },
      dayColor: { type: String },
      propBackgroundColor3: { type: String },
      propBorderRadius2: { type: String },
      dayLeft1: { type: [Object, Array, String, Number, Boolean] },
      dayColor1: { type: String },
      propBackgroundColor4: { type: String },
      dayLeft2: { type: [Object, Array, String, Number, Boolean] },
      dayColor2: { type: String },
      propBorderRadius3: { type: String },
      propBackgroundColor5: { type: String },
      dayLeft3: { type: [Object, Array, String, Number, Boolean] },
      dayColor3: { type: String },
      propBorderRadius4: { type: String },
      propBackgroundColor6: { type: String },
      dayLeft4: { type: [Object, Array, String, Number, Boolean] },
      dayColor4: { type: String },
      propBorderRadius5: { type: String },
      propBackgroundColor7: { type: String },
      dayLeft5: { type: [Object, Array, String, Number, Boolean] },
      dayColor5: { type: String },
    },
    computed: {
      hederDaysOfWeekStyle() {
        return {
          height: this.propHeight,
          width: this.propWidth,
          top: this.propTop,
          right: this.propRight,
          bottom: this.propBottom,
          backgroundColor: this.propBackgroundColor,
          left: this.propLeft,
        };
      },
      stateInRangeTodayStyle() {
        return {
          backgroundColor: this.propBackgroundColor1,
          cursor: this.propCursor,
          border: this.propBorder,
          padding: this.propPadding,
          borderRadius: this.propBorderRadius,
        };
      },
      div3Style() {
        return {
          left: this.propLeft1,
          color: this.propColor,
        };
      },
      stateInRangeTodayStyle1() {
        return {
          backgroundColor: this.propBackgroundColor2,
          borderRadius: this.propBorderRadius1,
        };
      },
      div2Style() {
        return {
          left: this.dayLeft,
          color: this.dayColor,
        };
      },
      stateInRangeTodayStyle2() {
        return {
          backgroundColor: this.propBackgroundColor3,
          borderRadius: this.propBorderRadius2,
        };
      },
      div2Style1() {
        return {
          left: this.dayLeft1,
          color: this.dayColor1,
        };
      },
      stateInRangeTodayStyle3() {
        return {
          backgroundColor: this.propBackgroundColor4,
        };
      },
      div2Style2() {
        return {
          left: this.dayLeft2,
          color: this.dayColor2,
        };
      },
      stateInRangeTodayStyle4() {
        return {
          borderRadius: this.propBorderRadius3,
          backgroundColor: this.propBackgroundColor5,
        };
      },
      div2Style3() {
        return {
          left: this.dayLeft3,
          color: this.dayColor3,
        };
      },
      stateInRangeTodayStyle5() {
        return {
          borderRadius: this.propBorderRadius4,
          backgroundColor: this.propBackgroundColor6,
        };
      },
      div2Style4() {
        return {
          left: this.dayLeft4,
          color: this.dayColor4,
        };
      },
      stateInRangeTodayStyle6() {
        return {
          borderRadius: this.propBorderRadius5,
          backgroundColor: this.propBackgroundColor7,
        };
      },
      div2Style5() {
        return {
          left: this.dayLeft5,
          color: this.dayColor5,
        };
      },
    },
  });
</script>
<style module>
  .hederdaysofweek {
    position: absolute;
    height: 15.73%;
    width: 99.04%;
    top: 0%;
    right: 0.46%;
    bottom: 84.27%;
    left: 0.49%;
    background-color: var(--light-light-100);
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    padding: 0px var(--padding-xs);
    box-sizing: border-box;
  }
</style>
