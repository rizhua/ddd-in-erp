import { useContext, useEffect, useState } from "react";
import { Link, Outlet, useNavigate, useSearchParams } from "react-router-dom";
import { message } from "antd";

import { model, api } from "@/service";
import styled from "styled-components";
import { UserContext } from "@/context";

const Container = styled.div`
    display: flex;
    min-height: 100vh;
`;

const Side = styled.aside`
    height: 100vh;
    background-color: #f6f6f3;
    position: relative;
    z-index: 100;

    &:hover > .close {

        /* .close { */
            display: flex;
            justify-content: center;
            align-items: center;
        /* } */
    }

    .close {
        display: none;
        cursor: pointer;
        width: 24px;
        height: 24px;
        color: #fff;
        position: absolute;
        top: 16px;
        right: -24px;
        border-top-right-radius: 12px;
        border-bottom-right-radius: 12px;
        background-color: rgba(0,0,0,.3);

        &:hover {
            width: 32px;
            right: -32px;
        }
    }

    .sidebar {
        display: flex;
        height: 100vh;
    }

    .nav-head {
        cursor: pointer;
        display: flex;
        gap: 8px;
        align-items: center;
        position: relative;

        .icon {
            display: flex;
            justify-content: center;
            align-items: center;
            width: 20px;
            height: 20px;
            border-radius: 2px;
            background-color: #ddd;
        }

        .text {
            align-items: center;
            max-width: 38px;
            white-space: nowrap;
            text-overflow: ellipsis;
            overflow: hidden;
        }
    }

    .nav-first {
        width: 118px;
    }

    .menu-item {
        text-decoration: none;
    }

    .menu-link {
        display: grid;
        grid-template-columns: 16px 1fr 16px;
        gap: 8px;
        cursor: pointer;
        margin: 8px;
        padding: 6px 8px;
        color: #666;
        border-radius: 4px;

        &:hover {
            background-color: #eee;
        }

        &.active {
            background-color: #eee;
        }
    }

    .menu-text {
        flex: 1;
    }

    .sub-menu {
        .circle {
            display: flex;
            justify-content: center;
            align-items: center;

            &::before {
                display: flex;
                justify-content: center;
                align-items: center;
                content: '';
                width: 4px;
                height: 4px;
                border-radius: 50%;
                background-color: #333;
            }
        }
    }

    .nav-second {
        width: 184px;
        box-shadow: inset 2px 0 0 #f0f0f0;
    }

    .nav-group-name {
        display: flex;
        padding: 8px 16px;
        color: #999;
    }

    .nav-group-item {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        gap: 8px;
        margin: 0;
        padding: 8px;

        a {
            padding: 4px 8px;
            border-radius: 4px;
            color: #666;

            &:hover {
                background-color: #eee;
            }

            &.active {
                background-color: #eee;
            }
        }
    }

    .row {
        display: flex;
        gap: 8px;

        .icon {
            display: flex;
            justify-content: center;
            align-items: center;
            width: 48px;
            height: 48px;
            font-size: 24px;
            border-radius: 4px;
            background-color: #f0f0f0;
        }
    }

    .cell {
        cursor: pointer;
    }

    .btn {
        cursor: pointer;
    }
`;

const Box = styled.div`
    position: absolute;
    left: 16px;
    top: 52px;
    min-width: 320px;
    border-radius: 8px;
    color: #666;
    background-color: #fff;
    box-shadow: 0 0 4px #ccc;
    z-index: 1000;

    a {
        color: #666;
        font-weight: 400;
    }

    .txt-core {
        font-weight: 500;
        font-size: 14px;
        color: #333;
    }

    .box-head {
        display: grid;
        grid-template-rows: 50px 50px;
        padding: 16px;

        img {
            width: 36px;
            height: 36px;
            border-radius: 4px;
        }
    }

    .box-body {
        background-color: #f9f9f9;

        .box-part {
            display: grid;
            gap: 8px;
            padding: 16px;
            border-top: 1px solid #e6e6e6;
        }
    }
`;

const Main = styled.main`
    flex: 1;
    height: 100vh;
    overflow-y: auto;

    .nav {
        display: flex;
        justify-content: flex-end;
        align-items: center;
        gap: 24px;
        height: 52px;
        padding: 0 16px;
    }

    .box {
        padding: 0 16px;
        min-height: calc(100vh - 52px);
    }
`;

