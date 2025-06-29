import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "./auth-context";

export const VerifiedRoutes = () => {
    const { user, loading } = useAuth();

    if (loading) {
        return null;
    }

    if (!user?.verified) {
        return <Navigate to="/verify" replace />;
    }

    return <Outlet />
}