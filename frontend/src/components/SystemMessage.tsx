type PropsType = {
    type: string;
    content: string;
};

const SystemMessage = ({ type, content }: PropsType) => {
    if (type !== "self") {
        return (
            <div className="flex flex-col gap-y-2">
                <div className="flex flex-row items-center justify-center">
                    <div className="relative ml-3 px-4 py-2 text-sm">{content}</div>
                </div>
            </div>
        );
    } else return <></>;
};

export default SystemMessage;
