<template>
  <footer
    :class="$style.scrollbarClassic"
    data-scroll-to="scrollbarClassic1"
    :onClick="onScrollbarClassic1Click"
  >
    <img :class="$style.arrowLeftIcon" alt="" src="/arrow-left.svg" />
    <img :class="$style.arrowRightIcon" alt="" src="/arrow-right.svg" />
    <div :class="$style.bar" />
  </footer>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "ScrollbarClassic2",
    methods: {
      onScrollbarClassic1Click() {
        const anchor = document.querySelector(
          "[data-scroll-to='scrollbarClassic1']"
        );
        if (anchor) {
          anchor.scrollIntoView({ block: "center", behavior: "smooth" });
        }
      },
    },
  });
</script>
<style module>
  .arrowLeftIcon {
    position: absolute;
    top: calc(50% - 3.5px);
    left: 6px;
    width: 7px;
    height: 7px;
  }
  .arrowRightIcon {
    position: absolute;
    top: calc(50% - 3.5px);
    right: 6px;
    width: 7px;
    height: 7px;
  }
  .bar {
    position: absolute;
    top: calc(50% + 6.5px);
    left: calc(50% - 13px);
    background-color: var(--scrollbutton);
    width: 13px;
    height: 25px;
    transform: rotate(-90deg);
    transform-origin: 0 0;
  }
  .scrollbarClassic {
    position: relative;
    background-color: var(--scrollbar);
    width: 1326px;
    height: 17px;
    cursor: pointer;
  }
</style>
