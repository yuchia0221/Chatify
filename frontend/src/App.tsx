import { Route, HashRouter as Router, Routes } from "react-router-dom";
import Layout from "./components/Layout";
import HomePage from "./pages/HomePage";
import LoginPage from "./pages/LoginPage";
import RegisterPage from "./pages/RegisterPage";
import PublicRoute from "./services/PublicRoute";
import RequireAuth from "./services/RequireAuth";

const App = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Layout />}>
                    {/* public routes */}
                    <Route element={<PublicRoute />}>
                        <Route path="/login" element={<LoginPage />} />
                        <Route path="/register" element={<RegisterPage />} />
                    </Route>

                    {/* private routes */}
                    <Route element={<RequireAuth />}>
                        <Route path="/" element={<HomePage />} />
                    </Route>
                </Route>
            </Routes>
        </Router>
    );
};

export default App;
