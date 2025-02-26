import { useEffect, useState } from "react";
import styled from "styled-components";

const Container = styled.div`
    display: flex;
    justify-content: flex-end;
    align-items: center;
    gap: 8px;
    padding: 16px 0;

    a.item {
        cursor: pointer;
        display: flex;
        justify-content: center;
        align-items: center;
        width: 32px;
        height: 32px;
        border-radius: 4px;
        color: var(--text-color);

        &:hover {
            background-color: #f0f0f0;
        }

        &.active {
            background-color: var(--primary-color);
            color: #fff;
        }
    }

    .prev, .next {
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 8px;
        min-width: 32px;
        min-height: 32px;
        border-radius: 4px;
        color: var(--text-color);
        font-weight: 600;
        background-color: transparent;

        &:hover {
            background-color: #f0f0f0;
        }

        &[disabled] {
            cursor: not-allowed;
            background-color: transparent;

            &:hover {
                background-color: transparent;
            }
        }
    }

    .options {
        display: flex;
        align-items: center;
        gap: 8px;
    }

    .total-display {
        display: flex;
        align-items: center;
        gap: 4px;

        i {
            font-style: normal;
            color: #f40;
        }
    }

    .next-input {
        display: flex;
        align-items: center;
        gap: 4px;
    }

    .btn-primary {
        min-width: 50px;
    }
`

export interface PaginationProps {
    current?: number;
    pageSize?: number;
    total?: number;
    showSizeChanger?: boolean;
    placement?: 'center' | 'left' | 'right';
    onChange: (page: number, pageSize: number) => void;
}

export function Pagination(props: PaginationProps) {
    const {current = 1, pageSize = 10, total = 0, showSizeChanger = false, placement = 'right'} = props;
    const [totalPage, setTotalPage] = useState(0);

    useEffect(() => {
        let totalPage = Math.ceil(total / pageSize);
        getPageCode(current, totalPage);
        setTotalPage(totalPage);
    }, [total, pageSize]);

    const [pageCode, setPageCode] = useState(new Array<number | string>());

    // 生成页码
    const getPageCode = (current: number, totalPage: number) => {
        let arr = new Array<number | string>();
        arr.push(1);
        if (current > 3) {
            arr.push('...');
        }
        let beg = Math.max(current - 3, 2);
        let end = Math.min(current + 3, totalPage - 1);
        if (end > beg) {
            for (let i = beg; i <= end; i++) {
                arr.push(i);
            }
        }
        if (totalPage > 1) {
            if (totalPage > 6 && current < totalPage - 6) {
                arr.push('...');
            }
            arr.push(totalPage);
        }
        setPageCode([...arr]);
    }
    
    // 页码改变
    const onChangePage = (current: number) => {
        if (current < 1) {
            current = 1;
        };
        
        if (current > totalPage) {
            current = totalPage;
        };
        props.onChange(current, pageSize);
        getPageCode(current, totalPage);
    }

    // 页码大小改变
    const onChangePageSize = (pageSize: number) => {
        props.onChange(current, pageSize);
    }

    // 输入页码
    const [inputPage, setInputPage] = useState(current);

    const onInput = (e: React.ChangeEvent<HTMLInputElement>) => {
        e.preventDefault();
        let tmp = e.target.value;
        if (!isNaN(parseInt(tmp))) {
            setInputPage(parseInt(tmp));
        } else {
            setInputPage(null!);
        }
    }

    // 输入页码回调
    const onConfirm = (e: React.FormEvent) => {
        e.preventDefault();
        let current = inputPage;
        if (inputPage < 1) {
            current = 1;
        }
        if (inputPage > totalPage) {
            current = totalPage;
        }
        onChangePage(current);
    }

    return <Container style={{justifyContent: placement === 'center' ? 'center' : placement === 'left' ? 'flex-start' : 'flex-end'}}>
            <button className="prev" onClick={() => onChangePage(current - 1)} disabled={current === 1}>
                <i>&lt;</i><span>上一页</span>
            </button>
            {pageCode.map((m, i) => (
                m == '...' ? <i key={i + 1} className='ellipsis'>...</i> :
                <a key={i + 1} className={current == m ? 'item active' : 'item'} onClick={() => onChangePage(Number(m))}>{m}</a>
            ))}
            <button className="next" onClick={() => onChangePage(current + 1)} disabled={current === totalPage}>
                <span>下一页</span><i>&gt;</i>
            </button>
            {showSizeChanger && <select value={pageSize} onChange={(e) => onChangePageSize(parseInt(e.target.value))}>
                {[5, 10, 20, 30, 40, 50].map((m) => (
                    <option key={m} value={m}>{m}</option>
                ))}
            </select>}
            <form className="options">
                <div className="total-display"><i>{current}</i>/{totalPage}</div>
                <div className="next-input">
                    到第
                    <input type="number" value={inputPage} onChange={onInput} />
                    页
                </div>
                <button className="btn-primary" type="submit" onClick={onConfirm}>确定</button>
            </form>
    </Container>
}