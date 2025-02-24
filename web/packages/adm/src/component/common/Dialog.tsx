import React, { useEffect, useState } from "react";
import styled from "styled-components";

interface StyledContainerProps {
    justifyContent?: 'center' | 'right' | 'left'
}

const Container =  styled.div<StyledContainerProps>`
    display: flex;
    /* justify-content: ${props => {
        switch (props.justifyContent) {
            case 'right':
                return 'flex-end';
            case 'left':
                return 'flex-start';
            default:
                return 'center';
        }
    }}; */
    align-items: center;
    position: absolute;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    z-index: 100;

    button {
        cursor: pointer;
    }

    &.mask {
        background-color: rgba(0, 0, 0, .5);
        z-index: 100;
    }
`
const Box = styled.div`
    display: grid;
    grid-template-rows: 50px 1fr auto;
    min-width: 360px;
    border-radius: 4px;
    background-color: #fff;

    .box-head {
        display: grid;
        grid-template-columns: 1fr auto;
        align-items: center;
        min-height: 50px;
        padding: 0 16px;

        h2 {
            margin: 0;
        }

        .icon {
            cursor: pointer;
            display: flex;
            justify-content: center;
            align-items: center;
        }
    }

    .box-body {
        padding: 16px;
        max-height: 100vh;
        overflow-y: auto;
    }

    .box-foot {
        display: flex;
        justify-content: flex-end;
        gap: 16px;
        padding: 16px;
        border-top: 1px solid #e0e0e0;
    }
`

export type DialogProps = {
    open: boolean;
    title: string;
    children: React.ReactNode;
    mask?: boolean;
    placement?: 'center' | 'left' | 'right';
    footer?: React.ReactNode;
    onClose?: () => void;
    onOk?: () => void;
}

export function Dialog(props: DialogProps) {
    const {placement = 'center', mask = true} = props;
    const [open, setOpen] = useState(false);

    const onClose = () => {
        props.onClose && props.onClose();
    }

    useEffect(() => {
        setOpen(props.open);
    }, [props.open]);

    return <Container className={mask ? 'mask' : ''} style={{display: open ? 'flex' : 'none', justifyContent: placement}} onClick={() => onClose()}>
        <Box onClick={(e) => e.stopPropagation()} style={{minHeight: placement != 'center' ? '100vh' : 'unset', borderRadius: placement != 'center' ? '0' : '4px'}}>
            <div className="box-head">
                <h2>{props.title}</h2>
                <a className="icon" onClick={() => onClose()}>
                    <svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M12.8536 2.85355C13.0488 2.65829 13.0488 2.34171 12.8536 2.14645C12.6583 1.95118 12.3417 1.95118 12.1464 2.14645L7.5 6.79289L2.85355 2.14645C2.65829 1.95118 2.34171 1.95118 2.14645 2.14645C1.95118 2.34171 1.95118 2.65829 2.14645 2.85355L6.79289 7.5L2.14645 12.1464C1.95118 12.3417 1.95118 12.6583 2.14645 12.8536C2.34171 13.0488 2.65829 13.0488 2.85355 12.8536L7.5 8.20711L12.1464 12.8536C12.3417 13.0488 12.6583 13.0488 12.8536 12.8536C13.0488 12.6583 13.0488 12.3417 12.8536 12.1464L8.20711 7.5L12.8536 2.85355Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
                    </svg>
                </a>
            </div>
            <div className="box-body">
                {props.children}
            </div>
            {!props.footer ? <div className="box-foot">
                <button className="btn btn-default" onClick={() => onClose()}>取消</button>
                <button className="btn btn-primary" onClick={() => props.onOk && props.onOk()}>确定</button>
            </div> : props.footer}
        </Box>
    </Container>
}