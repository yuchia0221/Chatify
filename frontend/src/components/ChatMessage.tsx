type PropsType = {
    type: string;
    content: string;
    username: string;
};

const ChatMessage = ({ type, content, username }: PropsType) => {
    if (type === "self") {
        return (
            <div className="flex flex-col gap-y-2">
                <div className="flex flex-row items-center justify-end">
                    <div className="relative rounded-xl bg-white px-4 py-2 text-sm shadow">{content}</div>
                    <div className="ml-3 flex h-10 w-10 flex-shrink-0 items-center justify-center rounded-full bg-blue-500 text-white">
                        {username[0].toUpperCase()}
                    </div>
                </div>
            </div>
        );
    } else {
        return (
            <div className="flex flex-col gap-y-2">
                <div className="flex flex-row items-center justify-start">
                    <div className="mr-3 flex h-10 w-10 flex-shrink-0 items-center justify-center rounded-full bg-blue-500 text-white">
                        {username[0].toUpperCase()}
                    </div>
                    <div className="relative  rounded-xl bg-white px-4 py-2 text-sm shadow">{content}</div>
                </div>
            </div>
        );
    }
};

export default ChatMessage;
