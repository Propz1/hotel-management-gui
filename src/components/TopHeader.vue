<template>
  <div :class="$style.topHeader">
    <div :class="$style.logoText">
      <b :class="$style.logo">Ormand</b>
      <b :class="$style.workit">
        <span>Workit</span>
        <span :class="$style.span">.</span>
      </b>
    </div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "TopHeader",
  });
</script>
<style module>
  .logo {
    position: relative;
    line-height: 120%;
  }
  .span {
    color: var(--color-blue);
  }
  .workit {
    align-self: stretch;
    flex: 1;
    position: relative;
    font-size: var(--h6-size);
    line-height: 100%;
    display: none;
    color: var(--light-light-100);
  }
  .logoText {
    width: 70px;
    height: 20px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    gap: var(--gap-xs);
  }
  .topHeader {
    align-self: stretch;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    padding: var(--padding-5xl) var(--padding-5xl) var(--padding-9xl);
    text-align: left;
    font-size: var(--font-size-9xl);
    color: var(--primary-blue-900);
    font-family: var(--font-ubuntu);
  }
</style>
