type PropsType = {
    message: string;
};

const AlertBar = ({ message }: PropsType) => {
    return (
        <div
            className="mb-4 w-4/5 rounded-lg bg-red-50 p-4 text-center text-sm font-medium text-red-800 md:w-2/5"
            role="alert"
        >
            {message}
        </div>
    );
};

export default AlertBar;
