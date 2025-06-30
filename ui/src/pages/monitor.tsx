import { useParams } from "react-router-dom"

export default function () {
    const params = useParams()

    return <div>
        {params.uuid}
    </div>
}