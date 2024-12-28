import { useEffect, useState } from "react";
import { Space, Table } from "antd";
import topbar from "topbar";

import { api, model } from "@/service";

export function List() {
  const columns = [
    {
      title: "编码",
      dataIndex: "code",
    },
    {
      title: "名称",
      dataIndex: "name",
    },
    {
      title: "单位",
      dataIndex: "unit",
    },
    {
      title: "创建时间",
      dataIndex: "createdAt",
    },
  ];
  const [req, setReq] = useState(new model.Request({ current: 1, pageSize: 10 }));
  const [position, setPosition] = useState({
    list: new Array<model.Product>(),
    total: 0,
  });

  const getPosition = async () => {
    topbar.show();
    let data = {
      ...req,
    }
    let res = await api.Product.find(data);
    if (res.code == 1000) {
      position.list = res.data.list;
      position.total = res.data.total;
    } else {
      position.list = [];
      position.total = 0;
    }
    topbar.hide();
  }

  useEffect(() => {
    getPosition();
  }, [req.current]);

  return <>
    <div className="box-head">
      <h1>商品库位</h1>
      <Space>
        <button className="btn btn-primary">货物入库</button>
      </Space>
    </div>
    <div className="box-body">
      <Table
        columns={columns}
        dataSource={position.list}
        rowKey="id"
        pagination={{ current: req.current, pageSize: req.pageSize, total: position.total }}
        onChange={(pagination) => {
          setReq({ ...req, current: pagination.current });
        }}
      />
    </div>
  </>
}