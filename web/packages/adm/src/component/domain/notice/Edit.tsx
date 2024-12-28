import { Form, Input } from "antd";
import styled from "styled-components";

const Container = styled.div`
    margin: 1rem;
    padding: 1rem;
    min-height: calc(100vh - 84px);
`

export function Edit() {
    let [newsForm] = Form.useForm();

    return <Container className="box">

        <div className="box-body">
            <Form form={newsForm} layout="vertical">
                <Form.Item name="title" label="标题">
                    <Input />
                </Form.Item>
                <Form.Item name="content" label="内容">
                    <Input />
                </Form.Item>
                <Form.Item name="type" label="类型">
                    <Input />
                </Form.Item>
                <Form.Item name="drafter" label="拟稿人">
                    <Input />
                </Form.Item>
                <Form.Item name="scope" label="公布范围">
                    <Input />
                </Form.Item>
                <Form.Item name="attached" label="附件">
                    <Input />
                </Form.Item>
                <Form.Item name="id" hidden>
                    <Input />
                </Form.Item>
            </Form>
        </div>
    </Container>
}