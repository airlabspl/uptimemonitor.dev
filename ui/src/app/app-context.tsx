import { createContext, useCallback, useContext, useEffect, useState } from "react";

type AppContextType = {
    setupFinished: boolean;
    reload: () => void;
    loading: boolean;
}

const AppContext = createContext<AppContextType | undefined>(undefined);

export const AppProvider = ({ children }: { children: React.ReactNode }) => {
    const [loading, setLoading] = useState<boolean>(true);
    const [setupFinished, setSetupFinished] = useState<boolean>(false);

    const reload = useCallback(() => {
        fetch(`/v1/app`, {
            method: "GET",
        })
            .then((response) => response.json())
            .then((data) => {
                setSetupFinished(data.setup_finished);
            })
            .finally(() => setLoading(false));
    }, []);

    useEffect(() => reload(), [reload]);

    return (
        <AppContext.Provider value={{
            setupFinished,
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