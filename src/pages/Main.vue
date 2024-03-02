<template>
  <div :class="$style.main"><SidebarDark8 /></div>

  <div>
    <div class="card">
        <Toolbar class="mb-4" :class="$style.toolbarTableStyle">
            <template #start>
                <Button label="Добавить" icon="pi pi-plus" severity="success" class="mr-2" style="min-width: 8rem; padding-left:0.7rem; padding-right:0.5rem; margin-left:0.4rem; margin-right: 0.4rem; height:35px; background-color: var(--indigo-600); color: var(--primary-color-text);" @click="openNew" />
                <!-- <Button label="Удалить" icon="pi pi-trash" severity="danger" style="min-width: 3.5rem; padding-left:1rem; padding-right:0.5rem; margin-left:0.4rem; margin-right: 0.4rem; height:35px; background-color: var(--red-700); color: var(--primary-color-text);" @click="confirmDeleteSelected" :disabled="!selectedHotels || !selectedHotels.length" /> -->
                <Button label="Экспорт" icon="pi pi-upload" severity="help" style="min-width: 8rem; padding-left:0.7rem; padding-right:0.5rem; margin-left:0.4rem; margin-right: 0.4rem; height:35px; background-color: var(--teal-800); color: var(--primary-color-text);" @click="exportCSV($event)"  />
              </template>

            <template #center>
                <!-- <FileUpload mode="basic" accept="image/*" :maxFileSize="1000000" label="Import" chooseLabel="Import" class="mr-2 inline-block" /> -->
               
            </template>
        </Toolbar>

        <Toolbar :class="$style.toolbar2TableStyle" style="padding-left: 0.1rem; min-height: 2rem;">     
                   <template #start>
                     <div :class="$style.buttonsToggle">

                       <!-- <ToggleButton v-model="selectionFrozen" class="p-inputtext-sm" onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="⃣" offLabel="⃣" style="min-width: 3rem; margin-left: 0.2rem; margin-right: 1rem; border-radius: 0.5rem"/>     -->
                       <ToggleButton v-model="editingFrozen" class="p-inputtext-sm" onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="🖍 | 🗑️" offLabel=" 🖍 | 🗑️ " style="min-width: 5rem; margin-left: 0.4rem; margin-right: 1rem; border-radius: 0.5rem"/>  
                       <ToggleButton v-model="statusFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Статус" offLabel="Статус" style="min-width: 5rem; margin-right: 1rem; border-radius: 0.5rem"/>
                       <ToggleButton v-model="codeFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="код" offLabel="код" style="min-width: 3.5rem; margin-left: 0.1rem; margin-right: 0.1rem; border-radius: 0.5rem" />
                       <ToggleButton v-model="nameFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Название" offLabel="Название" style="min-width: 4rem; margin-left: 0.5rem; margin-right: 0.5rem; border-radius: 0.5rem" />
                       <ToggleButton v-model="nameTelegramFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Название (телеграм)" offLabel="Название (телеграм)" style="min-width: 3.5rem; margin-left: 0.5rem; margin-right: 1rem; border-radius: 0.5rem" />
                       <ToggleButton v-model="starsFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Звезд" offLabel="Звезд" style="min-width: 5rem; margin-right: 0.5rem; border-radius: 0.5rem" />
                       <ToggleButton v-model="cityFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Город" offLabel="Город" style="min-width: 5rem; margin-left: 0.5rem; margin-right: 1rem; border-radius: 0.5rem" />
                       <ToggleButton v-model="cityTelegramFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Город (телеграм)" offLabel="Город (телеграм)" style="min-width: 10.2rem; margin-right: 1rem; border-radius: 0.5rem" />
                       <ToggleButton v-model="codeCityFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Код города" offLabel="Код города" style="min-width: 5.5rem; margin-right: 1rem; border-radius: 0.5rem" />
                       <ToggleButton v-model="orderCityFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Порядок городв" offLabel="Порядок городов" style="min-width: 6rem; margin-right: 1rem; border-radius: 0.5rem" />
                       <ToggleButton v-model="phoneFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="📞" offLabel="📞" style="min-width: 3rem; margin-right: 1rem; border-radius: 0.5rem"/>
                       <ToggleButton v-model="siteFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="🌐" offLabel="🌐" style="min-width: 3rem; margin-right: 1rem; border-radius: 0.5rem"/>
                       <ToggleButton v-model="photoCatalogFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Каталог фото" offLabel="Каталог фото" style="min-width: 5rem; margin-right: 1rem; border-radius: 0.5rem" />
                       <ToggleButton v-model="descriptionFrozen" class="p-inputtext-sm"  onIcon="pi pi-lock" offIcon="pi pi-lock-open" onLabel="Описание" offLabel="Описание" style="min-width: 5.4rem; margin-right: 1rem; border-radius: 0.5rem" />
           
                      </div>
                    </template>
                  </Toolbar> 

        <DataTable 
          :class="$style.hotelsTableStyle"
          resizableColumns columnResizeMode="expand" 
          :value="hotelsTable" 
          :frozenValue="lockedProps"
           ref="dt"
          v-model:filters="filters" filterDisplay="row" dataK="id" 
          :globalFilterFields="['status', 'stars']"  
          v-model:selection="selectedHotels" dataKey="id"  
          :loading="loading"  
          :rows="40" 
          scrollable scrollHeight="970px"
          :virtualScrollerOptions="{ itemSize: 36 }"
          > 

          <!-- paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown" :rowsPerPageOptions="[5,10,25]"
            currentPageReportTemplate="Showing {first} to {last} of {totalRecords} products"
          :paginator="true"  -->

            <!-- <template #header>
                <div class="flex flex-wrap gap-2 align-items-center justify-content-between">
                    <h4 class="m-0">Manage Products</h4>
                   <IconField iconPosition="left">
                        <InputIcon>
                            <i class="pi pi-search" />
                        </InputIcon>
                        <InputText v-model="filters['global'].value" placeholder="Search..." />
                    </IconField> 
                 </div>
            </template> -->

            <!-- <Column selectionMode="multiple" style="width: 3rem; padding-left: 1.5rem; padding-right: 2rem;" :exportable="false" :frozen="selectionFrozen"></Column> -->
            <Column :exportable="false" style="min-width:8rem" :frozen="editingFrozen">
                <template #body="slotProps">
                    <Button icon="pi pi-pencil" outlined rounded class="mr-2" @click="editHotel(slotProps.data)" />
                    <Button icon="pi pi-trash" outlined rounded severity="danger" @click="confirmDeleteHotel(slotProps.data)" />
                </template>
            </Column>
            <Column field="status" header="Статус" :showFilterMenu="false" :filterMenuStyle="{ width: '7rem' }" style="min-width: 7rem; padding-left: 0.3rem; padding-right: 3rem; text-align: left" :frozen="statusFrozen">
              <template #body="{ data }">
                  <Tag :value=formatStatus(data.status) :severity="getSeverity(formatStatus(data.status))" style="color: white; font-size: 1rem; font-weight:400; padding-left: 0.7rem; padding-right: 0.7rem;"/>
              </template>
              <template #filter="{ filterModel, filterCallback }">
                    <Dropdown v-model="filterModel.value" @change="filterCallback()" :options="statuses" :showClear="false" placeholder="" class="p-column-filter" style="min-width: 7rem">
                        <template #option="slotProps">
                            <Tag :value=formatStatus(slotProps.option) :severity="getSeverity(formatStatus(slotProps.option))" style="min-width: 7rem; "/>
                        </template>
                    </Dropdown>
              </template>
            </Column>
            <Column field="code" header="Код" style="min-width: 11.5rem; width: 10%; padding-left: 0.1rem; padding-right: 0.1rem; text-align: left" :frozen="codeFrozen"></Column>
            <Column field="name" header="Название" style="min-width: 16rem; width: 10%; padding-left: 0.1rem; padding-right: 0.1rem; text-align: left" :frozen="nameFrozen"></Column>
            <Column field="nameTelegram" header="Название (в телеграм)" style="min-width: 18rem; width: 10%; padding-left: 0.1rem; padding-right: 0.1rem; text-align: left" :frozen="nameTelegramFrozen"></Column>
            <Column field="stars" header="Звезд" :showFilterMenu="false" :filterMenuStyle="{ width: '4rem' }" style="min-width:4rem; padding-left: 0.1rem; padding-right: 0.1rem; text-align: center" :frozen="starsFrozen">
              <template #body="{ data }">
                  <Tag :value=formatStars(data.stars) :severity="getSeverity(formatStars(data.status))" style="background-color: white; font-size: 1rem; font-weight:400; padding-left: 0.7rem; padding-right: 0.7rem;"/>
              </template>
              <template #filter="{ filterModel, filterCallback }">
                    <Dropdown v-model="filterModel.value" @change="filterCallback()" :options="starshotels" :showClear="false" placeholder="" class="p-column-filter" style="min-width: 4rem; ">
                        <template #option="slotProps">
                            <Tag :value="slotProps.option" style="min-width: 6rem; color: white;"/>
                        </template>
                    </Dropdown>
              </template>
            </Column>
            <Column field="city" header="Город" style="min-width: 16rem; padding-left: 1rem; padding-right: 1rem; text-align: left" :frozen="cityFrozen"></Column>
            <Column field="cityTelegram" header="Название города (в телеграм)" style="min-width: 16rem; padding-left: 1rem; padding-right: 1rem; text-align: left" :frozen="cityTelegramFrozen"></Column>
            <Column field="codeCity" header="Код города" style="min-width:10rem; padding-left: 1rem; padding-right: 1rem" :frozen="codeCityFrozen"></Column>
            <Column field="orderCity" header="Порядок городов" style="min-width:3rem; padding-left: 1rem; padding-right: 1rem; text-align: center" :frozen="orderCityFrozen"></Column>
            <Column field="phone" header="Телефон 📞" style="min-width: 16rem; padding-left: 1.5rem; padding-right: 1.5rem" :frozen="phoneFrozen"></Column>
            <Column field="site" header="Сайт 🌐" style="min-width:16rem; padding-left: 1.5rem; padding-right: 1.5rem" :frozen="siteFrozen"></Column>
            <Column field="photoCatalog" header="Каталог для фото" style="min-width:16rem" :frozen="photoCatalogFrozen"></Column>
            <Column field="description" header="Описание" style="min-width:16rem" :frozen="descriptionFrozen"></Column>
            <Column field="updateDate" header="Дата изменения" style="min-width: 11.5rem; width: 10%; padding-left: 0.9rem; padding-right: 0.9rem; text-align: left" :frozen="codeFrozen">
              <template #body="{ data }">
                    {{ formatDateTime(data.updateDate) }}
                </template>
            </Column>
            <Column field="id" header="id" style="min-width: 11.5rem; width: 10%; padding-left: 0.9rem; padding-right: 0.9rem; text-align: left" :frozen="codeFrozen"></Column>


         
        </DataTable>
    </div>

    <Dialog v-model:visible="hotelDialog" :style="{width: '800px' }" header="Создание / редактирование гостиницы" :modal="true" class="p-fluid">
        
        <div class="field col">
                <label for="code" style="margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:1rem; font-weight: 600; font-size: large; color: var(--blue-900)">Код гостиницы</label>
                <InputText id="code" v-model="hotel.code" required="true" :class="{'p-invalid': submitted && !hotel.code}" style="padding:0.2rem; margin-left:0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; background-color: var(--surface-50); width: 780px;"/>
                <small class="p-error" v-if="submitted && !hotel.code">Заполните код гостиницы.</small>
        </div>
          
        <div class="field">
            <label for="name" style="margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:1rem; font-weight: 600; font-size: large; color: var(--blue-900)">Название</label>
            <InputText id="name" v-model.trim="hotel.name" required="true" autofocus :class="{'p-invalid': submitted && !hotel.name}" style="padding:0.2rem; margin-left:0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; background-color: var(--surface-50); width: 780px;"/>
            <small class="p-error" v-if="submitted && !hotel.name">Заполните название гостиницы.</small>
        </div>

        <div class="formgrid grid">
           
            <div class="field col">
                <label for="nameTelegram" style="margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 600; font-size: large; color: var(--blue-900)">Название гостиницы (для Телеграм)</label>
                <InputText id="nameTelegram" v-model="hotel.nameTelegram" required="true"  :class="{'p-invalid': submitted && !hotel.nameTelegram}" style="padding:0.2rem; margin-left:0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; background-color: var(--surface-50); width: 780px;"/>
                <small class="p-error" v-if="submitted && !hotel.nameTelegram">Заполните название гостиницы для телеграма.</small>
              </div>
            <div class="field col">
                <label for="city" style="margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 600; font-size: large;color: var(--blue-900)">Город</label>
                <InputText id="city" v-model="hotel.city" required="true"  :class="{'p-invalid': submitted && !hotel.city}" style="padding:0.2rem; margin-left:0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; background-color: var(--surface-50); width: 780px;"/>
                <small class="p-error" v-if="submitted && !hotel.city">Укажите город.</small>
              </div> 
           
            <div class="field col">
                <label for="cityTelegram" style="margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 600; font-size: large;color: var(--blue-900)">Название города (для Телеграм)</label>
                <InputText id="cityTelegram" v-model="hotel.cityTelegram" required="true"  :class="{'p-invalid': submitted && !hotel.cityTelegram}" style="padding:0.2rem; margin-left:0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; background-color: var(--surface-50); width: 780px;"/>
                <small class="p-error" v-if="submitted && !hotel.cityTelegram">Укажите город для Телеграм.</small>
              </div>
            <div class="field col">
                <label for="codeCity" style="margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 600; font-size: large;color: var(--blue-900)">Код города</label>
                <InputText id="codeCity" v-model="hotel.codeCity" style="padding:0.2rem; margin-left:0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; background-color: var(--surface-50); width: 780px;"/>
              </div>
            <div class="field col">
                <label for="orderCity" style="margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 600; font-size: large;color: var(--blue-900)">Порядковый номер города (для Телеграм)</label>
                <InputNumber id="orderCity" v-model="hotel.orderCity" required="true" :class="{'p-invalid': submitted && !hotel.orderCity}" style="padding:0.2rem; margin-left:0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; background-color: var(--surface-50); width: 780px;"/>
                <small class="p-error" v-if="submitted && !hotel.orderCity">Укажите порядковый номер города (небходим для телеграм).</small>
              </div>
            <div class="field col">
                <label for="stars" style="margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 600; font-size: large;color: var(--blue-900)">Звезд</label>
                <InputNumber id="stars" v-model="hotel.stars" required="true" :class="{'p-invalid': submitted && !hotel.stars}" style="padding:0.2rem; margin-left:0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; background-color: var(--surface-50); width: 780px;"/>
                <small class="p-error" v-if="submitted && !hotel.stars">Укажите количество звезд.</small>
              </div>
            <div class="field col">
                <label for="address" style="margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 600; font-size: large;color: var(--blue-900)">Адрес</label>
                <InputText id="address" v-model="hotel.address" required="true" :class="{'p-invalid': submitted && !hotel.address}" style="padding:0.2rem; margin-left:0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; background-color: var(--surface-50); width: 780px;"/>
                <small class="p-error" v-if="submitted && !hotel.address">Укажите адрес.</small>
              </div>
            <div class="field col">
                <label for="phone" style="margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 600; font-size: large;color: var(--blue-900)">Телефон</label>
                <InputText id="phone" v-model="hotel.phone" style="padding:0.2rem; margin-left:0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; background-color: var(--surface-50); width: 780px;"/>
              </div>
            <div class="field col">
                <label for="site" style="margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 600; font-size: large;color: var(--blue-900)">Сайт</label>
                <InputText id="site" v-model="hotel.site" style="padding:0.2rem; margin-left:0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; background-color: var(--surface-50); width: 780px;"/>
            </div>
            <div class="field col">
                <label for="photoCatalog" style="margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 600; font-size: large;color: var(--blue-900)">Каталог фотографий</label>
                <InputText id="photoCatalog" v-model="hotel.photoCatalog" required="true" :class="{'p-invalid': submitted && !hotel.photoCatalog}"  style="padding:0.2rem; margin-left:0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; background-color: var(--surface-50); width: 780px;"/>
                <small class="p-error" v-if="submitted && !hotel.photoCatalog">Укажите каталог, где будут храниться фотографии гостиницы.</small>
              </div>
        </div>

        <div class="field">
            <label for="description" style="margin-left: 0.5rem; margin-right:0.5rem; padding-top:0.1rem; padding-bottom:0.1rem; margin-bottom:0.5rem; font-weight: 600; font-size: large; color: var(--blue-900)">Описание гостиницы</label>
            <Textarea id="description" v-model="hotel.description" required="true" :class="{'p-invalid': submitted && !hotel.description}" rows="5" cols="40" style="padding-top:0.1rem; margin-left: 0.5rem; margin-right:0.5rem;padding-top:0.5rem;margin-bottom:0.5rem; width: 780px;"/>
            <small class="p-error" v-if="submitted && !hotel.photoCatalog">Заполните описание отеля.(для оформления, можно использовать теги HTML)</small>
          </div>

        <div class="field">
            <label class="mb-3" style="margin-left:0.5rem; margin-right:0.5rem; margin-top:0.5rem; margin-bottom:0.5rem; font-weight: 600; font-size: large; color: var(--blue-900)">Статус</label>
            <div class="formgrid grid">
                <div class="field-radiobutton col-6" style="padding:0.2rem; margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 400; font-size: large;">
                    <RadioButton id="status1" name="статус" value="valid" v-model="hotel.status" />
                    <label for="status1" style="padding:0.2rem; margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 500; font-size: large; color: var(--blue-900)">работает</label>
                </div>
                <div class="field-radiobutton col-6" style="padding:0.2rem; margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 400; font-size: large;">
                    <RadioButton id="status2" name="статус" value="not valid" v-model="hotel.status" />
                    <label for="status2" style="padding:0.2rem; margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 500; font-size: large; color: var(--blue-900)">отключен</label>
                </div>
            </div>
            <small class="p-error" v-if="submitted && !hotel.status">Установите статус гостиницы</small>
        </div>

        <template #footer style="margin-left: 0.5rem; margin-right:0.5rem;margin-top:0.5rem;margin-bottom:0.5rem; font-weight: 600; font-size: large;">
            <Button label="Отмена" icon="pi pi-times" text style="min-width: 7rem; padding:0.5rem; margin:0.4rem; height:35px; background-color: var(--red-600); color: var(--primary-color-text);" @click="hideDialog"/>
            <Button label="Сохранить" icon="pi pi-check" text style="min-width: 8rem; padding:0.5rem; margin:0.4rem; height:35px; background-color: var(--teal-700); color: var(--primary-color-text);" @click="saveHotel" />
        </template>
    </Dialog>

    <Dialog v-model:visible="deleteHotelDialog" :style="{width: '450px'}" header="Подтвердить удаление" :modal="true">
        <div class="confirmation-content">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="hotel">Вы уверены, что хотите удалить <b>{{hotel.name}}</b>?</span>
        </div>
        <template #footer>
            <Button label="Нет" icon="pi pi-times" text style="min-width: 5rem; padding:0.5rem; margin:0.4rem; height:35px; background-color: var(--teal-700); color: var(--primary-color-text);" @click="deleteHotelDialog = false"/>
            <Button label="Да" icon="pi pi-check" text style="min-width: 5rem; padding:0.5rem; margin:0.4rem; height:35px; background-color: var(--red-600); color: var(--primary-color-text);" @click="deleteHotel" />
        </template>
    </Dialog>

    <Dialog v-model:visible="deleteHotelsDialog" :style="{width: '450px'}" header="Подтвердить" :modal="true">
        <div class="confirmation-content">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="hotel">Вы уверены, что хотите удалить выбранные гостиницы?</span>
        </div>
        <template #footer>
            <Button label="Нет" icon="pi pi-times" text style="min-width: 5rem; padding:0.5rem; margin:0.4rem; height:35px; background-color: var(--red-600); color: var(--primary-color-text);"  @click="deleteHotelsDialog = false"/>
            <Button label="Да" icon="pi pi-check" text @click="deleteSelectedHotels" />
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
  import InputText from 'primevue/inputtext';
  import Textarea from 'primevue/textarea';
  import Dropdown from 'primevue/dropdown';
  import Dialog from 'primevue/dialog';
  import RadioButton from 'primevue/radiobutton';
  import Row from 'primevue/row';                  
  import { FilterMatchMode } from 'primevue/api';
  import axios from 'axios-https-proxy-fix';


  export default defineComponent({

    data() {
        return {
            hotelsTable: null,
            selectedHotels: null,
            hotelDialog: false,
            deleteHotelDialog: false,
            deleteHotelsDialog: false,
            hotel: {},
            starshotels: [1,2,3,4,5],
            statuses: ['valid', 'not valid'],
            filters: {
                status: { value: null, matchMode: FilterMatchMode.EQUALS },
                stars:  { value: null, matchMode: FilterMatchMode.EQUALS },
            },

            loading: false,

            selectionFrozen: null,    
            editingFrozen: null,   
            codeFrozen: null,
            nameFrozen: null,
            nameTelegramFrozen: null,
            starsFrozen: null,
            cityFrozen: null,
            cityTelegramFrozen: null,
            codeCityFrozen: null,
            orderCityFrozen: null,
            phoneFrozen: null,
            siteFrozen: null,
            statusFrozen: null,
            photoCatalogFrozen: null,
            descriptionFrozen: null,
          
            lockedProps: [],
            submitted: false,
        }
      },

      mounted() {
       this.loading = false;
       this.loadLazyData();
      },

    name: "Main",
    components: 
    { SidebarDark8,
      Button,
      Column,
      DataTable,
      InputNumber,
      InputText,
      Textarea,
      Dropdown,
      Dialog,
      RadioButton,
      Row,
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

      loadLazyData() {
        
        this.loading = true;

         axios
          .get('https://localhost:9090/getHotelsTable')
          .then((res) => {

           this.hotelsTable = res.data;
           this.totalRecords = res.data.length
           this.loading = false;
          })
          .catch((error) => {
                //console.log(error.res.data);      
          });
         
          this.loading = false;
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

      formatStars(value){
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

      formatStatus(value){
          var s = ""

          if (value === "valid") {
            s = "работает"
          }

          if (value === "not valid") {
            s = "отключен"
          }
         
          return s
      },

       getSeverity(status) {
            switch (status) {
              case 'работает':
                    return 'success';

              case 'valid':
                    return 'success';
                   
              case 'отключен':
                    return 'warning';

              case 'not valid':
                    return 'warning';
             }
        },


      formatCurrency(value) {
            if(value)
				return value.toLocaleString('en-US', {style: 'currency', currency: 'USD'});
			return;
        },

        openNew() {
            this.hotel = {};
            this.submitted = false;
            this.hotelDialog = true;
        },
        hideDialog() {
            this.hotelDialog = false;
            this.submitted = false;
        },
        saveHotel() {
            this.submitted = true;

           var n = this.hotel.name

			      if (this.hotel.name.trim()) {
                if (this.hotel.id) {
                     if (this.hotel.id.length == 0){
                       this.hotel.id = this.createId();
                     }
                    //this.hotel.status = this.hotel.status.value ? this.hotel.status.value: this.hotel.status;
                    this.hotelsTable[this.findIndexById(this.hotel.id)] = this.hotel;
                }else {
                    this.hotel.id = this.createId();
                    //this.hotel.status = this.hotel.status ? this.hotel.status.value : 'valid';
                    this.hotelsTable.push(this.hotel);
                }


                axios
                  .put('https://localhost:9090/updateHotel',

                    {
                     id:               this.hotel.id,
	                   code:             this.hotel.code,
                     name:             this.hotel.name,
	                   nameTelegram:     this.hotel.nameTelegram,
	                   city:             this.hotel.city,
                     cityTelegram:     this.hotel.cityTelegram,
	                   codeCity:         this.hotel.codeCity,
                     orderCity:        this.hotel.orderCity,
	                   stars:            this.hotel.stars,
                     status:           this.hotel.status,
	                   address:          this.hotel.address,
                     phone:            this.hotel.phone,
                     site:             this.hotel.site,
                     photoCatalog:     this.hotel.photoCatalog,
                     description:      this.hotel.description,
                     updateDate:       0,  
                    })

                  .then((res) => {

                    if (res.status == 200) {
                       this.$toast.add({severity:'success', summary: '👍', detail: 'Запись в базу данных гостиницы ' + n + ' выполнено успешно!', life: 3000});
                    };

                    this.loadLazyData()

                  })

                  .catch((error) => {

                     if (error.response.status == 400){ 
                        this.$toast.add({severity:'warn', summary: '', detail: 'Запись в базу данных гостиницы ' + n + ' не выполнена!' + ' Ошибка: ' + error.response.data, life: 8500});
                     }else if (error.response.status == 404){ 
                        this.$toast.add({severity:'warn', summary: '', detail: 'Запись в базу данных гостиницы ' + n + ' не выполнена!' + ' Ошибка: ' + error.response.data, life: 8500});
                     }else if (error.response.status == 500){ 
                        this.$toast.add({severity:'warn', summary: '', detail: 'Запись в базу данных гостиницы ' + n + ' не выполнена!' + ' Нет подлкючения к базе данных.', life: 4500});
                     };

                  });

                 this.hotelDialog = false;
                 this.hotel = {};
            }
        },




        editHotel(hotel) {
            this.hotel = {...hotel};
            this.hotelDialog = true;
        },

        confirmDeleteHotel(hotel) {
            this.hotel = hotel;
            this.deleteHotelDialog = true;
        },

        deleteHotel() {

          let name = this.hotel.name
          let code = this.hotel.code

          const param = {code: this.hotel.code};
        
            axios
             .delete('https://localhost:9090/deleteHotel',

              {params: param})

             .then((res) => {

                if (res.status == 200) {
                   this.$toast.add({severity:'success', summary: '', detail: 'Гостиница ' + name + ' удалена!', life: 3000});
                };

                this.loadLazyData()

              })

             .catch((error) => {

                  if (error.response.status == 404){ 
                     this.$toast.add({severity:'warn', summary: '', detail: 'Гостиница ' + name + '(код  ' + code + ') не удалена (не найдена)!', life: 5500});
                  }else if (error.response.status == 500){ 
                     this.$toast.add({severity:'warn', summary: '', detail: 'Гостиница ' + name + '(код ' + code + ') не удалена. Нет подключения к базе данных!', life: 5500});
                  }else {
                     this.$toast.add({severity:'warn', summary: '', detail: 'Гостиница ' + name + '(код ' + code + ') не удалена!', life: 5500});
                  };

              });


              this.deleteHotelDialog = false;
              this.hotel = {};
                   
        },




        findIndexById(id) {
            let index = -1;
            for (let i = 0; i < this.hotelsTable.length; i++) {
                if (this.hotelsTable[i].id === id) {
                    index = i;
                    break;
                }
            }

            return index;
        },
        createId() {
            let id = '';
            var chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
            for ( var i = 0; i < 5; i++ ) {
                id += chars.charAt(Math.floor(Math.random() * chars.length));
            }
            return id;
        },
        exportCSV() {
            this.$refs.dt.exportCSV();
        },
        confirmDeleteSelected() {
            this.deleteHotelsDialog = true;
        },
        deleteSelectedHotels() {
            this.hotelsTable = this.hotelsTable.filter(val => !this.selectedHotels.includes(val));
            this.deleteHotelsDialog = false;
            this.selectedHotels = null;
            this.$toast.add({severity:'success', summary: 'Successful', detail: 'Удалены выбранные гостиницы', life: 3000});
        },

        getStatusLabel(status) {
            switch (status) {
                case 'valid':
                    return 'success';

                case 'not valid':
                    return 'warning';

                default:
                    return null;
            }
        }
    },

  });

</script>


<style module>
  .main {
    position: relative;
    background-color: var(--color-whitesmoke-300);
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    justify-content: flex-start;
  }

  .hotelsTableStyle{
    display: grid;
    position: absolute;
    top: 109px;
    left: 249px; 
    width: 89.5%;
    height: 970px;
  }

  .toolbarTableStyle{
    position: absolute;
    top: 10px;
    left: 249px;
    width: 89.5%;
    height: 3rem;
    background: conic-gradient(from 145deg at 40%  75%, rgba(148,179,207,0.26) 0%, rgba(156, 184, 216, 0.75) 30%, transparent 50%, rgba(156, 184, 216, 0.75) 80%, rgba(115, 128, 216, 0.51) 100%) 15% 75%/175% 200%,
      radial-gradient(ellipse  at 55%  55%, rgb(178, 166, 222) 0%, rgb(9, 3, 22) 100%) 30% 30%/200% 170%;  

   }

   .toolbar2TableStyle{
    position: absolute;
    top:68px;
    left: 249px;
    width: 89.5%;
    height: 1rem;
    background: radial-gradient(ellipse  at 55%  55%, rgb(91, 171, 236) 0%, rgb(3, 14, 14) 100%) 30% 30%/5000% 50%;  
    background: #f1eded;
    
   }

   .buttonsToggle{
    position: absolute;
    left: 0px;
    width: 89.5%;
   }


</style>
