export class Address {
    id: number;
    userId: number;
    tag: string;
    contact: string;
    region: string | string[];
    detail: string;
    tel: string;
    default: boolean;
    orgId: number;
    updateAt?: string;
    createAt?: string;

    constructor() {
        this.id = 0;
        this.userId = 0;
        this.tag = '';
        this.contact = '';
        this.region = '';
        this.detail = '';
        this.tel = '';
        this.default = false;
        this.orgId = 0;
    }
}