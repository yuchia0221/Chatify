import { useEffect, useState } from "react";
import api from "../apis/api";
import useAuth from "../hooks/useAuth";
import ChatRoom from "./ChatRoom";
import ChatRoomButton from "./ChatRoomButton";
import CreateRoomForm from "./CreateRoomForm";

type ChatRoomType = {
    id: string;
    name: string;
};

export type MessageType = {
    content: string;
    username: string;
    room_id: string;
    send_time: string;
    type: "recv" | "self";
};

const Home = () => {
    const { user } = useAuth();
    const [clients, setClients] = useState<string[]>([]);
    const [conn, setConn] = useState<WebSocket | null>(null);
    const [chatRooms, setChatRooms] = useState<{ id: string; name: string }[]>([]);
    const [messages, setMessages] = useState<MessageType[] | null>(null);

    const getAllChatRooms = async () => {
        api.get("/websocket/rooms").then((res) => {
            setChatRooms(res.data.rooms);
        });
    };

    const getChatRoomMessages = async (roomId: string) => {
        api.get(`/websocket/messages/${roomId}`).then((res) => {
            let response = res.data.messages;
            response.forEach((element: MessageType) => {
                element.type = element.username === user.username ? "self" : "recv";
            });

            setMessages(response);
        });
    };

    const handleCreateRoom = async (roomName: string) => {
        try {
            const response = await api.post("/websocket/room", { name: roomName });
            const { room_id } = response.data;
            setChatRooms((prev) => {
                return prev === null ? [{ id: room_id, name: roomName }] : [...prev, { id: room_id, name: roomName }];
            });
        } catch (error: any) {
            return;
        }
    };

    const handleJoinRoom = async (roomId: string) => {
        const websocketUrl = process.env.REACT_APP_WEBSOCKET_URL;
        if (websocketUrl === undefined) {
            return;
        }

        await getChatRoomMessages(roomId);

        const ws = new WebSocket(`${websocketUrl}/joinRoom/${roomId}`);

        if (ws.OPEN) setConn(ws);
    };

    useEffect(() => {
        getAllChatRooms();
    }, []);

    useEffect(() => {
        if (conn === null) return;

        conn.onmessage = (message) => {
            const m: MessageType = JSON.parse(message.data);
            if (m.content.endsWith("joined the room")) {
                setClients((prev) => {
                    return prev === null ? [m.username] : [...prev, m.username];
                });
            }

            if (m.content.endsWith("left the chat")) {
                setClients((prev) => {
                    return prev === null ? [m.username] : clients.filter((user) => user !== m.username);
                });
                setMessages((prev) => {
                    return prev === null ? [m] : [...prev, m];
                });
                return;
            }

            user.username === m.username ? (m.type = "self") : (m.type = "recv");
            setMessages((prev) => {
                return prev === null ? [m] : [...prev, m];
            });
        };

        conn.onclose = () => {};
        conn.onerror = () => {};
        conn.onopen = () => {};

        // eslint-disable-next-line
    }, [conn]);

    return (
        <div className="flex overflow-hidden">
            <div className="mt-2 flex h-full w-1/5 flex-col gap-y-3 overflow-y-auto">
                <CreateRoomForm handleCreateRoom={handleCreateRoom} />
                {chatRooms ? (
                    chatRooms.map((chatRoom: ChatRoomType) => {
                        return (
                            <div className="flex w-full justify-center" key={chatRoom.id}>
                                <ChatRoomButton
                                    roomId={chatRoom.id}
                                    roomName={chatRoom.name}
                                    handleJoinRoom={handleJoinRoom}
                                />
                            </div>
                        );
                    })
                ) : (
                    <></>
                )}
            </div>
            {messages !== null ? <ChatRoom conn={conn} messages={messages} /> : <></>}
        </div>
    );
};

export default Home;
