<template>
  <div :class="$style.table">
    <div :class="$style.schedulecheckin">
      <header :class="$style.daysHorizontal">
        <div :class="$style.days">
          <Day
            numberDaySchedulerOccupancyLabel="1"
            dayWeekSchedulerOccupancyLabel="Пн"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="unset"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="2"
            dayWeekSchedulerOccupancyLabel="Вт"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="3"
            dayWeekSchedulerOccupancyLabel="Wed"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          />
          <div :class="$style.day">
            <div :class="$style.separator" />
            <div :class="$style.day1">
              <b :class="$style.b">4</b>
              <div :class="$style.mon">Thu</div>
              <div :class="$style.dayChild" />
            </div>
            <div :class="$style.separator" />
          </div>
          <Day
            numberDaySchedulerOccupancyLabel="5"
            dayWeekSchedulerOccupancyLabel="Frid"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="6"
            dayWeekSchedulerOccupancyLabel="Sut"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="7"
            dayWeekSchedulerOccupancyLabel="Sun"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="8"
            dayWeekSchedulerOccupancyLabel="Mon"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="9"
            dayWeekSchedulerOccupancyLabel="Tue"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="11"
            dayWeekSchedulerOccupancyLabel="Wed"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="4px"
            dayScheduleTableOccupancyBackgroundColor="#fff"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="12"
            dayWeekSchedulerOccupancyLabel="Thu"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="13"
            dayWeekSchedulerOccupancyLabel="Fri"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="4px"
            dayScheduleTableOccupancyBackgroundColor="#195ce5"
            divColor="#fff"
            monColor="#fff"
          /><Day
            numberDaySchedulerOccupancyLabel="14"
            dayWeekSchedulerOccupancyLabel="Sut"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="15"
            dayWeekSchedulerOccupancyLabel="Sun"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="16"
            dayWeekSchedulerOccupancyLabel="Mon"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="17"
            dayWeekSchedulerOccupancyLabel="Tue"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="18"
            dayWeekSchedulerOccupancyLabel="Wed"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="19"
            dayWeekSchedulerOccupancyLabel="Thu"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="20"
            dayWeekSchedulerOccupancyLabel="Thu"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="21"
            dayWeekSchedulerOccupancyLabel="Thu"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="19"
            dayWeekSchedulerOccupancyLabel="Thu"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="19"
            dayWeekSchedulerOccupancyLabel="Thu"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="19"
            dayWeekSchedulerOccupancyLabel="Thu"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="19"
            dayWeekSchedulerOccupancyLabel="Thu"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="19"
            dayWeekSchedulerOccupancyLabel="Thu"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="19"
            dayWeekSchedulerOccupancyLabel="Thu"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="19"
            dayWeekSchedulerOccupancyLabel="Thu"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="19"
            dayWeekSchedulerOccupancyLabel="Thu"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="19"
            dayWeekSchedulerOccupancyLabel="Thu"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          /><Day
            numberDaySchedulerOccupancyLabel="19"
            dayWeekSchedulerOccupancyLabel="Thu"
            dayScheduleTableOccupancyWidth="66px"
            dayScheduleTableOccupancyShrink="0"
            dayScheduleTableOccupancyMarginLeft="-1px"
            dayScheduleTableOccupancyBorderRadius="unset"
            dayScheduleTableOccupancyBackgroundColor="unset"
            divColor="#121d22"
            monColor="#9d9ca3"
          />
        </div>
      </header>
      <div :class="$style.table1">
        <div :class="$style.cells">
          <div :class="$style.table2">
            <img :class="$style.yellowLineIcon" alt="" src="/yellow-line.svg" />
            <div :class="$style.line">
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="unset"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
              <div :class="$style.cell" />
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
            </div>
            <div :class="$style.line">
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="unset"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
              <div :class="$style.cell" />
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
            </div>
            <div :class="$style.line">
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="unset"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
              <div :class="$style.cell" />
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
            </div>
            <div :class="$style.line">
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="unset"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
              <div :class="$style.cell" />
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
            </div>
            <div :class="$style.line">
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="unset"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
              <div :class="$style.cell" />
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
            </div>
            <div :class="$style.line">
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="unset"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
              <div :class="$style.cell" />
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
            </div>
            <div :class="$style.line">
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="unset"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
              <div :class="$style.cell" />
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
            </div>
            <div :class="$style.line">
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="unset"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
              <div :class="$style.cell" />
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
            </div>
            <div :class="$style.line">
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="unset"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
              <div :class="$style.cell" />
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
            </div>
            <div :class="$style.line9">
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="unset"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
              <div :class="$style.cell" />
              <Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fafafa"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              /><Typedefault
                cellSchedulerTableOccupancyBackgroundColor="#fff"
                cellSchedulerTableOccupancyHeight="60px"
                cellSchedulerTableOccupancyMarginLeft="-1px"
              />
            </div>
          </div>
        </div>
        <div :class="$style.fullDay">
          <div :class="$style.content">
            <b :class="$style.fullDay1">Full day</b>
            <div :class="$style.div">8:00 - 16:00</div>
          </div>
        </div>
        <div :class="$style.fullDay2" :onClick="openSelectDay">
          <div :class="$style.fullDayChild" />
          <div :class="$style.content">
            <b :class="$style.b1">Полный день</b>
            <div :class="$style.div">8:00 - 16:00</div>
          </div>
        </div>
        <div :class="$style.fullDay3">
          <div :class="$style.fullDayItem" />
          <div :class="$style.content">
            <b :class="$style.remote">Remote</b>
          </div>
        </div>
        <div :class="$style.fullDay4">
          <div :class="$style.fullDayInner" />
          <div :class="$style.content">
            <b :class="$style.remote">Удалено</b>
          </div>
        </div>
        <div :class="$style.fullDay5">
          <div :class="$style.fullDayChild" />
          <div :class="$style.content">
            <b :class="$style.b3">Полный день</b>
            <div :class="$style.div">8:00 - 16:00</div>
          </div>
        </div>
        <div :class="$style.fullDay6">
          <div :class="$style.fullDayChild1" />
          <div :class="$style.content5">
            <b :class="$style.remote">Morning</b>
          </div>
        </div>
        <div :class="$style.fullDay7">
          <div :class="$style.fullDayChild1" />
          <div :class="$style.content">
            <b :class="$style.remote">Morning</b>
          </div>
        </div>
        <div :class="$style.fullDay8">
          <div :class="$style.fullDayChild3" />
          <div :class="$style.content">
            <b :class="$style.remote">Morning</b>
          </div>
        </div>
        <div :class="$style.fullDay9">
          <div :class="$style.fullDayChild4" />
          <div :class="$style.content8">
            <b :class="$style.remote">Evening</b>
          </div>
        </div>
        <div :class="$style.fullDay10" :onClick="openSelectDay1">
          <div :class="$style.content">
            <b :class="$style.b4">Полный день</b>
            <div :class="$style.div">8:00 - 16:00</div>
          </div>
        </div>
        <div :class="$style.fullDay11">
          <div :class="$style.fullDayChild5" />
          <div :class="$style.content">
            <b :class="$style.b5">Свободно</b>
          </div>
        </div>
        <div :class="$style.fullDay12">
          <div :class="$style.fullDayChild6" />
          <div :class="$style.content">
            <b :class="$style.vacation">Vacation</b>
          </div>
        </div>
      </div>
      <ScrollbarClassic2 />
    </div>
    <div :class="$style.rowsguests">
      <div :class="$style.header">
        <div :class="$style.gradient" />
        <div :class="$style.header1">
          <div :class="$style.guestNameParent">
            <b :class="$style.guestName">Гостей</b>
            <div :class="$style.wrapper">
              <b :class="$style.remote">96</b>
            </div>
          </div>
          <div :class="$style.gradient1" />
        </div>
      </div>
      <div :class="$style.lefrColumn">
        <ScrollbarClassic1 />
        <ul :class="$style.namesguests" id="NamesGuests" Name Telephone>
          <div :class="$style.hederguestslist">
            <img
              :class="$style.hederguestslistChild"
              alt=""
              src="/rectangle-225.svg"
            />
            <div :class="$style.darkDesktopSwitchesSlParent">
              <DarkDesktopSwitchesSl
                guestsCheckBoxPosition="relative"
                guestsCheckBoxBorderRadius="4px"
                guestsCheckBoxBackgroundColor="#fff"
              />
              <b :class="$style.guests">Гости</b>
            </div>
          </div>
          <Guests1 /><Guests1
            propTop="calc(50% - 17px)"
            propHeight="33.4px"
            propWidth="unset"
          /><Guests1
            propTop="calc(50% - 18px)"
            propHeight="unset"
            propWidth="unset"
          /><Guests1
            propTop="calc(50% - 18px)"
            propHeight="unset"
            propWidth="unset"
          /><Guests1
            propTop="calc(50% - 18px)"
            propHeight="unset"
            propWidth="unset"
          /><Guests1
            propTop="calc(50% - 18px)"
            propHeight="unset"
            propWidth="unset"
          />
          <div :class="$style.guests1">
            <div :class="$style.guestsChild" />
            <div :class="$style.darkDesktopSwitchesSlGroup">
              <DarkDesktopSwitchesSl
                guestsCheckBoxPosition="relative"
                guestsCheckBoxBorderRadius="4px"
                guestsCheckBoxBackgroundColor="#fff"
              />
              <div :class="$style.frameChild" />
              <div :class="$style.nameParent">
                <div :class="$style.name">ФИО</div>
                <div :class="$style.tel">
                  <p :class="$style.p">+7</p>
                </div>
              </div>
            </div>
          </div>
          <Guests1
            propTop="calc(50% - 18px)"
            propHeight="unset"
            propWidth="52px"
          /><Guests /><Guests />
        </ul>
      </div>
    </div>
  </div>
  <div v-if="isSelectDayOpen">
    <PortalPopup
      overlayColor="rgba(113, 113, 113, 0.3)"
      placement="Centered"
      :relativeLayerRef="$refs[fullDayContainer1Ref]"
      :onOutsideClick="closeSelectDay"
    >
      <SelectDay :onClose="closeSelectDay" />
    </PortalPopup>
  </div>
  <div v-if="isSelectDay1Open">
    <PortalPopup
      overlayColor="rgba(113, 113, 113, 0.3)"
      placement="Centered"
      :relativeLayerRef="$refs[fullDayContainer9Ref]"
      :onOutsideClick="closeSelectDay1"
    >
      <SelectDay :onClose="closeSelectDay1" />
    </PortalPopup>
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import SelectDay from "./SelectDay.vue";
  import PortalPopup from "./PortalPopup.vue";
  import Day from "./Day.vue";
  import Typedefault from "./Typedefault.vue";
  import ScrollbarClassic2 from "./ScrollbarClassic2.vue";
  import ScrollbarClassic1 from "./ScrollbarClassic1.vue";
  import DarkDesktopSwitchesSl from "./DarkDesktopSwitchesSl.vue";
  import Guests1 from "./Guests1.vue";
  import Guests from "./Guests.vue";

  export default defineComponent({
    name: "Table",
    data() {
      return { isSelectDayOpen: false, isSelectDay1Open: false };
    },
    components: {
      SelectDay,
      PortalPopup,
      Day,
      Typedefault,
      ScrollbarClassic2,
      ScrollbarClassic1,
      DarkDesktopSwitchesSl,
      Guests1,
      Guests,
    },
    methods: {
      openSelectDay() {
        this.isSelectDayOpen = true;
      },
      closeSelectDay() {
        this.isSelectDayOpen = false;
      },
      openSelectDay1() {
        this.isSelectDay1Open = true;
      },
      closeSelectDay1() {
        this.isSelectDay1Open = false;
      },
    },
  });
