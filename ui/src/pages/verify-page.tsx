import { Button } from "@/components/ui/button";

export default function VerifyPage() {
    return (
        <div className="flex flex-col items-center justify-center h-screen">
            <h1 className="text-2xl font-bold mb-4">Email Verification Required</h1>
            <p className="text-sm text-muted-foreground mb-6">
                Please verify your email address to access this page.
            </p>
            <Button className="mb-4" asChild>
                <a href="/resend-verification">Resend Verification Email</a>
            </Button>
        </div>
    );
}