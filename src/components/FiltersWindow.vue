<template>
  <div :class="$style.filterswindow">
    <div :class="$style.dropdownFilters">
      <div :class="$style.header">
        <b :class="$style.titelfilters">Фильтры</b>
        <button :class="$style.clearfilters">
          <b :class="$style.clearfilter">Очистить фильтр</b>
        </button>
      </div>
      <div :class="$style.row">
        <v-select
          :class="$style.hotelVselect"
          variant="outlined"
          name="SelectHotelFilter"
          label="выберите отель"
          color="primary"
        />
        <div :class="$style.separate" />
        <div :class="$style.rowCity">
          <div :class="$style.title">
            <b :class="$style.city">Город</b>
            <img :class="$style.icon" alt="" src="/icon1.svg" />
          </div>
          <div :class="$style.citylist">
            <v-text-field
              :class="$style.searchInputVtextfield"
              color="primary"
              variant="filled"
              label="|Search"
            />
            <div :class="$style.cityParent">
              <div :class="$style.city1">
                <div :class="$style.name">
                  <input
                    :class="$style.checkbox"
                    :required="true"
                    :checked="true"
                    type="checkbox"
                  />
                  <div :class="$style.moscow">Москва</div>
                </div>
              </div>
              <div :class="$style.city2">
                <input
                  :class="$style.checkbox"
                  :required="true"
                  :checked="true"
                  type="checkbox"
                />
                <div :class="$style.moscow">Казань</div>
              </div>
              <div :class="$style.city3">
                <div :class="$style.cityChild" />
                <div :class="$style.name">
                  <input
                    :class="$style.checkbox"
                    :required="true"
                    :checked="true"
                    type="checkbox"
                  />
                  <div :class="$style.moscow">Нижний Новгород</div>
                </div>
              </div>
              <div :class="$style.cityWrapper">
                <div :class="$style.city4">
                  <div :class="$style.name">
                    <input
                      :class="$style.checkbox"
                      :checked="true"
                      type="checkbox"
                    />
                    <div :class="$style.moscow">Великий Новгород</div>
                  </div>
                </div>
              </div>
              <div :class="$style.city5">
                <input
                  :class="$style.checkbox4"
                  :checked="true"
                  type="checkbox"
                />
                <div :class="$style.moscow">Уфа</div>
              </div>
              <div :class="$style.city6">
                <div :class="$style.name">
                  <input
                    :class="$style.checkbox"
                    :checked="true"
                    type="checkbox"
                  />
                  <label :class="$style.label" htmlFor="CityLabel"
                    >Санкт-Петербург</label
                  >
                </div>
              </div>
              <v-btn
                :class="$style.vbtn"
                color="primary"
                variant="plain"
                width="240"
              />
            </div>
          </div>
        </div>
        <div :class="$style.separate" />
        <div :class="$style.shiftType">
          <b :class="$style.city">Shift Type</b>
          <img :class="$style.icon" alt="" src="/icon2.svg" />
        </div>
        <div :class="$style.separate" />
        <div :class="$style.status">
          <b :class="$style.city">Status</b>
          <img :class="$style.icon" alt="" src="/icon2.svg" />
        </div>
      </div>
    </div>
    <img :class="$style.frameIcon" alt="" src="/frame.svg" />
  </div>
</template>
<script>
  import { defineComponent } from "vue";

  export default defineComponent({
    name: "FiltersWindow",
  });
