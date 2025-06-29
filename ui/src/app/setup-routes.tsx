import { Navigate, Outlet } from "react-router-dom";
import { useApp } from "./app-context";

export const SetupRoutes = () => {
    const { setupFinished, loading } = useApp();

    if (loading) {
        return null;
    }

    if (!setupFinished) {
        return <Navigate to="/setup" />
    }

    return <Outlet />
}