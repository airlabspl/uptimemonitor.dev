import { createContext, useCallback, useContext, useEffect, useState, type ReactNode } from "react";

type User = {
    name: string;
    email: string;
    verified: boolean;
}

type AuthContextType = {
    user: User | null;
    loading: boolean;
    error: any;
    reload: () => void;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const [loading, setLoading] = useState(true);
    const [user, setUser] = useState<User | null>(null);
    const [error, setError] = useState<any>(null);

    const reload = useCallback(() => {
        fetch(`/v1/profile`, {
            credentials: 'include',
        })
            .then(res => res.json())
            .then(data => setUser(data))
            .finally(() => setLoading(false))
            .catch(err => setError(err));
    }, []);

    useEffect(() => reload(), [reload]);

    const value: AuthContextType = {
        user,
        loading,
        error,
        reload,
    };

    return (
        <AuthContext.Provider value={value}>
            {children}
        </AuthContext.Provider>
    );
}

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (context === undefined) {
        throw new Error("useAuth must be used within an AuthProvider");
    }
    return context;
};