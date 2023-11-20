<template>
  <button
    :class="$style.scrollbarClassic"
    data-scroll-to="scrollbarClassic"
    :onClick="onScrollbarClassicClick"
  >
    <img :class="$style.arrowBottomIcon" alt="" src="/arrow-bottom1.svg" />
    <div :class="$style.bar" />
    <img :class="$style.arrowTopIcon" alt="" src="/arrow-top.svg" />
  </button>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "ScrollbarClassic1",
    methods: {
      onScrollbarClassicClick() {
        const anchor = document.querySelector(
          "[data-scroll-to='scrollbarClassic']"
        );
        if (anchor) {
          anchor.scrollIntoView({ block: "center", behavior: "smooth" });
        }
      },
    },
  });
</script>
<style module>
  .arrowBottomIcon {
    position: absolute;
    bottom: 6px;
    left: calc(50% - 3.5px);
    width: 7px;
    height: 7px;
  }
  .bar {
    position: absolute;
    top: calc(50% - 12.5px);
    left: calc(50% - 6.5px);
    background-color: var(--scrollbutton);
    width: 13px;
    height: 25px;
  }
  .arrowTopIcon {
    position: absolute;
    top: 6px;
    left: calc(50% - 3.5px);
    width: 7px;
    height: 7px;
  }
  .scrollbarClassic {
    cursor: pointer;
    border: none;
    padding: 0;
    background-color: var(--scrollbar);
    position: absolute;
    top: 32px;
    left: 299px;
    width: 19px;
    height: 573px;
  }
</style>
