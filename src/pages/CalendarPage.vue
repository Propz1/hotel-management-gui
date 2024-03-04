<template>

        
  <div :class="$style.calendarpage">



   <div :class="$style.calendarpageChild"/>

    <div class="card">

                  <div :class="$style.buttonGETfromDB">
                    <Button type="button" label="Обновить ⬅ " icon="pi pi-database" iconPos="right" style="padding-right:0.5rem; margin-right: 0.4rem; height:35px; width:145px; backgroundColor: var(--primary-color); color: var(--primary-color-text)" @click="loadLazyData()"/>
                  </div>

                  <div :class="$style.buttonPUTtoDB">
                    <Button type="button" label="Записать ➞ " icon="pi pi-database" iconPos="right" style="padding-right:0.5rem; margin-right: 0.4rem; height:35px; width:145px; backgroundColor: var(--primary-color); color: var(--primary-color-text)" @click="updateDataForCalendar()"/>
                  </div>
    
                <Toolbar :class="$style.toolbarTableStyle" style="padding-left: 0.3rem; min-height: 3.5rem;">     
                   <template #start>
                    <!-- <Button type="button" icon="pi pi-chevron-left" rounded style="height:35px; width:35px; backgroundColor: var(--primary-color); color: var(--primary-color-text)" @click="previousYear(index)"/>    
                     <div id="appSelectYear">
                           <SwipeBox ref="myswipe" @onChange="yearChanged" speed="150">
                                 <div style="width: 160px; height: 20px; border: 0px solid black">
                                    <div v-for="image in imagesYear" :key="image.id">
                                      <img :src="image.url" style="margin-left: 1rem; margin-right: 1rem; width: 120px;  height: 20px"/>
                                    </div>
                                  </div> 
                           </SwipeBox>
                    </div>
                     <Button type="button" icon="pi pi-chevron-right" rounded style="margin-right: 0.5rem; height:35px; width:35px; backgroundColor: var(--primary-color); color: var(--primary-color-text)" @click="nextYear(index)"/> -->
                       
                      <div :class="$style.buttonsDatePicker">
                         <DatePicker id="datePicker" v-model="datePicerValues" @input="onUpdateDatePicker($event)"/>
                      </div>

                      <div :class="$style.selectHotelsStyle">
                        <div class="card flex flex-wrap gap-2 p-fluid">       
                          <MultiSelect class="w-full md:w-50rem" style="min-width:22rem"
                                v-model="selectedHotels" 
                                :options="groupedHotels" 
                                optionLabel="label"
                                optionGroupLabel="label" 
                                optionGroupChildren="items"
                                display="chip"
                                placeholder="Выберите отель" 
                                :selectionLimit=1 
                                :maxSelectedLabels=1
                                :showToggleAll=false 
                                @click="loadMultipleSelectHotels()" 
                               >
                           <template #optiongroup="slotProps">
                             <div class="flex align-items-center">
                               <img :alt="slotProps.option.label" src="https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png" :class="`flag flag-${slotProps.option.code.toLowerCase()} mr-2`" style="width: 18px" />
                               <div>{{ slotProps.option.label }}</div>
                             </div>
                            </template>
                          </MultiSelect>
                        </div>
                      </div>

                    <div :class="$style.buttonsMonth">
                       <Button  label="Январь"  style="color: white; font-size: 0.9rem; font-weight:400;padding-left: .62rem; padding-right: 0.62rem;padding-top: 0.82rem; padding-bottom: 0.82rem; " />
                       <Button  label="Февраль" style="color: white; font-size: 0.9rem; font-weight:400;padding-left: 0.5rem; padding-right: 0.5rem;padding-top: 0.80rem; padding-bottom: 0.80rem; " />
                       <Button  label="Март" style="color: white; font-size: 0.9rem; font-weight:400;padding-left: 1.22rem; padding-right: 1.22rem;padding-top: 0.80rem; padding-bottom: 0.80rem; " />
                       <Button  label="Апрель" style="color: white; font-size: 0.9rem; font-weight:400;padding-left: 0.75rem; padding-right: .75rem;padding-top: 0.80rem; padding-bottom: 0.80rem; " />
                       <Button  label="Май" style="color: white; font-size: 0.9rem; font-weight:400;padding-left: 1.29rem; padding-right: 1.29rem;padding-top: 0.80rem; padding-bottom: 0.80rem; " />
                       <Button  label="Июнь" style="color: white; font-size: 0.9rem; font-weight:400;padding-left: 0.65rem; padding-right: 0.65rem;padding-top: 0.80rem; padding-bottom: 0.80rem; " />
                       <Button  label="Июль" style="color: white; font-size: 0.9rem; font-weight:400;padding-left: 0.65rem; padding-right: 0.65rem;padding-top: 0.80rem; padding-bottom: 0.80rem; " />
                       <Button  label="Август" style="color: white; font-size: 0.9rem; font-weight:400;padding-left: 0.61rem; padding-right: 0.61rem;padding-top: 0.80rem; padding-bottom: 0.80rem; " />
                       <Button  label="Сентябрь" style="color: white; font-size: 0.9rem; font-weight:400;padding-left: 0.3rem; padding-right: 0.3rem;padding-top: 0.80rem; padding-bottom: 0.80rem; " />
                       <Button  label="Октябрь" style="color: white; font-size: 0.9rem; font-weight:400;padding-left: 0.6rem; padding-right: 0.6rem;padding-top: 0.80rem; padding-bottom: 0.80rem; " />
                       <Button  label="Ноябрь" style="color: white; font-size: 0.9rem; font-weight:400;padding-left: 0.62rem; padding-right: 0.62rem;padding-top: 0.80rem; padding-bottom: 0.80rem; " />
                       <Button  label="Декабрь" style="color: white; font-size: 0.9rem; font-weight:400;padding-left: 0.3rem; padding-right: 0.3rem;padding-top: 0.80rem; padding-bottom: 0.80rem; " />
                    </div>

                    </template>
                  </Toolbar>


                  <div :class="$style.calendarTableDatesStyle">
                         <sc
                           :schedule-data="scData"
                           :setting="setting"
                           @row-click-event="rowClickEvent"
                           @date-click-event="dateClickEvent"
                           @click-event="clickEvent"
                           @add-event="addEvent"
                           @move-event="moveEvent"
                           @edit-event="editEvent"
                           @delete-event="deleteEvent"
                          ></sc>
                  </div>
                
    </div>

    <SidebarDark7 />

    
 
   

