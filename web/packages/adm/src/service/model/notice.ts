export class Notice {
    id: number;
    title: string;
    content: string;
    scope: number;
    orgId?: number;
    attached?: string;
    type: number;
    drafter: string;
    updateAt?: string;
    createAt?: string;

    constructor() {
        this.id = 0;
        this.title = '';
        this.content = '';
        this.scope = 0;
        this.type = 0;
        this.drafter = '';
    }
}