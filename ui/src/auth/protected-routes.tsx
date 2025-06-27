import { useUser } from "./auth-context";
import { Navigate, Outlet } from "react-router-dom";

export const ProtectedRoutes = () => {
    const user = useUser();

    if (user === null) {
        return <Navigate to="/login" replace />;
    }

    return <Outlet />
}