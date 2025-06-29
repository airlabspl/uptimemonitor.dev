import { useApp } from "@/app/app-context";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Navigate, useNavigate } from "react-router-dom";
import { toast } from "sonner";

export default function SetupPage() {
    const { setupFinished, reload } = useApp();
    const navigate = useNavigate();

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const formData = new FormData(event.currentTarget);

        const res = await fetch(`/v1/setup`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                name: formData.get('name'),
                email: formData.get('email'),
                password: formData.get('password'),
                confirm_password: formData.get('confirm_password'),
            }),
        })

        if (!res.ok) {
            toast.error("Setup failed. Please check your inputs.");
            return;
        }

        reload();
        toast.success("Setup completed successfully!");
        navigate("/dashboard");
    };

    if (setupFinished) {
        return <Navigate to="/dashboard" replace />;
    }

    return <div className="flex flex-col items-center justify-center min-h-svh gap-6 p-6 md:p-10">
        <div className="flex flex-col w-full max-w-sm gap-6">
            <div className="flex flex-col gap-1 text-center">
                <h1 className="text-2xl font-bold">Setup Your Account</h1>
                <p className="text-sm text-muted-foreground">
                    Please complete the setup to start using the application.
                </p>
            </div>
            <form className="w-full max-w-sm flex flex-col gap-4" onSubmit={handleSubmit}>
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
        </div>
    </div>
}