import { model, http } from "..";

export class Category {
    /**
     * 新增发布分类
     */
    static create(data: model.Category): Promise<model.Response> {
        let url = '/category/create';
        return http.post(url, data);
    }

    /**
     * 删除发布分类
     */
    static delete(data: {id: number[]}): Promise<model.Response> {
        let url = '/category/delete';
        return http.post(url, data);
    }

    /**
     * 更新发布分类
     */
    static update(data: model.Category): Promise<model.Response> {
        let url = '/category/update';
        return http.post(url, data);
    }

    /**
     * 获取发布分类
     */
    static find(data: model.Request): Promise<model.Response> {
        let url = '/category/find';
        return http.post(url, data);
    }

    /**
     * 类目排序
     */
    static setSort(data: {id: number, sort: number}): Promise<model.Response> {
        let url = '/category/setSort';
        return http.post(url, data);
    }

    /**
     * 创建类目属性
     */
    static createAttribute(data: model.Attribute): Promise<model.Response> {
        let url = '/category/attribute/create';
        return http.post(url, data);
    }

    /**
     * 删除类目属性
     */
    static deleteAttribute(data: {id: number[]}): Promise<model.Response> {
        let url = '/category/attribute/delete';
        return http.post(url, data);
    }

    /**
     * 更新类目属性
     */
    static updateAttribute(data: model.Attribute): Promise<model.Response> {
        let url = '/category/attribute/update';
        return http.post(url, data);
    }

    /**
     * 获取类目属性
     */
    static findAttribute(data: model.Request): Promise<model.Response> {
        let url = '/category/attribute/find';
        return http.post(url, data);
    }
}