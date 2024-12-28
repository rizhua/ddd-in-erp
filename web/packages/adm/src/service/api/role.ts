import { model, http } from "../";

export class Role {
    /**
     * 新增角色
     * @param data 
     * @returns 
     */
    static create(data: model.Role): Promise<model.Response> {
        let url = '/role/create';
        return http.post(url, data);
    }

    /**
     * 删除角色
     * @param data 
     * @returns 
     */
    static delete(data: {id: number[]}): Promise<model.Response> {
        let url = '/role/delete';
        return http.post(url, data);
    }

    /**
     * 更新角色
     * @param data 
     * @returns 
     */
    static update(data: model.Role): Promise<model.Response> {
        let url = '/role/update';
        return http.post(url, data);
    }

    /**
     * 角色列表
     * @param data 
     * @returns 
     */
    static find(data: model.Request & {tree?: boolean}): Promise<model.Response> {
        let url = '/role/find';
        return http.post(url, data);
    }

    /**
     * 添加成员
     * @param params 
     * @returns 
     */
    static addUser(data: {roleId: number, userId: number[]}): Promise<model.Response> {
        let url = '/role/addUser';
        return http.post(url, data);
    }

    /**
     * 移除成员
     * @param params 
     * @returns 
     */
    static removeUser(data: {roleId: number, userId: number[]}): Promise<model.Response> {
        let url = '/role/removeUser';
        return http.post(url, data);
    }

    /**
     * 成员列表
     * @param data 
     * @returns 
     */
    static findUser(data: {roleID: number} & model.Request): Promise<model.Response> {
        let url = '/role/findUser';
        return http.post(url, data);
    }

    /**
     * 绑定、解绑节点
     * @param params 
     * @returns 
     */
    static bindNodeId(data: {roleId: number, nodeId: number[]}): Promise<model.Response> {
        let url = '/role/bindNodeId';
        return http.post(url, data);
    }

    /**
     * 某个角色的已绑定节点
     * @param params 
     * @returns 
     */
    static findNodeId(data: {roleId: number}): Promise<model.Response> {
        let url = '/role/findNodeId';
        return http.post(url, data);
    }
}