import { Checkbox, Form, Input } from "antd";

export function Subcribe() {

    return <div className="box">
        <div className="box-body">
            <h2>管理接收人</h2>
            <p>最多可新增5名接收人。接受财务和账号信息，如需接收产品消息请前往各个产品控制台单独设置。</p>
            <Form layout="inline">
                <Form.Item label="姓名">
                    <Input />
                </Form.Item>
                <Form.Item label="绑定邮箱">
                    <Input />
                </Form.Item>
                <Form.Item label="绑定手机">
                    <Input />
                </Form.Item>
            </Form>
            <h2>通知事件项</h2>
            <p>为了保障您在第一时间接收到有关财务、账号安全和系统公告等关键信息，本设置页面部分选项 (打钩图标) 默认激活且无法更改。其余可配置选项请依据您的具体需求激活或禁用。</p>
            <table className="table">
                <thead>
                    <tr>
                        <th colSpan={2}>事件项</th>
                        <th>站内信</th>
                        <th>短信</th>
                        <th>邮件</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td rowSpan={2}>财务</td>
                        <td>充值到账通知</td>
                        <td><Checkbox /></td>
                        <td><Checkbox /></td>
                        <td><Checkbox /></td>
                    </tr>
                    <tr>
                        <td>月账单通知</td>
                        <td><Checkbox /></td>
                        <td><Checkbox /></td>
                        <td><Checkbox /></td>
                    </tr>
                    <tr>
                        <td rowSpan={2}>账号</td>
                        <td>登录通知</td>
                        <td><Checkbox /></td>
                        <td><Checkbox /></td>
                        <td><Checkbox /></td>
                    </tr>
                    <tr>
                        <td>登录密码修改</td>
                        <td><Checkbox /></td>
                        <td><Checkbox /></td>
                        <td><Checkbox /></td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
}