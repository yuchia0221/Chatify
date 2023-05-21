import { Link } from "react-router-dom";
import useAuth from "../hooks/useAuth";

const Navbar = () => {
    const { auth } = useAuth();

    return (
        <nav className="border-gray-200">
            <div className="mx-auto flex max-w-screen-xl flex-wrap items-center justify-between p-4">
                <Link
                    to="/"
                    className="flex items-center gap-x-4 fill-black stroke-black hover:fill-blue-600 hover:stroke-blue-600 hover:text-blue-600"
                >
                    <svg
                        strokeWidth="0"
                        role="img"
                        viewBox="0 0 24 24"
                        height="1em"
                        width="1em"
                        xmlns="http://www.w3.org/2000/svg"
                    >
                        <title></title>
                        <path d="M23.849 14.91c-.24 2.94-2.73 5.22-5.7 5.19h-3.15l-6 3.9v-3.9l6-3.9h3.15c.93.03 1.71-.66 1.83-1.59.18-3 .18-6-.06-9-.06-.84-.75-1.47-1.56-1.53-2.04-.09-4.2-.18-6.36-.18s-4.32.06-6.36.21c-.84.06-1.5.69-1.56 1.53-.21 3-.24 6-.06 9 .09.93.9 1.59 1.83 1.56h3.15v3.9h-3.15a5.644 5.644 0 01-5.7-5.19c-.21-3.21-.18-6.39.06-9.6a5.57 5.57 0 015.19-5.1c2.1-.15 4.35-.21 6.6-.21s4.5.06 6.63.24a5.57 5.57 0 015.19 5.1c.21 3.18.24 6.39.03 9.57z"></path>
                    </svg>
                    <span className="self-center whitespace-nowrap text-2xl font-semibold">Chatify</span>
                </Link>
                <div className="block w-auto">
                    <ul className="flex flex-row gap-x-8 rounded-lg border-0 p-0 font-medium">
                        {auth.isLoggedIn ? (
                            <li>
                                <Link to="/" className="block rounded py-2 pl-3 pr-4 hover:text-blue-600">
                                    Home
                                </Link>
                            </li>
                        ) : (
                            <>
                                <li>
                                    <Link to="/login" className="block rounded py-2 pl-3 pr-4 hover:text-blue-600">
                                        Login
                                    </Link>
                                </li>
                                <li>
                                    <Link to="/register" className="block rounded py-2 pl-3 pr-4 hover:text-blue-600">
                                        Register
                                    </Link>
                                </li>
                            </>
                        )}
                    </ul>
                </div>
            </div>
        </nav>
    );
};

export default Navbar;
