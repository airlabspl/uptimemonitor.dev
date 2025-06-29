import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import React from "react"
import { Link, useNavigate, useParams } from "react-router-dom"
import { toast } from "sonner"

export default function ResetPassword() {
    return <div className="flex items-center justify-center p-6 md:p-10 min-h-svh">
        <div className="flex flex-col gap-6 text-center w-full max-w-sm">
            <div className="flex flex-col gap-1">
                <h1 className="text-2xl font-bold">Reset password</h1>
                <p className="text-sm text-muted-foreground">
                    Remember your password?{" "}
                    <Link to="/login" className="text-primary underline">
                        Back to log in
                    </Link>
                </p>
            </div>
            <ResetPasswordForm />
        </div>
    </div>
}

function ResetPasswordForm() {
    const params = useParams()
    const navigate = useNavigate()

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        const formData = new FormData(event.currentTarget)

        const res = await fetch(`/v1/auth/reset-password`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                token: params.token,
                password: formData.get("password"),
                confirm_password: formData.get("confirm_password"),
            })
        })

        if (!res.ok) {
            toast.error("Unable to reset pasword. Please try again later")
            return
        }

        toast.success("Password has been changed. You can now log in.")
        navigate('/login')
    }

    return <form onSubmit={handleSubmit} className="w-full max-w-sm flex flex-col gap-4">
        <div className="flex flex-col gap-1.5">
            <Label htmlFor="password">New Password</Label>
            <Input id="password" type="password" name="password" placeholder="••••••••" required />
        </div>
        <div className="flex flex-col gap-1.5">
            <Label htmlFor="confirm-password">Confirm New Password</Label>
            <Input id="confirm-password" type="password" name="confirm_password" placeholder="••••••••" required />
        </div>
        <div>
            <Button type="submit" className="w-full">
                Send reset link
            </Button>
        </div>
    </form>
}