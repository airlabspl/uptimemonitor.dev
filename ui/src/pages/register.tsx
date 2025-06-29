import { useApp } from "@/app/app-context"
import { useAuth } from "@/auth/auth-context"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Link, Navigate, useNavigate } from "react-router-dom"
import { toast } from "sonner"

export default function RegisterPage() {
    const { selfhosted } = useApp();

    if (selfhosted) {
        return <Navigate to="/login" replace />
    }

    return (
        <div className="min-h-svh flex items-center justify-center p-6 md:p-10">
            <div className="flex flex-col w-full max-w-sm gap-6">
                <div className="flex flex-col gap-1 text-center">
                    <h1 className="text-2xl font-bold">Create a new account</h1>
                    <p className="text-sm text-muted-foreground">
                        Already have an account?{" "}
                        <Link to="/login" className="text-primary underline">
                            Log in
                        </Link>
                    </p>
                </div>
                <RegisterForm />
            </div>
        </div>
    )
}

function RegisterForm() {
    const navigate = useNavigate();
    const { reload } = useAuth();

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        const formData = new FormData(event.currentTarget);

        const res = await fetch("/v1/auth/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                email: formData.get("email"),
                password: formData.get("password"),
                name: formData.get("name"),
                confirm_password: formData.get("confirm_password"),
            }),
        })

        if (!res.ok) {
            toast.error("Registration failed. Please check your email and password.");
            return;
        }

        reload();
        navigate("/dashboard");
    };

    return <form className="w-full max-w-sm flex flex-col gap-4" onSubmit={handleSubmit}>
        <div className="flex flex-col gap-1.5">
            <Label htmlFor="name">Name</Label>
            <Input id="name" type="text" name="name" placeholder="John Doe" autoFocus required />
        </div>
        <div className="flex flex-col gap-1.5">
            <Label htmlFor="email">Email</Label>
            <Input id="email" type="email" name="email" placeholder="email@example.com" required />
        </div>
        <div className="flex flex-col gap-1.5">
            <Label htmlFor="password">Password</Label>
            <Input id="password" type="password" name="password" placeholder="••••••••" required />
        </div>
        <div className="flex flex-col gap-1.5">
            <Label htmlFor="confirm-password">Confirm Password</Label>
            <Input id="confirm-password" type="password" name="confirm_password" placeholder="••••••••" required />
        </div>
        <div>
            <Button type="submit" className="w-full">Sign up</Button>
        </div>
    </form>
}