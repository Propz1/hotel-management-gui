<template>
  <div :class="$style.main"><SidebarDark8 /></div>

  <div>
    <div class="card">
        <Toolbar class="mb-4">
            <template #start>
                <Button label="New" icon="pi pi-plus" severity="success" class="mr-2" @click="openNew" />
                <Button label="Delete" icon="pi pi-trash" severity="danger" @click="confirmDeleteSelected" :disabled="!selectedProducts || !selectedProducts.length" />
            </template>

            <template #end>
                <FileUpload mode="basic" accept="image/*" :maxFileSize="1000000" label="Import" chooseLabel="Import" class="mr-2 inline-block" />
                <Button label="Export" icon="pi pi-upload" severity="help" @click="exportCSV($event)"  />
            </template>
        </Toolbar>

        <DataTable 
          :class="$style.hotelsTableStyle"
           ref="dt"
           v-model:selection="selectedHotels" dataKey="id"
          :value="hotelsTable" 
          :filters="filters"
          :paginator="true" :rows="10" 
            paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown" :rowsPerPageOptions="[5,10,25]"
            currentPageReportTemplate="Showing {first} to {last} of {totalRecords} products">
            <template #header>
                <div class="flex flex-wrap gap-2 align-items-center justify-content-between">
                    <h4 class="m-0">Manage Products</h4>
        <IconField iconPosition="left">
                        <InputIcon>
                            <i class="pi pi-search" />
                        </InputIcon>
                        <InputText v-model="filters['global'].value" placeholder="Search..." />
                    </IconField>
      </div>
            </template>

            <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column>
            <Column field="code" header="Code" sortable style="min-width:12rem"></Column>
            <Column field="name" header="Name" sortable style="min-width:16rem"></Column>
            <Column header="Image">
                <template #body="slotProps">
                    <img :src="`https://primefaces.org/cdn/primevue/images/product/${slotProps.data.image}`" :alt="slotProps.data.image" class="border-round" style="width: 64px" />
                </template>
            </Column>
            <Column field="price" header="Price" sortable style="min-width:8rem">
                <template #body="slotProps">
                    {{formatCurrency(slotProps.data.price)}}
                </template>
            </Column>
            <Column field="category" header="Category" sortable style="min-width:10rem"></Column>
            <Column field="rating" header="Reviews" sortable style="min-width:12rem">
                <template #body="slotProps">
                    <Rating :modelValue="slotProps.data.rating" :readonly="true" :cancel="false" />
                </template>
            </Column>
            <Column field="inventoryStatus" header="Status" sortable style="min-width:12rem">
                <template #body="slotProps">
                    <Tag :value="slotProps.data.inventoryStatus" :severity="getStatusLabel(slotProps.data.inventoryStatus)" />
                </template>
            </Column>
            <Column :exportable="false" style="min-width:8rem">
                <template #body="slotProps">
                    <Button icon="pi pi-pencil" outlined rounded class="mr-2" @click="editProduct(slotProps.data)" />
                    <Button icon="pi pi-trash" outlined rounded severity="danger" @click="confirmDeleteProduct(slotProps.data)" />
                </template>
            </Column>
        </DataTable>
    </div>

    <Dialog v-model:visible="productDialog" :style="{width: '450px'}" header="Product Details" :modal="true" class="p-fluid">
        <img v-if="product.image" :src="`https://primefaces.org/cdn/primevue/images/product/${product.image}`" :alt="product.image" class="block m-auto pb-3" />
        <div class="field">
            <label for="name">Name</label>
            <InputText id="name" v-model.trim="product.name" required="true" autofocus :class="{'p-invalid': submitted && !product.name}" />
            <small class="p-error" v-if="submitted && !product.name">Name is required.</small>
        </div>
        <div class="field">
            <label for="description">Description</label>
            <Textarea id="description" v-model="product.description" required="true" rows="3" cols="20" />
        </div>

        <div class="field">
    <label for="inventoryStatus" class="mb-3">Inventory Status</label>
    <Dropdown id="inventoryStatus" v-model="product.inventoryStatus" :options="statuses" optionLabel="label" placeholder="Select a Status">
      <template #value="slotProps">
        <div v-if="slotProps.value && slotProps.value.value">
                        <Tag :value="slotProps.value.value" :severity="getStatusLabel(slotProps.value.label)" />
                    </div>
                    <div v-else-if="slotProps.value && !slotProps.value.value">
                        <Tag :value="slotProps.value" :severity="getStatusLabel(slotProps.value)" />
                    </div>
        <span v-else>
          {{slotProps.placeholder}}
        </span>
      </template>
    </Dropdown>
  </div>

        <div class="field">
            <label class="mb-3">Category</label>
            <div class="formgrid grid">
                <div class="field-radiobutton col-6">
                    <RadioButton id="category1" name="category" value="Accessories" v-model="product.category" />
                    <label for="category1">Accessories</label>
                </div>
                <div class="field-radiobutton col-6">
                    <RadioButton id="category2" name="category" value="Clothing" v-model="product.category" />
                    <label for="category2">Clothing</label>
                </div>
                <div class="field-radiobutton col-6">
                    <RadioButton id="category3" name="category" value="Electronics" v-model="product.category" />
                    <label for="category3">Electronics</label>
                </div>
                <div class="field-radiobutton col-6">
                    <RadioButton id="category4" name="category" value="Fitness" v-model="product.category" />
                    <label for="category4">Fitness</label>
                </div>
            </div>
        </div>

        <div class="formgrid grid">
            <div class="field col">
                <label for="price">Price</label>
                <InputNumber id="price" v-model="product.price" mode="currency" currency="USD" locale="en-US" />
            </div>
            <div class="field col">
                <label for="quantity">Quantity</label>
                <InputNumber id="quantity" v-model="product.quantity" integeronly />
            </div>
        </div>
        <template #footer>
            <Button label="Cancel" icon="pi pi-times" text @click="hideDialog"/>
            <Button label="Save" icon="pi pi-check" text @click="saveProduct" />
        </template>
    </Dialog>

    <Dialog v-model:visible="deleteProductDialog" :style="{width: '450px'}" header="Confirm" :modal="true">
        <div class="confirmation-content">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="product">Are you sure you want to delete <b>{{product.name}}</b>?</span>
        </div>
        <template #footer>
            <Button label="No" icon="pi pi-times" text @click="deleteProductDialog = false"/>
            <Button label="Yes" icon="pi pi-check" text @click="deleteProduct" />
        </template>
    </Dialog>

    <Dialog v-model:visible="deleteProductsDialog" :style="{width: '450px'}" header="Confirm" :modal="true">
        <div class="confirmation-content">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="product">Are you sure you want to delete the selected products?</span>
        </div>
        <template #footer>
            <Button label="No" icon="pi pi-times" text @click="deleteProductsDialog = false"/>
            <Button label="Yes" icon="pi pi-check" text @click="deleteSelectedProducts" />
        </template>
    </Dialog>
