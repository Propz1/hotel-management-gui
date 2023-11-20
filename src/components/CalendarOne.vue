<template>
  <div :class="$style.calendarone">
    <Week3 /><Week4
      dateCode="19"
      modelCode="20"
      carCode="21"
      carCodeText="22"
      carCodeNumber="23"
      carCodeText2="24"
      carCodeImage="25"
      weekWidth="99.51%"
      weekHeight="16.4%"
      weekPosition="absolute"
      weekTop="66.55%"
      weekRight="0.49%"
      weekBottom="17.04%"
      weekLeft="0%"
      dayLeft="calc(50% - 10px)"
      dayLeft1="calc(50% - 10px)"
      dayLeft2="calc(50% - 10px)"
      dayLeft3="calc(50% - 10px)"
      dayLeft4="calc(50% - 10px)"
    /><Week2 /><Week1 /><Week /><HederDaysOfWeek
      monday="ПН"
      tuesday="ВТ"
      wednesday="СР"
      thursday="ЧТ"
      friday="ПТ"
      saturday="СБ"
      sunday="ВС"
    />
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import Week3 from "./Week3.vue";
  import Week4 from "./Week4.vue";
  import Week2 from "./Week2.vue";
  import Week1 from "./Week1.vue";
  import Week from "./Week.vue";
  import HederDaysOfWeek from "./HederDaysOfWeek.vue";

  export default defineComponent({
    name: "CalendarOne",
    components: { Week3, Week4, Week2, Week1, Week, HederDaysOfWeek },
  });
</script>
<style module>
  .calendarone {
    background-color: var(--light-light-100);
    width: 324px;
    height: 267px;
  }
</style>
