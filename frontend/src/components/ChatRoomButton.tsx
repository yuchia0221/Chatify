type PropsType = {
    roomId: string;
    roomName: string;
    handleJoinRoom: (roomId: string) => void;
};

const ChatRoomButton = ({ roomId, roomName, handleJoinRoom }: PropsType) => {
    return (
        <div
            className="flex w-11/12 cursor-pointer items-center gap-2 break-all rounded-md px-3 py-3 hover:text-blue-600"
            onClick={() => handleJoinRoom(roomId)}
        >
            <svg
                stroke="currentColor"
                fill="currentColor"
                strokeWidth="0"
                viewBox="0 0 16 16"
                height="1em"
                width="1em"
                className="w-1/6"
                xmlns="http://www.w3.org/2000/svg"
            >
                <path d="M14 1a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H4.414A2 2 0 0 0 3 11.586l-2 2V2a1 1 0 0 1 1-1h12zM2 0a2 2 0 0 0-2 2v12.793a.5.5 0 0 0 .854.353l2.853-2.853A1 1 0 0 1 4.414 12H14a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2H2z"></path>
            </svg>
            <div className="w-5/6 truncate text-ellipsis">{roomName}</div>
        </div>
    );
};

export default ChatRoomButton;
