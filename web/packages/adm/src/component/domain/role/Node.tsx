import { useEffect, useState } from "react";
import { Button, Form, Input, message, Space, Select, Table } from "antd";

import './role.less';
import { model, api } from "@/service";
import styled from "styled-components";

const Contianer = styled.div`
    
    .box-foot {
        padding: 16px 0;
    }
`

interface NodeProps {
    role: model.Role;
    parent: model.Role;
}

export function Node(props: NodeProps) {
    const columns = [
        {
            title: '名称',
            dataIndex: 'name',
        }, {
            title: '元数据',
            dataIndex: 'meta',
        }, {
            title: '类型',
            dataIndex: 'type',
        }
    ];
    const [nodeModel, setNodeModel] = useState({
        list: new Array<model.Node>(),
        selectedRowKeys: new Array<number>(),
    });

    const getNodeId = async () => {
        let data = {
            roleId: props.role.id,
        }
        let res = await api.Role.findNodeId(data);
        if (res.code == 1000) {
            nodeModel.selectedRowKeys = res.data;
        } else {
            nodeModel.selectedRowKeys = [];
        }
        setNodeModel({ ...nodeModel });
    }

    const getNode = async () => {
        let data: model.Request & { roleId: number } = {
            roleId: props.role.id,
        };
        let res = await api.Structure.findNode(data);
        if (res.code == 1000) {
            nodeModel.list = res.data;
        } else {
            nodeModel.list = [];
        }
        setNodeModel({...nodeModel});
    }

    useEffect(() => {
        if (!props.role.id) {
            return;
        }
        getNodeId();
        getNode();
    }, [props]);

    const selection = {
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

    // 绑定|解绑
    const saveNode = async () => {
        if (!props.role.id) {
            message.error('请先选择角色');
            return;
        }
        let data = {
            nodeId: nodeModel.selectedRowKeys,
            roleId: props.role.id,
        };
        let res = await api.Role.bindNodeId(data);
        if (res.code == 1000) {
            message.success("保存成功");
        } else {
            message.error(res.desc);
        }
    }


    return <Contianer>
        <div className="box-head">
            <div className="text">
                <span><label>分组:</label>{props.parent.name}</span>
                <span><label>角色:</label>{props.role.name}</span>
            </div>
        </div>
        <div className="box-body">
            <Table
                rowKey={(record) => record.id}
                columns={columns}
                dataSource={nodeModel.list}
                rowSelection={selection}
                pagination={false}
            />
            <Space>
            </Space>
        </div>
        <div className="box-foot">
                <Button type="primary" onClick={() => saveNode()}>保存</Button>
        </div>
    </Contianer>
}