<script setup>
  //import { ref } from "vue";
  import Toast from 'primevue/toast';
  import Button from 'primevue/button';
  import ToggleButton from 'primevue/togglebutton';
  import InputText from 'primevue/inputtext';
  import DataTable from 'primevue/datatable';
  import Column from 'primevue/column';
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

     <div class="containerGuestName">
        <span class="p-float-label p-input-filled" style="margin-right: 1rem">
       <InputText id="username" type="text" v-model="userName" />
       <label for="username">ФИО</label>
         </span>
        <div class="buttons">
         <Button  class="p-inputtext-sm" label="Записать" icon="pi pi-save" iconPos="right" severity="success" @click="showSuccess"  style="background-color:oldlace" ></Button>
        </div>
     </div>


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


         
           <Toolbar>
             <!-- <Toolbar style="border-radius: 3rem; background-image: linear-gradient(to right, var(--bluegray-500), var(--bluegray-800))"> -->
            <template #start>
                  <ToggleButton v-model="guestsNameFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="ФИО" offLabel="ФИО" style="margin-right: 1rem; border-radius: 1rem" />
                  <ToggleButton v-model="balanceFrozen"    class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Стоимость" offLabel="Стоимость" style="margin-right: 1rem; border-radius: 1rem"/>
             </template>
          </Toolbar>


          <div class="card p-fluid">
          
            <DataTable 
            :value="requisitionsTable" 
            class="mt-4"
            :frozenValue="lockedCustomers"
            lazy paginator :first="first"
            :rows="24"
            scrollable scrollHeight="1080px"
            :virtualScrollerOptions="{ itemSize: 38 }"
            v-model:filters="filters" ref="dt" dataKey="id"
            :totalRecords="totalRecords" :loading="loading" @page="onPage($event)" @sort="onSort($event)" @filter="onFilter($event)" filterDisplay="row"
            :globalFilterFields="['name','country.name', 'company', 'representative.name']"
            v-model:selection="selectedCustomers" :selectAll="selectAll" @select-all-change="onSelectAllChange" @row-select="onRowSelect" @row-unselect="onRowUnselect" tableStyle="min-width: 75rem"
            :pt="{
               table: { style: 'min-width: 50rem' },
               bodyrow: ({ props }) => ({
                class: [{ 'font-bold': props.frozenRow }]
                })
            }">

            <Column selectionMode="multiple" headerStyle="width: 3rem"></Column>
            <Column field="telegramID" header="Id" style="min-width: 100px; width: 2%"></Column>   
            <Column field="name" header="ФИО"  style="min-width: 100px; width: 11%" filterMatchMode="startsWith" sortable alignFrozen="right" :frozen="guestsNameFrozen">
              <template #filter="{filterModel,filterCallback}">
                    <InputText type="text" v-model="filterModel.value" @keydown.enter="filterCallback()" class="p-column-filter" placeholder="Search"/>
                </template>
            </Column>
            <Column field="country.name" header="Город" filterField="country.name" filterMatchMode="contains" sortable style="min-width: 100px; width: 10%">
                <template #body="{ data }">
                    <div class="flex align-items-center gap-2">
                        <img alt="flag" src="https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png" :class="`flag flag-${data.country.code}`" style="width: 24px" />
                        <span>{{ data.country.name }}</span>
                    </div>
                </template>
                <template #filter="{filterModel,filterCallback}">
                    <InputText type="text" v-model="filterModel.value" @keydown.enter="filterCallback()" class="p-column-filter" placeholder="Search"/>
                </template>
            </Column>
            <Column field="company" header="Гостиница" filterMatchMode="contains" sortable>
                <template #filter="{filterModel,filterCallback}">
                    <InputText type="text" v-model="filterModel.value" @keydown.enter="filterCallback()" class="p-column-filter" placeholder="Search"/>
                </template>
            </Column>
            <Column field="representative.name" header="Аватар" filterField="representative.name" sortable>
                <template #body="{ data }">
                    <div class="flex align-items-center gap-2">
                        <img :alt="data.representative.name" :src="`https://primefaces.org/cdn/primevue/images/avatar/${data.representative.image}`" style="width: 32px" />
                        <span>{{ data.representative.name }}</span>
                    </div>
                </template>
                <template #filter="{filterModel,filterCallback}">
                    <InputText type="text" v-model="filterModel.value" @keydown.enter="filterCallback()" class="p-column-filter" placeholder="Search"/>
                </template>
            </Column>
            <Column field="status" header="Status" style="min-width: 200px"></Column>
            <Column field="activity" header="Activity" style="min-width: 200px"></Column>
            <Column field="date" header="Date" style="min-width: 200px"></Column>
            <Column field="balance" header="Стоимость" style="min-width: 200px" alignFrozen="right" :frozen="balanceFrozen">
                 <!-- <template #body="{ data }">
                    <span class="font-bold">{{ formatCurrency(data.balance) }}</span>
                </template> -->
            </Column>
            <Column style="flex: 0 0 4rem">
                <template #body="{ data, frozenRow, index }">
                  <Button type="button" style="color:green" :icon="frozenRow ? 'pi pi-lock-open' : 'pi pi-lock'" :disabled="frozenRow ? false : lockedCustomers.length >= 2" text size="small" @click="toggleLock(data, frozenRow, index)" />
                </template>
            </Column>
        </DataTable>
	</div>

      </div>
    </div>
  
