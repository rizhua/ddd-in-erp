export class Bundle {
    id: number;
    name: string;
    term: number;
    quota: number;
    price: number;
    updateAt?: string;
    createAt?: string;

    constructor() {
        this.id = 0;
        this.name = '';
        this.term = 0;
        this.quota = 0;
        this.price = 0;
    }
}