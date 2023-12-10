<template>
  <div :class="$style.datePicker">
    <v-text-field
      :class="$style.daterangeoneVtextfield"
      name="DateRangeOne"
      label="Начало периода"
      color="#5d5fef"
      style="background-color: white;"
      type="date"
    />
    <v-text-field
      :class="$style.daterangetwoVtextfield"
      name="DateRangeTwo"
      label="Окончание периода"
      color="#5d5fef"
      style="background-color:white;"
      type="date"
    />


    <!-- <img
      :class="$style.arrowForwardMultimediaSvgreIcon"
      alt=""
      src="/arrowforwardmultimediasvgrepocom-1.svg"
    /> -->
    <!-- <button :class="$style.buttoncalendaricon" id="selectCalendarDateButton">
      <div :class="$style.shadow" />
      <div :class="$style.bg" />
      <img :class="$style.calendarIcon" alt="" src="/calendar-icon.svg" />
    </button> -->
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "DatePicker",
  });
</script>
<style module>

.datePicker {
    position: absolute;
    height: 5.19%;
    /* width: 16.23%; */
    width: 358px;
    /* top: 55.19%; */
    top: 70px;
    /* left: 52.99%; */
    left: 260px;
    right: 10.21%;
    bottom: 39.63%; 
    background-color: var(--light-light-100);
    overflow: hidden; 
    border-radius: 2rem;
    color:rgb(61, 56, 122);
  }
  .daterangetwoVtextfield {
    position: absolute;
    top: calc(50% - 29px);
    left: calc(50% - 0px);
  }
  .daterangeoneVtextfield {
    position: absolute;
    top: 0px;
    left: 0px; 
  }

  /* .arrowForwardMultimediaSvgreIcon {
    position: absolute;
    top: 21px;
    left: 183px;
    width: 15px;
    height: 15px;
    overflow: hidden;
  } */
  /* .shadow {
    position: absolute;
    height: calc(100% - 16px);
    width: calc(100% - 16px);
    top: 16px;
    right: 8px;
    bottom: 0px;
    left: 8px;
    background-color: var(--light-light-100);
  }
  .bg {
    position: absolute;
    height: 100%;
    width: 100%;
    top: 0px;
    right: 0px;
    bottom: 0px;
    left: 0px;
    background-color: var(--dark-dark-400);
  }
  .calendarIcon {
    position: absolute;
    top: 8px;
    left: 8px;
    width: 16px;
    height: 16px;
    overflow: hidden;
  }
  .buttoncalendaricon {
    cursor: pointer;
    border: none;
    padding: 0;
    background-color: transparent;
    position: absolute;
    top: calc(50% - 16px);
    right: 16px;
    width: 32px;
    height: 32px;
  } */

</style>
