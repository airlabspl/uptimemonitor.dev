import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Sheet, SheetContent, SheetDescription, SheetFooter, SheetHeader, SheetTitle, SheetTrigger } from "@/components/ui/sheet";
import { FormEvent, ReactNode } from "react";

export default function CreateMonitor({ children }: { children: ReactNode }) {
    const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        // const formData = new FormData(event.currentTarget)

        // fetch(`/v1/`)
    }
    return <Sheet>
        <SheetTrigger className="w-full cursor-pointer">
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
