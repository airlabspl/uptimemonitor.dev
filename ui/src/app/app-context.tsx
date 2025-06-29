import { createContext, useCallback, useContext, useEffect, useState } from "react";

type AppContextType = {
    setupFinished: boolean;
    selfhosted: boolean;
    reload: () => void;
    loading: boolean;
}

const AppContext = createContext<AppContextType | undefined>(undefined);

export const AppProvider = ({ children }: { children: React.ReactNode }) => {
    const [loading, setLoading] = useState<boolean>(true);
    const [data, setData] = useState<{
        setup_finished: boolean;
        selfhosted: boolean;
    } | null>(null);

    const reload = useCallback(() => {
        fetch(`/v1/app`, {
            method: "GET",
        })
            .then((response) => response.json())
            .then((data) => {
                setData(data);
            })
            .finally(() => setLoading(false));
    }, []);

    useEffect(() => reload(), [reload]);

    return (
        <AppContext.Provider value={{
            setupFinished: data?.setup_finished ?? false,
            selfhosted: data?.selfhosted ?? false,
            reload,
            loading,
        }}>
            {children}
        </AppContext.Provider>
    );
}

export const useApp = () => {
    const context = useContext(AppContext);
    if (context === undefined) {
        throw new Error("useApp must be used within an AppProvider");
    }
    return context;
};