</div>    
</template>


<script>
  import { defineComponent } from "vue";
  import { ref } from 'vue';
  import schedulerLite from "../components/schedulerLite.vue";
  import { FilterMatchMode } from 'primevue/api';
  import FieldsSmallLabel from "../components/FieldsSmallLabel.vue";
  import FieldsSmallDropdown from "../components/FieldsSmallDropdown.vue";
  import ButtonOutlinedSquareD1 from "../components/ButtonOutlinedSquareD1.vue";
  import SidebarDark7 from "../components/SidebarDark7.vue";
  import DatePicker from "../components/DatePicker.vue";
  import MultiSelect from 'primevue/multiselect';
  import InputNumber from 'primevue/inputnumber';
  import InputText from 'primevue/inputtext';
  import ToggleButton from 'primevue/togglebutton';
  import Column from 'primevue/column';
  import DataTable from 'primevue/datatable';
  import axios from 'axios-https-proxy-fix';
  import SwipeBox from '@shopid/vue3-swipe-box';
  import { PrimeIcons } from 'primevue/api';
  import 'primeicons/primeicons.css';


const month = ref({
  month: new Date().getMonth(),
  year: new Date().getFullYear(),
});

const formatMonthYear = (date) => {
  const month = date.getMonth();
  const year = date.getFullYear();
  let monthArray = ["Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь"];
  let monthText = monthArray[month];
  return `${monthText}, ${year}`;
};

const getDateDiff = (date1, date2) => {
        const diffTime = Math.abs(date2 - date1);
        return Math.ceil(diffTime / (1000 * 60 * 60 *24))-1;//
      };

