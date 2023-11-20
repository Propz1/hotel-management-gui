import CarService from '../../service/CarService';

export default {
    data() {
        return {
            columns: null,
            cars: null
        }
    },
    carService: null,
    created() {
        this.carService = new CarService();

        this.columns = [
            {field: 'vin', header: 'Vin'},
            {field: 'year', header: 'Year'},
            {field: 'brand', header: 'Brand'},
            {field: 'color', header: 'Color'}
        ];
    },
    mounted() {
        this.carService.getCarsSmall().then(data => this.cars = data);
    }
}