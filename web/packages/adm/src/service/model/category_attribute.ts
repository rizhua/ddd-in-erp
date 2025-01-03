export class CategoryAttribute {
    public id: number;
    public label: string;
    public value: string[];
    public type: string;
    public categoryId: number;
    public required: boolean;
    
    constructor() {
        this.id = 0;
        this.label = "";
        this.value = [];
        this.type = 'SELECT';
        this.categoryId = 0;
        this.required = false;
    }
}
