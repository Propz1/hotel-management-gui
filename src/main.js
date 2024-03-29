import { createApp } from "vue";
import { createRouter, createWebHistory } from "vue-router";
// import VueTableLite from "vue3-table-lite";
// import TableLite from './components/TableLite.vue';
// import Vue3EasyDataTable from 'vue3-easy-data-table';
// import 'vue3-easy-data-table/dist/style.css';
//import TableLite from "vue3-table-lite";
import PrimeVue from 'primevue/config';
import InputText from 'primevue/inputtext';
import InputNumber from 'primevue/inputnumber';
import Toast from 'primevue/toast';
import ToastService from 'primevue/toastservice';
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';
import ToggleButton from 'primevue/togglebutton';
import Calendar from 'primevue/calendar';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import ColumnGroup from 'primevue/columngroup';  
import Row from 'primevue/row';                   // optional
import Paginator from "primevue/paginator"
import Toolbar from 'primevue/toolbar';
import Dropdown from 'primevue/dropdown';
import RadioButton from 'primevue/radiobutton';
import Tag from 'primevue/tag';
import MultiSelect from 'primevue/multiselect';

import 'primevue/resources/themes/lara-light-indigo/theme.css'
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';
import App from "./App.vue";
import "vuetify/styles";
import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";

import Requisitions from "./pages/Requisitions.vue";
import Help from "./pages/Help.vue";
import Settings from "./pages/Settings.vue";
import Integrations from "./pages/Integrations.vue";
import Reports from "./pages/Reports.vue";
import Chat from "./pages/Chat.vue";
import Occupancy from "./pages/Occupancy.vue";
import CalendarPage from "./pages/CalendarPage.vue";
import Main from "./pages/Main.vue";
import "./global.css";


const routes = [
  {
    path: "/",
    name: "Requisitions",
    component: Requisitions,
  },
  {
    path: "/help",
    name: "Help",
    component: Help,
  },
  {
    path: "/settings",
    name: "Settings",
    component: Settings,
  },
  {
    path: "/integrations",
    name: "Integrations",
    component: Integrations,
  },
  {
    path: "/reports",
    name: "Reports",
    component: Reports,
  },
  {
    path: "/chat",
    name: "Chat",
    component: Chat,
  },
  {
    path: "/occupancy",
    name: "Occupancy",
    component: Occupancy,
  },
  {
    path: "/calendar",
    name: "CalendarPage",
    component: CalendarPage,
  },
  {
    path: "/main",
    name: "Main",
    component: Main,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((toRoute, fromRoute, next) => {
  const documentTitle =
    toRoute?.meta && toRoute?.meta?.title ? toRoute?.meta?.title : "Ormand";
  window.document.title = documentTitle;
  if (toRoute?.meta?.description) {
    addMetaTag(toRoute?.meta?.description);
  }
  next();
});

const addMetaTag = (value) => {
  let element = document.querySelector(`meta[name='description']`);

  if (element) {
    element.setAttribute("content", value);
  } else {
    element = `<meta name="description" content="${value}" />`;
    document.head.insertAdjacentHTML("beforeend", element);
  }
};


const vuetify = createVuetify({ components, directives });

const app = createApp(App)

app.use(router);
app.use(vuetify);
app.use(PrimeVue);
app.use(ToastService);


//app.component('ToastService', ToastService)
app.component('InputText', InputText)
app.component('Button', Button);
app.component('ToggleButton', ToggleButton);
app.component('Calendar', Calendar);
app.component('DataTable', DataTable);
app.component('Column', Column);
app.component('Row', Row);
app.component('ColumnGroup', ColumnGroup);
app.component('Toast', Toast);
app.component('Paginator', Paginator);
app.component('Toolbar',Toolbar);
app.component('Dropdown', Dropdown);
app.component('Dialog', Dialog);
app.component('RadioButton', RadioButton);
app.component('Tag', Tag);
app.component('MultiSelect', MultiSelect);
app.component('InputNumber', InputNumber);

app.mount("#app");






export default router;





