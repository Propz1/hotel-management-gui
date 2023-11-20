<template>
  <div :class="$style.scrollbarClassic">
    <TypeWindowsHorizontalFalse1
      dimensionCode="/typewindows-horizontalfalse-positionstart.svg"
      propLeft="16px"
      propWidth="17px"
      propHeight="150px"
      propTop="16px"
    /><TypeMacHorizontalFalsePo
      propLeft="281px"
      propTop="16px"
      propBorderRight="1px solid var(--color-gainsboro-100)"
      propBorderLeft="1px solid var(--color-gainsboro-100)"
      propWidth="12px"
      propHeight="150px"
      propBorderTop="unset"
      propBorderBottom="unset"
      propBottom="unset"
      propTop1="2px"
      propLeft1="calc(50% - 3px)"
      propTransform="unset"
      propTransformOrigin="unset"
      propRight="unset"
    /><TypeWindowsHorizontalFalse1
      dimensionCode="/typewindows-horizontaltrue-positionstart.svg"
      propLeft="99px"
      propWidth="150px"
      propHeight="17px"
      propTop="16px"
    /><TypeMacHorizontalFalsePo
      propLeft="365px"
      propTop="47px"
      propBorderRight="unset"
      propBorderLeft="unset"
      propWidth="150px"
      propHeight="12px"
      propBorderTop="1px solid var(--color-gainsboro-100)"
      propBorderBottom="1px solid var(--color-gainsboro-100)"
      propBottom="unset"
      propTop1="calc(50% + 3px)"
      propLeft1="2px"
      propTransform=" rotate(-90deg)"
      propTransformOrigin="0 0"
      propRight="unset"
    /><TypeWindowsHorizontalFalse1
      dimensionCode="/typewindows-horizontaltrue-positionmiddle.svg"
      propLeft="99px"
      propWidth="150px"
      propHeight="17px"
      propTop="41px"
    /><TypeMacHorizontalFalsePo
      propLeft="365px"
      propTop="72px"
      propBorderRight="unset"
      propBorderLeft="unset"
      propWidth="150px"
      propHeight="12px"
      propBorderTop="1px solid var(--color-gainsboro-100)"
      propBorderBottom="1px solid var(--color-gainsboro-100)"
      propBottom="unset"
      propTop1="calc(50% + 3px)"
      propLeft1="calc(50% - 12px)"
      propTransform=" rotate(-90deg)"
      propTransformOrigin="0 0"
      propRight="unset"
    /><TypeWindowsHorizontalFalse1
      dimensionCode="/typewindows-horizontaltrue-positionend.svg"
      propLeft="99px"
      propWidth="150px"
      propHeight="17px"
      propTop="66px"
    /><TypeMacHorizontalFalsePo
      propLeft="365px"
      propTop="97px"
      propBorderRight="unset"
      propBorderLeft="unset"
      propWidth="150px"
      propHeight="12px"
      propBorderTop="1px solid var(--color-gainsboro-100)"
      propBorderBottom="1px solid var(--color-gainsboro-100)"
      propBottom="unset"
      propTop1="calc(50% + 3px)"
      propLeft1="unset"
      propTransform=" rotate(-90deg)"
      propTransformOrigin="0 0"
      propRight="21px"
    /><TypeWindowsHorizontalFalse1
      dimensionCode="/typewindows-horizontalfalse-positionmiddle.svg"
    /><TypeMacHorizontalFalsePo
      propLeft="309px"
      propTop="16px"
      propBorderRight="1px solid var(--color-gainsboro-100)"
      propBorderLeft="1px solid var(--color-gainsboro-100)"
      propWidth="12px"
      propHeight="150px"
      propBorderTop="unset"
      propBorderBottom="unset"
      propBottom="unset"
      propTop1="calc(50% - 12px)"
      propLeft1="calc(50% - 3px)"
      propTransform="unset"
      propTransformOrigin="unset"
      propRight="unset"
    /><TypeWindowsHorizontalFalse /><TypeMacHorizontalFalsePo />
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import TypeWindowsHorizontalFalse1 from "./TypeWindowsHorizontalFalse1.vue";
  import TypeMacHorizontalFalsePo from "./TypeMacHorizontalFalsePo.vue";
  import TypeWindowsHorizontalFalse from "./TypeWindowsHorizontalFalse.vue";

  export default defineComponent({
    name: "ScrollbarClassic",
    components: {
      TypeWindowsHorizontalFalse1,
      TypeMacHorizontalFalsePo,
      TypeWindowsHorizontalFalse,
    },
  });
</script>
<style module>
  .scrollbarClassic {
    border-radius: var(--br-8xs);
    border: 1px dashed var(--color-mediumslateblue-100);
    box-sizing: border-box;
    width: 799px;
    height: 419px;
    overflow: hidden;
  }
</style>
