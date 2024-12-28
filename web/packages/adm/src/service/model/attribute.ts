export class Attribute {
    public id: number;
    public label: string;
    public value: string[];
    public multi: boolean;
    public required: boolean;

    constructor() {
        this.id = 0;
        this.label = "";
        this.value = [];
        this.multi = false;
        this.required = false;
    }
}