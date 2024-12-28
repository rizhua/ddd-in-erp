export class Config {
    id: number;
    code: string;
    data: string;
    remark?: string;
    createAt?: string;
    updateAt?: string;

    constructor() {
        this.id = 0;
        this.code = '';
        this.data = '';
    }
}