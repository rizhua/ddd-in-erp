import { useEffect, useState } from 'react';
import styled from 'styled-components';
import { Space, Tabs, TabsProps, Modal, Form, message, Select, Dropdown, MenuProps, Input, Button, Empty } from 'antd';
import topbar from 'topbar';

import { Emp, Node } from './';
import { model, api } from '@/service';
import { Dialog } from '@/component/common';


const Container = styled.div`
    display: flex;
    min-height: calc(100vh - 52px);
`

const Side = styled.div`
    min-width: 260px;
    min-height: 100%;
    margin-bottom: 0;
    border-right: 1px solid #ddd;
    margin-right: 16px;
    padding-right: 16px;

    .btn-group {
        display: grid;
        grid-template-columns: repeat(2, 80px);
        justify-content: center;
        gap: 1rem;
        margin-bottom: 16px;
    }

    .level-1, .level-2 {
        display: flex;
        justify-content: space-between;
        line-height: 34px;
        padding: 2px 8px;
        margin-bottom: 8px;

        a {
            cursor: pointer;
            width: 34px;
            text-align: center;
        }
    }

    .level-2 {
        padding-left: 24px;
        border-radius: 4px;
        cursor: pointer;

        :hover {
            background-color: #eee;
        }
    }

    .active {
        background-color: #eee;
    }
`;

const Main = styled.div`
    flex: 1;

    .ant-empty {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100%;
    }
`;

export function List() {
    const [role, setRole] = useState({
        info: new model.Role(),
        list: new Array<model.Role>(),
        modal: {
            open: false,
            title: '',
            type: 'group'
        },
        parent: new model.Role(),
    });

    const getRole = async () => {
        topbar.show();
        let data = {
            current: 1,
            pageSize: 100,
            tree: true,
        };
        let res = await api.Role.find(data);
        if (res.code == 1000) {
            role.list = res.data;
        } else {
            role.list = [];
        }
        setRole({ ...role });
        topbar.hide();
    }

    const [roleForm] = Form.useForm();

    const disRole = (type?: string, item?: model.Role) => {
        role.modal.open = !role.modal.open;
        if (type != undefined) {
            role.modal.type = type;
            roleForm.resetFields();
            switch(type) {
                case 'group':
                    role.modal.title = !!item ? '编辑分组' : '新建分组';
                    break;
                case 'role':
                    role.modal.title = !!item ? '编辑角色' : '新建角色';
                    break;
            }
            if (!!item) {
                roleForm.setFieldsValue(item);
            }
        }
        setRole({ ...role });
    }

    const onRole = async () => {
        let data = await roleForm.validateFields();
        let res = new model.Response();
        if (data.id > 0) {
            res = await api.Role.update(data);
        } else {
            res = await api.Role.create(data);
        }
        if (res.code == 1000) {
            getRole();
            disRole();
        } else {
            message.error(res.desc);
        }
    }

    const optRole = (item: model.Role, t: model.Role) => {
        role.info = item;
        role.parent = t;
        setRole({ ...role });
    }

    useEffect(() => {
        getRole();
    }, []);

    const items: TabsProps['items'] = [
        {
            key: '1',
            label: `成员`,
            children: <Emp parent={role.parent} role={role.info} />,
        }, {
            key: '2',
            label: `节点`,
            children: <Node parent={role.parent} role={role.info} />,
        },
    ];

    return <Container>
        <Side>
            <div className="btn-group">
                <Button type="primary" onClick={() => disRole('group')}>新建分组</Button>
                <Button type="primary" onClick={() => disRole('role')}>新建角色</Button>
                <Dialog
                    title={role.modal.title}
                    open={role.modal.open}
                    onOk={onRole}
                    onClose={() => disRole()}
                    // width={360}
                >
                    <Form form={roleForm}>
                        <Form.Item name="id" hidden>
                            <Input />
                        </Form.Item>
                        <Form.Item label="名称" name="name" rules={[{ required: true, message: '请输入名称' }]}>
                            <Input type="text" />
                        </Form.Item>
                        {role.modal.type == 'role' && <Form.Item label="分组" name="parentId" rules={[{ required: true, message: '请选择分组' }]}>
                            <Select options={role.list} fieldNames={{ label: 'name', value: 'id' }} />
                        </Form.Item>}
                    </Form>
                </Dialog>
            </div>
            {role.list.map((m: model.Role) => {
                let arr = [];
                let items: MenuProps['items'] = [{
                        label: '删除',
                        key: '2',
                        onClick: () => {
                            Modal.confirm({
                                title: '删除分组',
                                content: '确定删除该分组？',
                                onOk: async () => {
                                    let res = await api.Role.delete({id: [m.id]});
                                    if (res.code == 1000) {
                                       getRole();
                                    }
                                }
                            });
                        }
                    }, {
                        label: '编辑',
                        key: '3',
                        onClick: () => disRole('group', m)
                    }
                ];
                let dt = <div className="level-1" key={m.id}>
                    {m.name}
                    <Dropdown menu={{ items: items }} trigger={['click']}>
                        <a className="iconfont" onClick={(e) => e.preventDefault()}>&#xe696;</a>
                    </Dropdown>
                </div>
                arr.push(dt);
                if (m.children && m.children.length > 0) {
                    let a = m.children.map((n) => {
                        let items: MenuProps['items'] = [
                            {
                                label: '删除',
                                key: '2',
                                onClick: () => {
                                    Modal.confirm({
                                        title: '删除角色',
                                        content: '确定删除该角色？',
                                        onOk: async () => {
                                            let res = await api.Role.delete({id : [n.id]});
                                            if (res.code == 1000) {
                                               getRole();
                                            }
                                        }
                                    })
                                }
                            }, {
                                label: '编辑',
                                key: '3',
                                onClick: () => disRole('role', n)
                            }
                        ];
                        let tmp = new model.Role();
                        Object.assign(tmp, n);
                        return <div className={"level-2" + ((n.id == role.info.id) ? ' active' : '')} key={n.id} onClick={() => optRole(tmp, m)} style={{ marginLeft: `${n.level * 16}px` }}>
                            {n.name}
                            <Dropdown menu={{ items: items }} trigger={['click']}>
                                <a onClick={(e) => e.preventDefault()}>
                                    <Space className="iconfont">&#xe696;</Space>
                                </a>
                            </Dropdown>
                        </div>;
                    });
                    arr = arr.concat(a);
                }
                return arr;
            })}
        </Side>
        <Main>
            {role.info.id > 0 ? <Tabs defaultActiveKey="1" items={items} /> : <Empty description={false} />}
        </Main>
    </Container>
}