<script setup>
  //import { ref } from "vue";
  import Toast from 'primevue/toast';
  import Button from 'primevue/button';
  import ToggleButton from 'primevue/togglebutton';
  import InputText from 'primevue/inputtext';
  import DataTable from 'primevue/datatable';
  import Column from 'primevue/column';
  import Dropdown from 'primevue/dropdown';
  import Tag from 'primevue/tag';
  import Calendar from 'primevue/calendar';
// import ColumnGroup from 'primevue/columngroup';   // optional
// import Row from 'primevue/row';                   // optional


 //const toast = ToastService.useToast()
 
  //const userName = ref()
   // set default theme folder-name to currentTheme
   //const currentTheme = ref('lara-light-purple')
  

</script>

<template>
  <!-- <Toast position="top-center" group="headless" @close="visible = false"></Toast> -->
<!-- <Toast></Toast> -->

<div class="wrapperColumns">
    
  <Toast/>
 
  <div :class="$style.requisitions"><SidebarDark /></div>
  
    <div class="wrapperRows">

     <!-- <div class="containerGuestName">
        <span class="p-float-label p-input-filled" style="margin-right: 1rem">
       <InputText id="userName" type="text" v-model="userName" />
       <label for="userName">ФИО</label>
         </span>
        <div class="buttons">
         <Button  class="p-inputtext-sm" label="Записать" icon="pi pi-save" iconPos="right" severity="success" @click="showSuccess"  style="background-color:oldlace" ></Button>
        </div>
     </div> -->


      <div class="requisitionDataTableContainer">

          <!-- <DataTable :value="products" showGridlines paginator :rows="5" :rowsPerPageOptions="[5, 10, 20, 50]"  tableStyle="min-width: 50rem">
            <Column field="id" header="№" style="width: 25%"></Column>
            <Column field="name" header="Name"></Column>
            <Column field="country.name" header="Country"></Column>
            <Column field="company" header="Company"></Column>
            <Column field="status" header="Статус"></Column>
            <Column field="verified" header="Подтверждена"></Column>
            <Column field="representative.name" header="Representative"></Column>
          </DataTable> -->


         
           <Toolbar :class="$style.toolbarTableStyle" style="min-height: 2rem">
             <!-- <Toolbar style="border-radius: 3rem; background-image: linear-gradient(to right, var(--bluegray-500), var(--bluegray-800))"> -->
                <template #start>
                  <ToggleButton v-model="editingFrozen" class="p-inputtext-sm" onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="✍🏼" offLabel="✍🏼" style="min-width: 3rem; margin-left: 0.4rem; margin-right: 1rem; border-radius: 0.5rem"/>
                  <ToggleButton v-model="requisitionNumberFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="№" offLabel="№" style="min-width: 3.5rem; margin-left: 0.5rem; margin-right: 1rem; border-radius: 1rem" />
                  <ToggleButton v-model="applicationSubmissionTimeFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Дата заявки" offLabel="Дата заявки" style="min-width: 8.8rem; margin-right: 1rem; border-radius: 1rem" />
                  <ToggleButton v-model="rejectButtonFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="👎" offLabel="👎" style="min-width: 3.5rem; margin-left: 1rem; margin-right: 1rem; border-radius: 1rem" />
                  <ToggleButton v-model="confirmButtonFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="👍" offLabel="👍" style="min-width: 3.5rem; margin-right: 1rem; border-radius: 1rem" />
                  <ToggleButton v-model="statusFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Статус" offLabel="Статус" style="min-width: 6rem; margin-left: 1rem; margin-right: 1rem; border-radius: 1rem" />
                  <ToggleButton v-model="applicationStatusDateFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Время статуса" offLabel="Время статуса" style="min-width: 10.2rem; margin-right: 1rem; border-radius: 1rem" />
                  <ToggleButton v-model="checkInDateFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Заезд" offLabel="Заезд" style="min-width: 5.5rem; margin-right: 1rem; border-radius: 1rem" />
                  <ToggleButton v-model="checkOutDateFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Выезд" offLabel="Выезд" style="min-width: 5.8rem; margin-right: 1rem; border-radius: 1rem" />
                  <ToggleButton v-model="costFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Стоимость" offLabel="Стоимость" style="min-width: 8rem; margin-right: 1rem; border-radius: 1rem"/>
                  <ToggleButton v-model="cashbackFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Кешбэк" offLabel="Кешбэк" style="min-width: 6rem; margin-right: 1rem; border-radius: 1rem"/>
                  <ToggleButton v-model="cashbackNoCheckInFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Кешбэк (незаезд)" offLabel="Кешбэк (незаезд)" style="min-width: 11.5rem; margin-right: 1rem; border-radius: 1rem"/>
                  <ToggleButton v-model="userNameFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="ФИО" offLabel="ФИО" style="min-width: 5rem; margin-right: 1rem; border-radius: 1rem" />
                  <ToggleButton v-model="userEmailFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Email" offLabel="Email" style="min-width: 5.4rem; margin-right: 1rem; border-radius: 1rem" />
                  <ToggleButton v-model="userPhoneFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Телефон" offLabel="Телефон" style="min-width: 7rem; margin-right: 1rem; border-radius: 1rem" />
                  <ToggleButton v-model="hotelCityFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Город" offLabel="Город" style="min-width: 5.5rem; margin-right: 1rem; border-radius: 1rem" />
                  <ToggleButton v-model="hotelNameFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Гостиница" offLabel="Гостиница" style="min-width: 8rem; margin-right: 1rem; border-radius: 1rem" />
                </template>

             <template #end>

              </template>
          </Toolbar>


                      <!-- v-model:selection="selectedRequisitions" :selectOne="selectOne" @select-all-change="onSelectAllChange" @row-select="onRowSelect" @row-unselect="onRowUnselect" tableStyle="min-width: 75rem" -->

                      

          <div class="card p-fluid">
          
            <DataTable 
            :value="requisitionsTable" 
            :frozenValue="lockedRequisitions"
            class="mt-4"
            resizableColumns columnResizeMode="expand"    
            lazy paginator :first="first"  
            :rows="40"
            scrollable scrollHeight="1080px"
            :virtualScrollerOptions="{ itemSize: 36 }"
            v-model:filters="filters" ref="dt" 
            :totalRecords="totalRecords" :loading="loading" @page="onPage($event)" @sort="onSort($event)" @filter="onFilter($event)" filterDisplay="row"
            :globalFilterFields="['cost','user.email','user.phone','anotherCity','anotherHotel','hotel.city', 'hotel.name', 'hotel.stars','representative.name']"
             v-model:selection="selectedRequisitions" :selectAll="selectAll" @select-all-change="onSelectAllChange" @row-select="onRowSelect($event)" @row-unselect="onRowUnselect" tableStyle="min-width: 75rem" 
             v-model:editingRows="editingRows" editMode="row" dataKey="id" @row-edit-save="onRowEditSave" 
             :pt="{
               table: { style: 'min-width: 40rem' },
              //  bodyrow: ({ props }) => ({
              //   class: [{ 'font-bold': props.frozenRow }]
              //   }),
              column: {
                    bodycell: ({ state }) => ({
                        style:  state['d_editing']&&'padding-top: 0.6rem; padding-bottom: 0.6rem'
                    })
                }
            }"
              >


            <!-- <Column style="flex: 0 0 5rem">
                <template #body="{ data, frozenRow, index }">
                  <Button type="button" style="color:green" :icon="frozenRow ? 'pi pi-lock' : 'pi pi-lock-open'" :disabled="frozenRow ? false : lockedRequisitions.length >= 2" text size="small" @click="toggleLock(data, frozenRow, index)" />
                </template>
            </Column> -->

            <!-- <Column selectionMode="single" headerStyle="width: 4rem"></Column> --> 

           
            
            <Column :rowEditor="true" style="min-width: 4rem; margin-left: 2rem; margin-right: 1rem;" bodyStyle="text-align:center" alignFrozen="left" :frozen="editingFrozen"></Column>

            <Column field="requisitionNumber" header="№" headerStyle="width: 2%; margin-left: 2rem; text-align: center" sortable: false style="min-width: 85px; width: 2%; padding-left: 0.1rem; padding-right: 0.1rem; text-align: center; font-weight:820" alignFrozen="left" :frozen="requisitionNumberFrozen">
     
            </Column>  
            <!-- <Column field="applicationSubmissionTime" header="Дата заявки" style="min-width: 200px"></Column> -->
            <Column field="applicationSubmissionTime" header="Дата заявки" filterField="applicationSubmissionTime" dataType="date" style="min-width: 10rem; padding-left: 0.5rem; padding-right: 0.5rem" bodyStyle="text-align:center" alignFrozen="left" :frozen="applicationSubmissionTimeFrozen">
                <template #body="{ data }">
                    {{ formatDateTime(data.applicationSubmissionTime) }}
                </template>
                <!-- <template #filter="{ filterModel }">
                    <Calendar v-model="filterModel.value" dateFormat="dd.mm.yy" placeholder="dd.mm.yyyy" mask="99.99.9999" />
                </template> -->
            </Column>

            <Column header="Отклонить" headerStyle="width: 3rem; text-align: center" bodyStyle="text-align: center; overflow: visible" style="min-width: 3rem; width: 1%; padding-left: 0.2rem; padding-right: 0.2rem" alignFrozen="left" :frozen="rejectButtonFrozen">
                <template #body="{ data, index }">
                    <Button type="button" icon="pi pi-thumbs-down" rounded style="height:35px; width:35px; backgroundColor: var(--primary-color); color: var(--primary-color-text)" @click="rejectBookingRequest(data, index)"/>
                </template>
            </Column>
            <Column header="Подтвердить" headerStyle="width: 3rem; text-align: center" bodyStyle="text-align: center; overflow: visible" style="min-width: 3rem; width: 1%; padding-left: 0.2rem; padding-right: 0.2rem" alignFrozen="left" :frozen="confirmButtonFrozen">
                <template #body="{ data, index }">
                    <Button type="button" icon="pi pi-thumbs-up" rounded style="height:35px; width:35px; backgroundColor: var(--primary-color); color: var(--primary-color-text)" @click="confirmeBookingRequest(data, index)"/>
                </template>
            </Column>
            
            <Column field="status"  header="Статус" headerStyle="text-align: center"  :filterMenuStyle="{ width: '5rem'}" style="min-width: 3rem; padding-left: 0.3rem; padding-right: 0.3rem; text-align: left" alignFrozen="left" :frozen="statusFrozen" >
                <template #body="{ data }">
                  <Tag :value=formatEmojiStatus(data.status) style="background-color: white; color: white; font-size: 1rem; font-weight:400; padding-left: 0.7rem; padding-right: 0.7rem;"/>  
                  <Tag :value=formatStatus(data.status) :severity="getSeverity(formatStatus(data.status))" style="color: white; font-size: 1rem; font-weight:400; padding-left: 0.7rem; padding-right: 0.7rem;"/>
                </template>
                <!-- <template #filter="{ filterModel }">
                    <Dropdown v-model="filterModel.value" :options="statuses" placeholder="Выберите статус" class="p-column-filter" showClear>
                        <template #option="slotProps">
                            <Tag :value="slotProps.option" :severity="getSeverity(slotProps.option)" />
                        </template>
                    </Dropdown>
                </template> -->
            </Column>

            <!-- <Column field="applicationStatusDate" header="Время статуса" style="min-width: 200px"></Column> -->
            <Column field="applicationStatusDate" header="Время статуса 🕓" headerStyle="text-align: center" filterField="applicationStatusDate" dataType="date" style="min-width: 12rem; margin-right: 2rem;" bodyStyle="text-align:left"  alignFrozen="left" :frozen="applicationStatusDateFrozen" >
                <template #body="{ data }">
                    {{ formatDateTime(data.applicationStatusDate) }}
                </template>
                <!-- <template #filter="{ filterModel }">
                    <Calendar v-model="filterModel.value" dateFormat="dd.mm.yy" placeholder="dd.mm.yyyy" mask="99.99.9999" />
                </template> -->
            </Column>

             <Column field="checkInDate" header="Заезд" filterField="checkInDate" dataType="date" style="min-width: 10rem;" bodyStyle="text-align:left" alignFrozen="left" :frozen="checkInDateFrozen">
                <template #body="{ data }">
                    {{ formatDate(data.checkInDate) }}
                </template>
                <!-- <template #filter="{ filterModel }">
                    <Calendar v-model="filterModel.value" dateFormat="dd.mm.yy" placeholder="dd.mm.yyyy" mask="99.99.9999" />
                </template> -->
            </Column>

            <Column field="checkOutDate" header="Выезд" filterField="checkOutDate" dataType="date" style="min-width: 9.5rem;" bodyStyle="text-align:left" alignFrozen="left" :frozen="checkOutDateFrozen">
                <template #body="{ data }">
                    {{ formatDate(data.checkOutDate) }}
                </template>
                <!-- <template #filter="{ filterModel }">
                    <Calendar v-model="filterModel.value" dateFormat="dd.mm.yy" placeholder="dd.mm.yyyy" mask="99.99.9999" />
                </template> -->
            </Column>
            
            <Column field="cost" header="Стоимость" style="min-width: 6rem; width: 2%; background-color: beige; font-weight:700" bodyStyle="text-align:center" alignFrozen="left" :frozen="costFrozen">
                 <!-- <template #body="{ data }">
                    <span class="font-bold">{{ formatCurrency(data.balance) }}</span>
                </template> -->
                <template #editor="{ data, field }">
                    <InputText v-model="data[field]" />
                </template>
            </Column>

            <Column field="cashback" header="Кешбэк" style="min-width: 6rem; width: 2%; padding-left: 1rem; padding-right: 1rem;" bodyStyle="text-align:center" alignFrozen="left" :frozen="cashbackFrozen">
              <template #body="{ data }">
                    <span class="font-bold">{{ formatCurrency(data.cashback) }}</span>
                </template>
                <template #editor="{ data, field }">
                    <InputText v-model="data[field]" />
                </template>
            </Column>

            <Column field="cashbackNoCheckIn" header="Кешбэк (незаезд)" style="min-width: 6rem; width: 2%; padding-left: 1rem; padding-right: 1rem;" bodyStyle="text-align:center" alignFrozen="left" :frozen="cashbackNoCheckInFrozen">
                 <template #body="{ data }">
                    <span class="font-bold">{{ formatCurrency(data.cashbackNoCheckIn) }}</span>
                </template>
                <template #editor="{ data, field }">
                    <InputText v-model="data[field]" />
                </template>
            </Column>

            <Column field="guests" header="Человек" filterMatchMode="contains" style="min-width: 120px; padding-left: 1rem; padding-right: 1rem;">
                <template #body="{ data }">
                   <Tag :value=formatPerson(data.guests)  style="background-color: white;"/>
                </template>
            </Column>

            <Column field="user.name" header="ФИО"  style="min-width: 100px; width: 12%; padding-left: 1rem; padding-right: 1rem;" filterMatchMode="startsWith" alignFrozen="left" :frozen="userNameFrozen">
              <!-- <template #filter="{filterModel,filterCallback}">
                    <InputText type="text" v-model="filterModel.value" @keydown.enter="filterCallback()" class="p-column-filter" placeholder="Search"/>
                </template> -->
            </Column>

            <Column field="user.email" header="Email 📧"  style="min-width: 150px; width: 11%; padding-left: 1.5rem;color: rgb(0, 34, 128)" filterMatchMode="startsWith" alignFrozen="left" :frozen="userEmailFrozen">
              <!-- <template #filter="{filterModel,filterCallback}">
                    <InputText type="text" v-model="filterModel.value" @keydown.enter="filterCallback()" class="p-column-filter" placeholder="Search"/>
                </template> -->
            </Column>

            <Column field="user.phone" header="Телефон 📞"  style="min-width: 120px; width: 12%; padding-left: 1.5rem" filterMatchMode="startsWith" alignFrozen="left" :frozen="userPhoneFrozen">
              <!-- <template #filter="{filterModel,filterCallback}">
                    <InputText type="text" v-model="filterModel.value" @keydown.enter="filterCallback()" class="p-column-filter" placeholder="Search"/>
                </template> -->
            </Column>

            <Column field="hotel.city" header="Город" filterField="hotel.city" filterMatchMode="contains" style="min-width: 180px; padding-left: 1rem; padding-right: 1rem" alignFrozen="left" :frozen="hotelCityFrozen">    
              <template #body="{ data }">
                    <div class="flex align-items-center gap-2">
                        <img alt="flag" src="https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png" :class="`flag flag-${data.hotel.code}`" style="width: 24px" />
                        <span>{{ data.hotel.city }}</span>
                    </div>
                </template>
                <!-- <template #filter="{filterModel,filterCallback}">
                    <InputText v-model="filterModel.value" type="text" @input.enter="filterCallback()" class="p-column-filter" placeholder="Поиск"/>
                </template> -->
            </Column>

            <Column field="hotel.name" header="Гостиница" filterMatchMode="contains" style="min-width: 120px; padding-left: 1rem; padding-right: 1rem" alignFrozen="left" :frozen="hotelNameFrozen">
                <!-- <template #filter="{filterModel,filterCallback}">
                    <InputText type="text" v-model="filterModel.value" @keydown.enter="filterCallback()" class="p-column-filter" placeholder="Search"/>
                </template> -->
            </Column>

            <Column field="hotel.stars" header="Звезд" filterMatchMode="contains" style="min-width: 3px; padding-left: 1rem; padding-right: 1rem" bodyStyle="text-align:center">
                 <template #body="{ data }">
                   <Tag :value=formatHotelStars(data.hotel.stars)  style="background-color: white;"/>
                </template>
            </Column>

            <Column field="anotherHotel" header="Другая гостиница" filterMatchMode="contains" style="min-width: 120px; padding-left: 1rem; padding-right: 1rem">
                <template #editor="{ data, field }">
                    <InputText v-model="data[field]" />
                </template>
            </Column>

            <Column field="anotherCity" header="Другой город" filterMatchMode="contains" style="min-width: 120px; padding-left: 1rem; padding-right: 1rem">
                <template #editor="{ data, field }">
                    <InputText v-model="data[field]" />
                </template>
            </Column>

            <Column field="costLimit" header="Лимит на проживание" filterMatchMode="contains" style="min-width: 120px; padding-left: 1rem; padding-right: 1rem"></Column>

            <Column field="regularCustomer" header="Постоянный клиент" filterMatchMode="contains" style="min-width: 120px; padding-left: 1rem; padding-right: 1rem; text-align: center;">
                <template #body="{ data }">
                   <Tag :value=formatRegularCustomer(data.regularCustomer)  style="background-color: white;"/>
                </template>
            </Column>

            <Column field="identificationPerson" header="Тип клиента" filterMatchMode="contains" style="min-width: 120px; padding-left: 1rem; padding-right: 1rem; text-align: left;">
                <template #body="{ data }">
                   <Tag :value=formatIdentificationPerson(data.identificationPerson)  style="background-color: white; color: rgb(0, 0, 0); font-size: 0.9rem;"/>
                   <Tag :value=formatEmojiIdentificationPerson(data.identificationPerson) style="background-color: white; color: rgb(0, 0, 0); font-size: 1.1rem;"/>
                </template>
            </Column>

            <Column field="comment" header="Комментарий" filterMatchMode="contains" style="min-width: 120px; padding-left: 1rem; padding-right: 1rem"></Column>

            <Column field="user.telegramId" header="Id клиента" style="min-width: 2rem; width: 3%; padding-left: 1rem; padding-right: 1rem"></Column>   

            
            
      
           
            <!-- <Column field="representative.name" header="Аватар" filterField="representative.name" sortable>
                <template #body="{ data }">
                    <div class="flex align-items-center gap-2">
                        <img :alt="data.representative.name" :src="`https://primefaces.org/cdn/primevue/images/avatar/${data.representative.image}`" style="width: 32px" />
                        <span>{{ data.representative.name }}</span>
                    </div>
                </template>
                <template #filter="{filterModel,filterCallback}">
                    <InputText type="text" v-model="filterModel.value" @keydown.enter="filterCallback()" class="p-column-filter" placeholder="Search"/>
                </template>
            </Column> -->

            <!-- <ColumnGroup type="footer">
            <Row>
                <Column footer="Всего стоимость:" :colspan="3" footerStyle="text-align:right"/>
                <Column :footer="thisTotalCosts" />
            </Row>
        </ColumnGroup> --> 
            
        </DataTable>
	</div>

      </div>
    </div>
  
