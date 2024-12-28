import { Table } from "antd";
import { useEffect, useState } from "react";

import { api, model } from "@/service";
import topbar from "topbar";

export function List() {
  const columns = [
    {
      title: '商品编号',
      dataIndex: 'id',
      key: 'id',
    },
    {
      title: '商品名称',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: '商品价格',
      dataIndex: 'price',
      key: 'price',
    },
    {
     title: '商品库存',
    }
  ];
  const [req, setReq] = useState(new model.Request({current: 1, pageSize: 10}));
  const [product, setProduct] = useState({
    list: new Array<model.Product>(),
    total: 0,
  });

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
    setProduct({...product});
    topbar.hide();
  };

  useEffect(() => {
    getProduct();
  }, [req.current]);

  return (
    <>
      <div className="box-head">
        <h1>分销商品</h1>
      </div>
      <div className="box-body">
        <Table columns={columns} dataSource={product.list} />
      </div>
    </>
  );
}