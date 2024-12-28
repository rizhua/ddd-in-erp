import { useEffect, useState } from "react";
import { Button, Form, Input, message, Modal, Popconfirm, Select, Space, Table, Image } from "antd";

import { api, model } from "@/service";
import topbar from "topbar";

export function List() {
  const columns = [{
    title: '名称',
    dataIndex: 'name',
    render: (f: string, r: model.Brand) => {
      return <Space>
        <Image src={r.logo} height={40} />
        <span>{f}</span>
      </Space>
    }
  }, {
    title: '创建时间',
    dataIndex: 'createAt',
  }, {
    title: '操作',
    dataIndex: 'action',
    width: 100,
    render: (f: null, r: model.Brand) => {
      return <Space>
        <a className="iconfont" onClick={() => disBrand(r)} title="编辑">&#xe640;</a>
        <Popconfirm
          title="确定要删除吗？"
          onConfirm={() => delBrand(r)}
        >
          <a className="iconfont" onClick={() => delBrand(r)} title="删除">&#xe618;</a>
        </Popconfirm>
      </Space>
    }
  }];
  const [req, setReq] = useState(new model.Request({current:1, pageSize: 10}));
  const [brand, setBrand] = useState({
    info: new model.Brand(),
    list: new Array<model.Brand>(),
    total: 0,
    modal: {
      title: '新增品牌',
      open: false,
    }
  });

  const getBrand = async () => {
    topbar.show();
    let data = {
      ...req,
    }
    let res = await api.Brand.find(data);
    if (res.code == 1000) {
      brand.list = res.data.list;
      brand.total = res.data.total;
    } else {
      brand.list = [];
      brand.total = 0;
    }
    setBrand({ ...brand });
    topbar.hide();
  }

  useEffect(() => {
    getBrand();
  }, [req.current]);

  const [brandForm] = Form.useForm();

  const disBrand = (item?: model.Brand) => {
    brandForm.resetFields();
    Object.assign(brand.info, new model.Brand());
    if (!!item) {
      brandForm.setFieldsValue(item);
      Object.assign(brand.info, item);
      brand.modal.title = '编辑品牌';
    } else {
      brand.modal.title = '新增品牌';
    }
    setBrand({ ...brand });
    brand.modal.open = !brand.modal.open;
    console.log(brand.info);
  }

  const onBrand = async () => {
    let data = {
      ...brandForm.getFieldsValue(),
      logo: brand.info.logo,
    }
    let res: model.Response;
    if (data.id > 0) {
      res = await api.Brand.update(data);
    } else {
      res = await api.Brand.create(data);
    }
    if (res.code == 1000) {
      disBrand();
      getBrand();
    } else {
      message.error(res.desc);
    }
  }

  const delBrand = async (item: model.Brand) => {
    let data = {
      id: [item.id],
    }
    let res = await api.Brand.delete(data);
    if (res.code == 1000) {
      message.success('删除成功');
      getBrand();
    } else {
      message.error('删除失败');
    }
  }

  const beforeUpload = (file: any) => {
    const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
    if (!isJpgOrPng) {
      message.error('You can only upload JPG/PNG file!');
    }
    const isLt2M = file.size / 1024 / 1024 < 2;
    if (!isLt2M) {
      message.error('Image must smaller than 2MB!');
    }
    return isJpgOrPng && isLt2M;
  }

  const onUpload = async (f: any) => {
    const fileObj = f.target.files && f.target.files[0];
    if (!fileObj) {
      return;
    }
    let formData = new FormData();
    formData.append('file', fileObj);
    let res = await api.File.upload(formData);
    if (res.code == 1000) {
      brand.info.logo = res.data;
    } else {
      message.error(res.desc);
    }
  }


  return <>
    <div className="box-head">
      <h1>品牌管理</h1>
      <Form layout="inline" onFinish={onBrand} initialValues={{ field: 'name' }}>
        <Form.Item name="field">
          <Select>
            <Select.Option value="name">名称</Select.Option>
          </Select>
        </Form.Item>
        <Form.Item name="value">
          <Input type="text" placeholder="请输入关键字" />
        </Form.Item>
        <Form.Item>
          <Button type="primary">搜索</Button>
        </Form.Item>
        <Button type="primary" onClick={() => disBrand()}>新增品牌</Button>
      </Form>
      <Modal
        title={brand.modal.title}
        open={brand.modal.open}
        width={360}
        footer={<Space size={16}>
          <Button type="default" onClick={() => disBrand()}>取消</Button>
          <Button type="primary" htmlType="submit" onClick={() => onBrand()}>确定</Button>
        </Space>}
        onCancel={() => disBrand()}
      >
        <Form form={brandForm} onFinish={onBrand}>
          <Form.Item name="id" hidden={true}>
            <Input type="number" />
          </Form.Item>
          <Form.Item name="name" label="名称">
            <Input />
          </Form.Item>
          <Form.Item label="图标">
            {!!!brand.info.logo ? <input type="file" onChange={(e) => onUpload(e)} /> : <Image src={brand.info.logo} height={40} />}
          </Form.Item>
        </Form>
      </Modal>
    </div>
    <div className="box-body">
      <Table
        columns={columns}
        dataSource={brand.list}
        rowKey="id"
        pagination={{ ...req, total: brand.total }}
        onChange={(e) => setReq(e)}
      />
    </div>
  </>
}