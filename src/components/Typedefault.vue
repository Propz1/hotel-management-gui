<template>
  <div :class="$style.typedefault" :style="typedefaultStyle" />
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "Typedefault",
    props: {
      cellSchedulerTableOccupancyBackgroundColor: { type: String },
      cellSchedulerTableOccupancyHeight: {
        type: [Object, Array, String, Number, Boolean],
      },
      cellSchedulerTableOccupancyMarginLeft: {
        type: [Object, Array, String, Number, Boolean],
      },
    },
    computed: {
      typedefaultStyle() {
        return {
          backgroundColor: this.cellSchedulerTableOccupancyBackgroundColor,
          height: this.cellSchedulerTableOccupancyHeight,
          marginLeft: this.cellSchedulerTableOccupancyMarginLeft,
        };
      },
    },
  });
</script>
<style module>
  .typedefault {
    background-color: var(--color-gray-200);
    border: 1px solid var(--color-whitesmoke-200);
    box-sizing: border-box;
    width: 66px;
  }
</style>
