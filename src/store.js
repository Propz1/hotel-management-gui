import { useToast } from "primevue/usetoast";
        
export const StoreExample = defineStore("StoreExample", () => {
  const toast = useToast();

  const onActionGetProducts = async () => {
     return API_EXAMPLE.onApiGetProducts().then(() => {

       toast.add({
          severity: "success",
          summary: "Success Message",
          detail: "Order submitted",
          life: 3000,
       });
     });
  };

 return {
   onActionGetProducts,
 };
});