import { model, http } from "../";

export class User {

    static signIn(data: model.SignIn): Promise<model.Response> {
        let url = '/user/signIn';
        return http.post(url, data);
    }

    static signUp(data: {}): Promise<model.Response> {
        let url = '/user/signUp';
        return http.post(url, data);
    }

    static forget(data: {}): Promise<model.Response> {
        let url = '/user/forget';
        return http.post(url, data);
    }

    static reset(data: {}): Promise<model.Response> {
        let url = '/user/reset';
        return http.post(url, data);
    }

    static parse(): Promise<model.Response> {
        let url = '/user/parse';
        return http.post(url, {});
    }

    /**
     * 退出登录
     * @param data 
     * @returns 
     */
    static logout(): Promise<model.Response> {
        let url = '/user/logout';
        return http.post(url, {});
    }

    /**
     * 用户列表
     * @param data 
     * @returns 
     */
    static find(data: model.Request): Promise<model.Response> {
        let url = '/user/find';
        return http.post(url, data);
    }

    /**
     * 供职组织
     * @param data 
     * @returns 
     */
    static work(): Promise<model.Response> {
        let url = '/user/work';
        return http.post(url, {});
    }
}