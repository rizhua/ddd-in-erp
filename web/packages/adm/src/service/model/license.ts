export class License {
    id: number;
    code: string;
    bizId: number;
    createAt?: string;

    constructor() {
        this.id = 0;
        this.code = '';
        this.bizId = 0;
    }
}