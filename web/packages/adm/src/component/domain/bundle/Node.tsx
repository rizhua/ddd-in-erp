import { message, Table } from "antd";
import { useEffect, useState } from "react";
import styled from "styled-components";

import { api, model } from "@/service";

const Container = styled.div`
    position: absolute;
    top: 0;
    right: 0;
    min-width: 50vw;
    height: 100vh;
    background-color: #fff;
    z-index: 100;
    box-shadow: -1px 0px 2px #eee;
    overflow: hidden;

    .box-head {
        display: grid;
        grid-template-columns: 50px 1fr;
        align-items: center;

        .title {
            margin: 0;
            padding: 0 1rem;
        }

        .close {
            cursor: pointer;
            display: flex;
            justify-content: center;
        }
    }

    .box-body {
        padding: 1rem;
        height: calc(100vh - 100px);
        box-sizing: border-box;
        overflow-y: auto;
    }

    .box-foot {
        height: 50px;
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 1rem;
        padding: 0 1rem;

        .btn-default, .btn-primary {
            cursor: pointer;
            width: 100px;
            height: 34px;
            border: 1px solid transparent;
            border-radius: 4px;
        }

        .btn-primary {
            color: #fff;
            background-color: var(--primary-color);
        }
    }
`

type NodeProps = {
    bundle: model.Bundle,
    open: boolean,
    onClose: () => void,
}

export function Node(props: NodeProps) {
    const columns = [
        {
            title: '名称',
            dataIndex: 'name',
        },
        {
            title: '类型',
            dataIndex: 'type',
        },
        {
            title: '元数据',
            dataIndex: 'meta',
        },
    ];

    const [nodeModel, setNodeModel] = useState({
        list: new Array<model.Node>(),
        selectedRowKeys: new Array<number>(),
    });

    const getNode = async () => {
        let data = {
            current: 1,
            pageSize: 1000,
        };
        let res = await api.Node.find(data);
        if (res.code == 1000) {
            nodeModel.list = res.data.list;
        } else {
            nodeModel.list = [];
        }
        setNodeModel({ ...nodeModel });
    }

    const getNodeId = async () => {
        let data = {
            id: props.bundle.id,
        };
        let res = await api.Bundle.findNodeId(data);
        if (res.code == 1000) {
            nodeModel.selectedRowKeys = res.data;
        } else {
            nodeModel.selectedRowKeys = [];
        }
        setNodeModel({ ...nodeModel });
    }

    let selection = {
        selectedRowKeys: nodeModel.selectedRowKeys,
        onSelect: (selectedRow: model.Node, selected: boolean) => {
            let set = new Set(nodeModel.selectedRowKeys);
            let doRecord = (data: model.Node) => {
                if (selected) {
                    set.add(data.id);
                } else {
                    set.delete(data.id);
                }
                if (!!selectedRow.path && selected) {
                    let ids = selectedRow.path.split(',');
                    ids.forEach((e: any) => {
                        e = Number(e);
                        set.add(e);
                    });
                }
                if (!!data.children) {
                    data.children.forEach((e: model.Node) => {
                        doRecord(e);
                    });
                }
            }
            doRecord(selectedRow);
            nodeModel.selectedRowKeys = Array.from(set);
            setNodeModel({...nodeModel});
        },
        hideSelectAll: true,
    }

    useEffect(() => {
        if (props.open) {
            getNode();
            getNodeId();
        }
    }, [props.open]);

    const close = () => {
        props.onClose();
    }

    const saveNode = async () => {
        let data = {
            bundleId: props.bundle.id,
            nodeId: nodeModel.selectedRowKeys,
        }
        let res = await api.Bundle.bindNodeId(data);
        if (res.code == 1000) {
            message.success('保存成功');
            props.onClose();
        } else {
            message.error(res.desc);
        }
    }

    return props.open ? <Container>
        <div className="box-head">
            <a className="close iconfont" onClick={() => close()}>&#xe6b3;</a>
        </div>
        <div className="box-body">
            <h1>套餐权限</h1>
            <Table
                rowKey={(record) => record.id}
                columns={columns}
                dataSource={nodeModel.list}
                rowSelection={selection}
                pagination={false}
            />
        </div>
        <div className="box-foot">
            <button className="btn-default" onClick={() => close()}>取消</button>
            <button className="btn-primary" onClick={() => saveNode()}>保存</button>
        </div>
    </Container> : null
}