import { Button, Form, Input, Select, Space, Table } from "antd";
import { useEffect, useState } from "react";

import { api, model } from "@/service";
import topbar from "topbar";
import { useNavigate } from "react-router-dom";

const columns = [{
  title: "商品名",
  dataIndex: "name",
  render: (f: string, r: model.Product) => {
    return <a>
      <img src="/images/product/1.png" alt="" />
      <span>{f}</span>
    </a>
  }
}, {
  title: "价格(元)",
  dataIndex: "lowPrice",
}, {
  title: "库存",
  dataIndex: "stock",
}, {
  title: "销量",
  dataIndex: "saleCount",
}, {
  title: "排序",
  dataIndex: "sort",
}, {
  title: "商品状态",
  dataIndex: "status",
}, {
  title: "创建时间",
  dataIndex: "createAt",
}, {
  title: "操作",
  dataIndex: "action",
  render: (f: undefined, r: model.Product) => {
    return <Space>
      <a href="#">编辑</a>
      <a href="#">编辑</a>
    </Space>
  }
}];

export function List() {
  const [product, setProduct] = useState({
    list: new Array<model.Product>(),
    total: 0,
  });
  const [req, setReq] = useState(new model.Request({current: 1, pageSize: 10}));

  const getProduct = async () => {
    topbar.show();
    let data = {
      ...req,
    }
    let res = await api.Product.find(data);
    if (res.code == 1000) {
      product.list = res.data.list;
      product.total = res.data.total;
    } else {
      product.list = [];
      product.total = 0;
    }
    setProduct({ ...product });
    topbar.hide();
  }

  useEffect(() => {
    getProduct();
  }, [req.current]);

  const navigate = useNavigate();

  const toPublish = () => {
    navigate("/product/publish");
  }

  return <>
    <div className="box-head">
      <h1>自产商品</h1>
      <Form layout="inline" initialValues={{ field: "name" }}>
        <Form.Item name="field">
          <Select>
            <Select.Option value="name">商品名</Select.Option>
          </Select>
        </Form.Item>
        <Form.Item name="value">
          <Input />
        </Form.Item>
        <Form.Item>
          <Button type="primary">搜索</Button>
        </Form.Item>
        <Button type="primary" onClick={toPublish}>发布商品</Button>
      </Form>
    </div>
    <div className="box-body">
      <Table columns={columns} dataSource={product.list} pagination={{ ...req, total: product.total }} onChange={(e) => setReq(e)} />
    </div>
  </>
}