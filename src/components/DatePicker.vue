<template>
  <div :class="$style.datePicker">
    <v-text-field
      :class="$style.daterangeoneVtextfield"
      name="DateRangeOne"
      label="Начало периода"
      color="#5d5fef"
      style="background: conic-gradient(from 145deg at 40%  75%, rgba(148,179,207,0.26) 0%, rgba(156, 184, 216, 0.75) 30%, transparent 50%, rgba(156, 184, 216, 0.75) 80%, rgba(115, 128, 216, 0.51) 100%) 15% 75%/175% 200%,
                         radial-gradient(ellipse  at 55%  55%, rgb(178, 166, 222) 0%, rgb(9, 3, 22) 100%) 30% 30%/200% 170%;"
      type="date"
    />
    <v-text-field
      :class="$style.daterangetwoVtextfield"
      name="DateRangeTwo"
      label="Окончание периода"
      color="#5d5fef"      
      style="background: conic-gradient(from 145deg at 40%  75%, rgba(148,179,207,0.26) 0%, rgba(156, 184, 216, 0.75) 30%, transparent 50%, rgba(156, 184, 216, 0.75) 80%, rgba(115, 128, 216, 0.51) 100%) 15% 75%/175% 200%,
                         radial-gradient(ellipse  at 55%  55%, rgb(178, 166, 222) 0%, rgb(9, 3, 22) 100%) 30% 30%/200% 170%;"
      type="date"
    />

    <!-- style="background: linear-gradient(329deg, rgba(2,0,36,1) 0%, rgba(15,176,214,0) 16%, rgba(195,225,232,1) 28%, rgba(217,221,222,1) 63%);" -->

    <!-- style="background: linear-gradient(188deg, rgba(2,0,36,1) 0%, rgba(74,47,210,0) 32%, rgba(200,39,194,0) 55%, rgba(217,221,222,1) 83%);" -->



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
    height: 47px;
    /* height: 5.19%; */
    /* width: 16.23%; */
    width: 354px;
    /* top: 55.19%; */
    top: 0px;
    /* left: 52.99%; */
    left: 0px;
    /* right: 10.21%; */
    bottom: 0px; 
/*     background-color: var(--light-light-100); */ 

    overflow: hidden; 
    border-radius: 2rem;
    /* color:rgb(61, 56, 122); */
    color:var(--primary-color-text);
  }
  .daterangeoneVtextfield {
    position: absolute;
    top: -5.5px;
    left: 0px; 
  }
  .daterangetwoVtextfield {
    position: absolute;
    top: calc(50% - 30px);
    left: calc(50% + 20px);
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