</div>



   <!-- <div>
//      <table-lite
//     :is-loading="table.isLoading"
//     :columns="table.columns"
//     :rows="table.rows"
//     :total="table.totalRecordCount"
//     :sortable="table.sortable"
//     :messages="table.messages"
//     @do-search="doSearch"
//     @is-finished="table.isLoading = false"
//   ></table-lite>
// </div> -->



</template>

<script>
import { ref, onMounted } from 'vue';
//import { FilterMatchMode } from 'primevue/api';
import { FilterMatchMode, FilterOperator } from 'primevue/api';
import { defineComponent } from "vue";
import SidebarDark from "../components/SidebarDark.vue";
import { CustomerService } from '../components/demo/components/CustomerService';
import axios from 'axios-https-proxy-fix';
//import axios from 'axios';
  //import ToastService from 'primevue/toastservice';
  //import { inject } from 'vue';


//   const proxy = {
//   host: 'https://localhost',
//   port: 9090,
//   // auth: {
//   //   username: 'some_login',
//   //   password: 'some_pass'
//   // }
// };


// const customers = ref();
// const filters = ref({
//     global: { value: null, matchMode: FilterMatchMode.CONTAINS },
//     name: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
//     'country.name': { value: null, matchMode: FilterMatchMode.STARTS_WITH },
//     representative: { value: null, matchMode: FilterMatchMode.IN },
//     status: { value: null, matchMode: FilterMatchMode.EQUALS },
//     verified: { value: null, matchMode: FilterMatchMode.EQUALS }
// });
// const representatives = ref([
//     { name: 'Amy Elsner', image: 'amyelsner.png' },
//     { name: 'Anna Fali', image: 'annafali.png' },
//     { name: 'Asiya Javayant', image: 'asiyajavayant.png' },
//     { name: 'Bernardo Dominic', image: 'bernardodominic.png' },
//     { name: 'Elwin Sharvill', image: 'elwinsharvill.png' },
//     { name: 'Ioni Bowcher', image: 'ionibowcher.png' },
//     { name: 'Ivan Magalhaes', image: 'ivanmagalhaes.png' },
//     { name: 'Onyama Limba', image: 'onyamalimba.png' },
//     { name: 'Stephen Shaw', image: 'stephenshaw.png' },
//     { name: 'XuXue Feng', image: 'xuxuefeng.png' }
// ]);
// const statuses = ref(['unqualified', 'qualified', 'new', 'negotiation', 'renewal', 'proposal']);
// const loading = ref(true);

