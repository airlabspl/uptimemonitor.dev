import { createContext, useContext, useEffect, useState, type ReactNode } from "react";

type User = {
    id: string;
    name: string;
    email: string;
}

type AuthContextType = {
    user: User | null;
    loading: boolean;
    error: any;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const [loading, setLoading] = useState(true);
    const [user, setUser] = useState<User | null>(null);
    const [error, setError] = useState<any>(null);

    useEffect(() => {
        fetch(`/v1/profile`)
            .then(res => res.json())
            .then(data => setUser(data))
            .finally(() => setLoading(false))
            .catch(err => setError(err));
    }, []);

    const value: AuthContextType = {
        user,
        loading,
        error,
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