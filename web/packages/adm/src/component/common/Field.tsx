import styled from "styled-components";

const Input = styled.input`
    
`

type InputProps = {
    type: string
}

export function Field(props: InputProps) {
    const { type = 'text' } = props

    return <Input type={type} />
}