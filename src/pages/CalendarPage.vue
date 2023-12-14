<template>
  <div :class="$style.calendarpage">

    <div :class="$style.calendarpageChild"/>

    <div class="card">
        <DataTable 
                  v-model:filters="filters"
                  dataKey="id" filterDisplay="row"
                  :loading="loading"
                  :value="calendarTable"
                  showGridlines
                  resizableColumns columnResizeMode="expand"
                  scrollable scrollHeight="963px"
                  :virtualScrollerOptions="{}"
                  :rows="40" 
                  :class="$style.calendarTableDatesStyle"
                  tableStyle="min-width: 50rem">

                
            <template #header>
                <!-- <div class="flex flex-wrap align-items-center justify-content-between gap-2">
                    <span class="text-xl text-900 font-bold">Таблица цен</span>
                    <Button icon="pi pi-refresh" rounded raised />
                </div> -->


                
                <Toolbar :class="$style.toolbarTableStyle" style="min-height: 2rem;">
                   <template #start>
                       <ToggleButton v-model="dateFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Дата" offLabel="Дата" style="margin-right: 1rem; border-radius: 0.5rem" />
                       <ToggleButton v-model="hotelFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Отель" offLabel="Отель" style="margin-right: 1rem; border-radius: 0.5rem"/>
                       <ToggleButton v-model="cityFrozen"  class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Город" offLabel="Город" style="margin-right: 1rem; border-radius: 0.5rem"/>
                    </template>


                   <template #center  style="min-height: 1.5rem;">
                    <Button icon="pi pi-print" class="mr-2" />
                     <span class="p-input-icon-left-center">
                        <i class="pi pi-search" style="padding: 0.5rem;"/>
                        <InputText placeholder="Поиск" style="min-height: 10px; "/>
                     </span>
                   </template>

                    <template #end>
                      <ToggleButton v-model="installationDateFrozen"  class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Дата обновления" offLabel="Дата обновления" style="margin-right: 1rem; border-radius: 0.5rem"/>
                     </template>
                  </Toolbar>
                 
            </template>
            
            <Column field="date" header="Дата" style="min-width: 3rem; padding: 1rem" alignFrozen="left" :frozen="dateFrozen">
              <template #body="{ data }">
                    {{ formatDate(data.date) }}
                </template>
            </Column>
            <Column field="quantity" header="Кол-во" style="min-width: 2rem;padding: 1rem"></Column>

            <Column field="hotelName" header="Отель" filterField="hotelName" :showFilterMenu="false" :filterMenuStyle="{ width: '14rem' }" style="min-width: 100px; width: 2%; padding: 1rem" alignFrozen="left" :frozen="hotelFrozen">
              <template #body="{ data }">
                    <div class="flex align-items-center gap-2">
                        <span>{{ data.hotelName }}</span>
                    </div>
              </template>
              <!-- <template #filter="{ filterModel, filterCallback }">
                    <MultiSelect  id="hotelsSelectInTable" v-model="filterModel.value" @change="filterCallback()" :options="groupedHotelsFoTable" optionLabel="label" ptionGroupLabel="label" optionGroupChildren="items" display="chip" placeholder="Выберите отель" class="p-column-filter" style="min-width: 14rem" :maxSelectedLabels="1">
                        <template #option="slotProps">
                            <div class="flex align-items-center gap-2">
                                <img :alt="slotProps.option.label" src="https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png" :class="`flag flag-${slotProps.option.code.toLowerCase()} mr-2`" style="width: 18px" />
                                <span>{{ slotProps.option.label }}</span>
                            </div>
                        </template>
                    </MultiSelect>
                </template> -->
            </Column>

            <Column field="hotelCity" header="Город" style="min-width: 100px; width: 2%; padding: 1rem" alignFrozen="left" :frozen="cityFrozen" >
             
            </Column>
            <Column field="hotelStars" header="Звезд" style="min-width: 1rem; padding: 1rem">
                <!-- <template #body="slotProps">
                    <Rating :modelValue="slotProps.data.hotelStars" readonly :cancel="false" />
                </template> -->
              </Column>
            <Column field="priceRoom" header="Цена /номер" style="min-width: 100px; width: 2%; padding: 1rem"></Column>
                <template #body="slotProps">
                    {{ formatCurrency(slotProps.data.priceRoom) }}
                </template>
            <Column field="priceBedTwin" header="Цена /место (2-х)" style="min-width: 100px; width: 2%; padding: 1rem"></Column>
            <Column field="priceBedTriple" header="Цена / место (3-х)" style="min-width: 100px; width: 2%; padding: 1rem"></Column>
            <Column field="priceBedQuadriple" header="Цена / место (4-х)" style="min-width: 100px; width: 2%; padding: 1rem"></Column>
            <Column field="cashback" header="Cashback" style="min-width: 100px; width: 2%; padding: 1rem"></Column>
            <Column field="cashbackPercent" header="Cashback (%)" style="min-width: 100px; width: 2%; padding: 1rem"></Column>
            <Column field="cashbackNoCheckIn" header="Cashback не заезд" style="min-width: 100px; width: 2%; padding: 1rem"></Column>
            <Column field="cashbackNoCheckInPercent" header="Cashback не заезд (%)" style="min-width: 100px; width: 2%; padding: 1rem"></Column>
            <Column field="installationDate" header="Дата обновления" style="min-width: 8rem; padding: 1rem" alignFrozen="right" :frozen="installationDateFrozen" >
              <template #body="{ data }">
                    {{ formatDateTime(data.installationDate) }}
                </template></Column>
            <Column field="administratorID" header="id автора" style="min-width: 8rem; width: 2%; padding: 1rem"></Column>
            <Column field="administratorName" header="Автор" style="min-width: 15rem; width: 2%; padding: 1rem"></Column>
            <!-- <Column header="Image">
                <template #body="slotProps">
                    <img :src="`https://primefaces.org/cdn/primevue/images/product/${slotProps.data.image}`" :alt="slotProps.data.image" class="w-6rem shadow-2 border-round" />
                </template>
            </Column> -->
            <!-- <Column field="category" header="Category"></Column>
            <Column header="Status">
                <template #body="slotProps">
                    <Tag :value="slotProps.data.inventoryStatus" :severity="getSeverity(slotProps.data)" />
                </template>
            </Column> -->
            <template #footer> Всего {{ calendarTable ? calendarTable.length : 0 }} строк. </template>
        </DataTable>
    </div>




     
    <!-- <div :class="$style.div">Отель</div> -->
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



    <!-- <div>
       <label  for="datePickert" class="font-bold block mb-2"> Выберите период</label>
       <DatePicker id="datePicker" v-model="Values" @input="onUpdateDatePicker($event)"/>
    </div> -->
   


   

    <DatePicker id="datePicker" v-model="datePicerValues" @input="onUpdateDatePicker($event)"/>

    
      <!-- <div class="card flex flex-wrap gap-2 p-fluid"> -->
 
    <!-- <div :class="$style.div0">
       <label  for="datePickert" class="font-bold block mb-2"> Выберите период</label>
    </div>  -->

      <div :class="$style.div1">
        <div class="card flex flex-wrap gap-2 p-fluid">
         <!-- <label  for="hotelsSelect" class="font-bold block mb-2"> Выберите отель </label> -->
        <MultiSelect id="hotelsSelect" v-model="selectedHotels" @click="loadMultipleSelectHotels()" :options="groupedHotels" optionLabel="label" optionGroupLabel="label" optionGroupChildren="items" display="chip" placeholder="Выберите отель" class="w-full md:w-50rem" style="min-width:25rem">
            <template #optiongroup="slotProps">
                <div class="flex align-items-center">
                  <img :alt="slotProps.option.label" src="https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png" :class="`flag flag-${slotProps.option.code.toLowerCase()} mr-2`" style="width: 18px" />
                  <div>{{ slotProps.option.label }}</div>
                </div>
            </template>
        </MultiSelect>
      </div>
    </div>


   
    <div :class="$style.div2">
        <label  for="quantityNumbersInput" class="font-bold block mb-2" style="padding-right: 0.5rem;"> Кол-во номеров </label>
        <InputText id="quantityNumbersInput" v-model="quantityNumbers" type="number" inputId="withoutgrouping" :useGrouping="false" style="min-width: 1px; max-width: 100px; max-height: 2px; background-color:white; border-radius: 0.5rem ; padding: 1rem;"/>
     </div> 
  
   
      <!-- <div class="card flex flex-wrap gap-2 p-fluid"> -->
        <!-- <div :class="$style.div_main"> -->

          
           <div :class="$style.div3"> 
              <label  for="price1" class="font-bold block mb-2" style="padding-right: 0.5rem;"> Цена / номер (руб.)</label>
              <InputText id="price1" v-model="price_one" type="number" inputId="withoutgrouping" :useGrouping="false" style="min-width: 100px; max-width: 100px;; max-height: 2px; background-color:white; border-radius: 0.5rem ; padding: 1rem; "/>
           </div>

           <div :class="$style.div4">
              <label  for="price2" class="font-bold block mb-2" style="padding-right: 0.5rem;"> Цена / место (2-х) (руб.)</label>
              <InputText id="price2" v-model="price_two" type="number" inputId="withoutgrouping" :useGrouping="false" style="min-width: 100px; max-width: 100px; max-height: 2px; background-color:white; border-radius: 0.5rem ; padding: 1rem; "/> 
           </div>
           
           <div :class="$style.div5">
             <label  for="price3" class="font-bold block mb-2" style="padding-right: 0.5rem;"> Цена / место (3-х) (руб.)</label>
             <InputText id="price3" v-model="price_three" type="number" inputId="withoutgrouping" :useGrouping="false" style="min-width: 100px; max-width: 100px; max-height: 2px; background-color:white; border-radius: 0.5rem ; padding: 1rem; "/>
           </div>
           
           <div :class="$style.div6">
             <label  for="price4" class="font-bold block mb-2" style="padding-right: 0.5rem;"> Цена за место (4-х) (руб.)</label>
             <InputText id="price4" v-model="price_four" type="number" inputId="withoutgrouping" :useGrouping="false" style="min-width: 100px; max-width: 100px; max-height: 2px; background-color:white; border-radius: 0.5rem ; padding: 1rem; "/>
          </div>

        </div>
  


    <div :class="$style.div7">
      <label  for="cashbackRUBInput" class="font-bold block mb-2" style="padding-right: 0.5rem;"> Cashback (руб.)</label>
      <InputText id="cashbackRUBInput" v-model="cashbackRUB" type="number" inputId="withoutgrouping" :useGrouping="false" style="min-width: 100px; max-width: 100px; max-height: 2px; background-color:white; border-radius: 0.5rem ; padding: 1rem;"/>
    </div> 

    <div :class="$style.div8">
      <label  for="cashbackPercentInput" class="font-bold block mb-2" style="padding-right: 0.5rem;"> Cashback (%)</label>
      <InputText id="cashbackPercentInput" v-model="cashbackPercent" type="number" inputId="minmax" :min="0" :max="100" :useGrouping="false" style="min-width: 100px; max-width: 100px; max-height: 2px; background-color:white; border-radius: 0.5rem ; padding: 1rem;"/>
    </div> 

    <div :class="$style.div9">
      <label  for="cashbackNoCheckInRUBInput" class="font-bold block mb-2" style="padding-right: 0.5rem;"> Cashback не заезд (руб.)</label>
      <InputText id="cashbackNoCheckInRUBInput" v-model="cashbackNoCheckInRUB" type="number" inputId="withoutgrouping" :useGrouping="false" style="min-width: 100px; max-width: 100px; max-height: 2px; background-color:white; border-radius: 0.5rem ; padding: 1rem;"/>
    </div> 

    <div :class="$style.div10">
       <label  for="cashbackNoCheckInPercentInput" class="font-bold block mb-2" style="padding-right: 0.5rem;"> Cashback не заезд (%)</label>
       <InputText id="cashbackNoCheckInPercentInput" v-model="cashbackNoCheckInPercent" type="number" inputId="minmax" :min="0" :max="100" :useGrouping="false" style="min-width: 100px; max-width: 100px; max-height: 2px; background-color:white; border-radius: 0.5rem ; padding: 1rem;"/>
    </div> 
    
   


       <ButtonOutlinedSquareD1
          buttonText="Установить"
          propTop="620px"
          propLeft="380px"
          propCursor="pointer"
          propBackgroundColor="transparent"
          propWidth="111px"
          propDisplay="inline-block"
          @click="setupDataForCalendar()"
          style="border-radius: 1rem;"
        />

     <div :class="$style.lineTop"/>

     <div :class="$style.lineBottom"/>

     <div :class="$style.lineRight"/>