var startSetting = {
             startDate: "2023/1/3 00:00",
             endDate: "2026/1/1 24:00",
             resourceNames:["quantity", "price_room", "price_bed_twin", "price_bed_triple", "price_bed_quadriple", "cashback", "cashback_percent", "cashback_no_check_in","cashback_no_check_in_percent"],
             weekdayText: ["ВС", "ПН", "ВТ", "СР", "ЧТ", "ПТ", "СБ"],
             unit: 1440, // Minutes 
             borderW: 1, // Px
             dateDivH: 26, // Px 26
             timeDivH: 26, // Px
             unitDivW: 54, // Px //width of the cells
             titleDivW: 8, // Percent
             rowH: 57, // Px
       };


  export default defineComponent({

    data() {
        return {

           scData:  [],
           setting: startSetting,
           month,
           formatMonthYear,

           showDate: new Date(),

          //  imagesYear: [
          //   { id: '2024', url: "./year2024.svg" },
          //   { id: '2025', url: "./year2025.svg" },
          //   { id: '2026', url: "./year2026.svg" },
          //   { id: '2027', url: "./year2027.svg" },
          //   { id: '2028', url: "./year2028.svg" },
          //   { id: '2029', url: "./year2029.svg" },
          //   { id: '2030', url: "./year2030.svg" },
          //   { id: '2031', url: "./year2031.svg" },
          //   { id: '2032', url: "./year2032.svg" },
          //   { id: '2033', url: "./year2033.svg" },
          //   { id: '2034', url: "./year2034.svg" },
          //   { id: '2035', url: "./year2035.svg" },
          //   { id: '2036', url: "./year2036.svg" },
          //   { id: '2037', url: "./year2037.svg" },
          //   { id: '2038', url: "./year2038.svg" },
          //   { id: '2039', url: "./year2039.svg" },
          //   { id: '2040', url: "./year2040.svg" },
          //  ],

          // calendarTable: {
          //   date:0,
	        //   installationDate:0,
          //   quantity: 0,
	        //   hotelName:"",
          //   hotelCity:"",
	        //   hotelAddress:"",
          //   hotelStars:"",
          //   priceRoom:0,
          //   priceBedTwin:0,
	        //   priceBedTriple:0,
	        //   priceBedQuadriple:0,
          //   cashback:0,
          //   cashbackPercent:0,
          //   cashbackNoCheckIn:0,
          //   cashbackNoCheckInPercent:0,
          //   administratorID:0,
          //   administratorName:0,
          // },
          
            installationDate:0,
            cashbackNoCheckInPercent:0,
            cashbackNoCheckInRUB:0,
            cashbackPercent:0,
            cashbackRUB:0,
            price_one:0,
            price_two:0,
            price_three:0,
            price_four:0,
            quantityNumbers:0,
            selectedHotels: [],
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
      InputText,
      MultiSelect,
      FieldsSmallLabel,
      FieldsSmallDropdown,
      ButtonOutlinedSquareD1,
      SidebarDark7,
      DatePicker,
      ToggleButton,
      Column,
      DataTable,
      PrimeIcons,
      SwipeBox, 
      sc: schedulerLite,
    },

    mounted() {

      this.loading = false;
      this.loadLazyData();
      this.updateDataForCalendar();
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

      onUpdateSetting() {
        this.setting.startDate = "2022/1/1 00:00"
        this.setting.endDate = "2025/1/1 00:00"
      },

      dateClickEvent(date) {
      console.log("------");
      console.log("DateClickEvent:");
      console.log("Date:" + date);
    },
    rowClickEvent(rowIndex, text) {
      console.log("------");
      console.log("RowClickEvent:");
      console.log("RowIndex:" + rowIndex);
      console.log("RowTitle:" + text);
    },
    
    clickEvent(startDate, endDate, text, other) {
      console.log("------");
      console.log("ClickEvent:");
      console.log("StartDate:" + startDate);
      console.log("EndDate:" + endDate);
      console.log("ContentText:" + text);
      
      if (other) {
        console.log("OtherData:");
        console.log(other);
      }
    },
    addEvent(rowIndex, startDate, endDate) {
      console.log("------");
      console.log("AddEvent:");
      console.log("RowIndex:" + rowIndex);
      console.log("StartDate:" + startDate);
      console.log("EndDate:" + endDate);
    },
    moveEvent(status, newStartDate, newEndDate) {
      console.log("------");
      console.log("MoveEvent:");
      if (status == 1) {
        console.log("NewStartDate:" + newStartDate);
        console.log("NewEndDate:" + newEndDate);
      } else if (status == 2) {
        console.log("Has other event, can't move.");
      } else {
        console.log("Not businessDay, can't move.");
      }
    },
    editEvent(newStartDate, newEndDate) {
      console.log("------");
      console.log("EditEvent:");
      console.log("NewStartDate:" + newStartDate);
      console.log("NewEndDate:" + newEndDate);
    },
    deleteEvent(row, index) {
      console.log("------");
      console.log("DeleteEvent:");
      console.log("Row:" + row);
      console.log("Index:" + index);
    },
    
      setShowDate(d) {
				this.showDate = d;
			},

      updateDataForCalendar(){

        this.loading = true;

         var schedule = [];
         var i;
         var j;
         var index;
   
         if (this.selectedHotels.length > 0 ) {

             for (i = 0; i <= this.scData.length - 1; i++) {
               for (j = 0; j <= this.scData[i].schedule.length - 1; j++) {

                   schedule.push({
                             name:   this.scData[i].schedule[j].resourceName,
                             value:  this.scData[i].schedule[j].text,
                         startYear:  new Date(this.scData[i].schedule[j].start).getFullYear(),
                        startMonth:  new Date(this.scData[i].schedule[j].start).getMonth() + 1,
                          startDay:  new Date(this.scData[i].schedule[j].start).getDate(),
                              diff:  getDateDiff(
                                       new Date(this.scData[i].schedule[j].start),
                                       new Date(this.scData[i].schedule[j].end)
                                     ),    
                    });

                }
              }

   
           for (index = this.selectedHotels.length - 1; index >= 0; index--) {
               var jsonSelectedHotels = [{name: this.selectedHotels[index].label, address: this.selectedHotels[index].value, city: this.selectedHotels[index].city}];
            }

          var yearStartDate = new Date(this.setting.startDate).getFullYear()
          var monthStartDate = new Date(this.setting.startDate).getMonth() + 1
          var dayStartDate = new Date(this.setting.startDate).getDate()

          var yearEndDate = new Date(this.setting.endDate).getFullYear()
          var monthEndDate = new Date(this.setting.endDate).getMonth() + 1
          var dayEndDate = new Date(this.setting.endDate).getDate()-1 


           axios
            .put('https://localhost:9090/updateDataForCalendar',
             {
               installationDate:  0,
               hotels:            jsonSelectedHotels,
               resources:         schedule,
               startDay:          dayStartDate,
               startMonth:        monthStartDate,
               startYear:         yearStartDate,
               endDay:            dayEndDate,
               endMonth:          monthEndDate,
               endYear:           yearEndDate,
            })

      
           .then((res) => {
              this.loadLazyData()
            })
               
            .catch((error) => {
              //console.log(error.res.data);
            });

            this.loading = false;
    
        }
     },

    
      loadLazyData() {

        var today = new Date();

        var yearStartDate 
        var monthStartDate 
        var dayStartDate 

        var yearEndDate
        var monthEndDate 
        var dayEndDate 

        
        if (this.datePicerValues.startDate === null && this.datePicerValues.endDate === null ){  
            this.datePicerValues.startDate = new Date(today.getFullYear(), today.getMonth(), today.getDate());
            this.datePicerValues.endDate = new Date(today.getFullYear() + 1, today.getMonth(), today.getDate());
        }

        yearStartDate = new Date(this.datePicerValues.startDate).getFullYear()
        monthStartDate = new Date(this.datePicerValues.startDate).getMonth() + 1
        dayStartDate = new Date(this.datePicerValues.startDate).getDate()

        yearEndDate = new Date(this.datePicerValues.endDate).getFullYear()
        monthEndDate = new Date(this.datePicerValues.endDate).getMonth() + 1
        dayEndDate = new Date(this.datePicerValues.endDate).getDate() //for period overlaps


        this.setting = {
                          startDate: `${yearStartDate}/${monthStartDate}/${dayStartDate} 00:00`,
                          endDate: `${yearEndDate}/${monthEndDate}/${dayEndDate} 24:00`,
                          resourceNames:["quantity", "price_room", "price_bed_twin", "price_bed_triple", "price_bed_quadriple", "cashback", "cashback_percent", "cashback_no_check_in","cashback_no_check_in_percent"],
                          weekdayText: ["ВС", "ПН", "ВТ", "СР", "ЧТ", "ПТ", "СБ"],
                          unit: 1440, // Minutes 
                          borderW: 1, // Px
                          dateDivH: 26, // Px 26
                          timeDivH: 26, // Px
                          unitDivW: 54, // Px //width of the cells
                          titleDivW: 8, // Percent
                          rowH: 57, // Px
                        };

        //var jsonSelectedHotels = []     
        var h  


         console.log("loadLazyData():  this.selectedHotels.length =>" + this.selectedHotels.length)

      if (this.selectedHotels.length > 0 ) {

          // jsonSelectedHotels.push({
          //                        name: this.selectedHotels[0].label,
          //                        address: this.selectedHotels[0].value,
          //                        city: this.selectedHotels[0].city,
          //                      });
                              
         h = this.selectedHotels[0].label

         console.log("loadLazyData():  this.selectedHotels[0].label =>" + this.selectedHotels[0].label)
         
        const param = {startDay:   dayStartDate, 
                       startMonth: monthStartDate, 
                       startYear:  yearStartDate,
                       endDay:     dayEndDate,
                       endMonth:   monthEndDate,
                       endYear:    yearEndDate,
                       hotels:     h};

        
        this.loading = true;

   
                 axios
                .get('https://localhost:9090/getCalendarTable', {params: param})
                .then((res) => {

                  this.scData = res.data;
               
                  })
                .catch((error) => {
                //console.log(error.res.data);
                });

                this.loading = false;

      }

     

    },

        formatDateTime(value) {     

        const date = new Date(value/1000000)

        date.setDate(date.getDate())

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

       let date = new Date(value/1000000)
     
       date.setDate(date.getDate())

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


      


      loadMultipleSelectHotels() {
                 axios
                .get('https://localhost:9090/getGroupedHotels')
                .then((res) => {
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
    top: 75px;
    left: 255px;
    background-color: #40409a;
    width: 0px;
    height: 0px;
  }

  .toolbarTableStyle{
    position: absolute;
    top: 50px;
    left: 249px;
    width: 100%;
    height: 3rem;
    background: conic-gradient(from 145deg at 40%  75%, rgba(148,179,207,0.26) 0%, rgba(156, 184, 216, 0.75) 30%, transparent 50%, rgba(156, 184, 216, 0.75) 80%, rgba(115, 128, 216, 0.51) 100%) 15% 75%/175% 200%,
      radial-gradient(ellipse  at 55%  55%, rgb(178, 166, 222) 0%, rgb(9, 3, 22) 100%) 30% 30%/200% 170%;  

   }

   .toolbar2TableStyle{
    position: absolute;
    top:70px;
    left: 249px;
    width: 100%;
    height: 1rem;
    background: black;
    background:rgba(91, 108, 128, 0.75);
    background: #f1eded;
    
   }

   .buttonsDatePicker{
    position: absolute;
    left: 10px;
    top:3px;
   }

   .buttonsMonth{
    position: absolute;
    left: 730px;
   }

  .calendarTableDatesStyle{
    display: grid;
    position: absolute;
    top: 120.9px;
    left: 249px;
    background-color: white;
    width: 89.5%;
    height: 1010px;
  }
  
  .selectHotelsStyle{
    position: absolute;
    top: 9px;
    /* left: 1127px; */
    left: 370px;
    line-height:150%;
    font-size: var(--label-large-label-14-size);
    line-height: 35px;
    font-weight: 500;
    text-align: center;
  }

  .buttonGETfromDB{
    position: absolute;
    top: 5px;
    left: 250px;
  }

  .buttonPUTtoDB{
    position: absolute;
    top: 5px;
    left: 420px;
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

  .sample {
  width: 10px;
  height: 10px;
  margin: 5px;
  border: 1px solid black;
}

.cant-res {
  background-color: #999 !important;
}

.reserved {
  background-color: #ec920a !important;
}

</style>

