import { useEffect, useState } from "react";
import { Button, Form, Input, Select, Space, Table, TablePaginationConfig } from "antd";

import { model, api } from "@/service";
import styled from "styled-components";
import { Link } from "react-router-dom";


const Container = styled.div`
    margin: 1rem;
    padding: 1rem;
    min-height: calc(100vh - 84px);
`;

const columns = [{
    title: '标题',
    dataIndex: 'title',
}, {
    title: '内容',
    dataIndex: 'content',
}, {
    title: '公布范围',
    dataIndex: 'scope',
    render: (v: number) => {
        let m = new Map([
            [0, '草稿'],
            [1, '对内'],
            [2, '对外'],
            [3, '不限']
        ]);
        return m.has(v) ? m.get(v) : '-';
    }
}, {
    title: '类型',
    dataIndex: 'type',
    render: (v: number) => {
        let m = new Map([
            [0, '消息'],
            [1, '公文'],
            [2, '公告'],
        ]);
        return m.has(v) ? m.get(v) : '-';
    }
}, {
    title: '拟稿人',
    dataIndex: 'drafter',
}, {
    title: '附件',
    dataIndex: 'attached',
}, {
    title: '操作',
    dataIndex: 'action',
    width: '80px',
    render: (v: null, r: model.Notice) => {
        return <Space size={16}>
            <Link className="iconfont" to={`/news/edit/${r.id}`}>&#xe640;</Link>
            <a className="iconfont">&#xe618;</a>
        </Space>
    }
}];

export function List() {
    const [req, setReq] = useState(new model.Request());
    const [noticeModel, setNoticeModel] = useState({
        info: new model.Notice(),
        list: new Array<model.Notice>(),
        total: 0,

    });
    const [loading, setLoading] = useState(false);

    const getNotice = async () => {
        setLoading(true);
        let data = {
            ...req,
        };
        let res = await api.System.findNotice(data);
        if (res.code == 1000) {
            noticeModel.list = res.data.list;
            noticeModel.total = res.data.total;
        } else {
            noticeModel.list = [];
            noticeModel.total = 0;
        }
        setNoticeModel({ ...noticeModel });
        setLoading(false);
    }

    useEffect(() => {
        getNotice();
    }, []);

    const onTable = (pagination: TablePaginationConfig) => {
        setReq({
            ...pagination,
        });
    }

    return <Container className="box">
        <div className="box-head">
            <Form layout="inline">
                <Form.Item>
                    <Select options={[{ label: '标题', value: 'title' }]} value="title" />
                </Form.Item>
                <Form.Item>
                    <Input />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" onClick={getNotice}>搜索</Button>
                </Form.Item>
            </Form>
            <Button type="primary"><Link to="/news/edit">新增消息</Link></Button>
        </div>
        <div className="box-body">
            <Table
                rowKey={(record) => record.id}
                columns={columns}
                dataSource={noticeModel.list}
                pagination={{ current: req.current, pageSize: req.pageSize, total: noticeModel.total }}
                loading={loading}
                onChange={onTable}
            />
        </div>
    </Container>
}