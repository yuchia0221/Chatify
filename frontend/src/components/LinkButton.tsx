import React from "react";
import { Link } from "react-router-dom";

type PropsType = {
    path: string;
    name: string;
};

const LinkButton = ({ path, name }: PropsType) => {
    return (
        <Link to={path}>
            <span className="cursor-pointer text-blue-600 transition duration-200 ease-in-out hover:font-bold">
                {name}
            </span>
        </Link>
    );
};

export default LinkButton;
