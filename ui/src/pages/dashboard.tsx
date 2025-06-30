import { DataTable } from "@/components/data-table";
import { SectionCards } from "@/components/section-cards";
import data from "@/app/dashboard/data.json"

export default function Dashboard() {
    return (
        <>
            <SectionCards />
            <DataTable data={data} />
        </>
    )
}