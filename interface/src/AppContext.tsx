import React, { createContext, useEffect, useState } from 'react';
import { User } from './types/models';


export interface AppContextType {
    currentUser: User | null;
    setCurrentUser: (user: User | null) => void;
    isLoggedIn: boolean
}

type Props = {
    children: JSX.Element
}



const AppContext = createContext<AppContextType | null>(null);

export const AppProvider = (props: Props) => {
    const [currentUser, setCurrentUser] = useState<User | null>(null);

    const loadUser = () => {
        const val = localStorage.getItem("user")
        if (!val) {
            setCurrentUser(null)
            return
        }
        setCurrentUser(JSON.parse(val))
    }

    const isLoggedIn = currentUser != null

    useEffect(() => {
        loadUser()
    }, [])

   

    return <AppContext.Provider value={
        {
            currentUser: currentUser,
            setCurrentUser,
            isLoggedIn,
        }}>{props.children}</AppContext.Provider>
}

export default AppContext;