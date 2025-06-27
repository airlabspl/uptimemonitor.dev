import { useUser } from "./auth-context";
import { Navigate, Outlet } from "react-router-dom";

export const GuestRoutes = () => {
    const user = useUser();

    if (user !== null) {
        return <Navigate to="/" replace />;
    }

    return <Outlet />
}