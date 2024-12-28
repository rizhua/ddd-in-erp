export class Product {
    public id: number;
    public code: string;
    public name: string;
    public lowPrice: number;
    public saleCount: number;
    public rateCount: number;
    public barCode: string;
    public media: string;
    public detail: string;
    public createdAt?: string;
    public updatedAt?: string;

    constructor() {
        this.id = 0;
        this.code = '';
        this.name = '';
        this.lowPrice = 0;
        this.saleCount = 0;
        this.rateCount = 0;
        this.barCode = '';
        this.media = '';
        this.detail = '';
    }
}