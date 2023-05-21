import { useState } from "react";
import ChatMessage from "./ChatMessage";
import { MessageType } from "./Home";
import SystemMessage from "./SystemMessage";

type PropsType = {
    conn: WebSocket | null;
    messages: MessageType[];
};

const ChatRoom = ({ conn, messages }: PropsType) => {
    const [message, setMessage] = useState<string>("");

    const handleSendMessage = () => {
        if (message === "" || conn === null) {
            return;
        }

        conn.send(message);
        setMessage("");
    };

    return (
        <div className="flex h-screen w-4/5 flex-auto flex-col px-6">
            <div className="flex h-full flex-auto flex-shrink-0 flex-col gap-y-4 rounded-2xl bg-gray-100 p-4">
                <div className="flex h-4/5 flex-col gap-y-2 overflow-x-auto p-2">
                    {messages.map((message: MessageType, index: number) => {
                        if (message.content.endsWith("joined the room") || message.content.endsWith("left the room")) {
                            return <SystemMessage key={index} type={message.type} content={message.content} />;
                        } else {
                            return (
                                <ChatMessage
                                    key={index}
                                    type={message.type}
                                    content={message.content}
                                    username={message.username}
                                />
                            );
                        }
                    })}
                </div>
                <div className="flex h-14 w-full flex-row items-center rounded-xl bg-white px-2 py-4">
                    <div className="flex-grow">
                        <div className="relative w-full">
                            <input
                                type="text"
                                className="flex h-10 w-full rounded-xl border pl-4 focus:outline-none"
                                value={message}
                                onChange={(event) => setMessage(event.target.value)}
                            />
                        </div>
                    </div>
                    <div className="ml-4">
                        <button
                            className="flex flex-shrink-0 items-center justify-center rounded-xl bg-blue-500 px-4 py-1 text-white hover:bg-blue-600"
                            onClick={handleSendMessage}
                        >
                            <span>Send</span>
                            <span className="ml-2">
                                <svg
                                    className="-mt-px h-4 w-4 rotate-45 transform"
                                    fill="none"
                                    stroke="currentColor"
                                    viewBox="0 0 24 24"
                                    xmlns="http://www.w3.org/2000/svg"
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        strokeWidth="2"
                                        d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"
                                    ></path>
                                </svg>
                            </span>
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default ChatRoom;
