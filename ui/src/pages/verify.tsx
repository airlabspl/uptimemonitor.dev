import { Button } from "@/components/ui/button";
import { FormEvent, useState } from "react";
import { toast } from "sonner";

export default function VerifyPage() {
    const [resent, setResent] = useState(false)

    const handleResend = async (event: FormEvent<HTMLButtonElement>) => {
        event.preventDefault()

        if (resent) return;

        const res = await fetch('/v1/auth/resend-verification', {
            method: "POST"
        })

        if (!res.ok) {
            toast.error("We were unable to resend verification email. Please try again later")
            return
        }

        setResent(true)
        toast.success("The verification email has been resent")
    }
    return (
        <div className="flex flex-col items-center justify-center h-screen">
            <h1 className="text-2xl font-bold mb-4">Email Verification Required</h1>
            <p className="text-sm text-muted-foreground mb-6">
                Please verify your email address to access this page.
            </p>

            <Button className="mb-4" onClick={handleResend} disabled={resent}>
                Resend Verification Email
            </Button>
        </div>
    );
}