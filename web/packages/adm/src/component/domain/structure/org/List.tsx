import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { Button, Badge } from '@radix-ui/themes';

import { model, api } from "@/service";
import topbar from "topbar";
import { Pagination } from "@/component/common";

export function List() {
    const [req, setReq] = useState(new model.Request({ current: 1, pageSize: 1 }));
    const [org, setOrg] = useState({
        list: new Array<model.Org>(),
        total: 0,
    });
    const queryForm = useForm();

    const getOrg = async () => {
        topbar.show();
        let data: model.Request = {
            ...req,
        };
        let tmp = queryForm.getValues();
        if (!!tmp.value) {
            data.queryBy = [{ field: tmp.field, value: tmp.value }];
        }
        let res = await api.Structure.list(data);
        if (res.code == 1000) {
            org.list = res.data.list || [];
            org.total = 50;//res.data.total;
        } else {
            org.list = [];
            org.total = 0;
        }
        setOrg({ ...org });
        topbar.hide();
    }

    useEffect(() => {
        getOrg();
    }, [req.current]);

    return <>
        <div className="box-head">
            <h1>组织列表</h1>
            <form className="form-inline" onSubmit={queryForm.handleSubmit(getOrg)}>
                <div className="form-item">
                    <select {...queryForm.register('field')}>
                        <option value="name">名称</option>
                        <option value="contact">联系人</option>
                    </select>
                </div>
                <div className="form-item">
                    <input {...queryForm.register('value')} type="text" />
                </div>
                <Button type="submit">搜索</Button>
            </form>
        </div>
        <div className="box-body">
            <table className="table">
                <thead>
                    <tr>
                        <th>组织名称</th>
                        <th>联系人</th>
                        <th>联系电话</th>
                        <th>通讯地址</th>
                        <th>创建时间</th>
                        <th>状态</th>
                        <th>操作</th>
                    </tr>
                </thead>
                <tbody>
                    {org.list.map((item: model.Org) => <tr key={item.id}>
                        <td>{item.fullName}</td>
                        <td>{item.contact}</td>
                        <td>{item.tel}</td>
                        <td>{item.address}</td>
                        <td>{item.createAt}</td>
                        <td>
                            <Badge color={item.status == 1 ? "green" : "red"}>{item.status == 1 ? "启用" : "禁用"}</Badge>
                        </td>
                        <td>
                            <a className="iconfont" title="免密登录">&#xe61f;</a>
                        </td>
                    </tr>)}
                </tbody>
            </table>
            <Pagination current={req.current} pageSize={req.pageSize} total={org.total} onChange={(current: number, pageSize: number) => setReq({ ...req, current, pageSize })} />
        </div>
    </>
}