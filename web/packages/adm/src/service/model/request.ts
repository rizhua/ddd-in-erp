export type OrderBy = {
    field: string,
    asc: boolean
}

export type QueryBy = {
    field: string,
    value: any
}

export class Request {
    current?: number;
    pageSize?: number;
    orderBy?: OrderBy[];
    queryBy?: QueryBy[];

    constructor(pagination?: { current: number, pageSize: number }) {
        if (!!pagination) {
            this.current = pagination.current;
            this.pageSize = pagination.pageSize;
        }
    }
}