<!-- </div>  -->
  <!-- </div> -->
</template>
<script>
  import { defineComponent } from "vue";
  import { FilterMatchMode } from 'primevue/api';
  import FieldsSmallLabel from "../components/FieldsSmallLabel.vue";
  import FieldsSmallDropdown from "../components/FieldsSmallDropdown.vue";
  import ButtonOutlinedSquareD1 from "../components/ButtonOutlinedSquareD1.vue";
  import SidebarDark7 from "../components/SidebarDark7.vue";
  // import CalendarOne1 from "../components/CalendarOne1.vue";
  import DatePicker from "../components/DatePicker.vue";
  import MultiSelect from 'primevue/multiselect';
  import InputNumber from 'primevue/inputnumber';
  import ToggleButton from 'primevue/togglebutton';
  import Column from 'primevue/column';
  import DataTable from 'primevue/datatable';
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
          calendarTable: {
            date:0,
	          installationDate:0,
            quantity: 0,
	          hotelName:"",
            hotelCity:"",
	          hotelAddress:"",
            hotelStars:"",
            priceRoom:0,
            priceBedTwin:0,
	          priceBedTriple:0,
	          priceBedQuadriple:0,
            cashback:0,
            cashbackPercent:0,
            cashbackNoCheckIn:0,
            cashbackNoCheckInPercent:0,
            administratorID:0,
            administratorName:0,
          },
          
            cashbackNoCheckInPercent:0,
            cashbackNoCheckInRUB:0,
            cashbackPercent:0,
            cashbackRUB:0,
            price_one:0,
            price_two:0,
            price_three:0,
            price_four:0,
            quantityNumbers:0,
            selectedHotels: null,
            groupedHotelsFoTable: null,
            groupedHotels: null,
            datePicerValues:{
              startDate: null, 
                endDate: null,  
            },
            loading: false,
            totalRecords: 0,
            dateFrozen: null,
            hotelFrozen: null,
            cityFrozen: null,
            installationDateFrozen: null,

            filters: {
                hotelCity: { value: null, matchMode: FilterMatchMode.IN },
            },
           
        };
    },



    name: "CalendarPage",
    components: {
      InputNumber,
      MultiSelect,
      FieldsSmallLabel,
      FieldsSmallDropdown,
      ButtonOutlinedSquareD1,
      SidebarDark7,
      DatePicker,
      ToggleButton,
      Column,
      DataTable,
    },



    mounted() {

      this.loading = true;

      this.loadLazyData();

      //  axios
      //  .get('https://localhost:9090/getCalendarTable')
      //  .then((res) => {
      //    // assign state posts with response data
      //      this.calendarTable = res.data;
      // // this.totalRecords = res.data.length
      //      this.loading = false;
      //  })
      // .catch((error) => {
      //  // console.log(error.res.data);
      //  });



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

      loadLazyData() {
        
        this.loading = true;

        setTimeout(() => {
                 axios
                .get('https://localhost:9090/getCalendarTable')
                .then((res) => {

                 this.calendarTable = res.data;
                 this.totalRecords = res.data.length
                 this.loading = false;
              })
                .catch((error) => {
                //console.log(error.res.data);
            });

                 this.loading = false;

         }, Math.random() * 1 + 2);

      },


        formatDateTime(value) {     

        const date = new Date(value/1000000);

         return date.toLocaleDateString('ru-RU', {
           day: '2-digit',
           month: '2-digit',
           year: 'numeric',
           hour: 'numeric',
           minute: 'numeric',
          second: 'numeric',
          }); 
        },


       formatDate(value) {

       let date = new Date(value/1000000);

        return date.toLocaleDateString('ru-RU', {
           day: '2-digit',
           month: '2-digit',
           year: 'numeric',
         });
      },

      formatCurrency(value) {
            return value.toLocaleString('ru-RU', { style: 'currency', currency: 'RUB' });
        },

      onUpdateDatePicker(event){

        if (event.target.name == "DateRangeOne"){
          this.datePicerValues.startDate = event.target.value
        }

        if (event.target.name == "DateRangeTwo"){
          this.datePicerValues.endDate = event.target.value
        }

        console.log("Начало периода:" + this.datePicerValues.startDate)
        console.log("Конец периода:" + this.datePicerValues.endDate)

      },


      setupDataForCalendar(){

            if (this.datePicerValues.startDate != null && this.datePicerValues.endDate != null ){

              var jsonSelectedHotels = []
              var index;

              for (index = this.selectedHotels.length - 1; index >= 0; --index) {
                  console.log(this.selectedHotels[index]);

                  jsonSelectedHotels.push({
                                 name: this.selectedHotels[index].label,
                                 address: this.selectedHotels[index].value,
                                 city: this.selectedHotels[index].city,
                               });

              }


  

              // var yearstartdate = new Date(this.datePicerValues.startDate).getFullYear()
              // var monthstartdate = new Date(this.datePicerValues.startDate).getMonth()
              // var daystartdate = new Date(this.datePicerValues.startDate).getDate()

              // var sdate = new Date(yearstartdate, monthstartdate, daystartdate);
              // var offsetInMs = (sdate.getTimezoneOffset() * -60 * 60 *3)  // 

              // var sday = new Date(sdate.getTime() + offsetInMs);

              var sday = new Date(this.datePicerValues.startDate)
              var eday = new Date(this.datePicerValues.endDate)


             // console.log("Начало периода sday => " + sday)
             // console.log("Начало периода offsetInMs => " + offsetInMs)
             // console.log("Начало периода sdate.getTimezoneOffset() => " + sdate.getTimezoneOffset())
             // console.log("Начало периода sday.getTime() => " + sday.getTime())
            //  console.log("Начало периода sday.getTime()*1000000 => " + sday.getTime()*1000000)


              var round = Math.round;

        
               axios
                 .put('https://localhost:9090/updateDataForCalendar',
                 {
                  installationDate:         0,
                  priceRoom:                round(this.price_one),
                  priceBedTwin:             round(this.price_two),
                  priceBedTriple:           round(this.price_three),
                  priceBedQuadriple:        round(this.price_four),
                  numbersQuantity:          round(this.quantityNumbers),
                  cashback:                 round(this.cashbackRUB),
                  cashbackPercent:          round(this.cashbackPercent),
                  cashbackNoCheckIn:        round(this.cashbackNoCheckIn),
                  cashbackNoCheckInPercent: round(this.cashbackNoCheckInPercent),
                  startDate:                sday.getTime()*1000000,
                  endDate:                  eday.getTime()*1000000,
                  jsonSelectedHotels,

                 })
              

                 .then((res) => {

                  console.log(res.data)

                 //this.groupedHotels = res.data;
                 })
                   .catch((error) => {
                //console.log(error.res.data);
                 });
            }


            
      },


      loadMultipleSelectHotels() {
              console.log("Пытаюсь загрузить в groupedHotels")


                 axios
                .get('https://localhost:9090/getGroupedHotels')
                .then((res) => {

                  console.log(res.data)

                 this.groupedHotels = res.data;
                 this.groupedHotelsForTable = res.data;
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
    top: 50px;
    left: 255px;
    background-color: #40409a;
    width: 1650px;
    height: 1.8px;
  }

  .lineTop{
    position: absolute;
    top: 0px;
    left: 240px;
    background-color: #b4b4d2;
    width: 1680px;
    height: 0.5px;
  }

  .lineBottom{
    position: absolute;
    top: 1079px;
    left: 240px;
    background-color: #b4b4d2;
    width: 1680px;
    height: 0.5px;
  }

  .lineRight{
    position: absolute;
    top: 0px;
    left: 1920px;
    background-color: #b4b4d2;
    width: 0.5px;
    height: 1080px;
  }

  .toolbarTableStyle{
    background: conic-gradient(from 145deg at 40%  75%, rgba(148,179,207,0.26) 0%, rgba(156, 184, 216, 0.75) 30%, transparent 50%, rgba(156, 184, 216, 0.75) 80%, rgba(115, 128, 216, 0.51) 100%) 15% 75%/175% 200%,
      radial-gradient(ellipse  at 55%  55%, rgb(178, 166, 222) 0%, rgb(9, 3, 22) 100%) 30% 30%/200% 170%;
}

  .calendarTableDatesStyle{
    position: absolute;
    top: 60px;
    left: 650px;
    background-color: #40409a;
    width: 1272px;
    height: 1020px;
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

  .div_main{
    position: absolute;
    top: 150px;
    /* left: 1091px; */
    left: 260px;
    max-width: 300px;
    line-height: 150%;
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

  .div0 {
    position: absolute;
    top: 30px;
    left: 380px;
    font-size: var(--label-large-label-14-size);
    line-height: 24px;
    font-weight: 500;
    text-align: center;
  }
  
  .div1 {
    position: absolute;
    top: 12px;
    /* left: 1127px; */
    left: 255px;
    line-height: 150%;
    font-size: var(--label-large-label-14-size);
    line-height: 24px;
    font-weight: 500;
    text-align: center;
  }
  .div2 {
    position: absolute;
    top: 150px;
    left: 315px;
    line-height: 150%;
    font-size: var(--label-large-label-14-size);
    line-height: 24px;
    font-weight: 500;
    text-align: center;
  }
  .div3 {
    position: absolute;
    top: 200px;
    left: 300px;
    /* max-width: 200px; */
    line-height: 150%;
    font-size: var(--label-large-label-14-size);
    line-height: 24px;
    font-weight: 500;
    text-align: center;
  }
  .div4 {
    position: absolute;
    top: 250px;
    left: 265px;
    line-height: 150%;
    font-size: var(--label-large-label-14-size);
    line-height: 24px;
    font-weight: 500;
    text-align: center;
  }
  .div5 {
    position: absolute;
    top: 300px;
    /* left: 1166px; */
    left: 265px;
    line-height: 150%;
    font-size: var(--label-large-label-14-size);
    line-height: 24px;
    font-weight: 500;
    text-align: center;
  }
  .div6 {
    position: absolute;
    top: 350px;
    left: 255px;
    line-height: 150%;
    font-size: var(--label-large-label-14-size);
    line-height: 24px;
    font-weight: 500;
    text-align: center;
  }
  .div7 {
    position: absolute;
    top: 400px;
    left: 325px;
    line-height: 150%;
    font-size: var(--label-large-label-14-size);
    line-height: 24px;
    font-weight: 500;
    text-align: center;
  }
  .div8 {
    position: absolute;
    top: 450px;
    left: 345px;
    font-size: var(--label-large-label-14-size);
    line-height: 24px;
    font-weight: 500;
    text-align: center;
  }

  .div9 {
    position: absolute;
    top: 500px;
    left: 260px;
    font-size: var(--label-large-label-14-size);
    line-height: 24px;
    font-weight: 500;
    text-align: center;
  }

  .div10 {
    position: absolute;
    top: 550px;
    left: 280px;
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

   .datarangeoneVtextfield {
    position: absolute;
    top: calc(45% - 29px);
    left: calc(45% - 0px);
  }

.input__box {
    display: flex;
    flex-direction: column;
    position: relative;
}

.name__input{
    background: #fff;
    padding: 1px;
    border: 1px solid #6c6c6c;
    color: #000;

}

.name__input_label{
    position: absolute;
    color: #6c6c6c;
    top: -10px;
    left: 20px;
    background: #fff;
    padding: 0 5px;
}


</style>
