import { useState } from "react";

type PropsType = {
    handleCreateRoom: (name: string) => void;
};

const CreateRoomForm = ({ handleCreateRoom }: PropsType) => {
    const [roomName, setRoomName] = useState("");

    const handleOnChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setRoomName(event.target.value);
    };

    const handleOnClick = () => {
        handleCreateRoom(roomName);
        setRoomName("");
    };

    return (
        <form className="mx-2 flex flex-row gap-x-2">
            <label htmlFor="createRoom" className="inline-block text-gray-700"></label>
            <input
                type="text"
                className="block w-full rounded-md border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 outline-none"
                id="createRoom"
                placeholder="Create Room"
                value={roomName}
                onChange={handleOnChange}
                required={true}
            />
            <button
                type="submit"
                onClick={handleOnClick}
                className="rounded-lg bg-black fill-white stroke-white p-3 text-sm font-medium hover:bg-blue-600 hover:fill-white hover:stroke-white focus:outline-none"
            >
                <svg
                    strokeWidth="2"
                    viewBox="0 0 24 24"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    className="h-4 w-4"
                    height="1em"
                    width="1em"
                    xmlns="http://www.w3.org/2000/svg"
                >
                    <line x1="12" y1="5" x2="12" y2="19"></line>
                    <line x1="5" y1="12" x2="19" y2="12"></line>
                </svg>
                <span className="sr-only">Add Room</span>
            </button>
        </form>
    );
};

export default CreateRoomForm;
