<template>
  <div :class="$style.buttons">
    <div :class="$style.buttonSolidRoundedDef">
      <div :class="$style.default">Default</div>
    </div>
    <ButtonSolidRoundedHov /><ButtonSolidRoundedAct /><ButtonSolidRoundedDis /><ButtonSolidRoundedFoc /><ButtonOutlinedRounded4 /><ButtonOutlinedRounded3 /><ButtonOutlinedRounded2 /><ButtonOutlinedRounded1 /><ButtonOutlinedRounded /><ButtonOutlinedSquareF /><ButtonOutlinedSquareD1
      buttonText="Default"
    /><ButtonOutlinedSquareH /><ButtonOutlinedSquareA /><ButtonOutlinedSquareD /><ButtonSolidSquareDefa /><ButtonSolidSquareHove /><ButtonSolidSquareActi /><ButtonSolidSquareDisa /><ButtonSolidSquareFocu />
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import ButtonSolidRoundedHov from "./ButtonSolidRoundedHov.vue";
  import ButtonSolidRoundedAct from "./ButtonSolidRoundedAct.vue";
  import ButtonSolidRoundedDis from "./ButtonSolidRoundedDis.vue";
  import ButtonSolidRoundedFoc from "./ButtonSolidRoundedFoc.vue";
  import ButtonOutlinedRounded4 from "./ButtonOutlinedRounded4.vue";
  import ButtonOutlinedRounded3 from "./ButtonOutlinedRounded3.vue";
  import ButtonOutlinedRounded2 from "./ButtonOutlinedRounded2.vue";
  import ButtonOutlinedRounded1 from "./ButtonOutlinedRounded1.vue";
  import ButtonOutlinedRounded from "./ButtonOutlinedRounded.vue";
  import ButtonOutlinedSquareF from "./ButtonOutlinedSquareF.vue";
  import ButtonOutlinedSquareD1 from "./ButtonOutlinedSquareD1.vue";
  import ButtonOutlinedSquareH from "./ButtonOutlinedSquareH.vue";
  import ButtonOutlinedSquareA from "./ButtonOutlinedSquareA.vue";
  import ButtonOutlinedSquareD from "./ButtonOutlinedSquareD.vue";
  import ButtonSolidSquareDefa from "./ButtonSolidSquareDefa.vue";
  import ButtonSolidSquareHove from "./ButtonSolidSquareHove.vue";
  import ButtonSolidSquareActi from "./ButtonSolidSquareActi.vue";
  import ButtonSolidSquareDisa from "./ButtonSolidSquareDisa.vue";
  import ButtonSolidSquareFocu from "./ButtonSolidSquareFocu.vue";

  export default defineComponent({
    name: "Buttons",
    components: {
      ButtonSolidRoundedHov,
      ButtonSolidRoundedAct,
      ButtonSolidRoundedDis,
      ButtonSolidRoundedFoc,
      ButtonOutlinedRounded4,
      ButtonOutlinedRounded3,
      ButtonOutlinedRounded2,
      ButtonOutlinedRounded1,
      ButtonOutlinedRounded,
      ButtonOutlinedSquareF,
      ButtonOutlinedSquareD1,
      ButtonOutlinedSquareH,
      ButtonOutlinedSquareA,
      ButtonOutlinedSquareD,
      ButtonSolidSquareDefa,
      ButtonSolidSquareHove,
      ButtonSolidSquareActi,
      ButtonSolidSquareDisa,
      ButtonSolidSquareFocu,
    },
  });
</script>
<style module>
  .default {
    position: relative;
    line-height: 24px;
    font-weight: 500;
  }
  .buttonSolidRoundedDef {
    position: absolute;
    top: 120px;
    left: 0px;
    background-color: var(--primary-main-color);
    height: 32px;
    overflow: hidden;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    padding: var(--padding-5xs) var(--padding-base);
    box-sizing: border-box;
  }
  .buttons {
    width: 487px;
    height: 212px;
    text-align: center;
    font-size: var(--label-large-label-14-size);
    color: var(--light-light-100);
    font-family: var(--body-body-16);
  }
</style>
