<!-- <DataTable :value="cars">
    <Column field="vin" header="Vin"></Column>
    <Column field="year" header="Year"></Column>
    <Column field="brand" header="Brand"></Column>
    <Column field="color" header="Color"></Column>
</DataTable> 
Column components can be dynamically generated using a v-for as well.

<DataTable :value="cars">
    <Column v-for="col of columns" :field="col.field" :header="col.header" :key="col.field"></Column>
</DataTable> -->

<DataTable :value="cars">
    <template #header>
        <div>
            <Button icon="pi pi-refresh" style="float: left"/>
            List of Cars
        </div>
    </template>
    <Column field="vin" header="Vin"></Column>
    <Column field="year" header="Year"></Column>
    <Column field="brand" header="Brand">
        <template #body="slotProps">
            <img :src="'demo/images/car/' + slotProps.data.brand + '.png'" :alt="slotProps.data.brand"  width="48px"/>
        </template>
    </Column>
    <Column field="color" header="Color"></Column>
    <Column headerStyle="width: 8em" bodyStyle="text-align: center">
        <template #header>
            <Button type="button" icon="pi pi-cog"></Button>
        </template>
        <template #body="slotProps">
            <Button type="button" icon="pi pi-search" class="p-button-success" style="margin-right: .5em"></Button>
            <Button type="button" icon="pi pi-pencil" class="p-button-warning"></Button>
        </template>
    </Column>
    <template #footer>
        In total there are {{cars ? cars.length : 0 }} cars.
    </template>
</DataTable>


