import { Link } from "react-router-dom";

export default function HomePage() {
    return (
        <div className="flex flex-col items-center justify-center h-screen">
            <h1 className="text-4xl font-bold mb-4">Welcome to the Home Page</h1>
            <p className="text-lg">This is a protected route.</p>

            <Link to="/dashboard" className="mt-4 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600">
                Go to Dashboard
            </Link>
        </div>
    );
}