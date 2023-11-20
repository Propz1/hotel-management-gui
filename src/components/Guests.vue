<template>
  <div :class="$style.guests">
    <div :class="$style.darkDesktopSwitchesSlParent">
      <DarkDesktopSwitchesSl
        guestsCheckBoxPosition="relative"
        guestsCheckBoxBorderRadius="4px"
        guestsCheckBoxBackgroundColor="#fff"
      />
      <div :class="$style.frameChild" />
      <div :class="$style.name">ФИО</div>
    </div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import DarkDesktopSwitchesSl from "./DarkDesktopSwitchesSl.vue";

  export default defineComponent({
    name: "Guests",
    components: { DarkDesktopSwitchesSl },
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
  .darkDesktopSwitchesSlParent {
    position: absolute;
    top: calc(50% - 17px);
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
