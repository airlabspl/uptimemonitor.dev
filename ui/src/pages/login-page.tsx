import { useAuth } from "@/auth/auth-context"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useNavigate } from "react-router-dom"
import { toast } from "sonner"

export default function LoginPage() {
    return (
        <div className="min-h-svh flex items-center justify-center p-6 md:p-10">
            <LoginForm />
        </div>
    )
}

function LoginForm() {
    const navigate = useNavigate();
    const { reload } = useAuth();

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        const formData = new FormData(event.currentTarget);

        const res = await fetch("/v1/auth/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                email: formData.get("email"),
                password: formData.get("password"),
            }),
        })

        if (!res.ok) {
            toast.error("Login failed. Please check your email and password.");
            return;
        }

        reload();
        toast.success("Logged in successfully!");
        navigate("/dashboard");
    };

    return <form className="w-full max-w-sm flex flex-col gap-4" onSubmit={handleSubmit}>
        <div className="flex flex-col gap-1.5">
            <h1 className="text-2xl font-semibold">Log in</h1>
            <p className="text-sm text-muted-foreground">Enter your email and password to log in.</p>
        </div>
        <div className="flex flex-col gap-1.5">
            <Label htmlFor="email">Email</Label>
            <Input id="email" type="email" name="email" placeholder="email@example.com" autoFocus required />
        </div>
        <div className="flex flex-col gap-1.5">
            <div className="flex items-center justify-between gap-4">
                <Label htmlFor="password">Password</Label>
                <a href="/forgot-password" className="text-sm text-primary underline">
                    Forgot password?
                </a>
            </div>
            <Input id="password" type="password" name="password" placeholder="••••••••" required />
        </div>
        <div>
            <Button type="submit" className="w-full">Log in</Button>
        </div>
    </form>
}