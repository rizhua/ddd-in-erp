import { useEffect, useState } from "react";
import { Button, Dropdown, Empty, Form, Input, message, Modal, Popconfirm, Space, Tooltip, Tree, TreeSelect } from "antd";
import styled from "styled-components";

import { model, api } from "@/service";
import topbar from "topbar";
import * as attribute from "./attribute";
import { Dialog } from "@radix-ui/themes";

const Container = styled.div`
  display: flex;
  justify-content: space-between;
`;
const Side = styled.div`
  min-width: 280px;
  padding-right: 16px;
  margin-right: 16px;
  border-right: 1px solid #f0f0f0;
  min-height: calc(100vh - 52px);

  .box-foot {
    display: flex;
    justify-content: center;
    padding: 16px 0;

    button {
      width: 100%;
      max-width: 200px;
    }
  }

  .ant-tree-title {
        display: grid;
        grid-template-columns: 1fr 25px;
        align-items: center;
    }
`
const Main = styled.div`
  flex: 1;
  padding: 0 20px;

  .ant-empty {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
  }

  
`


export function List() {
  document.title = '商品类目';
  const [category, setCategory] = useState({
    info: new model.Category(),
    list: new Array<model.Category>(),
    dialog: {
      open: false,
      title: '',
    }
  });
  const [queryBy] = Form.useForm();

  const [editable, setEditable] = useState<number[]>([]);
  const selCell = (item: model.Category) => {
    let s = new Set([...editable]);
    s.add(item.id);
    setEditable([...s]);
  }

  const getCategory = async () => {
    topbar.show();
    let data: model.Request = {
      current: 1,
      pageSize: 1000,
    };
    if (!!queryBy.getFieldValue('value')) {
      data.queryBy = [{ field: queryBy.getFieldValue('field'), value: queryBy.getFieldValue('value') }];
    }
    let res = await api.Category.find(data);
    if (res.code == 1000) {
      category.list = res.data.list;
    } else {
      category.list = [];
    }
    setCategory({ ...category });
    topbar.hide();
  }

  const [categoryForm] = Form.useForm();

  const disCategory = (item?: model.Category, isSub?: boolean) => {
    category.dialog.open = !category.dialog.open;
    let tmp = new model.Category();
    category.dialog.title = '新增分类';
    if (!!item) {
      if (!!isSub) {
        category.dialog.title = '新增分类';
        tmp.parentId = item.id;
      } else {
        category.dialog.title = '编辑分类';
        Object.assign(tmp, item);
      }
    }
    if (tmp.parentId == 0) {
      tmp.parentId = null!;
    }
    categoryForm.setFieldsValue(tmp);
    setCategory({ ...category });
  }

  const onCategory = async () => {
    let valid = await categoryForm.validateFields().catch(e => console.log(e));
    if (!valid) {
      return;
    }
    let data = {
      ...categoryForm.getFieldsValue(),
    };
    let res: model.Response;
    if (data.id > 0) {
      res = await api.Category.update(data);
    } else {
      res = await api.Category.create(data);
    }
    if (res.code == 1000) {
      disCategory();
      getCategory();
    } else {
      message.error(res.desc);
    }
  }

  const delCategory = async (item: model.Category) => {
    let data = {
      id: [item.id],
    };
    let res = await api.Category.delete(data);
    if (res.code == 1000) {
      message.success('删除成功');
      getCategory();
    } else {
      message.error(res.desc);
    }
  }

  useEffect(() => {
    getCategory();
  }, []);

  const optCategory = (v: model.Category) => {
    category.info = v;
    setCategory({ ...category });
    console.log(v);
  }

  return <Container>
    <Side>
      <Tree
        treeData={category.list}
        blockNode
        fieldNames={{ title: 'name', key: 'id' }}
        onSelect={(k, v) => optCategory(v.node)}
        showLine
        titleRender={(node: any) => {
          return <>
            <span>{node.name}</span>
            <Dropdown menu={{
              items: [{
                key: '1',
                label: (<span onClick={() => disCategory(node)}>编辑类目</span>)
              }, {
                key: '2',
                label: (<span onClick={() => disCategory(node, true)}>新增子类</span>)
              }, {
                key: '3',
                label: (<span onClick={() => delCategory(node)}>删除</span>)
              }]
            }} trigger={['click']}>
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
      <div className="box-foot">
        <Button type="primary" onClick={() => disCategory()}>新增分类</Button>
      </div>
      <Modal
        open={category.dialog.open}
        title={category.dialog.title}
        width={360}
        footer={
          <Space>
            <Button onClick={() => disCategory()}>取消</Button>
            <Button type="primary" onClick={onCategory}>确定</Button>
          </Space>
        }
        onCancel={() => disCategory()}
      >
        <Form form={categoryForm} layout="vertical">
          <Form.Item name="id" hidden>
            <Input />
          </Form.Item>
          <Form.Item
            name="name"
            label="名称"
            rules={[{ required: true, message: '请输入节点名称!' }]}
          >
            <Input />
          </Form.Item>
          <Form.Item name="parentId" label="上级节点">
            <TreeSelect treeData={category.list} fieldNames={{ label: 'name', value: 'id' }} allowClear />
          </Form.Item>
        </Form>
      </Modal>
    </Side>
    <Main>
      {category.info.id > 0 ? <attribute.List category={category.info} /> : <Empty description={false} />}
    </Main>
  </Container>
}