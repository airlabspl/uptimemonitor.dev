import { fetcher } from "@/api/fetcher";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Link } from "react-router-dom";
import useSWR from "swr"

type Monitor = {
    uuid: string;
    url: string;
}

export default function MonitorsTable() {
    const { data, isLoading, error } = useSWR<{ monitors: Monitor[] }>(`/v1/monitors`, fetcher)

    if (isLoading) return <div>Loading</div>
    if (error) return <div>error</div>

    return (
        <Table>
            <TableHeader className="bg-muted sticky top-0 z-10">
                <TableRow>
                    <TableHead>Url</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody className="**:data-[slot=table-cell]:first:w-8">
                {data?.monitors.map(monitor => (
                    <TableRow key={monitor.uuid}>
                        <TableCell>
                            <Link to={`/m/${monitor.uuid}`}>
                                {monitor.url}
                            </Link>
                        </TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </Table>
    )
}