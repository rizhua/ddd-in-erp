import { Table } from "antd";
import { useEffect, useState } from "react";

import { api, model } from "@/service";
import topbar from "topbar";

export function Comment() {
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
  const [comment, setComment] = useState({
    list: new Array<model.Product>(),
    total: 0,
  });

  const getComment = async () => {
    topbar.show();
    let data = {
      ...req,
    }
    let res = await api.Product.find(data);
    if (res.code == 1000) {
      comment.list = res.data.list;
      comment.total = res.data.total;
    } else {
      comment.list = [];
      comment.total = 0;
    }
    setComment(comment);
    topbar.hide();
  }

  useEffect(() => {
    getComment();
  }, [req.current]);

  return <>

  <div className="box-body">
    <Table
      columns={columns}
      dataSource={comment.list}
      rowKey="id"
      pagination={{current: req.current, pageSize: req.pageSize, total: comment.total}}
      onChange={(pagination) => {
        setReq({
          ...req,
          current: pagination.current,
        })
      }}
    />
  </div>
  </>
}