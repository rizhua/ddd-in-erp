import { model } from "..";
import { http } from "../http";

export class File {

    static upload(data:{}) : Promise<model.Response> {
        let url = '/file/upload';
        return http.post(url, data);
    }
}