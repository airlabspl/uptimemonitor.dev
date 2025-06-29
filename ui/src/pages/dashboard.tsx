import { useAuth } from "@/auth/auth-context";
import { Button } from "@/components/ui/button";

export default function Dashboard() {
    const { user, logout } = useAuth();

    return <div>
        <h1 className="text-2xl font-bold">Welcome, {user?.name}!</h1>
        <p className="text-sm text-muted-foreground">
            This is your dashboard. You can manage your settings and view your data here.
        </p>

        <Button className="mt-4" onClick={() => logout()}>
            Logout
        </Button>
    </div>
}