</div>



</template>
<script>
  import { defineComponent } from "vue";
  import SidebarDark8 from "../components/SidebarDark8.vue";
  import Button from 'primevue/button';
  import Column from 'primevue/column';
  import DataTable from 'primevue/datatable';
  import InputNumber from 'primevue/inputnumber';
  import Dropdown from 'primevue/dropdown';
  import Dialog from 'primevue/dialog';
  import RadioButton from 'primevue/radiobutton';
  import axios from 'axios-https-proxy-fix';
  import { FilterMatchMode } from 'primevue/api';

  export default defineComponent({

    data() {
        return {
            hotelsTable: null,
            selectedHotels: null,
            productDialog: false,
            deleteProductDialog: false,
            deleteProductsDialog: false,
            product: {},
            
            filters: {},
            submitted: false,

        }
      },


    name: "Main",
    components: 
    { SidebarDark8,
      Button,
      Column,
      DataTable,
      InputNumber,
      Dropdown,
      Dialog,
      RadioButton,
     },




    methods: {
      onLinkCalendarPageClick() {
        this.$router.push("/calendar");
      },
      onLinkRequisitionPageClick() {
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
    },
  });
</script>


<style module>
  .main {
    position: relative;
    background-color: var(--color-whitesmoke-300);
    width: 100%;
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    justify-content: flex-start;
  }

  .hotelsTableStyle{
    display: grid;
    position: absolute;
    top: 68px;
    left: 249px;
    background-color: #40409a;
    width: 89.5%;
    height: 1010px;
  }


</style>
