import { useAuth } from "@/auth/auth-context";

export default function Dashboard() {
    const { user } = useAuth();

    return <div>
        <h1 className="text-2xl font-bold">Welcome, {user?.name}!</h1>
        <p className="text-sm text-muted-foreground">
            This is your dashboard. You can manage your settings and view your data here.
        </p>
    </div>
}