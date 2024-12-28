export class Dept {
    id: number;
    name: string;
    parentId: number;
    mgrId: number;
    mgrName: string;
    orgId: number;
    deleted?: string;

    constructor() {
        this.id = 0;
        this.name = '';
        this.parentId = 0;
        this.mgrId = 0;
        this.mgrName = '';
        this.orgId = 0;
    }
}