// onMounted(() => {
//     CustomerService.getCustomersMedium().then((data) => {
//             customers.value = getCustomers(data);
//             loading.value = false;
//         });
// });

// const getCustomers = (data) => {
//     return [...(data || [])].map((d) => {
//         d.date = new Date(d.date);

//         return d;
//     });
// };
// const formatDate = (value) => {
//     return value.toLocaleDateString('en-US', {
//         day: '2-digit',
//         month: '2-digit',
//         year: 'numeric'
//     });
// };
// const formatCurrency = (value) => {
//     return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' });
// };
// const getSeverity = (status) => {
//     switch (status) {
//         case 'unqualified':
//             return 'danger';

//         case 'qualified':
//             return 'success';

//         case 'new':
//             return 'info';

//         case 'negotiation':
//             return 'warning';

//         case 'renewal':
//             return null;
//     }
// }




  // const PrimeVueToastSymbol = Symbol();
  // const toast = inject(PrimeVueToastSymbol);
  // const onClickButton = () => {
  //        toast.add({severity: 'success', summary: 'PrimeTime', detail: userName.value});
  //     }


  // import { useToast } from 'primevue/usetoast';

  // const toast = useToast();

  //   const showSuccess = () => {
  //        toast.add({severity: 'success', summary: 'PrimeTime', detail: userName.value});
  //     }



  //import { text } from "body-parser";
  //import { StoreExample } from "../store.js";
  // import { useToast} from 'primevue/usetoast';
  // const toast = useToast();

  //const toast = useToast();

  // const onClickButton = () => {
  //    ToastService.useToast().add({severity: 'success', summary: 'PrimeTime', detail: text.value});
  //    //toast.add({severity: 'success', summary: 'PrimeTime', detail: text.value});
  // }



