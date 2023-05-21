import Home from "../components/Home";
import Navbar from "../components/Navbar";

const HomePage = () => {
    return (
        <div className="flex h-screen flex-col">
            <Navbar />
            <Home />
        </div>
    );
};

export default HomePage;
