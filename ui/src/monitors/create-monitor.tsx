import { Sheet, SheetContent, SheetDescription, SheetHeader, SheetTitle, SheetTrigger } from "@/components/ui/sheet";
import { ReactNode } from "react";

export default function CreateMonitor({ children }: { children: ReactNode }) {
    return <Sheet>
        <SheetTrigger className="w-full cursor-pointer">
            {children}
        </SheetTrigger>
        <SheetContent>
            <SheetHeader>
                <SheetTitle>Are you absolutely sure?</SheetTitle>
                <SheetDescription>
                    This action cannot be undone. This will permanently delete your account
                    and remove your data from our servers.
                </SheetDescription>
            </SheetHeader>
        </SheetContent>
    </Sheet>
}