import { model, http } from "../";

export class Bundle {
    static create(data: model.Bundle): Promise<model.Response> {
        let url = '/bundle/create';
        return http.post(url, data);
    }

    static delete(data: {id: number[]}): Promise<model.Response> {
        let url = '/bundle/delete';
        return http.post(url, data);
    }

    static update(data: model.Bundle): Promise<model.Response> {
        let url = '/bundle/update';
        return http.post(url, data);
    }

    static find(data: model.Request): Promise<model.Response> {
        let url = '/bundle/find';
        return http.post(url, data);
    }

    static bindNodeId(data: {bundleId: number, nodeId: number[]}): Promise<model.Response> {
        let url = '/bundle/bindNodeId';
        return http.post(url, data);
    }

    static findNodeId(data: {id: number}): Promise<model.Response> {
        let url = '/bundle/findNodeId';
        return http.post(url, data);
    }

    static findLicense(data: model.Request): Promise<model.Response> {
        let url = '/bundle/license/find';
        return http.post(url, data);
    }
}