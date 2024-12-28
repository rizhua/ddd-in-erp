export class Node {
    id: number;
    icon: string;
    name: string;
    meta: string;
    type: number;
    parentId?: number;
    path: string;
    sort: number;
    scheme: string;
    status: number;
    leaf?: boolean;
    level?: number;
    expanded?: boolean;
    createAt?: string;
    updateAt?: string;
    children?: Node[];

    constructor() {
        this.id = 0;
        this.icon = '';
        this.name = '';
        this.meta = '';
        this.path = '';
        this.type = 3;
        this.sort = 255;
        this.scheme = '';
        this.status = 0;
    }
}