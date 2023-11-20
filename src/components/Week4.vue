<template>
  <div :class="$style.week" :style="weekStyle">
    <AtomDate
      carModelNumber="01"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 9px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="01"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 9px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="01"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 9px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="01"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 9px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="01"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 9px)"
      divDisplay="inline-block"
    /><AtomDate
      carModelNumber="01"
      atomDatePosition="relative"
      divColor="#000"
      divLeft="calc(50% - 9px)"
      divDisplay="inline-block"
    /><AtomDate
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
    name: "Week4",
    components: { AtomDate },
    props: {
      dateCode: { type: String },
      modelCode: { type: String },
      carCode: { type: String },
      carCodeText: { type: String },
      carCodeNumber: { type: String },
      carCodeText2: { type: String },
      carCodeImage: { type: String },
      weekWidth: { type: [Object, Array, String, Number, Boolean] },
      weekHeight: { type: [Object, Array, String, Number, Boolean] },
      weekPosition: { type: String },
      weekTop: { type: [Object, Array, String, Number, Boolean] },
      weekRight: { type: [Object, Array, String, Number, Boolean] },
      weekBottom: { type: [Object, Array, String, Number, Boolean] },
      weekLeft: { type: [Object, Array, String, Number, Boolean] },
      dayLeft: { type: [Object, Array, String, Number, Boolean] },
      dayLeft1: { type: [Object, Array, String, Number, Boolean] },
      dayLeft2: { type: [Object, Array, String, Number, Boolean] },
      dayLeft3: { type: [Object, Array, String, Number, Boolean] },
      dayLeft4: { type: [Object, Array, String, Number, Boolean] },
    },
    computed: {
      weekStyle() {
        return {
          width: this.weekWidth,
          height: this.weekHeight,
          position: this.weekPosition,
          top: this.weekTop,
          right: this.weekRight,
          bottom: this.weekBottom,
          left: this.weekLeft,
        };
      },
      div1Style() {
        return {
          left: this.dayLeft,
        };
      },
      div1Style1() {
        return {
          left: this.dayLeft1,
        };
      },
      div1Style2() {
        return {
          left: this.dayLeft2,
        };
      },
      div1Style3() {
        return {
          left: this.dayLeft3,
        };
      },
      div1Style4() {
        return {
          left: this.dayLeft4,
        };
      },
    },
  });
</script>
<style module>
  .week {
    width: 210px;
    height: 35px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    padding: 0px var(--padding-xs);
    box-sizing: border-box;
  }
</style>
