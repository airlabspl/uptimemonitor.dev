import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import React, { useState } from "react"
import { Link } from "react-router-dom"
import { toast } from "sonner"

export default function ResetPasswordLink() {
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
            <ResetPasswordLinkForm />
        </div>
    </div>
}

function ResetPasswordLinkForm() {
    const [success, setSuccess] = useState(false)

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        const formData = new FormData(event.currentTarget)

        const res = await fetch(`/v1/auth/password-reset-link`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                email: formData.get("email"),
            })
        })

        if (!res.ok) {
            toast.error("Unable to send a password reset link. Please try again later")
            return
        }

        toast.success("Password reset link has been sent to your e-mail address")
        setSuccess(true)
    }

    return <form onSubmit={handleSubmit} className="w-full max-w-sm flex flex-col gap-4">
        <div className="flex flex-col gap-1.5">
            <Label htmlFor="email">E-mail address</Label>
            <Input id="email" type="email" name="email" autoFocus placeholder="email@example.com" required disabled={success} />
        </div>
        <div>
            <Button type="submit" className="w-full" disabled={success}>
                Send reset link
            </Button>
        </div>
    </form>
}