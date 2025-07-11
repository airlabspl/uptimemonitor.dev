import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "./auth-context";

export const ProtectedRoutes = () => {
    const { user, loading } = useAuth();

    if (loading) {
        return null;
    }

    if (user === null) {
        return <Navigate to="/login" replace />;
    }

    return <Outlet />
}