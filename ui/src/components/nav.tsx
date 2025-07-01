import { IconCirclePlusFilled, IconDashboard } from "@tabler/icons-react"

import {
  SidebarGroup,
  SidebarGroupContent,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar"
import CreateMonitor from "@/monitors/create-monitor"
import { Link } from "react-router-dom"

export function NavMain() {
  return (
    <SidebarGroup>
      <SidebarGroupContent className="flex flex-col gap-2">
        <SidebarMenu>
          <SidebarMenuItem className="flex items-center gap-2">
            <CreateMonitor>
              <SidebarMenuButton
                tooltip="New Monitor"
                className="bg-primary text-primary-foreground hover:bg-primary/90 hover:text-primary-foreground active:bg-primary/90 active:text-primary-foreground min-w-8 duration-200 ease-linear"
              >
                <IconCirclePlusFilled />
                <span>New Monitor</span>
              </SidebarMenuButton>
            </CreateMonitor>
          </SidebarMenuItem>
        </SidebarMenu>
        <SidebarMenu>
          <Link to="/dashboard">
            <SidebarMenuItem>
              <SidebarMenuButton tooltip="Dashboard">
                <IconDashboard />
                <span>Dashboard</span>
              </SidebarMenuButton>
            </SidebarMenuItem>
          </Link>
        </SidebarMenu>
      </SidebarGroupContent>
    </SidebarGroup>
  )
}
