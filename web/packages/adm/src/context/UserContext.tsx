import { createContext, useReducer, ReactNode, useEffect } from "react";

import { api, model } from "@/service";

export const UserContext = createContext({
    state: new model.User(),
    dispatch: new Function(),
});

function reducer(state: model.User, action: {type: string, payload: any}) {
    switch(action.type) {
        case 'login': 
            localStorage.setItem('token', action.payload.token);
            state = action.payload.info;
            return state;
        case 'setUser': 
            return action.payload;
        case 'logout': 
            localStorage.clear();
            api.User.logout();
            return new model.User();
        default: 
            throw Error('Unknow action: ' + action.type);
    }
}

export function UserProvider(props: { children: ReactNode}) {
    const [state, dispatch] = useReducer(reducer, new model.User());

    const getUser = async () => {
        let res = await api.User.parse();
        if (res.code == 1000) {
            dispatch({
                type: 'setUser',
                payload: res.data,
            });
        }
    }

    useEffect(() => {
        let token = localStorage.getItem('token');
        if (!!token) {
            getUser();
        }
    }, []);

    return (
        <UserContext.Provider value={{state, dispatch}}>
            {props.children}
        </UserContext.Provider>
    )
}