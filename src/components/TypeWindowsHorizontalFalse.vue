<template>
  <div :class="$style.typewindowsHorizontalfalse">
    <img :class="$style.arrowBottomIcon" alt="" src="/arrow-bottom.svg" />
    <div :class="$style.thumb">
      <div :class="$style.vspacer">
        <p :class="$style.blankLine">&nbsp;</p>
        <p :class="$style.blankLine">&nbsp;</p>
        <p :class="$style.blankLine">&nbsp;</p>
        <p :class="$style.blankLine">&nbsp;</p>
      </div>
    </div>
    <img :class="$style.arrowTopIcon" alt="" src="/arrow-top.svg" />
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "TypeWindowsHorizontalFalse",
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
  .blankLine {
    margin: 0;
  }
  .vspacer {
    position: relative;
    display: inline-block;
    width: 13px;
  }
  .thumb {
    position: absolute;
    bottom: 14px;
    left: calc(50% - 6.5px);
    background-color: var(--scrollbutton);
    width: 13px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
  }
  .arrowTopIcon {
    position: absolute;
    top: 6px;
    left: calc(50% - 3.5px);
    width: 7px;
    height: 7px;
  }
  .typewindowsHorizontalfalse {
    position: absolute;
    top: 16px;
    left: 66px;
    background-color: var(--gray-100);
    width: 17px;
    height: 150px;
    text-align: left;
    font-size: var(--font-size-9xs);
    color: var(--scrollbutton);
    font-family: var(--font-roboto);
  }
</style>
