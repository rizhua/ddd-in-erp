import { model, http } from "../";

export class System {
    static sendSms(data: {mobile: string}): Promise<model.Response> {
        let url = '/system/sms/send';
        return http.post(url, data);
    }

    static createNotice(data: model.Notice): Promise<model.Response> {
        let url = '/system/notice/create';
        return http.post(url, data);
    }

    static updateNotice(data: model.Notice): Promise<model.Response> {
        let url = '/system/notice/update';
        return http.post(url, data);
    }

    static findNotice(data: model.Request): Promise<model.Response> {
        let url = '/system/notice/list';
        return http.post(url, data);
    }
}