// import TableLite from "../components/TableLite.vue";
//  // import TableLite from "vue3-table-lite";

// // Fake Data for 'asc' sortable
// const sampleData1 = (offst, limit) => {
//   offst = offst + 1;
//   let data = [];
//   for (let i = offst; i <= limit; i++) {
//     data.push({
//       id: i,
//       name: "TEST" + i,
//       email: "test" + i + "@example.com",
//     });
//   }
//   return data;
// };

// // Fake Data for 'desc' sortable
// const sampleData2 = (offst, limit) => {
//   let data = [];
//   for (let i = limit; i > offst; i--) {
//     data.push({
//       id: i,
//       name: "TEST" + i,
//       email: "test" + i + "@example.com",
//     });
//   }
//   return data;
// };


  export default defineComponent({

    data() {
        return {

            editingRows: [],
            statuses: ['new', 'confirmed', 'rejected', 'cancelled'],

            lockedRequisitions: [
                // {
                //     user: {
                //       telegramId: 5135,
                //       name: 'Иван Грозный',
                //       email: 'sfdf@mail.ru', 
                //       phone: '+79218568586',
                //     },
                //     hotel: {
                //         name: 'Лучший',
                //         code: 'somthing',
                //         city: 'Москва',
                //         codeCity: 'MOW',
                //         address: 'ул. Преображения дом 1',
                //         stars: 1,
                //         phone: '+78123546688'
                //     },
                //     status: 'новая',
                //     checkInDaty: '2019-05-05',
                //     checkOutDay: '2019-05-05',
                //     applicationSubmissionTime: '2019-05-05',
                //     applicationStatusDate: '2019-05-05',
                //     cost: 0,
           
                //     // representative: {
                //     //     name: 'Amy Elsner',
                //     //     image: 'amyelsner.png'
                //     // }
                // }
            ],
          
            requisitionNumberFrozen: null,
            applicationSubmissionTimeFrozen: null,
            rejectButtonFrozen: null,
            confirmButtonFrozen: null,
            statusFrozen: null,
            applicationStatusDateFrozen: null,
            checkInDateFrozen: null,
            checkOutDateFrozen: null,
            costFrozen: null,
            userNameFrozen: null,
            userEmailFrozen: null,
            userPhoneFrozen: null,
            hotelCityFrozen: null,
            hotelNameFrozen: null,
            hotelStarsFrozen: null,
            anotherHotelFrozen: null,
            anotherCityFrozen: null,
            editingFrozen: null,
            cashbackFrozen: null,
            cashbackNoCheckInFrozen: null,
            loading: false,
            totalRecords: 0,
            requisitionsTable: null,
            selectedRequisitions: null,
            selectAll: false,
            first: 0,
            filters: null, 
            //{
            //     'user.name': {value: '', matchMode: 'contains'},
            //     'user.email': {value: '', matchMode: FilterMatchMode.CONTAINS},
            //     'user.phone': {value: '', matchMode: 'contains'},
            //     'hotel.city': {value: '', matchMode: 'contains'},
            //     'hotel.name': {value: '', matchMode: 'contains'},
            //     'hotel.stars': {value: '', matchMode: 'contains'},
             //   'checkInDate': { operator: FilterOperator.AND, constraints: [{ value: null, matchMode: FilterMatchMode.DATE_IS }] },
             //    'checkOutDate': { operator: FilterOperator.AND, constraints: [{ value: null, matchMode: FilterMatchMode.DATE_IS }] },
            //     'applicationSubmissionTime': { operator: FilterOperator.AND, constraints: [{ value: null, matchMode: FilterMatchMode.DATE_IS }] },
            //     'applicationStatusDate': { operator: FilterOperator.AND, constraints: [{ value: null, matchMode: FilterMatchMode.DATE_IS }] },
            // },

            lazyParams: {},
            columns: [
                {field: 'user.telegramId', header: 'Id пользователя'},
                {field: 'user.name', header: 'ФИО'},
                {field: 'user.phone', header: 'Телефон 📞'},
                {field: 'user.email', header: 'Email 📧'},
                {field: 'anotherCity', header: 'Другой город'},
                {field: 'anotherHotel', header: 'Другая гостиница'},
                {field: 'hotel.name', header: 'Гостиница'},
                {field: 'hotel.stars', header: 'Звезд'},
                {field: 'hotel.code', header: 'Гостиница (код)'},
                {field: 'hotel.city', header: 'Город'},
                {field: 'hotel.codeCity', header: 'Код \nгорода'},
                {field: 'hotel.address', header: 'Адрес \nгостиницы'},
                {field: 'hotel.phone', header: 'Телефон гостиницы'},
                {field: 'hotel.site', header: 'Сайт гостиницы'},
                {field: 'status',      header:'Статус'},
                {field: 'requisitionNumber', header:'№'},
                {field: 'checkInDate', header:'Заезд'},
                {field: 'checkOutDate', header:'Выезд'},
                {field: 'applicationSubmissionTime', header:'Дата заявки'},
                {field: 'applicationStatusDate', header:'Время статуса 🕓'},
                {field: 'cost', header:'Стоимость'},
                {field: 'cashback', header:'Кешбэк'},
                {field: 'cashbackNoCheckIn', header:'Кешбэк (незаезд)'},
                {field: 'guests', header:'Человек'},
                {field: 'regularCustomer', header:'Постоянный клиент'},
                {field: 'costLimit', header:'Лимит на проживание'},
                {field: 'identificationPerson', header:'Тип клиента'},
                {field: 'comment', header:'Комментарий'},
            ]
        
        };
    },

    created() {
        this.initFilters();
    },

      mounted() {

        this.loading = true;

        this.lazyParams = {
            first: this.$refs.dt.first,
            rows: this.$refs.dt.rows,
            sortField: null,
            sortOrder: null,
            filters: this.filters
        };

       this.loadLazyData();
   
    },

    name: "Requisitions",
    components: { SidebarDark, InputText, Button, ToggleButton, Toast, DataTable, Column, Dropdown, Tag, Calendar},
   
    // setup() {
    // // Table config
    // const table = reactive({
    //   isLoading: false,
    //   columns: [
    //     {
    //       label: "ID",
    //       field: "id",
    //       width: "3%",
    //       sortable: true,
    //       isKey: true,
    //     },
    //     {
    //       label: "Name",
    //       field: "name",
    //       width: "10%",
    //       sortable: true,
    //     },
    //     {
    //       label: "Email",
    //       field: "email",
    //       width: "15%",
    //       sortable: true,
    //     },
    //   ],
    //   rows: [],
    //   totalRecordCount: 0,
    //   sortable: {
    //     order: "id",
    //     sort: "asc",
    //   },
    // });

    // /**
    //  * Search Event
    //  */
    // const doSearch = (offset, limit, order, sort) => {
    //   table.isLoading = true;
    //   setTimeout(() => {
    //     table.isReSearch = offset == undefined ? true : false;
    //     if (offset >= 10 || limit >= 20) {
    //       limit = 20;
    //     }
    //     if (sort == "asc") {
    //       table.rows = sampleData1(offset, limit);
    //     } else {
    //       table.rows = sampleData2(offset, limit);
    //     }
    //     table.totalRecordCount = 20;
    //     table.sortable.order = order;
    //     table.sortable.sort = sort;
    //   }, 600);
    // };

    // // First get data
    // doSearch(0, 10, "id", "asc");

    // return {
    //   table,
    //   doSearch,
    // };
 // },

    methods: {
      onLinkMainPageClick() {
        this.$router.push("/main");
      },
      onLinkCalendarPageClick() {
        this.$router.push("/calendar");
      },
      onLinkRequisitionPageClick() {
        this.$router.push("/");
      },
      onLinkOccupancyPageContainerClick() {
        this.$router.push("/occupancy");
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
      // onClickButton() {
      //    toast.add({severity: 'success', summary: 'Готово', detail: text.value});
      // },
      // fillRequisitionsTable(){
      //   CustomerService.getCustomersMedium().then((data) => (this.customers = data));
      // }

      // formatCurrency(value) {
      //       return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' });
      //   },

      onRowEditSave(event) {

            let { newData, index } = event;

            this.requisitionsTable[index] = newData;

            const param = {requisitionNumber: this.requisitionsTable[index].requisitionNumber,
               status: "",
               anotherHotel: this.requisitionsTable[index].anotherHotel,
               anotherCity: this.requisitionsTable[index].anotherCity,
               cost: this.requisitionsTable[index].cost, 
               cashback: this.requisitionsTable[index].cashback,
               cashbackNoCheckIn: this.requisitionsTable[index].cashbackNoCheckIn};


          axios
            .patch('https://localhost:9090/bookingRequest', null,

             {params: param})

            .then((res) => {

             if (res.status == 200) {
               this.loadLazyData()
             }

           })
              .catch((error) => {
                  // console.log(error.res.data);  
           });

      },

      confirmeBookingRequest(data, index){

        if (this.requisitionsTable[index].status === 'new' || this.requisitionsTable[index].status === 'rejected') {
          
          const param = {requisitionNumber: this.requisitionsTable[index].requisitionNumber, 
                         status: 'confirmed',  
                         anotherHotel: this.requisitionsTable[index].anotherHotel,
                         anotherCity: this.requisitionsTable[index].anotherCity,
                         cost: this.requisitionsTable[index].cost, 
                         cashback: this.requisitionsTable[index].cashback,
                         cashbackNoCheckIn: this.requisitionsTable[index].cashbackNoCheckIn};


                    axios
                    .patch('https://localhost:9090/bookingRequest', null,

                    {params: param})

                    .then((res) => {

                      if (res.status == 200) {
                        //this.requisitionsTable[index].status = 'confirmed'

                        this.loadLazyData()
                      }
          
                  })
                    .catch((error) => {
                    // console.log(error.res.data);
                });
        }
      },

      rejectBookingRequest(data, index){

           if (this.requisitionsTable[index].status === 'new' || this.requisitionsTable[index].status === 'confirmed') {
  
           const param = {requisitionNumber: this.requisitionsTable[index].requisitionNumber,
                          status: 'rejected',      
                          anotherHotel: this.requisitionsTable[index].anotherHotel,
                          anotherCity: this.requisitionsTable[index].anotherCity,
                          cost: this.requisitionsTable[index].cost,
                          cashback: this.requisitionsTable[index].cashback,
                          cashbackNoCheckIn: this.requisitionsTable[index].cashbackNoCheckIn};

            axios
            .patch('https://localhost:9090/bookingRequest', null,

            {params: param})

            .then((res) => {

              if (res.status == 200) {
                this.loadLazyData()
              }
  
          })
            .catch((error) => {
            // console.log(error.res.data);
           });
        }    
      },

      formatDate(value) {

        let date = new Date(value/1000000);

            return date.toLocaleDateString('ru-RU', {
                day: '2-digit',
                month: '2-digit',
                year: 'numeric',
            });
        },

        formatDateTime(value) {     

          const date = new Date(value/1000000);

          //  console.log(typeof value); 
          //  console.log(value);
          //  console.log(typeof date); 
          //  console.log(date instanceof Date); // shows 'true'

    
          return date.toLocaleDateString('ru-RU', {
           day: '2-digit',
           month: '2-digit',
           year: 'numeric',
           hour: 'numeric',
           minute: 'numeric',
           second: 'numeric',
          });
        },


        formatHotelStars(value){
          var s = ""

          if (value === 1) {
            s = "⭐"
          }

          if (value === 2) {
            s = "⭐⭐"
          }

          if (value === 3) {
            s = "⭐⭐⭐"
          }

          if (value === 4) {
            s = "⭐⭐⭐⭐"
          }

          if (value === 5) {
            s = "⭐⭐⭐⭐⭐"
          }
         
          return s
        },

        formatEmojiStatus(value){
          var s = ""

          if (value === "new") {
            s = "🔵"
          }

          if (value === "confirmed") {
            s = "✔️"
          }

          if (value === "rejected") {
            s = "⛔"
          }

          if (value === "cancelled") {
            s = "❌"
          }
         
          return s
        },

        formatStatus(value){
          var s = ""

          if (value === "new") {
            s = "новая"
          }

          if (value === "confirmed") {
            s = "подтверждена"
          }

          if (value === "rejected") {
            s = "отклонена"
          }

          if (value === "cancelled") {
            s = "отменена пользователем"
          }
         
          return s
        },

      
        formatEmojiIdentificationPerson(value) {
         
         var eip = ""

         if (value === "Турист") {
           eip = " 🧳"
         }

         if (value === "В деловой поездке") {
           eip = " 💼"
         }

         if (value === "Представитель компании") {
           eip =  " 👨‍💼🏦"
         }

         return eip
       },

        formatIdentificationPerson(value) {
         
          var ip = ""

          if (value === "Турист") {
            ip = "Турист "
          }

          if (value === "В деловой поездке") {
            ip = "В деловой поездке "
          }

          if (value === "Представитель компании") {
            ip =  "Представитель компании  "
          }

          return ip
        },

        formatRegularCustomer(regularCustomer){

          var rc = ""

          if (regularCustomer === true) {
            rc = "✔️"
          }

          return rc
        },
        
        formatPerson(value){
          var p = ""

          if (value === 1) {
            p = "🙎‍♂️"
          }

          if (value === 2) {
            p = "🙎‍♂️🙎‍♂️"
          }

          if (value === 3) {
            p = "🙎‍♂️🙎‍♂️🙎‍♂️"
          }

          if (value === 4) {
            p = "🙎‍♂️🙎‍♂️🙎‍♂️🙎‍♂️"
          }
         
          return p
        },

        formatCurrency(value) {
            return value.toLocaleString('ru-RU', {style: 'currency', currency: 'RUB'});
        },

       

        getColorPerson(guests) {

           switch (guests) {
            case '🙎‍♂️':
              return 'danger';
            
           default:
              return null;
            }
        },


        getSeverity(status) {
            switch (status) {
                case 'отменена пользователем':
                    return 'danger';
                   
                case 'cancelled':
                    return 'danger';

                case 'подтверждена':
                    return 'success';

                case 'confirmed':
                    return 'success';

                case 'новая':
                    return 'info';

                case 'new':
                    return 'info';

                case 'отклонена':
                    return 'warning';

                case 'rejected':
                    return 'warning';

                case 'renewal':
                    return null;
            }
        },

     

        clearFilter() {
            this.initFilters();
        },

        initFilters() {
            this.filters = {
                global: { value: null, matchMode: FilterMatchMode.CONTAINS },
                'user.name': {value: '', matchMode: 'contains'},
                'user.email': {value: '', matchMode: FilterMatchMode.CONTAINS},
                'user.phone': {value: '', matchMode: 'contains'},
                'hotel.city': {value: '', matchMode: 'contains'},
                'hotel.name': {value: '', matchMode: 'contains'},
                'hotel.stars': {value: '', matchMode: 'contains'},
                status: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.EQUALS }] },
               'checkInDate': { operator: FilterOperator.AND, constraints: [{ value: 0, matchMode: FilterMatchMode.DATE_IS }] },
                'checkOutDate': { operator: FilterOperator.AND, constraints: [{ value: 0, matchMode: FilterMatchMode.DATE_IS }] },
                applicationSubmissionTime: { operator: FilterOperator.AND, constraints: [{ value: null, matchMode: FilterMatchMode.DATE_IS }] },
                applicationStatusDate: { operator: FilterOperator.AND, constraints: [{ value: null, matchMode: FilterMatchMode.DATE_IS }] },
            };
        },

      loadLazyData() {
        
            this.loading = true;
            this.lazyParams = { ...this.lazyParams, first: event?.first || this.first };

            setTimeout(() => {
                     axios
                    .get('https://localhost:9090/getRequisitions')
                    .then((res) => {

                     this.requisitionsTable = res.data;
                     this.totalRecords = res.data.length
                     this.loading = false;
                  })
                    .catch((error) => {
                    //console.log(error.res.data);
                });
             //    CustomerService.getCustomers({ lazyEvent: JSON.stringify(this.lazyParams) }).then((data) => {
            //         this.customers = data.customers;
            //         //this.customers = data;
            //         this.totalRecords = data.totalRecords;
                     this.loading = false;
            //     });
           }, Math.random() * 2 + 3);

        },
        onPage(event) {
            this.lazyParams = event;
            this.loadLazyData(event);
        },
        onSort(event) {
            this.lazyParams = event;
            this.loadLazyData(event);
        },
        onFilter(event) {
            this.lazyParams.filters = this.filters;
            this.loadLazyData(event);
        },
        onSelectAllChange(event) {
            const selectAll = event.checked;

            if (selectAll) {
                     this.selectAll = true;
                     this.selectedRequisitions = data.requisitionsTable
            }
            else {
                this.selectAll = false;
                this.selectedRequisitions = this.requisitionsTable;
            }
        },
        onRowSelect(event) {
           // this.selectAll = this.selectedRequisitions.length === this.totalRecords

           console.log("Индекс" + event.index)

          
  
           this.selectAll = false
         //  this.selectedRequisitions = this.requisitionsTable[event.index]
        },

        onRowUnselect() {
            this.selectAll = false;
            this.selectedRequisitions = []
        },

        toggleLock(data, frozen, index) {
            // if (frozen) {
            //     this.lockedRequisitions = this.lockedRequisitions.filter((c, i) => i !== index);
            //     this.requisitionsTable.push(data);
            // } else {
            //     this.requisitionsTable = this.requisitionsTable.filter((c, i) => i !== index);
            //     this.lockedRequisitions.push(data);
            // }

            // this.requisitionsTable.sort((val1, val2) => {
            //     return val1.id < val2.id ? -1 : 1;
            // });
        },  

      },

     async handleUpdateRequisitions() {

        try {
          const response = await axios.get('https://localhost:9090/getRequisitions', {proxy});

          console.log(response)

          this.requisitionsTable = response.data;
          this.totalRecords = response.data.length;
          this.loading = false;

        } catch (error) {
          
          console.error(error);      
        }

    },

    computed: {

    thisTotalCosts() {
            let total = 0;
            // for(let req of this.requisitionsTable) {
            //     total += req.cost;
            // }

            return this.formatCurrency(total);
        }
      }
    
  });