export function Console() {
    const [menu, setMenu] = useState({
        list: new Array<model.Node>(),
        sub: new Array<model.Node>(),
    });
    const navigate = useNavigate();
    let [searchParams] = useSearchParams();
    let path = new Set(searchParams.get('path')?.split(','));
    let userContext = useContext(UserContext);

    const [org, setorg] = useState({
        list: new Array<model.Org>(),
    });

    const work = async () => {
        let res = await api.User.work();
        if (res.code == 1000) {
            org.list = res.data;
        } else {
            org.list = [];
        }
        setorg({ ...org });
        getMenu();
    }

    let getMenu = async () => {
        let data = {
            meta: 'rzp'
        };
        let res = await api.Node.permission(data);
        if (res.code == 1000) {
            menu.list = res.data;
            menu.list.forEach(v => {
                if (path.has(v.id.toString())) {
                    setSubMenu(v);
                    return;
                }
            });
        } else {
            menu.list = [];
        }
        setMenu({ ...menu });
    }

    const setSubMenu = (node: model.Node) => {
        let url = '';
        menu.sub = [];
        switch (node.meta) {
            case 'A':
                node.children?.forEach(v => {
                    if (path.has(v.id.toString())) {
                        setSubMenu(v);
                        return;
                    }
                });
                return;
            case 'B':
                if (!!!node.children || node.children.length == 0) {
                    return;
                }
                menu.sub = node.children;
                let tmp = menu.sub[0].children;
                if (!!tmp && tmp.length > 0) {
                    url = tmp[0].meta + '?path=' + tmp[0].path;
                }
                break;
            default:
                url = node.meta + '?path=' + node.path;
        }
        setMenu({ ...menu });
        navigate(url);
    }

    const onExpand = (index: number) => {
        let node = menu.list[index];
        node.expanded = !node.expanded;
        setMenu({ ...menu });
    }

    // 退出登录 
    const logout = async () => {
        let res = await api.User.logout();
        if (res.code == 1000) {
            localStorage.removeItem('token');
            navigate('/auth/signIn');
        } else {
            message.error(res.desc);
        }
    }

    useEffect(() => {
        work();
    }, []);

    const [open, setOpen] = useState(false);

    // 切换组织
    const onOrg = async (org: model.Org) => {
        let data = {
            id: org.id,
        }
        let res = await api.Structure.switch(data);
        if (res.code == 1000) {
            location.reload();
        } else {
            message.error(res.desc);
        }
    }

    return <Container>
        <Side>
            <a className="close iconfont">&#xe682;</a>
            <div className="sidebar">
                <ul className="nav-first">
                    <li className="menu-item">
                        <a className="nav-head menu-link" onClick={() => setOpen(!open)}>
                            <i className="icon">{userContext.state.nickname.slice(0,1).toUpperCase()}</i>
                            <div className="text">{userContext.state.nickname}</div>
                            <i className="iconfont">&#xe66e;</i>
                        </a>
                        {open && <Box>
                            <div className="box-head">
                                <div className="row">
                                    <div className="icon">{userContext.state.nickname.slice(0,1).toUpperCase()}</div>
                                    <div>
                                        <div className="txt-core">{userContext.state.nickname}</div>
                                        <span>Free Plan.1 member</span>
                                    </div>
                                </div>
                                <div className="row">
                                    <button className="btn btn-default">设置</button>
                                    <button className="btn btn-default">邀请</button>
                                </div>
                            </div>
                            <div className="box-body">
                                <div className="box-part">
                                    <div className="cell">
                                        <span>{userContext.state.email || userContext.state.mobile}</span>
                                        <a className="iconfont">&#xe625;</a>
                                    </div>
                                    {org.list.map((m) => <div className="cell" key={m.id} onClick={() => onOrg(m)}>
                                        <div className="cell-text">
                                            <div className="icon">H</div>
                                            <span>{m.fullName}</span>
                                        </div>
                                        {m.id == userContext.state.org.id && <i className="iconfont">&#xe6de;</i>}
                                    </div>
                                    )}
                                </div>
                                <div className="box-part">
                                    <a onClick={logout}>退出登录</a>
                                </div>
                            </div>
                        </Box>}
                    </li>
                    {menu.list.map((v, k) => v.leaf || v.meta == 'B' ?
                        <li className="menu-item" key={v.id}>
                            <a className={"menu-link" + (path.has(v.id.toString()) ? ' active' : '')} onClick={() => setSubMenu(v)}>
                                <i className="iconfont" dangerouslySetInnerHTML={{ __html: v.icon }}></i>
                                <span className="menu-text">{v.name}</span>
                            </a>
                        </li>
                        : <li className="menu-item" key={v.id}>
                            <a className="menu-link" onClick={() => onExpand(k)}>
                                <i className="iconfont" dangerouslySetInnerHTML={{ __html: v.icon }}></i>
                                <span>{v.name}</span>
                                <i className="iconfont" dangerouslySetInnerHTML={{ __html: v.expanded ? '&#xe66c;' : '&#xe66e;' }}></i>
                            </a>
                            {v.expanded && <ul className="sub-menu">
                                {v.children?.map(n =>
                                    <li className="menu-item" key={n.id}>
                                        <a className={"menu-link" + (path.has(n.id.toString()) ? ' active' : '')} onClick={() => setSubMenu(n)}>
                                            <i className="circle"></i>
                                            <span>{n.name}</span>
                                        </a>
                                    </li>
                                )}
                            </ul>}
                        </li>
                    )}
                </ul>
                {menu.sub.length > 0 && <div className="nav-second">
                    {menu.sub.map(x => !x.leaf && <div className="nav-group" key={x.name}>
                        <div className="nav-group-name">{x.name}</div>
                        <div className="nav-group-item">
                            {x.children?.map(y => <Link className={path.has(y.id.toString()) ? ' active' : ''} to={y.meta + '?path=' + y.path} key={y.id}>{y.name}</Link>)}
                        </div>
                    </div>)}
                    <div className="nav-group-name">&nbsp;</div>
                    <div className="nav-group-item">
                        {menu.sub.map(x => x.leaf && <Link className={path.has(x.id.toString()) ? ' active' : ''} to={x.meta + '?path=' + x.path} key={x.id}>{x.name}</Link>)}
                    </div>
                </div>}
            </div>
        </Side>
        <Main>
            <nav className="nav">
                <Link className="nav-item iconfont" to="/news/notice">&#xe649;</Link>
            </nav>
            <div className="box">
                <Outlet />
            </div>
        </Main>
    </Container>
}