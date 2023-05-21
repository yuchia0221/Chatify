import React from "react";
import { Navigate, Outlet, useLocation } from "react-router-dom";
import useAuth from "../hooks/useAuth";

const PublicRoute = () => {
    const { auth } = useAuth();
    const location = useLocation();
    const navigate_link = location.state?.from ? location.state.from.pathname : "/";

    return auth?.isLoggedIn ? (
        <Navigate to={navigate_link} state={{ from: location?.state?.from?.pathname || location }} replace />
    ) : (
        <Outlet />
    );
};

export default PublicRoute;