</script>
<style module>
  .separator {
    position: relative;
    border-radius: var(--br-12xs);
    background-color: var(--color-whitesmoke-200);
    width: 1px;
    height: 40px;
  }
  .b {
    position: absolute;
    top: calc(50% - 22px);
    left: calc(50% - 5px);
    line-height: 26px;
  }
  .mon {
    position: absolute;
    top: calc(50% + 4px);
    left: calc(50% - 10px);
    font-size: var(--text-xs-regular-size);
    line-height: 18px;
    color: var(--color-darkgray-200);
  }
  .dayChild {
    position: absolute;
    bottom: -2px;
    left: calc(50% - 23px);
    background-color: var(--color-royalblue-200);
    width: 45px;
    height: 3px;
  }
  .day1 {
    flex: 1;
    position: relative;
    height: 54px;
  }
  .day {
    background-color: var(--light-light-100);
    width: 66px;
    flex-shrink: 0;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    padding: var(--padding-10xs) 0px;
    box-sizing: border-box;
    gap: var(--gap-11xs);
    margin-left: -1px;
  }
  .days {
    position: absolute;
    top: 0px;
    left: 1px;
    border-bottom: 1px solid var(--color-whitesmoke-200);
    box-sizing: border-box;
    width: 1300px;
    overflow-x: auto;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
  }
  .daysHorizontal {
    position: relative;
    width: 1326px;
    height: 60px;
    overflow-x: auto;
    flex-shrink: 0;
    text-align: left;
    font-size: var(--label-text-large-size);
    color: var(--primary-footer);
    font-family: var(--text-xs-regular);
  }
  .yellowLineIcon {
    position: relative;
    width: 1334px;
    height: 32px;
  }
  .cell {
    background-color: var(--light-light-100);
    border: 1px solid var(--color-whitesmoke-200);
    box-sizing: border-box;
    width: 66px;
    height: 60px;
    margin-left: -1px;
  }
  .line {
    width: 1956px;
    height: 60px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    margin-top: -1px;
  }
  .line9 {
    flex: 1;
    width: 1956px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    margin-top: -1px;
  }
  .table2 {
    position: absolute;
    top: 0px;
    left: 0px;
    width: 976px;
    height: 622px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
  }
  .cells {
    position: absolute;
    top: 0px;
    left: 0px;
    width: 1326px;
    height: 602px;
    overflow: hidden;
  }
  .fullDay1 {
    position: relative;
    line-height: 16px;
    display: inline-block;
    width: 58.1px;
    mix-blend-mode: normal;
  }
  .div {
    position: relative;
    line-height: 16px;
    color: var(--color-dimgray-100);
  }
  .content {
    position: absolute;
    top: 6px;
    left: 6px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
  }
  .fullDay {
    position: absolute;
    top: 264px;
    left: 50px;
    border-radius: var(--br-9xs);
    background-color: var(--color-teal-300);
    border: 1px solid var(--color-teal-200);
    box-sizing: border-box;
    width: 318px;
    height: 52px;
    overflow: hidden;
  }
  .fullDayChild {
    position: absolute;
    top: 0px;
    left: 0px;
    border-radius: var(--br-9xs);
    background-color: var(--color-teal-300);
    border: 1px solid var(--color-teal-200);
    box-sizing: border-box;
    width: 317px;
    height: 52px;
  }
  .b1 {
    position: relative;
    line-height: 16px;
    display: inline-block;
    width: 86px;
  }
  .fullDay2 {
    position: absolute;
    top: 205px;
    left: 583.3px;
    width: 370.8px;
    height: 52px;
    overflow: hidden;
    mix-blend-mode: normal;
    cursor: pointer;
  }
  .fullDayItem {
    position: absolute;
    top: 0px;
    left: 0px;
    border-radius: var(--br-9xs);
    background-color: var(--color-orange-300);
    border: 1px solid var(--color-orange-200);
    box-sizing: border-box;
    width: 122px;
    height: 52px;
  }
  .remote {
    position: relative;
    line-height: 16px;
  }
  .fullDay3 {
    position: absolute;
    top: 205px;
    left: 1114.3px;
    width: 142.7px;
    height: 52px;
    overflow: hidden;
    mix-blend-mode: normal;
    color: var(--color-orange-100);
  }
  .fullDayInner {
    position: absolute;
    top: 0px;
    left: 0px;
    background-color: var(--color-orange-300);
    border: 1px solid var(--color-orange-200);
    box-sizing: border-box;
    width: 188px;
    height: 52px;
  }
  .fullDay4 {
    position: absolute;
    top: 205px;
    left: 202px;
    width: 219.9px;
    height: 52px;
    overflow: hidden;
    mix-blend-mode: normal;
    color: var(--color-orange-100);
  }
  .b3 {
    position: relative;
    line-height: 16px;
    display: inline-block;
    width: 95px;
  }
  .fullDay5 {
    position: absolute;
    top: 264px;
    left: 583.3px;
    width: 370.8px;
    height: 52px;
    overflow: hidden;
    mix-blend-mode: normal;
  }
  .fullDayChild1 {
    position: absolute;
    top: 0px;
    left: 0px;
    border-radius: var(--br-9xs);
    background-color: var(--color-royalblue-400);
    border: 1px solid var(--color-royalblue-300);
    box-sizing: border-box;
    width: 187px;
    height: 26px;
  }
  .content5 {
    position: absolute;
    top: calc(50% - 8px);
    left: 6px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
  }
  .fullDay6 {
    position: absolute;
    top: 323px;
    left: 583.3px;
    width: 218.7px;
    height: 26px;
    overflow: hidden;
    mix-blend-mode: normal;
    color: var(--color-royalblue-100);
  }
  .fullDay7 {
    position: absolute;
    top: 146px;
    left: 1114.3px;
    width: 218.7px;
    height: 26px;
    overflow: hidden;
    mix-blend-mode: normal;
    color: var(--color-royalblue-100);
  }
  .fullDayChild3 {
    position: absolute;
    top: 0px;
    left: 0px;
    background-color: var(--color-royalblue-400);
    border: 1px solid var(--color-royalblue-300);
    box-sizing: border-box;
    width: 187px;
    height: 26px;
  }
  .fullDay8 {
    position: absolute;
    top: 87px;
    left: 1114.3px;
    width: 218.7px;
    height: 26px;
    overflow: hidden;
    mix-blend-mode: normal;
    color: var(--color-royalblue-100);
  }
  .fullDayChild4 {
    position: absolute;
    top: 0px;
    left: 0px;
    border-radius: var(--br-9xs);
    background-color: var(--color-gray-800);
    border: 1px solid var(--color-gray-700);
    box-sizing: border-box;
    width: 253px;
    height: 26px;
  }
  .content8 {
    position: absolute;
    top: calc(50% - 8px);
    left: 6px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
  }
  .fullDay9 {
    position: absolute;
    top: 408px;
    left: 658.1px;
    width: 295.9px;
    height: 26px;
    overflow: hidden;
    mix-blend-mode: normal;
    color: var(--base-pure-black);
  }
  .b4 {
    position: relative;
    line-height: 16px;
    display: inline-block;
    width: 90px;
  }
  .fullDay10 {
    position: absolute;
    top: 28px;
    left: 583px;
    border-radius: var(--br-9xs);
    background-color: var(--color-teal-300);
    border: 1px solid var(--color-teal-200);
    box-sizing: border-box;
    width: 187px;
    height: 52px;
    overflow: hidden;
    cursor: pointer;
  }
  .fullDayChild5 {
    position: absolute;
    top: 0px;
    left: 0px;
    background-color: var(--color-red-400);
    border: 1px solid var(--color-red-300);
    box-sizing: border-box;
    width: 123px;
    height: 52px;
  }
  .b5 {
    position: relative;
    line-height: 16px;
    display: inline-block;
    width: 66px;
  }
  .fullDay11 {
    position: absolute;
    top: 146px;
    left: 202px;
    width: 143.9px;
    height: 52px;
    overflow: hidden;
    mix-blend-mode: normal;
    color: var(--color-red-200);
  }
  .fullDayChild6 {
    position: absolute;
    top: 0px;
    left: 0px;
    background-color: var(--color-red-400);
    border: 1px solid var(--color-red-300);
    box-sizing: border-box;
    width: 294px;
    height: 52px;
  }
  .vacation {
    position: relative;
    line-height: 16px;
    display: inline-block;
    width: 58.1px;
  }
  .fullDay12 {
    position: absolute;
    top: 472px;
    left: 2px;
    width: 343.9px;
    height: 52px;
    overflow: hidden;
    mix-blend-mode: normal;
    color: var(--color-red-200);
  }
  .table1 {
    position: relative;
    width: 1333px;
    height: 602px;
  }
  .schedulecheckin {
    position: absolute;
    top: 4px;
    left: 320px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
    gap: var(--gap-12xs);
  }
  .gradient {
    position: absolute;
    top: 59px;
    left: 295px;
    background-color: var(--color-whitesmoke-200);
    border-top: 1px solid var(--color-whitesmoke-200);
    box-sizing: border-box;
    width: 59px;
    height: 31px;
    transform: rotate(-90deg);
    transform-origin: 0 0;
  }
  .guestName {
    position: relative;
    line-height: 18px;
  }
  .wrapper {
    border-radius: var(--br-11xl);
    border: 1px solid var(--color-whitesmoke-200);
    box-sizing: border-box;
    width: 30px;
    height: 22px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: var(--padding-3xs);
    font-size: var(--text-xs-regular-size);
    color: var(--color-darkgray-100);
  }
  .guestNameParent {
    position: absolute;
    top: 19px;
    left: 21px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-3xs);
  }
  .gradient1 {
    position: absolute;
    top: 759px;
    left: 299px;
    background-color: var(--color-whitesmoke-200);
    border-top: 1px solid var(--color-whitesmoke-200);
    box-sizing: border-box;
    width: 700px;
    height: 30px;
    transform: rotate(-90deg);
    transform-origin: 0 0;
  }
  .header1 {
    position: absolute;
    top: 0px;
    left: 0px;
    background-color: var(--light-light-100);
    border: 1px solid var(--color-whitesmoke-200);
    box-sizing: border-box;
    width: 299px;
    height: 60px;
    overflow: hidden;
  }
  .header {
    position: absolute;
    top: 0px;
    left: 0px;
    width: 326px;
    height: 60px;
  }
  .hederguestslistChild {
    position: absolute;
    top: 0px;
    left: 13px;
    width: 299px;
    height: 32px;
  }
  .guests {
    position: relative;
    line-height: 22px;
    text-transform: uppercase;
  }
  .darkDesktopSwitchesSlParent {
    position: absolute;
    top: calc(50% - 11px);
    left: calc(50% - 97px);
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-sm);
  }
  .hederguestslist {
    position: relative;
    width: 299px;
    height: 32px;
    font-size: var(--text-xs-regular-size);
    color: var(--color-darkseagreen);
  }
  .guestsChild {
    position: absolute;
    top: 2px;
    left: 0px;
    background-color: var(--light-light-100);
    border: 1px solid var(--color-whitesmoke-200);
    box-sizing: border-box;
    width: 299px;
    height: 60px;
  }
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
  .p {
    margin-block-start: 0;
    margin-block-end: 14px;
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
  .darkDesktopSwitchesSlGroup {
    position: absolute;
    top: calc(50% - 32px);
    left: 20px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: var(--gap-sm);
  }
  .guests1 {
    position: relative;
    width: 299px;
    height: 64px;
    margin-top: -1px;
  }
  .namesguests {
    margin: 0;
    position: absolute;
    top: 0px;
    left: 0px;
    height: 605px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
  }
  .lefrColumn {
    position: absolute;
    top: 59px;
    left: 0px;
    width: 326px;
    height: 605px;
    overflow: hidden;
    color: var(--base-pure-black);
  }
  .rowsguests {
    position: absolute;
    top: 7px;
    left: -13px;
    width: 326px;
    height: 668px;
    overflow: hidden;
    font-size: var(--label-large-label-14-size);
    color: var(--primary-footer);
  }
  .table {
    position: absolute;
    top: 356px;
    left: 266px;
    border: 1px solid var(--color-whitesmoke-200);
    box-sizing: border-box;
    width: 1653px;
    height: 685px;
    overflow: hidden;
    text-align: left;
    font-size: var(--text-xs-regular-size);
    color: var(--color-teal-100);
    font-family: var(--text-xs-regular);
  }
</style>
