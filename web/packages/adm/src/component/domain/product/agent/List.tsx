import { Flex, Link, Table } from "@radix-ui/themes";
import { useEffect, useState } from "react";

import { api, model } from "@/service";
import topbar from "topbar";
import { toast } from "react-toastify";

export function List() {
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

  const toPublish = (item: model.Product) => {
    window.location.href = `/product/publish/${item.id}`;
  }

  const offShelf = (item: model.Product) => {
    toast.error('该商品无法下架');
  }

  return <>
      <div className="box-head">
        <h1>分销商品</h1>
      </div>
      <div className="box-body">
        <Table.Root>
            <Table.Header>
              <Table.Row>
                  <Table.ColumnHeaderCell>商品编号</Table.ColumnHeaderCell>
                  <Table.ColumnHeaderCell>商品名称</Table.ColumnHeaderCell>
                  <Table.ColumnHeaderCell>商品价格</Table.ColumnHeaderCell>
                  <Table.ColumnHeaderCell>商品库存</Table.ColumnHeaderCell>
                  <Table.ColumnHeaderCell>操作</Table.ColumnHeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {product.list.map(m => <Table.Row key={m.id}>
                <Table.Cell>{m.id}</Table.Cell>
                <Table.Cell>{m.name}</Table.Cell>
                <Table.Cell>{m.name}</Table.Cell>
                <Table.Cell>{m.name}</Table.Cell>
                <Table.Cell width="100px">
                  <Flex gap="3" justify="center">
                    <Link onClick={() => toPublish(m)}>编辑</Link>
                    <Link onClick={() => offShelf(m)}>下架</Link>
                  </Flex>
                </Table.Cell>
              </Table.Row>)}
            </Table.Body>
          </Table.Root>
      </div>
    </>
}