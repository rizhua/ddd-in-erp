export class Role {
    id: number;
    name: string;
    orgId: number;
    parentId: number;
    createAt?: string;
    updateAt?: string;
    children?: Role[];
    level: number;

    constructor() {
        this.id = 0;
        this.name = '';
        this.orgId = 0;
        this.parentId = 0;
        this.level = 1;
    }
}
