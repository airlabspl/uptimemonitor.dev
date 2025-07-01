import { fetcher } from "@/api/fetcher";
import { SidebarGroup, SidebarGroupContent, SidebarGroupLabel, SidebarMenu, SidebarMenuButton, SidebarMenuItem } from "@/components/ui/sidebar";
import { Link } from "react-router-dom";
import useSWR from "swr"

type Monitor = {
    uuid: string;
    url: string;
}

export default function NavMonitors() {
    const { data, isLoading, error } = useSWR<{ monitors: Monitor[] }>(`/v1/monitors`, fetcher)

    if (isLoading) return <div>Loading</div>
    if (error) return <div>error</div>

    return (
        <SidebarGroup>
            <SidebarGroupLabel>Monitors</SidebarGroupLabel>
            <SidebarGroupContent>
                <SidebarMenu>
                    {(data?.monitors ?? []).map(m => (
                        <SidebarMenuItem key={m.uuid}>
                            <SidebarMenuButton asChild>
                                <Link to={`/m/${m.uuid}`}>
                                    <span>{m.url}</span>
                                </Link>
                            </SidebarMenuButton>
                        </SidebarMenuItem>
                    ))}
                </SidebarMenu>
            </SidebarGroupContent>
        </SidebarGroup>
    )
}