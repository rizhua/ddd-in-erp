export enum EmpStatus {
    '待入职' = 0,
    '试用期',
    '已转正',
    '已离职',
}

export class Emp {
    id: number;
    userId: number;
    name: string;
    number: string;
    position: string;
    grade: string;
    mobile: string;
    tel: string;
    address: string;
    email: string;
    deptId: number;
    joinTime?: string;
    quitTime?: string;
    status: EmpStatus;
    orgId: number;

    constructor() {
        this.id = 0;
        this.userId = 0;
        this.name = '';
        this.number = '';
        this.position = '';
        this.grade = '';
        this.mobile = '';
        this.tel = '';
        this.address = '';
        this.email = '';
        this.deptId = 0;
        this.status = 0;
        this.orgId = 0;
    }
}