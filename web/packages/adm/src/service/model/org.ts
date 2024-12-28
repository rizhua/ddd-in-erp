export class Org {
    id: number;
    icon: string;
    code: string;
    name: string;
    fullName: string;
    industry: string;
    capacity: number;
    contact: string;
    tel: string;
    address: string;
    ownerId: number;
    license: string;
    status: number;
    updateAt?: string;
    createAt?: string;

    constructor() {
        this.id = 0;
        this.icon = '';
        this.code = '';
        this.name = '';
        this.fullName = '';
        this.industry = '';
        this.capacity = 0;
        this.contact = '';
        this.tel = '';
        this.address = '';
        this.ownerId = 0;
        this.license = '';
        this.status = 0;
    }
}
