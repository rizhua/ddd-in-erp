import { model, http } from "..";

export class Product {
    /**
     * 新增商品
     */
    static create(data: model.Product): Promise<model.Response> {
        let url = '/product/create';
        return http.post(url, data);
    }

    /**
     * 更新商品
     */
    static update(data: model.Product): Promise<model.Response> {
        let url = '/product/update';
        return http.post(url, data);
    }

    /**
     * 商品列表
     */
    static find(data: model.Request): Promise<model.Response> {
        let url = '/product/find';
        return http.post(url, data);
    }

    /**
     * 新增属性
     */
    static createAttribute(data: model.Attribute): Promise<model.Response> {
        let url = '/product/attribute/create';
        return http.post(url, data);
    }

    /**
     * 删除属性
     */
    static deleteAttribute(data: {id: number[]}): Promise<model.Response> {
        let url = '/product/attribute/delete';
        return http.post(url, data);
    }

    /**
     * 更新属性
     */
    static updateAttribute(data: model.Attribute): Promise<model.Response> {
        let url = '/product/attribute/update';
        return http.post(url, data);
    }

    /**
     * 属性列表
     */
    static findAttribute(data: model.Request): Promise<model.Response> {
        let url = '/product/attribute/find';
        return http.post(url, data);
    }

    /**
     * 新增发布分类
     */
    static createCategory(data: model.Category): Promise<model.Response> {
        let url = '/product/category/create';
        return http.post(url, data);
    }

    /**
     * 删除发布分类
     */
    static deleteCategory(data: {id: number[]}): Promise<model.Response> {
        let url = '/product/category/delete';
        return http.post(url, data);
    }

    /**
     * 更新发布分类
     */
    static updateCategory(data: model.Category): Promise<model.Response> {
        let url = '/product/category/update';
        return http.post(url, data);
    }

    /**
     * 获取发布分类
     */
    static findCategory(data: model.Request): Promise<model.Response> {
        let url = '/product/category/find';
        return http.post(url, data);
    }

    /**
     * 类目排序
     */
    static setCategorySort(data: {id: number, sort: number}): Promise<model.Response> {
        let url = '/product/category/setSort';
        return http.post(url, data);
    }
}