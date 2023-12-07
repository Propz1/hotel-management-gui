<template>
  <div :class="$style.calendarpage">

    <ButtonOutlinedSquareD1
      buttonText="Установить"
      propTop="731px"
      propLeft="1275px"
      propCursor="pointer"
      propBackgroundColor="transparent"
      propWidth="111px"
      propDisplay="inline-block"
    />
    <div :class="$style.calendarpageChild" />
    <div :class="$style.div">Отель</div>
    <SidebarDark7 />
    <!-- <div :class="$style.monthandyaertwo">
      <b :class="$style.textmonth">Апрель</b>
      <b :class="$style.textyear">2023</b>
    </div>
    <div :class="$style.monthandyaerone">
      <b :class="$style.textmonth">Март</b>
      <b :class="$style.textyear">2023</b>
    </div> -->
    <!-- <CalendarOne1
      calendarDayTextSize="23"
      propLeft="660px"
      calendarOneBackgroundColor="rgba(255, 255, 255, 0.2)"
      weekColor="#9d9ca3"
      calendarOneBackgroundColor1="rgba(255, 255, 255, 0.2)"
      weekColor1="#9d9ca3"
      calendarOneBorderRadius="unset"
      calendarOneBackgroundColor2="rgba(255, 255, 255, 0.2)"
      weekColor2="#9d9ca3"
      calendarOneBorderRadius1="unset"
      calendarOneBackgroundColor3="rgba(255, 255, 255, 0.2)"
      weekColor3="#9d9ca3"
      calendarOneBorderRadius2="unset"
      calendarOneBackgroundColor4="rgba(255, 255, 255, 0.2)"
      weekColor4="#9d9ca3"
      calendarOneBorderRadius3="2px"
      calendarOneBackgroundColor5="#ef5da8"
      calendarOneBackgroundColor6="rgba(68, 153, 238, 0.75)"
      calendarOneCursor="pointer"
      calendarOneBorder="none"
      calendarOnePadding="0"
      calendarOneBorderRadius4="30px"
      weekColor5="#000"
      calendarOneBackgroundColor7="rgba(68, 153, 238, 0.75)"
      calendarOneBorderRadius5="30px"
      weekColor6="#000"
      calendarOneBackgroundColor8="rgba(68, 153, 238, 0.75)"
      calendarOneBorderRadius6="30px"
      weekColor7="#000"
      calendarOneBackgroundColor9="unset"
      weekColor8="#000"
    /> -->
    <!-- <CalendarOne1
      calendarDayTextColor="25"
      calendarDayTextSize="19"
    /> -->
    <DatePicker />
    <div :class="$style.div1">Выберите отель
    <MultiSelect v-model="selectedHotels" @click="loadMultipleSelectHotels()" :options="groupedHotels" optionLabel="label" optionGroupLabel="label" optionGroupChildren="items" display="chip" placeholder="Выберите отель" class="w-full md:w-50rem">
       <template #optiongroup="slotProps">
          <div class="flex align-items-center">
            <img :alt="slotProps.option.label" src="https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png" :class="`flag flag-${slotProps.option.code.toLowerCase()} mr-2`" style="width: 18px" />
            <div>{{ slotProps.option.label }}</div>
          </div>
       </template>
   </MultiSelect>
    </div>
    <div :class="$style.div2">Кол-во номеров</div>
    <div :class="$style.div3">Цена за номер (руб.)</div>
    <div :class="$style.div4">Кэшбек (руб.)</div>
    <div :class="$style.div5">Кэшбек (%)</div>
    <div :class="$style.div6">Кэшбек не заезд (руб.)</div>
    <div :class="$style.div7">Кэшбек не заезд (%)</div>
    <div :class="$style.div8">Выберите период</div>
    <button :class="$style.arrowRightCalendar">
      <img :class="$style.layer1Icon" alt="" src="/layer-1.svg" />
    </button>
    <button :class="$style.arrowLeftCalendar">
      <img :class="$style.layer1Icon" alt="" src="/layer-11.svg" />
    </button>
  </div>
