type PropsType = {
    message: string;
};

const SuccessBar = ({ message }: PropsType) => {
    return (
        <div
            className="mb-4 w-4/5 rounded-lg bg-green-50 p-4 text-center text-sm font-medium text-green-800 md:w-2/5"
            role="alert"
        >
            {message}
        </div>
    );
};

export default SuccessBar;
