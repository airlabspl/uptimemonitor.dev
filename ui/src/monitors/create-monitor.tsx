import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Sheet, SheetContent, SheetDescription, SheetFooter, SheetHeader, SheetTitle, SheetTrigger } from "@/components/ui/sheet";
import { FormEvent, ReactNode, useState } from "react";
import { useNavigate } from "react-router-dom";
import { toast } from "sonner";
import { mutate } from "swr";

export default function CreateMonitor({ children }: { children: ReactNode }) {
    const navigate = useNavigate()
    const [open, setOpen] = useState(false)

    const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        const formData = new FormData(event.currentTarget)

        const res = await fetch(`/v1/monitors`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify({
                url: formData.get("url"),
            })
        })

        if (!res.ok) {
            toast.error("Could not create a monitor")
            return
        }

        const data = await res.json()

        toast.success("Monitor created")
        mutate(`/v1/monitors`)
        navigate(`/m/${data.uuid}`)
        setOpen(false)
    }
    return <Sheet onOpenChange={o => setOpen(o)} open={open}>
        <SheetTrigger className="w-full cursor-pointer" onToggle={() => setOpen(o => !o)}>
            {children}
        </SheetTrigger>
        <SheetContent>
            <form onSubmit={handleSubmit} className="h-full flex flex-col">
                <SheetHeader>
                    <SheetTitle>Add url to monitor</SheetTitle>
                    <SheetDescription>
                        We will be monitoring the given url periodically for you and notify when it is down.
                    </SheetDescription>
                </SheetHeader>
                <div className="flex-1 px-4">
                    <div className="flex flex-col gap-1">
                        <Label htmlFor="url">Website URL</Label>
                        <Input id="url" name="url" type="url" required autoFocus placeholder="https://example.com/page/123" />
                    </div>
                </div>
                <SheetFooter>
                    <Button type="submit" className="w-full">
                        Create monitor
                    </Button>
                </SheetFooter>
            </form>
        </SheetContent>
    </Sheet >
}
