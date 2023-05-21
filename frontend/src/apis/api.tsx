import axios from "axios";

export default axios.create({
    baseURL: process.env.REACT_APP_BACKEND_URL,
    headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
    },
    withCredentials: true,
});
