import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "./auth-context";

export const GuestRoutes = () => {
    const { user } = useAuth();

    if (user !== null) {
        return <Navigate to="/dashboard" replace />;
    }

    return <Outlet />
}