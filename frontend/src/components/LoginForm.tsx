import React, { useState } from "react";
import useAuth from "../hooks/useAuth";
import AlertBar from "./AlertBar";
import LinkButton from "./LinkButton";
import SubmitButton from "./SubmitButton";

const LoginForm = () => {
    const { handleAuthLogIn } = useAuth();
    const [message, setMessage] = useState("");

    const handleOnSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const username = event.currentTarget.username.value;
        const password = event.currentTarget.password.value;
        const response = await handleAuthLogIn(username, password);

        if (response.status >= 400 && response.status < 500) {
            setMessage("Invalid username or password");
        } else if (response.status >= 500 && response.status < 600) {
            setMessage("There is something wrong with the server. Please try again later.");
        }
    };

    return (
        <div className="mt-24 flex flex-col content-center items-center justify-center">
            {message.length > 0 && <AlertBar message={message} />}
            <div className="flex w-4/5 flex-col gap-y-4 rounded-lg bg-white p-6 py-4 shadow-xl drop-shadow-md md:w-2/5">
                <div className="text-center text-lg font-bold">Sign in your account</div>
                <form onSubmit={handleOnSubmit}>
                    <div className="mb-3 md:mb-6">
                        <label htmlFor="username" className="mb-2 inline-block text-gray-700">
                            Username
                        </label>
                        <input
                            type="text"
                            className="block w-full rounded-md border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900"
                            id="username"
                            placeholder="Enter username"
                            required={true}
                        />
                    </div>
                    <div className="mb-3 md:mb-6">
                        <label htmlFor="password" className="mb-2 inline-block text-gray-700">
                            Password
                        </label>
                        <input
                            type="password"
                            className="block w-full rounded-md border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900"
                            id="password"
                            placeholder="Password"
                            required={true}
                        />
                    </div>
                    <div className="flex justify-center justify-items-center">
                        <SubmitButton buttonName={"Sign in"} />
                    </div>
                </form>
                <div className="inline-flex w-full items-center justify-center">
                    <hr className="h-px w-full border-0 bg-gray-200" />
                    <span className="absolute left-1/2 -translate-x-1/2 bg-white px-3 text-gray-800">or</span>
                </div>
                <div className="mt-2 text-center text-gray-800">
                    Not a member?&nbsp;&nbsp;&nbsp;
                    <LinkButton path="/register" name="Register" />
                </div>
            </div>
        </div>
    );
};

export default LoginForm;
