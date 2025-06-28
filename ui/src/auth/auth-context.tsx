import { fetcher } from "@/api/fetcher";
import { createContext, useContext, type ReactNode } from "react";
import useSWR from "swr/immutable";

type User = {
    id: string;
    name: string;
    email: string;
}

const AuthContext = createContext<User | null>(null);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const { data, error, isLoading } = useSWR(`/v1/profile`, fetcher, {
        revalidateOnFocus: false,
        revalidateIfStale: false,
        revalidateOnReconnect: false,
    })

    if (isLoading) return <div>Loading...</div>;

    const user = error ? null : data;

    return (
        <AuthContext.Provider value={user}>
            {children}
        </AuthContext.Provider>
    );
}

export const useUser = () => {
    const context = useContext(AuthContext);
    if (context === undefined) {
        throw new Error("useUser must be used within an AuthProvider");
    }
    return context;
};