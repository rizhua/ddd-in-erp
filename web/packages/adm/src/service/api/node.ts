import { model, http } from "../";

export class Node {
    /**
     * 创建节点
     * @param data
     */
    static create(data: model.Node): Promise<model.Response> {
        let url = '/node/create';
        return http.post(url, data);
    }

    /**
     * 删除节点
     * @param data
     */
    static delete(data: { id: number[] }): Promise<model.Response> {
        let url = '/node/delete';
        return http.post(url, data);
    }

    /**
     * 更新节点
     * @param data
     */
    static update(data: model.Node): Promise<model.Response> {
        let url = '/node/update';
        return http.post(url, data);
    }

    static findNodeId(data: {roleId: number}): Promise<model.Response> {
        let url = '/role/findNodeId';
        return http.post(url, data);
    }

    /**
     * 查询节点
     * @param data
     */
    static find(data: model.Request): Promise<model.Response> {
        let url = '/node/find';
        return http.post(url, data);
    }

    /**
     * 节点排序
     * @param data
     */
    static setSort(data: { id: number, sort: number }): Promise<model.Response> {
        let url = '/node/setSort';
        return http.post(url, data);
    }

    /**
     * 修改状态
     * @param data
     */
    static setStatus(data: { id: number, status: number }): Promise<model.Response> {
        let url = '/node/setStatus';
        return http.post(url, data);
    }

    /**
     * 权限
     * @param data
     */
    static permission(data: {meta: string}): Promise<model.Response> {
        let url = '/node/permission';
        return http.post(url, data);
    }
}