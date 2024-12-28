import { useContext, useEffect, useState } from "react";
import { Dropdown, Form, Input, message, Modal, Select, Tree, TreeSelect } from "antd";
import modal from "antd/es/modal";
import styled from "styled-components";
import topbar from "topbar";

import { model, api } from "@/service";
import { UserContext } from "@/context";
import { Emp } from "./Emp";

const Container = styled.div`
    display: grid;
    grid-template-columns: 280px 1fr;
    min-height: calc(100vh - 52px);

    .side {
        margin-right: 16px;
        padding-right: 16px;
        border-right: 1px solid #eee;
    }

    .side-head {
        display: grid;
        grid-template-columns: 24px 1fr;
        align-items: center;
        margin-bottom: 16px;

        a {
            cursor: pointer;
        }

        img {
            width: 24px;
            height: 24px;
            border-radius: 50%;
        }

        .icon {
            width: 24px;
            height: 24px;
            background: #eee;
            border-radius: 2px;
            display: flex;
            justify-content: center;
            align-items: center;
            color: #999;
            font-size: 16px;
            font-weight: bold;
        }

        .text {
            padding: 0 8px;
            font-weight: 500;
        }
    }

    .side-foot {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 16px;

        .btn-primary {
            cursor: pointer;
            width: 80%;
        }
    }

    .ant-tree-title {
        display: grid;
        grid-template-columns: 1fr 25px;
        align-items: center;
    }

`;

export function List() {
    const userContext = useContext(UserContext);
    const [dept, setDept] = useState({
        info: new model.Dept(),
        list: new Array<model.Dept>(),
        modal: { title: '', open: false },
    });

    const getDept = async () => {
        topbar.show();
        let res = await api.Structure.findDept({});
        if (res.code == 1000) {
            dept.list = res.data;
        } else {
            dept.list = [];
        }
        setDept({ ...dept });
        topbar.hide();
    }

    const [deptForm] = Form.useForm();

    const disDept = (item?: model.Dept, isEdit?: boolean) => {
        dept.modal.open = !dept.modal.open;
        deptForm.resetFields();
        if (!!item) {
            if (isEdit) {
                dept.modal.title = '编辑部门';
                if (item.parentId == 0) {
                    item.parentId = null!;
                }
                deptForm.setFieldsValue(item);
            } else {
                dept.modal.title = '新增部门';
                deptForm.setFieldValue('parentId', item.id);
            }
            item.mgrId == 0 && deptForm.setFieldValue('mgrId', null);
        }
        setDept({ ...dept });
    }

    const optDept = (v: model.Dept) => {
        dept.info = v;
        setDept({ ...dept });
    }

    const onDept = async () => {
        let data = {
            ...deptForm.getFieldsValue(),
        };
        let res: model.Response;
        if (data.id > 0) {
            res = await api.Structure.updateDept(data);
        } else {
            res = await api.Structure.createDept(data);
        }
        if (res.code == 1000) {
            disDept();
            getDept();
        } else {
            message.error(res.desc);
        }
    }

    const delDept = async (item: model.Dept) => {
        modal.confirm({
            title: '删除部门',
            content: `确定要删除部门【${item.name}】吗？`,
            okText: '确认',
            cancelText: '取消',
            onOk: async () => {
                let data = {
                    id: [item.id]
                };
                let res = await api.Structure.deleteDept(data);
                if (res.code == 1000) {
                    message.success('删除成功');
                    getDept();
                } else {
                    message.error(res.desc);
                }
            }
        });
    }

    useEffect(() => {
        getDept();
        return () => {
            dept.info = new model.Dept();
            setDept({ ...dept });
        }
    }, []);

    const [emp, setEmp] = useState({
        list: new Array<model.Emp>(),
    });

    const getEmp = async (kw: string) => {
        if (!kw) {
            return;
        }
        let data = {
            currPage: 1,
            pageSize: 10,
            queryBy: [
                { field: 'keyword', value: kw }
            ]
        };
        let res = await api.Structure.findEmp(data);
        if (res.code == 1000) {
            emp.list = res.data.list;
        } else {
            emp.list = [];
        }
        setEmp({ ...emp });
    }

    return <Container>
        <aside className="side">
            <div className="side-head">
                <div className="icon">
                    {userContext.state.org.icon ? <img src={userContext.state.org.icon} /> : 'H'}
                </div>
                <span className="text">{userContext.state.org.name}</span>
            </div>
            <Tree
                defaultExpandAll
                treeData={dept.list}
                blockNode
                onSelect={(k, v) => optDept(v.node)}
                fieldNames={{ title: 'name', key: 'id' }}
                showLine
                titleRender={(node: any) => {
                    return <>
                        <span>{node.name}({node.empCount}人)</span>
                        <Dropdown menu={{
                            items: [{
                                key: '1',
                                label: (<span onClick={() => disDept(node, true)}>编辑部门</span>)
                            }, {
                                key: '2',
                                label: (<span onClick={() => disDept(node, false)}>新增子部门</span>)
                            }, {
                                key: '3',
                                label: (<span onClick={() => delDept(node)}>删除</span>)
                            }]
                        }} trigger={['hover']}>
                            <a
                                className="popover iconfont"
                                onClick={e => e.preventDefault()}
                            >
                                &#xe6c5;
                            </a>
                        </Dropdown>
                    </>
                }}
            />
            {/* 新增、编辑部门 */}
            <Modal
                title={dept.modal.title}
                open={dept.modal.open}
                onCancel={() => disDept()}
                onOk={onDept}
                width={360}
                forceRender
            >
                <Form form={deptForm} layout="vertical">
                    <Form.Item name="id" hidden>
                        <Input />
                    </Form.Item>
                    <Form.Item name="name" label="部门名称">
                        <Input />
                    </Form.Item>
                    <Form.Item name="parentId" label="上级部门">
                        <TreeSelect treeData={dept.list} fieldNames={{ label: 'name', value: 'id' }} allowClear />
                    </Form.Item>
                    <Form.Item name="mgrId" label="主管">
                        <Select
                            showSearch
                            onSearch={(v) => getEmp(v)}
                            options={emp.list.map(m => ({label: m.name, value: m.id}))}
                        />
                    </Form.Item>

                </Form>
            </Modal>
            <div className="side-foot">
                <button className="btn btn-primary" onClick={() => { deptForm.resetFields(); setDept({ ...dept, modal: { title: '新增部门', open: true } }) }}>新增部门</button>
            </div>
        </aside>
        <main>
            <Emp dept={dept.info} />
        </main>
    </Container>
}