</template>
<script>
  import { defineComponent } from "vue";
  import FieldsSmallLabel from "../components/FieldsSmallLabel.vue";
  import FieldsSmallDropdown from "../components/FieldsSmallDropdown.vue";
  import ButtonOutlinedSquareD1 from "../components/ButtonOutlinedSquareD1.vue";
  import SidebarDark7 from "../components/SidebarDark7.vue";
  import CalendarOne1 from "../components/CalendarOne1.vue";
  import DatePicker from "../components/DatePicker.vue";
  import MultiSelect from 'primevue/multiselect';
  import axios from 'axios-https-proxy-fix';

  const proxy = {
  host: 'https://localhost',
  port: 9090,
  // auth: {
  //   username: 'some_login',
  //   password: 'some_pass'
  // }
};

  export default defineComponent({

    data() {
        return {
            selectedHotels: null,
            groupedHotels: null,
            //     {
            //         label: 'Москва',
            //         code: 'DE',
            //         items: [
            //             { label: 'Berlinnnnnnnnnnnnnnnnn', value: 'Berlin' },
            //             { label: 'Frankfurt', value: 'Frankfurt' },
            //             { label: 'Hamburg', value: 'Hamburg' },
            //             { label: 'Munich', value: 'Munich' }
            //         ]
            //     },
            //     {
            //         label: 'Санкт-Петербург',
            //         code: 'SPB',
            //         items: [
            //             { label: 'Chicago', value: 'Chicago' },
            //             { label: 'Los Angeles', value: 'Los Angeles' },
            //             { label: 'New York', value: 'New York' },
            //             { label: 'San Francisco', value: 'San Francisco' }
            //         ]
            //     },
            //     {
            //         label: 'УФА',
            //         code: 'UFA',
            //         items: [
            //             { label: 'Kyoto', value: 'Kyoto' },
            //             { label: 'Osaka', value: 'Osaka' },
            //             { label: 'Tokyo', value: 'Tokyo' },
            //             { label: 'Yokohama', value: 'Yokohama' }
            //         ]
            //     }
            // ]
        };
    },



    name: "CalendarPage",
    components: {
      MultiSelect,
      FieldsSmallLabel,
      FieldsSmallDropdown,
      ButtonOutlinedSquareD1,
      SidebarDark7,
      DatePicker,
    },
    methods: {
      onLinkMainPageContainerClick() {
        this.$router.push("/main");
      },
      onLinkRequisitionsPageContainerClick() {
        this.$router.push("/");
      },
      onLinkOccupancyPageContainerClick() {
        this.$router.push("/requisitions");
      },
      onLinkChatPageContainerClick() {
        this.$router.push("/chat");
      },
      onLinkReportsPageContainerClick() {
        this.$router.push("/reports");
      },
      onLinkIntegrationPageContainerClick() {
        this.$router.push("/integrations");
      },
      onLinkHelpPageContainerClick() {
        this.$router.push("/help");
      },
      onLinkSettingsPageContainerClick() {
        this.$router.push("/settings");
      },

      loadMultipleSelectHotels() {
              console.log("Пытаюсь загрузить в groupedHotels")


                 axios
                .get('https://localhost:9090/getGroupedHotels')
                .then((res) => {

                  console.log(res.data)

                 this.groupedHotels = res.data;
                 })
                .catch((error) => {
                //console.log(error.res.data);
            });
        },
    },
  });
</script>
<style module>
  .calendarpageChild {
    position: absolute;
    top: 18px;
    left: 288px;
    background-color: #d9d9d9;
    width: 1586px;
    height: 42px;
  }
  .div {
    position: absolute;
    top: 28px;
    left: 333px;
    font-size: var(--label-large-label-14-size);
    line-height: 24px;
    font-weight: 500;
    color: var(--buttonon4);
    text-align: center;
  }
  .textmonth {
    position: absolute;
    top: 0px;
    left: 0px;
    letter-spacing: 0.15px;
    line-height: 32px;
    display: inline-block;
    width: 88px;
  }
  .textyear {
    position: absolute;
    top: 0px;
    left: 108px;
    letter-spacing: 0.15px;
    line-height: 32px;
  }
  .monthandyaertwo {
    position: absolute;
    top: 132px;
    left: 741px;
    width: 208px;
    height: 32px;
    font-size: var(--h6-size);
  }
  .monthandyaerone {
    position: absolute;
    top: 133px;
    left: 360px;
    width: 208px;
    height: 32px;
    font-size: var(--h6-size);
  }
  .div1 {
    position: absolute;
    top: 91px;
    left: 1127px;
    line-height: 150%;
  }
  .div2 {
    position: absolute;
    top: 160px;
    left: 1128px;
    line-height: 150%;
  }
  .div3 {
    position: absolute;
    top: 231px;
    left: 1091px;
    line-height: 150%;
  }
  .div4 {
    position: absolute;
    top: 303px;
    left: 1147px;
    line-height: 150%;
  }
  .div5 {
    position: absolute;
    top: 382px;
    left: 1166px;
    line-height: 150%;
  }
  .div6 {
    position: absolute;
    top: 455px;
    left: 1072px;
    line-height: 150%;
  }
  .div7 {
    position: absolute;
    top: 529px;
    left: 1091px;
    line-height: 150%;
  }
  .div8 {
    position: absolute;
    top: 609px;
    left: 1131px;
    font-size: var(--label-large-label-14-size);
    line-height: 24px;
    font-weight: 500;
    text-align: center;
  }
  .layer1Icon {
    position: absolute;
    height: 100%;
    width: 100%;
    top: -100%;
    right: 0%;
    bottom: 100%;
    left: 0%;
    max-width: 100%;
    overflow: hidden;
    max-height: 100%;
  }
  .arrowRightCalendar {
    cursor: pointer;
    border: none;
    padding: 0;
    background-color: transparent;
    position: absolute;
    top: 128px;
    left: 655px;
    width: 40px;
    height: 40px;
    overflow: hidden;
    transform: rotate(-90deg);
    transform-origin: 0 0;
  }
  .arrowLeftCalendar {
    cursor: pointer;
    border: none;
    padding: 0;
    background-color: transparent;
    position: absolute;
    top: 128px;
    left: 568px;
    width: 40px;
    height: 40px;
    overflow: hidden;
    transform: rotate(-90deg);
    transform-origin: 0 0;
  }
  .calendarpage {
    position: relative;
    background-color: var(--color-whitesmoke-300);
    width: 100%;
    height: 1080px;
    text-align: left;
    font-size: var(--label-text-large-size);
    color: var(--base-pure-black);
    font-family: var(--body-body-16);
  }
</style>