</script>
<style module>
  .titelfilters {
    position: relative;
    line-height: 26px;
  }
  .clearfilter {
    position: relative;
    font-size: var(--text-xs-regular-size);
    text-decoration: underline;
    line-height: 22px;
    font-family: var(--text-xs-regular);
    color: var(--color-royalblue-200);
    text-align: left;
  }
  .clearfilters {
    cursor: pointer;
    border: none;
    padding: 0;
    background-color: transparent;
    width: 112px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
  }
  .header {
    width: 240px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
  }
  .hotelVselect {
    align-self: stretch;
  }
  .separate {
    align-self: stretch;
    position: relative;
    border-radius: var(--br-12xs);
    background-color: var(--color-whitesmoke-200);
    height: 1px;
  }
  .city {
    position: relative;
    line-height: 18px;
  }
  .icon {
    position: relative;
    width: 18px;
    height: 18px;
    overflow: hidden;
    flex-shrink: 0;
  }
  .title {
    align-self: stretch;
    background-color: var(--light-light-100);
    height: 24px;
    overflow: hidden;
    flex-shrink: 0;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
  }
  .searchInputVtextfield {
    border: none;
    background-color: transparent;
    align-self: stretch;
    position: relative;
  }
  .checkbox {
    border-radius: var(--br-9xs);
    background-color: var(--light-light-100);
    width: 18px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }
  .moscow {
    position: relative;
    line-height: 22px;
  }
  .name {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-3xs);
  }
  .city1 {
    position: absolute;
    top: 0px;
    left: 0px;
    width: 240px;
    height: 34px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
    mix-blend-mode: normal;
  }
  .city2 {
    position: absolute;
    top: 40px;
    left: 0px;
    background-color: var(--light-light-100);
    width: 240px;
    height: 34px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-3xs);
    mix-blend-mode: normal;
  }
  .cityChild {
    position: relative;
    background-color: var(--color-gray-600);
    width: 3px;
    height: 34px;
  }
  .city3 {
    position: absolute;
    top: 80px;
    left: -16px;
    background-color: rgba(112, 112, 112, 0.05);
    width: 272px;
    height: 34px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-smi);
  }
  .city4 {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
    mix-blend-mode: normal;
  }
  .cityWrapper {
    position: absolute;
    top: 120px;
    left: 0px;
    width: 240px;
    height: 34px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
  }
  .checkbox4 {
    position: relative;
    border-radius: var(--br-9xs);
    background-color: var(--light-light-100);
    border: 1px solid var(--color-gainsboro-300);
    box-sizing: border-box;
    width: 18px;
    height: 18px;
  }
  .city5 {
    position: absolute;
    top: 160px;
    left: 0px;
    width: 240px;
    height: 34px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-3xs);
    mix-blend-mode: normal;
    color: var(--base-pure-black);
  }
  .label {
    cursor: pointer;
    position: relative;
    line-height: 22px;
  }
  .city6 {
    position: absolute;
    top: 200px;
    left: 0px;
    width: 240px;
    height: 34px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
    mix-blend-mode: normal;
  }
  .vbtn {
    position: absolute;
    top: 240px;
    left: 0px;
  }
  .cityParent {
    align-self: stretch;
    position: relative;
    height: 262px;
  }
  .citylist {
    align-self: stretch;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
    gap: var(--gap-base);
    font-size: var(--text-xs-regular-size);
    color: var(--color-darkgray-200);
  }
  .rowCity {
    align-self: stretch;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
    gap: var(--gap-3xs);
    color: var(--primary-footer);
  }
  .shiftType {
    align-self: stretch;
    height: 24px;
    overflow: hidden;
    flex-shrink: 0;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    mix-blend-mode: normal;
  }
  .status {
    align-self: stretch;
    height: 24px;
    overflow: hidden;
    flex-shrink: 0;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    padding: 2px 0px;
    box-sizing: border-box;
    gap: 178px;
  }
  .row {
    width: 240px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-base);
    font-size: var(--label-large-label-14-size);
    color: var(--color-darkgray-200);
  }
  .dropdownFilters {
    flex: 1;
    background-color: var(--light-light-100);
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
    padding: var(--padding-xl) var(--padding-base);
    gap: var(--gap-5xl);
  }
  .frameIcon {
    position: relative;
    width: 24px;
    height: 212px;
    overflow: hidden;
    flex-shrink: 0;
    margin-left: -104px;
  }
  .filterswindow {
    width: 272px;
    overflow: hidden;
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    justify-content: flex-end;
    max-width: 100%;
    max-height: 100%;
    text-align: left;
    font-size: var(--label-text-large-size);
    color: var(--primary-footer);
    font-family: var(--text-xs-regular);
  }
</style>
