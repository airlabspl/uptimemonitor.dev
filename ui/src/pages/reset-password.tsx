import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import React from "react"
import { Link } from "react-router-dom"

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
    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
    }

    return <form onSubmit={handleSubmit} className="w-full max-w-sm flex flex-col gap-4">
        <div className="flex flex-col gap-1.5">
            <Label htmlFor="email">E-mail address</Label>
            <Input id="email" type="email" name="email" autoFocus placeholder="email@example.com" required />
        </div>
        <div>
            <Button type="submit" className="w-full">
                Send reset link
            </Button>
        </div>
    </form>
}