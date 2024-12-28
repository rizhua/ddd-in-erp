import { model, http } from "..";

export class Structure {
    /**
     * 获取组织列表
     * @param data 
     * @returns 
     */
    static list(data: model.Request): Promise<model.Response> {
        let url = '/structure/org/find';
        return http.post(url, data);
    }

    /**
     * 切换组织
     * @param data 
     * @returns 
     */
    static switch(data: {id: number}): Promise<model.Response> {
        let url = '/structure/org/switch';
        return http.post(url, data);
    }

    static findNode(data: {roleId: number}): Promise<model.Response> {
        let url = '/structure/node/find';
        return http.post(url, data);
    }

    /**
     * 获取部门列表
     */
    static findDept(data: model.Request): Promise<model.Response> {
        let url = '/structure/dept/find';
        return http.post(url, data);
    }

    /**
     * 新增部门
     */
    static createDept(data: model.Dept): Promise<model.Response> {
        let url = '/structure/dept/create';
        return http.post(url, data);
    }

    /**
     * 删除部门
     */
    static deleteDept(data: {id: number[]}): Promise<model.Response> {
        let url = '/structure/dept/delete';
        return http.post(url, data);
    }

    /**
     * 更新部门
     */
    static updateDept(data: model.Dept): Promise<model.Response> {
        let url = '/structure/dept/update';
        return http.post(url, data);
    }

    /**
     * 添加员工
     */
    static createEmp(data: model.Dept): Promise<model.Response> {
        let url = '/structure/emp/create';
        return http.post(url, data);
    }

    /**
     * 更新员工
     */
    static updateEmp(data: model.Dept): Promise<model.Response> {
        let url = '/structure/emp/update';
        return http.post(url, data);
    }

    /**
     * 获取员工
     */
    static findEmp(data: model.Request): Promise<model.Response> {
        let url = '/structure/emp/find';
        return http.post(url, data);
    }
}