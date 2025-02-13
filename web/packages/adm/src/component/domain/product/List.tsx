import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Table, Flex, Select, Button, TextField, Link } from '@radix-ui/themes';

import { api, model } from "@/service";
import topbar from "topbar";


export function List() {
  const [product, setProduct] = useState({
    list: new Array<model.Product>(),
    total: 0,
  });
  const [req, setReq] = useState({ 
    ...new model.Request({ current: 1, pageSize: 10 }),
  });
  const [query, setQuery] = useState({
    field: 'name',
    value: '',
  });

  const getProduct = async (e?: React.FormEvent<HTMLFormElement>) => {
    if (!!e) {
      e.preventDefault();
    }
    topbar.show();
    let data = {
      ...req,
      queryBy: new Array(),
    }
    if (!!query.value) {
      data.queryBy.push({ field: query.field, value: query.value });
    } else {
      data.queryBy = undefined!;
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

  const onInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    setQuery({
      ...query,
      [e.target.name]: e.target.value,
    })
  }

  useEffect(() => {
    getProduct();
  }, [req.current]);

  const navigate = useNavigate();

  const toPublish = (item?: model.Product) => {
    let url = "/product/publish";
    if (!!item) {
      url += `?id=${item.id}`;
    }
    navigate(url);
  }

  // 下架
  const offShelf = async (item: model.Product) => {
    let data = {
      id: item.id
    }
    let res = await api.Product.offShelf(data);
    if (res.code == 1000) {
      getProduct();
    }
  }

  const onValueChange = (value: string) => {
    setQuery({
      ...query,
      field: value
    })
  }

  return <>
    <div className="box-head">
      <h1>自产商品</h1>
      <Flex gap="3">
        <form className="form-inline" onSubmit={getProduct}>
            <Select.Root onValueChange={onValueChange} defaultValue="name" name="field" size="2">
            <Select.Trigger />
              <Select.Content>
                <Select.Item value="name">商品名</Select.Item>
                <Select.Item value="price">价格</Select.Item>
              </Select.Content>
            </Select.Root>
            <TextField.Root name="value" onChange={onInput}></TextField.Root>
            <Button type="submit">搜索</Button>
        </form>
        <Button onClick={()=>toPublish()}>发布商品</Button>
      </Flex>
    </div>
    <div className="box-body">
      <Table.Root>
        <Table.Header>
          <Table.Row>
            <Table.ColumnHeaderCell>商品名</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>价格(元)</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>库存</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>销量</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>商品状态</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>创建时间</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell justify="center">操作</Table.ColumnHeaderCell>
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {product.list.map(m => <Table.Row key={m.id}>
            <Table.Cell>{m.name}</Table.Cell>
            <Table.Cell>{m.lowPrice}</Table.Cell>
            <Table.Cell>{m.name}</Table.Cell>
            <Table.Cell>{m.saleCount}</Table.Cell>
            <Table.Cell>{m.name}</Table.Cell>
            <Table.Cell>{m.createdAt}</Table.Cell>
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