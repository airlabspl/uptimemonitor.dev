import { IconDashboard } from "@tabler/icons-react"

import {
  SidebarGroup,
  SidebarGroupContent,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar"
import { NavLink } from "react-router-dom"

export function NavMain() {
  return (
    <SidebarGroup>
      <SidebarGroupContent className="flex flex-col gap-2">
        <SidebarMenu>
          <NavLink to="/dashboard">
            {({ isActive }) => (
              <SidebarMenuItem >
                <SidebarMenuButton tooltip="Dashboard" isActive={isActive}>
                  <IconDashboard />
                  <span>Dashboard</span>
                </SidebarMenuButton>
              </SidebarMenuItem>
            )}
          </NavLink>
        </SidebarMenu>
      </SidebarGroupContent>
    </SidebarGroup>
  )
}
