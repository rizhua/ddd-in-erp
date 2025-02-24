import { useEffect, useState } from "react";
import { Form } from "antd";
import { Select, TextField, Button, Badge, Table } from '@radix-ui/themes';

import { model, api } from "@/service";
import topbar from "topbar";

export function List() {
    const [req, setReq] = useState(new model.Request());
    const [org, setOrg] = useState({
        list: new Array<model.Org>(),
        total: 0,
    });
    const [searchForm] = Form.useForm();

    const getOrg = async () => {
        topbar.show();
        let data: model.Request = {
            ...req,
        };
        if (!!searchForm.getFieldValue('val')) {
            data.queryBy = [{
                ...searchForm.getFieldsValue()
            }];
        }
        let res = await api.Structure.list(data);
        if (res.code == 1000) {
            org.list = res.data.list;
            org.total = res.data.total;
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
            <Form
                form={searchForm}
                onFinish={getOrg}
                initialValues={{ col: 'name' }}
                layout="inline"
            >
                <Form.Item name="col">
                    <Select.Root>
                        <Select.Trigger />
                        <Select.Content>
                            <Select.Item value="name">名称</Select.Item>
                            <Select.Item value="contact">联系人</Select.Item>
                        </Select.Content>
                    </Select.Root>
                </Form.Item>
                <Form.Item name="val">
                    <TextField.Root />
                </Form.Item>
                <Button type="submit">搜索</Button>
            </Form>
        </div>
        <div className="box-body">
            <Table.Root>
                <Table.Header>
                    <Table.Row>
                        <Table.ColumnHeaderCell>组织名称</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>联系人</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>联系电话</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>通讯地址</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>创建时间</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>状态</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>操作</Table.ColumnHeaderCell>
                    </Table.Row>
                </Table.Header>
                <Table.Body>
                    {org.list.map((item: model.Org) => <Table.Row key={item.id}>
                        <Table.Cell>{item.fullName}</Table.Cell>
                        <Table.Cell>{item.contact}</Table.Cell>
                        <Table.Cell>{item.tel}</Table.Cell>
                        <Table.Cell>{item.address}</Table.Cell>
                        <Table.Cell>{item.createAt}</Table.Cell>
                        <Table.Cell>
                            <Badge color={item.status == 1 ? "green" : "red"}>{item.status == 1 ? "启用" : "禁用"}</Badge>
                        </Table.Cell>
                        <Table.Cell>
                            <a className="iconfont" title="免密登录">&#xe61f;</a>
                        </Table.Cell>
                    </Table.Row>)}
                </Table.Body>
            </Table.Root>
        </div>
    </>
}