export class Brand {
    id: number;
    name: string;
    logo: string;
    letter: string;
    sort: number;
    status: number;
    updateAt?: string;
    createAt?: string;

    constructor() {
        this.id = 0;
        this.name = '';
        this.logo = '';
        this.letter = '';
        this.sort = 0;
        this.status = 0;
    }
}