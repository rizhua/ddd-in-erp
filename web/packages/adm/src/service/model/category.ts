export class Category {
    id: number;
    name: string;
    parentId: number;
    sort: number;
    children: Category[];
    
    constructor() {
        this.id = 0;
        this.name = '';
        this.parentId = 0;
        this.sort = 0;
        this.children = [];
    }
}