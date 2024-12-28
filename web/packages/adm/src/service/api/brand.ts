import { model, http } from "../";

export class Brand {
    /**
     * 新增品牌
     * @param data 
     * @returns 
     */
    static create(data: model.Role): Promise<model.Response> {
        let url = '/brand/create';
        return http.post(url, data);
    }

    /**
     * 删除品牌
     * @param data 
     * @returns 
     */
    static delete(data: {id: number[]}): Promise<model.Response> {
        let url = '/brand/delete';
        return http.post(url, data);
    }

    /**
     * 更新品牌
     * @param data 
     * @returns 
     */
    static update(data: model.Role): Promise<model.Response> {
        let url = '/brand/update';
        return http.post(url, data);
    }

    /**
     * 品牌列表
     * @param data 
     * @returns 
     */
    static find(data: model.Request): Promise<model.Response> {
        let url = '/brand/find';
        return http.post(url, data);
    }
}