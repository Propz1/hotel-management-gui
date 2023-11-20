<template>
  <div :class="$style.selectDay">
    <div :class="$style.selectdayeditbuttonParent">
      <img
        :class="$style.selectdayeditbuttonIcon"
        alt=""
        src="/selectdayeditbutton.svg"
      />
      <img
        :class="$style.selectdayeditbuttonIcon"
        alt=""
        src="/selectdaydeletebutton.svg"
      />
      <button :class="$style.selectdayclosebutton">
        <div :class="$style.selectdayclosebuttonChild" />
        <div :class="$style.group">
          <div :class="$style.rectangleCopy" />
          <div :class="$style.rectangleCopy1" />
          <div :class="$style.rectangleCopy2" />
        </div>
      </button>
    </div>
    <div :class="$style.frame">
      <div :class="$style.fullDay">
        <div :class="$style.fullDayChild" />
        <b :class="$style.b">Полный день</b>
      </div>
      <div :class="$style.div">8:00 – 16:00</div>
    </div>
    <div :class="$style.date">
      <img :class="$style.selectdayeditbuttonIcon" alt="" src="/icon3.svg" />
      <div :class="$style.april2023">8 April 2023 - 13 April 2023</div>
    </div>
    <div :class="$style.name">
      <img :class="$style.selectdayeditbuttonIcon" alt="" src="/icon4.svg" />
      <b :class="$style.april2023">ФИО</b>
    </div>
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "SelectDay",
  });
</script>
<style module>
  .selectdayeditbuttonIcon {
    position: relative;
    width: 24px;
    height: 24px;
    overflow: hidden;
    flex-shrink: 0;
  }
  .selectdayclosebuttonChild {
    position: absolute;
    height: 80.83%;
    width: 80.83%;
    top: 10%;
    right: 9.17%;
    bottom: 9.17%;
    left: 10%;
  }
  .rectangleCopy {
    position: absolute;
    height: 14.12%;
    width: 100%;
    top: 30%;
    right: -30%;
    bottom: 55.88%;
    left: 30%;
    background-color: var(--base-pure-black);
    transform: rotate(-45deg);
    transform-origin: 0 0;
  }
  .rectangleCopy1 {
    position: absolute;
    height: 14.12%;
    width: 42.94%;
    top: -40.59%;
    right: 16.47%;
    bottom: 126.47%;
    left: 40.59%;
    background-color: var(--base-pure-black);
    transform: rotate(45deg);
    transform-origin: 0 0;
  }
  .rectangleCopy2 {
    position: absolute;
    height: 14.12%;
    width: 42.94%;
    top: 0%;
    right: -23.53%;
    bottom: 85.88%;
    left: 80.59%;
    background-color: var(--base-pure-black);
    transform: rotate(45deg);
    transform-origin: 0 0;
  }
  .group {
    position: absolute;
    height: 70.83%;
    width: 70.83%;
    top: 50%;
    right: 29.17%;
    bottom: -20.83%;
    left: 0%;
    transform: rotate(-45deg);
    transform-origin: 0 0;
  }
  .selectdayclosebutton {
    cursor: pointer;
    border: none;
    padding: 0;
    background-color: transparent;
    position: relative;
    width: 24px;
    height: 24px;
  }
  .selectdayeditbuttonParent {
    position: absolute;
    top: 20px;
    right: 20px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-end;
    gap: 20px;
  }
  .fullDayChild {
    position: relative;
    border-radius: 50%;
    background-color: var(--color-teal-100);
    width: 8px;
    height: 8px;
  }
  .b {
    position: relative;
    line-height: 22px;
  }
  .fullDay {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-3xs);
    mix-blend-mode: normal;
  }
  .div {
    position: relative;
    font-size: var(--label-large-label-14-size);
    line-height: 18px;
    font-family: var(--text-xs-regular);
    color: var(--color-dimgray-100);
  }
  .frame {
    position: absolute;
    top: 20px;
    left: 20px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
    gap: 6px;
    font-size: var(--label-text-large-size);
  }
  .april2023 {
    position: relative;
    line-height: 18px;
  }
  .date {
    position: absolute;
    top: 86px;
    left: 20px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-3xs);
  }
  .name {
    position: absolute;
    top: 130px;
    left: 20px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-smi);
    font-family: var(--text-xs-regular);
  }
  .selectDay {
    background-color: var(--light-light-100);
    box-shadow: 0px 0px 20px rgba(112, 112, 112, 0.36);
    width: 340px;
    height: 189px;
    max-width: 100%;
    max-height: 100%;
    overflow: auto;
    text-align: left;
    font-size: var(--label-large-label-14-size);
    color: var(--base-pure-black);
    font-family: var(--font-ubuntu);
  }
</style>