</div>



  <!-- <div>
     <table-lite
    :is-loading="table.isLoading"
    :columns="table.columns"
    :rows="table.rows"
    :total="table.totalRecordCount"
    :sortable="table.sortable"
    :messages="table.messages"
    @do-search="doSearch"
    @is-finished="table.isLoading = false"
  ></table-lite>
</div> -->



</template>

<script>
import { ref, onMounted } from 'vue';
import { FilterMatchMode } from 'primevue/api';
import { defineComponent } from "vue";
import SidebarDark from "../components/SidebarDark.vue";
import { CustomerService } from '../components/demo/components/CustomerService';
  //import ToastService from 'primevue/toastservice';
  //import { inject } from 'vue';


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

            lockedCustomers: [
                {
                    id: 5135,
                    name: 'Geraldine Bisset',
                    country: {
                        name: 'France',
                        code: 'fr'
                    },
                    company: 'Bisset Group',
                    status: 'proposal',
                    date: '2019-05-05',
                    activity: 0,
                    representative: {
                        name: 'Amy Elsner',
                        image: 'amyelsner.png'
                    }
                }
            ],
          
            balanceFrozen: false,
            guestsNameFrozen: false,
            loading: false,
            totalRecords: 0,
            requisitionsTable: null,
            selectedCustomers: null,
            selectAll: false,
            first: 0,
            filters: {
                'name': {value: '', matchMode: 'contains'},
                'country.name': {value: '', matchMode: 'contains'},
                'company': {value: '', matchMode: 'contains'},
                'representative.name': {value: '', matchMode: 'contains'},
            },
            lazyParams: {},
            columns: [
                {field: 'name', header: 'Name'},
                {field: 'country.name', header: 'Country'},
                {field: 'company', header: 'Company'},
                {field: 'representative.name', header: 'Representative'}
            ]
        
        };
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

        // CustomerService.getCustomersMedium().then((data) => {
        //     this.customers = data;
        // });
    },

    name: "Requisitions",
    components: { SidebarDark, InputText, Button, ToggleButton, Toast, DataTable, Column},
   
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

      loadLazyData() {
            this.loading = true;
            this.lazyParams = { ...this.lazyParams, first: event?.first || this.first };

            setTimeout(() => {
                CustomerService.getCustomers({ lazyEvent: JSON.stringify(this.lazyParams) }).then((data) => {
                    this.customers = data.customers;
                    //this.customers = data;
                    this.totalRecords = data.totalRecords;
                    this.loading = false;
                });
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
                CustomerService.getCustomers().then(data => {
                    this.selectAll = true;
                    this.selectedCustomers = data.customers;
                });
            }
            else {
                this.selectAll = false;
                this.selectedCustomers = [];
            }
        },
        onRowSelect() {
            this.selectAll = this.selectedCustomers.length === this.totalRecords
        },
        onRowUnselect() {
            this.selectAll = false;
        },

        toggleLock(data, frozen, index) {
            if (frozen) {
                this.lockedCustomers = this.lockedCustomers.filter((c, i) => i !== index);
                this.requisitionsTable.push(data);
            } else {
                this.requisitionsTable = this.requisitionsTable.filter((c, i) => i !== index);
                this.lockedCustomers.push(data);
            }

            this.requisitionsTable.sort((val1, val2) => {
                return val1.id < val2.id ? -1 : 1;
            });
        },  

      },

      async handleUpdateRequisitions() {

        try {
          const response = await axios.get('http://localhost:9090/users');
        
          this.requisitionsTable = response.data;
          //this.totalRecords = data.totalRecords;
          this.loading = false;

        } catch (error) {
          
          console.error(error);      
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
</style>

<style scoped>

.wrapperColumns{
  display: grid;
  grid-template-columns: 10% 90% ; 

  /* grid-template-columns: 1fr 2fr 1fr */
}

.wrapperRows{
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
      display: grid;
    }
    .buttons{
      /* background-color: #6e6ef7;
      color: #f8f8f9; */
    }
  </style>
