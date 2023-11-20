<template>
  <div :class="$style.guests">
    <div :class="$style.darkDesktopSwitchesSlParent" :style="frameDivStyle">
      <DarkDesktopSwitchesSl
        guestsCheckBoxPosition="relative"
        guestsCheckBoxBorderRadius="4px"
        guestsCheckBoxBackgroundColor="#fff"
      />
      <div :class="$style.frameChild" />
      <div :class="$style.nameParent">
        <div :class="$style.name">ФИО</div>
        <div :class="$style.tel" :style="telStyle">+7</div>
      </div>
    </div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import DarkDesktopSwitchesSl from "./DarkDesktopSwitchesSl.vue";

  export default defineComponent({
    name: "Guests1",
    components: { DarkDesktopSwitchesSl },
    props: {
      propTop: { type: [Object, Array, String, Number, Boolean] },
      propHeight: { type: [Object, Array, String, Number, Boolean] },
      propWidth: { type: [Object, Array, String, Number, Boolean] },
    },
    computed: {
      frameDivStyle() {
        return {
          top: this.propTop,
          height: this.propHeight,
        };
      },
      telStyle() {
        return {
          width: this.propWidth,
        };
      },
    },
  });
</script>
<style module>
  .frameChild {
    position: relative;
    border-radius: 50%;
    background: linear-gradient(#c4c4c4, #c4c4c4), #d9d9d9;
    width: 34px;
    height: 34px;
  }
  .name {
    position: relative;
    line-height: 20px;
  }
  .tel {
    position: relative;
    font-size: var(--font-size-3xs);
    line-height: 14px;
    color: var(--color-gray-400);
  }
  .nameParent {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
    gap: var(--gap-11xs);
  }
  .darkDesktopSwitchesSlParent {
    position: absolute;
    top: 10px;
    left: 20px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-sm);
  }
  .guests {
    position: relative;
    background-color: var(--light-light-100);
    border: 1px solid var(--color-whitesmoke-200);
    box-sizing: border-box;
    width: 299px;
    height: 60px;
    margin-top: -1px;
    text-align: left;
    font-size: var(--label-large-label-14-size);
    color: var(--base-pure-black);
    font-family: var(--text-xs-regular);
  }
</style>
