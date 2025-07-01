import { SidebarTrigger } from "@/components/ui/sidebar"
import CreateMonitor from "@/monitors/create-monitor"
import { Button } from "./ui/button"
import { IconCirclePlusFilled } from "@tabler/icons-react"

export function SiteHeader() {
  return (
    <header className="flex h-(--header-height) shrink-0 items-center gap-2 border-b transition-[width,height] ease-linear group-has-data-[collapsible=icon]/sidebar-wrapper:h-(--header-height)">
      <div className="flex w-full items-center justify-between gap-1 px-2 lg:gap-2 lg:px-2">
        <SidebarTrigger className="-ml-1" />
        <div>
          <CreateMonitor>
            <Button
              className="bg-primary text-primary-foreground hover:bg-primary/90 hover:text-primary-foreground active:bg-primary/90 active:text-primary-foreground min-w-8 duration-200 ease-linear"
            >
              <IconCirclePlusFilled />
              <span>New Monitor</span>
            </Button>
          </CreateMonitor>
        </div>
      </div>
    </header>
  )
}
