import React, { ReactNode, createContext, useState } from "react";
import api from "../apis/api";

type AuthState = {
    isLoggedIn: boolean;
};

type UserState = {
    username: string;
    display_name: string;
};

type AuthLoginProps = {
    status: number;
    data: string;
};

type AuthContextProps = {
    user: UserState;
    auth: AuthState;
    setAuth: React.Dispatch<React.SetStateAction<AuthState>>;
    handleAuthLogIn: (username: string, password: string) => Promise<AuthLoginProps>;
};

const AuthContext = createContext<AuthContextProps>({} as AuthContextProps);

export const AuthProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [user, setUser] = useState<UserState>({ username: "", display_name: "" });
    const [auth, setAuth] = useState<AuthState>({ isLoggedIn: false });

    const handleAuthLogIn = async (username: string, password: string): Promise<AuthLoginProps> => {
        try {
            const response = await api.post("/auth/login", { username, password });
            setAuth({ isLoggedIn: true });
            setUser((prev) => {
                return { ...prev, username: response.data.username, display_name: response.data.display_name };
            });
            return { status: response.status, data: response.data };
        } catch (error: any) {
            return { status: error.response.status, data: error.response.data };
        }
    };

    const stateValues: AuthContextProps = { user, auth, setAuth, handleAuthLogIn };

    return <AuthContext.Provider value={stateValues}>{children}</AuthContext.Provider>;
};

export default AuthContext;