</script>

<style module>
  .requisitions {
    position: relative;
    /* background-color: var(--color-whitesmoke-300); */
    width: 100%;
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    justify-content: flex-start;
  }


  .toolbarTableStyle{
    height: 4%;
    left: 240px;

     background: conic-gradient(from 145deg at 40%  75%, rgba(148,179,207,0.26) 0%, rgba(156, 184, 216, 0.75) 30%, transparent 50%, rgba(156, 184, 216, 0.75) 80%, rgba(115, 128, 216, 0.51) 100%) 15% 75%/175% 200%,
        radial-gradient(ellipse  at 55%  55%, rgb(178, 166, 222) 0%, rgb(9, 3, 22) 100%) 30% 30%/200% 170%;
   }



</style>

<style scoped>

.wrapperColumns{
  /* position: absolute; */
  /* height: 1080px; */
  display: grid;
  grid-template-columns: 7% 93% ; 

  /* grid-template-columns: 1fr 2fr 1fr */
}

.wrapperRows{
  position: absolute;
  left: 240px;
  width: 90%;
  display: grid;
  grid-template-rows: 7% 93% ; 

  /* grid-template-columns: 1fr 2fr 1fr */
}

 .containerGuestName {
     
      /* position: relative; */
      display: flex;
      /* position:relative;
      top: 91px;
      left: 1127px; */



       /*line-height: 150%; */


      /* align-items: flex-start;
      justify-content: flex-start;
      flex-direction: column;
      min-height: 100vh; */
      /* top: 5%;
      left: 50%; */
     /* transform: translateX(50px) */
    };

    .requisitionDataTableContainer{
      position: absolute;
      left: 240px;
      height: 100%; 
      display: grid;
    }

    .headerTable {
      background:#1f2936; 
      color:white; 
      padding:10px;
    }


    /* .buttons{
      background-color: #6e6ef7;
      color: #f8f8f9; 
    } */
  </style>
